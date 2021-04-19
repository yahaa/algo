package metrics

import (
	"strings"
	"sync"
)

const (
	aggValue  = "*"
	rootValue = "root"
)

// node tree
type node struct {
	value string
	// isAggre 节点是否是聚合节点
	isAggre bool
	// children 节点的下一层
	children map[string]*node
}

// AggreTree 聚合路径前缀树
type AggreTree struct {
	sync.RWMutex
	rootValue string
	root      map[string]*node
	// 每一层最多能有多少个节点，超过了这个值该层就会被聚合为一个 aggValue 字符串
	maxChild int
}

// Query query path
func (t *AggreTree) Query(path string) string {
	t.RLock()
	defer t.RUnlock()

	var (
		parts  = t.split(path)
		index  = 0
		curs   = t.root
		result = make([]string, 0)
	)

	for index < len(parts) {
		nx, ok := curs[parts[index]]
		if ok {
			result = append(result, parts[index])
			curs = nx.children
			index++
			continue
		}

		agg, ok := curs[aggValue]
		if ok {
			result = append(result, aggValue)
			curs = agg.children
			index++
			continue
		}

		break
	}

	if index < len(parts) {
		result = append(result, parts[index:]...)
		// todo async insert
		go t.doInsert(curs, parts[index:])
	}

	// 去掉 root 节点
	return "/" + strings.Join(result[1:], "/")
}

// aggre 对孩子过多的层级做聚合
func (t *AggreTree) aggre(children map[string]*node) {
	newChildren := make(map[string]*node)
	if len(children) >= t.maxChild {
		// 遍历当前层所有节点
		for key, child := range children {
			for kk, cc := range child.children {
				// 保存当前孩子的所有孩子节点
				newChildren[kk] = cc
			}
			// 当前孩子的孩子已经保存，可以删除当前孩子了
			delete(children, key)
		}

		aggnode := &node{
			value:    aggValue,
			isAggre:  true,
			children: newChildren,
		}

		children[aggnode.value] = aggnode
	}
	// 递归进行下一层聚合
	for _, child := range children {
		t.aggre(child.children)
	}
}

// split 统一给每一个 path 设置上 root value
func (t *AggreTree) split(path string) []string {
	path = strings.Trim(path, "/")
	path = t.rootValue + "/" + path
	return strings.Split(path, "/")
}

// insert parts of path into tree
func (t *AggreTree) doInsert(curs map[string]*node, parts []string) {
	t.Lock()
	defer t.Unlock()

	for _, p := range parts {
		nx := &node{
			value:    p,
			children: make(map[string]*node),
		}
		curs[p] = nx

		curs = nx.children
	}

	t.aggre(t.root)
}

// NewAggreTree new
func NewAggreTree(max int) *AggreTree {
	tree := &AggreTree{
		rootValue: rootValue,
		root: map[string]*node{
			rootValue: {
				value:    rootValue,
				children: make(map[string]*node),
			},
		},
		maxChild: max,
	}

	return tree
}
