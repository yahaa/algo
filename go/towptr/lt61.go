package towptr

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	var n int
	for p := head; p != nil; p = p.Next {
		n++
	}

	k = k % n

	var s, p = head, head
	var end = head

	for p != nil && k > 0 {
		if p.Next == nil {
			end = p
		}

		p = p.Next

		k--
		if k < 0 {
			s = s.Next
		}
	}

	res := s.Next

	s.Next = nil

	end.Next = head

	return res
}
