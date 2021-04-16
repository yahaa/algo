package metrics

import (
	"strings"
	"sync"
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
	sync.RWMutex
	root  *Node
	paths chan string
	// 每一层最多能有多少个节点，超过了这个值该层就会被聚合为一个 Aggv 字符串
	maxChild int
}

// Query query path in prefix tree return the converted path
func (t *PathTree) Query(path string) string {
	t.RLock()
	defer t.RUnlock()

	var (
		parts      = t.split(path)
		needInsert = false
		isAgg      = t.root.IsAgg
		children   = t.root.Children
		result     = make([]string, 0)
	)

	for _, p := range parts {
		if isAgg {
			p = Aggv
		}

		node, ok := children[p]
		if !ok {
			needInsert = true
			break
		}

		result = append(result, node.Value)
		children = node.Children
		isAgg = node.IsAgg
	}

	if !needInsert {
		return strings.Replace(strings.Join(result, "/"), t.root.Value, "", 1)
	}

	t.paths <- path
	return path
}

func (t *PathTree) asyncInsert() {
	for path := range t.paths {
		t.doInsert(path)
	}
}

func (t *PathTree) aggrega(r *Node) {
	if r == nil {
		return
	}

	children := make(map[string]*Node)
	// 孩子过多需要聚合
	if len(r.Children) > t.maxChild {
		for childv, child := range r.Children {
			for vv, cc := range child.Children {
				// 保存当前孩子的所有孩子节点
				children[vv] = cc
			}
			// 当前孩子的孩子已经保存，可以删除当前孩子了
			delete(r.Children, childv)
		}

		aggNode := &Node{
			Value:    Aggv,
			IsAgg:    true,
			Children: children,
		}

		r.Children[aggNode.Value] = aggNode
	}

	// 递归进行下层
	for _, child := range r.Children {
		t.aggrega(child)
	}
}

// split 统一给每一个 path 设置上 root value
func (t *PathTree) split(path string) []string {
	path = strings.Trim(path, "/")
	path = t.root.Value + "/" + path
	return strings.Split(path, "/")
}

// insert parts of path into tree
func (t *PathTree) doInsert(path string) {
	t.Lock()
	defer t.Unlock()

	var (
		parts  = t.split(path)
		childs = t.root.Children
		isAgg  = t.root.IsAgg
	)

	for _, p := range parts {
		if isAgg {
			p = Aggv
		}

		node, ok := childs[p]
		if !ok {
			node = &Node{
				Value:    p,
				Children: make(map[string]*Node),
			}
			childs[p] = node
		}

		childs = node.Children
		isAgg = node.IsAgg
	}

	t.aggrega(t.root)
}

// NewTree new
func NewTree(max int) *PathTree {
	tree := &PathTree{
		root: &Node{
			Value:    "=root=",
			Children: make(map[string]*Node),
		},
		paths:    make(chan string, 10000),
		maxChild: max,
	}

	go tree.asyncInsert()

	return tree
}
