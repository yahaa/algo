package offer

func rightSideView(root *TreeNode) []int {
	lasts := make([]int, 0)

	var dfs func(root *TreeNode, level int)

	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if len(lasts) < level {
			lasts = append(lasts, root.Val)
		} else {
			lasts[level-1] = root.Val
		}

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 1)

	return lasts
}
