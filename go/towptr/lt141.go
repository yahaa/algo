package towptr

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	s, q := head, head.Next

	for q != nil && q.Next != nil && q.Next.Next != nil {
		if q == s {
			return true
		}

		q = q.Next.Next
		s = s.Next
	}

	return false
}
