package dfs

type Node struct {
	Val       int
	Neighbors []*Node
}

// leetcode 133
func cloneGraph(node *Node) *Node {
	var visit func(node *Node, mp map[int]*Node) *Node

	visit = func(node *Node, mp map[int]*Node) *Node {
		if node == nil {
			return nil
		}

		if _, ok := mp[node.Val]; !ok {
			n := &Node{Val: node.Val}
			mp[node.Val] = n
			for _, ne := range node.Neighbors {
				n.Neighbors = append(n.Neighbors, visit(ne, mp))
			}
		}

		return mp[node.Val]
	}

	return visit(node, make(map[int]*Node))
}
