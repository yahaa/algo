package tree

// Node node int value
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// IsUnivalTree leetcode 965
func IsUnivalTree(root *Node) bool {
	if root == nil {
		return true
	}

	var dfs func(r *Node, pVal int) bool

	dfs = func(r *Node, pVal int) bool {
		if r == nil {
			return true
		}
		return r.Val == pVal && dfs(r.Left, r.Val) && dfs(r.Right, r.Val)
	}

	return dfs(root, root.Val)
}

// leetcode 589
type NNode struct {
	Val      int
	Children []*NNode
}

// preorder
func preorder(root *NNode) []int {
	res := make([]int, 0)
	var dfs func(root *NNode)

	dfs = func(root *NNode) {
		if root == nil {
			return
		}

		res = append(res, root.Val)

		for _, c := range root.Children {
			dfs(c)
		}
	}

	dfs(root)
	return res
}
