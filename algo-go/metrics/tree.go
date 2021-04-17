package metrics

import (
	"strings"
)

const (
	Aggv = "{}"
)

// Node tree
type Node struct {
	Value string
	// IsAgg 节点是否是聚合节点
	IsAgg bool
	// Children 节点的下一层
	Children map[string]*Node
}

// PathTree prefix tree
type PathTree struct {
	rootValue string
	root      map[string]*Node
	// 每一层最多能有多少个节点，超过了这个值该层就会被聚合为一个 Aggv 字符串
	maxChild int
}

// Query query path in prefix tree return the converted path
func (t *PathTree) Query(path string) string {
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
			curs = nx.Children
			index++
			continue
		}

		agg, ok := curs[Aggv]
		if ok {
			result = append(result, Aggv)
			curs = agg.Children
			index++
			continue
		}

		break
	}

	if index < len(parts) {
		result = append(result, parts[index:]...)
		// todo async insert
		t.doInsert(curs, parts[index:])
		// todo async aggraga
		t.aggrega(t.root)
	}

	// 去掉 root 节点
	return "/" + strings.Join(result[1:], "/")
}

// aggrega 对孩子过多的层级做聚合
func (t *PathTree) aggrega(children map[string]*Node) {
	newChildren := make(map[string]*Node)
	if len(children) >= t.maxChild {
		// 遍历当前层所有节点
		for key, child := range children {
			for kk, cc := range child.Children {
				// 保存当前孩子的所有孩子节点
				newChildren[kk] = cc
			}
			// 当前孩子的孩子已经保存，可以删除当前孩子了
			delete(children, key)
		}

		aggNode := &Node{
			Value:    Aggv,
			IsAgg:    true,
			Children: newChildren,
		}

		children[aggNode.Value] = aggNode
	}
	// 递归进行下一层聚合
	for _, child := range children {
		t.aggrega(child.Children)
	}
}

// split 统一给每一个 path 设置上 root value
func (t *PathTree) split(path string) []string {
	path = strings.Trim(path, "/")
	path = t.rootValue + "/" + path
	return strings.Split(path, "/")
}

// insert parts of path into tree
func (t *PathTree) doInsert(curs map[string]*Node, parts []string) {
	for _, p := range parts {
		nx := &Node{
			Value:    p,
			Children: make(map[string]*Node),
		}
		curs[p] = nx

		curs = nx.Children
	}
}

// NewTree new
func NewTree(max int) *PathTree {
	tree := &PathTree{
		rootValue: "root",
		root: map[string]*Node{
			"root": {
				Value:    "root",
				Children: make(map[string]*Node),
			},
		},
		maxChild: max,
	}

	return tree
}
