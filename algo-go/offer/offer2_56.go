package offer

func findTarget(root *TreeNode, k int) bool {
	var dfs func(r *TreeNode) bool
	m := make(map[int]bool)
	dfs = func(r *TreeNode) bool {
		if r == nil {
			return false
		}

		d := k - r.Val

		if _, v := m[d]; v {
			return true
		}

		m[r.Val] = true

		return dfs(r.Left) || dfs(r.Right)
	}

	return dfs(root)
}
