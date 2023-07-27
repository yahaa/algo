package offer

func isSameTree(a *TreeNode, b *TreeNode) bool {
	if b == nil {
		return true
	}

	if a == nil || a.Val != b.Val {
		return false
	}

	return isSameTree(a.Left, b.Left) && isSameTree(a.Right, b.Right)
}

func isSubStructure(a *TreeNode, b *TreeNode) bool {
	if a == nil || b == nil {
		return false
	}

	return isSameTree(a, b) || isSubStructure(a.Left, b) || isSubStructure(a.Right, b)
}
