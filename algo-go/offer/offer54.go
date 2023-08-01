package offer

func kthLargest(root *TreeNode, k int) int {
	inOrder := make([]int, 0)

	var dfs func(root *TreeNode)

	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		inOrder = append(inOrder, root.Val)
		dfs(root.Right)

	}

	dfs(root)

	return inOrder[len(inOrder)-k]
}
