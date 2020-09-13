package offer

func getKthFromEnd(head *ListNode, k int) *ListNode {
	p, res := head, head

	for p != nil {
		p = p.Next

		if k > 0 {
			k--
		} else {
			res = res.Next
		}
	}

	return res
}
