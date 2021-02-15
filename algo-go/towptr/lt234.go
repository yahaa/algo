package towptr

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	s, q := head, head
	for q != nil && q.Next != nil && q.Next.Next != nil {
		s = s.Next
		q = q.Next.Next
	}

	var p, n *ListNode = nil, s.Next
	for n != nil {
		t := n.Next
		n.Next = p
		p = n
		n = t
	}

	s.Next = p

	left, right := head, p

	for left != right && right != nil {
		if left.Val != right.Val {
			return false
		}

		left = left.Next
		right = right.Next
	}

	return true
}
