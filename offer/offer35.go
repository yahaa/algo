package offer

// Node 链表节点
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// offer35 leetcode 138
func copyRandomList(head *Node) *Node {
	mp := make(map[*Node]*Node)
	p := head
	for p != nil {
		n := &Node{Val: p.Val}
		mp[p] = n

		p = p.Next
	}

	p = head

	for p != nil {
		mp[p].Next = mp[p.Next]
		mp[p].Random = mp[p.Random]
		p = p.Next
	}

	return mp[head]
}
