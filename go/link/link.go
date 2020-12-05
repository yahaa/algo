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

func deleteDuplicates(head *ListNode) *ListNode {
	var fakeHead = &ListNode{Val: 0}
	var pre = fakeHead
	var cur = head

	for cur != nil {
		var count = 1

		var n = cur.Next

		for n != nil && n.Val == cur.Val {
			count++
			n = n.Next
		}

		if count == 1 {
			next := cur.Next
			cur.Next = nil

			pre.Next = cur
			pre = cur

			cur = next
		} else {
			cur = n
		}
	}
	return fakeHead.Next
}

func partition(head *ListNode, x int) *ListNode {
	var (
		fakeLessHead = &ListNode{}
		lessTail     = fakeLessHead
		fakeBigHead  = &ListNode{}
		bigTail      = fakeBigHead
	)

	for head != nil {
		t := head
		head = head.Next

		t.Next = nil

		if t.Val < x {
			lessTail.Next = t
			lessTail = lessTail.Next
		} else {
			bigTail.Next = t
			bigTail = bigTail.Next
		}
	}

	lessTail.Next = fakeBigHead.Next

	return fakeLessHead.Next
}

func orderInsert(node *ListNode, head *ListNode) *ListNode {
	var fakeHead = &ListNode{Next: head}
	var pre = fakeHead
	var index = fakeHead.Next

	for index != nil {
		// find the first node which Val >= node.Val
		if index.Val >= node.Val {
			break
		}

		index = index.Next
		pre = pre.Next
	}

	pre.Next = node
	node.Next = index

	return fakeHead.Next
}
