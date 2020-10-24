package tree

// IncreasingBST leetcode 897
func IncreasingBST(root *Node) *Node {
	var (
		inOrder func(root *Node)
		preNode = &Node{}
		res     = preNode
	)

	inOrder = func(root *Node) {
		if root == nil {
			return
		}

		inOrder(root.Left)
		root.Left = nil
		preNode.Right = root
		preNode = root
		inOrder(root.Right)
	}

	inOrder(root)
	return res.Right
}
