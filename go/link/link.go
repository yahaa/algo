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

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// copyRandomList leetcode 138
func copyRandomList(head *Node) *Node {
	var cur = head
	var mp = make(map[*Node]*Node)

	for cur != nil {
		n := &Node{Val: cur.Val}

		mp[cur] = n

		cur = cur.Next
	}

	cur = head

	for cur != nil {
		mp[cur].Next = mp[cur.Next]
		mp[cur].Random = mp[cur.Random]
		cur = cur.Next
	}

	return mp[head]
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head, c := &ListNode{}, 0
	p := head
	for l1 != nil && l2 != nil {
		val := (l1.Val + l2.Val + c) % 10
		c = (l1.Val + l2.Val + c) / 10

		p.Next = &ListNode{Val: val}
		p = p.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	var last *ListNode

	if l1 != nil {
		last = l1
	}

	if l2 != nil {
		last = l2
	}

	for last != nil {
		val := (last.Val + c) % 10
		c = (last.Val + c) / 10

		p.Next = &ListNode{Val: val}
		p = p.Next
		last = last.Next
	}

	if c != 0 {
		p.Next = &ListNode{Val: c}
	}

	return head.Next
}