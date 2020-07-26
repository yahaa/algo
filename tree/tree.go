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
