package offer

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	l := mirrorTree(root.Left)
	r := mirrorTree(root.Right)

	root.Left = r
	root.Right = l
	return root
}
