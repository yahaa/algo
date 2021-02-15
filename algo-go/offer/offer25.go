package offer

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var res, pre *ListNode

	for l1 != nil || l2 != nil {
		var t *ListNode

		if l1 == nil && l2 != nil {
			t = l2
			l2 = l2.Next
		} else if l2 == nil && l1 != nil {
			t = l1
			l1 = l1.Next
		} else if l1.Val <= l2.Val {
			t = l1
			l1 = l1.Next
		} else {
			t = l2
			l2 = l2.Next
		}

		if res == nil {
			res = t
			pre = t
		} else {
			pre.Next = t
			pre = t
		}
	}
	return res
}
