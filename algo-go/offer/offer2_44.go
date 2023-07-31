package offer

func largestValues(root *TreeNode) []int {
	ans := make([]int, 0)

	var dfs func(root *TreeNode, level int)

	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if len(ans) < level {
			ans = append(ans, root.Val)
		} else {
			ans[level-1] = max(ans[level-1], root.Val)
		}

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 1)

	return ans
}
