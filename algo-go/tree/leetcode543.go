package tree

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func diameterOfBinaryTree(root *TreeNode) int {
	var dfs func(root *TreeNode) int
	ans := 1

	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := dfs(root.Left)
		right := dfs(root.Right)

		ans = max(ans, left+right+1)

		return max(left, right) + 1
	}

	dfs(root)

	return ans - 1
}
