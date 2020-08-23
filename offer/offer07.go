package offer

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	var root = &TreeNode{
		Val: preorder[0],
	}

	rootIndex := 0

	for i := 0; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			rootIndex = i
			break
		}
	}

	root.Left = buildTree(preorder[1:1+rootIndex], inorder[:rootIndex])
	root.Right = buildTree(preorder[1+rootIndex:], inorder[rootIndex+1:])

	return root

}
