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
