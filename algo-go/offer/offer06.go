package offer

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	res := make([]int, 0)
	var dfs func(head *ListNode)

	dfs = func(head *ListNode) {
		if head == nil {
			return
		}

		dfs(head.Next)
		res = append(res, head.Val)
	}

	dfs(head)

	return res
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode = nil

	for head != nil {
		next := head.Next

		head.Next = pre
		pre = head

		head = next
	}

	return pre
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	res := reverseList2(head.Next)

	head.Next.Next = head
	head.Next = nil

	return res
}
