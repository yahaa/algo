package link

// ListNode 单链表 Node
type ListNode struct {
	Val  int
	Next *ListNode
}

// MiddleNode leetcode 876
func MiddleNode(head *ListNode) *ListNode {
	res, count := head, 2
	for head != nil {
		head = head.Next
		count--
		if count == 0 {
			res = res.Next
			count = 2
		}
	}
	return res
}
