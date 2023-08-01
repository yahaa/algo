package offer

func increasingBST(root *TreeNode) *TreeNode {
	var head = &TreeNode{}
	var p = head
	var dfs func(root *TreeNode)

	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)

		p.Right = root
		root.Left = nil
		p = root

		dfs(root.Right)
	}

	dfs(root)

	return head.Right
}
