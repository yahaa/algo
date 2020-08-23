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
