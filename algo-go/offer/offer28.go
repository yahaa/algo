package offer

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var dfs func(p, q *TreeNode) bool

	dfs = func(p, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}

		if p != nil && q == nil || p == nil && q != nil {
			return false
		}

		if p.Val != q.Val {
			return false
		}

		return dfs(p.Left, q.Right) && dfs(p.Right, q.Left)
	}

	return dfs(root.Left, root.Right)
}
