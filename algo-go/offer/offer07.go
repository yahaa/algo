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

// buildTree2  offer 07 先序遍历中序遍历重建二叉树
func buildTree2(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}

	i := 0

	for i = 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}

	root.Left = buildTree2(preorder[1:i+1], inorder[:i])
	root.Right = buildTree2(preorder[i+1:], inorder[i+1:])

	return root
}
