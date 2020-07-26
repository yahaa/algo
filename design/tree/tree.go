package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type node struct {
	value *TreeNode
	next  *node
}

// 单链表实现栈
type stack struct {
	head *node
}

func (s *stack) Push(n node) {
	n.next = s.head
	s.head = &n
}

func (s *stack) Pop() *node {
	if s.head == nil {
		return nil
	}

	head := s.head
	s.head = s.head.next

	return head
}

func (s *stack) Top() *node {
	return s.head
}

type BSTIterator struct {
	s *stack
}

func Constructor(root *TreeNode) BSTIterator {
	it := BSTIterator{s: &stack{}}
	it.leftInOrder(root)
	return it
}

func (i *BSTIterator) leftInOrder(root *TreeNode) {
	if root == nil {
		return
	}

	n := node{value: root}

	i.s.Push(n)

	i.leftInOrder(root.Left)
}

/** @return the next smallest number */
func (i *BSTIterator) Next() int {
	n := i.s.Pop()

	if n == nil {
		return -1
	}

	i.leftInOrder(n.value.Right)
	return n.value.Val

}

/** @return whether we have a next smallest number */
func (i *BSTIterator) HasNext() bool {
	if i.s.Top() == nil {
		return false
	}

	return true
}
