package offer

func maxPathSum(root *TreeNode) int {
	ans := -999999
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		lv := dfs(root.Left)
		rv := dfs(root.Right)

		ans = max(ans, lv+rv+root.Val)
		return max(max(lv, rv)+root.Val, 0)
	}

	dfs(root)

	return ans
}
