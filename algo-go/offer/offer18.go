package offer

func deleteNode(head *ListNode, val int) *ListNode {
	fake := &ListNode{Next: head}
	pre, cur := fake, head

	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
			break
		}

		cur = cur.Next
		pre = pre.Next
	}

	return fake.Next
}
