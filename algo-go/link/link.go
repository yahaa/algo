package link

// TreeNode 二叉搜索树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

// leetcode 876
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

// Node xxx
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

// leetcode 109
// 1. classic problem that use fast and slow points to find the mid node of the linkedlist
// 2. classic problem ths use dfs to build the binary tree
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	var (
		fake       = &ListNode{Next: head}
		slow, fast = fake, fake
	)

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	mid := slow.Next
	slow.Next = nil

	return &TreeNode{
		Val:   mid.Val,
		Left:  sortedListToBST(fake.Next),
		Right: sortedListToBST(mid.Next),
	}
}

// leetcode 141
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}

	return false
}

// leetcode  142 环形连标二
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast, hasCycle := head, head, false

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			hasCycle = true
			break
		}
	}

	if !hasCycle {
		return nil
	}

	for head != slow {
		head = head.Next
		slow = slow.Next
	}

	return head
}

// leetcode 143
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	slow, fast := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	left, right := head, reverList(slow.Next)
	slow.Next = nil

	for left != nil && right != nil {
		leftNext := left.Next
		rightNext := right.Next

		left.Next = right
		right.Next = leftNext

		left = leftNext
		right = rightNext
	}
}

func reverList(head *ListNode) *ListNode {
	var (
		cur = head
		pre *ListNode
	)

	for cur != nil {
		next := cur.Next
		cur.Next = pre

		pre = cur
		cur = next
	}

	return pre
}

func insertionSortList(head *ListNode) *ListNode {
	ascList, cur := &ListNode{}, head

	for cur != nil {
		p := ascList
		for p.Next != nil && cur.Val >= p.Next.Val {
			p = p.Next
		}

		t1 := p.Next
		t2 := cur.Next

		p.Next = cur
		cur.Next = t1

		cur = t2
	}

	return ascList.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	fast, slow := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	right := sortList(slow.Next)

	slow.Next = nil

	left := sortList(head)

	fake := &ListNode{}
	cur := fake

	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}
		cur = cur.Next
	}

	if left != nil {
		cur.Next = left
	}
	if right != nil {
		cur.Next = right
	}

	return fake.Next
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	a, b := headA, headB

	for a != b {
		if a != nil {
			a = a.Next
		} else {
			a = headB
		}

		if b != nil {
			b = b.Next
		} else {
			b = headA
		}
	}

	return a
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	old, even, i := &ListNode{}, &ListNode{}, 1

	oldCur, evenCur := old, even

	for head != nil {
		t := head.Next
		if i%2 == 1 {
			oldCur.Next = head
			oldCur = oldCur.Next
			oldCur.Next = nil
		} else {
			evenCur.Next = head
			evenCur = evenCur.Next
			evenCur.Next = nil
		}

		head = t
		i++
	}

	oldCur.Next = even.Next

	return old.Next
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tail := head

	for tail != nil {
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				tail.Next = l1
				l1 = l1.Next
			} else {
				tail.Next = l2
				l2 = l2.Next
			}
		} else if l1 != nil {
			tail.Next = l1
			break
		} else if l2 != nil {
			tail.Next = l2
			break
		}
		tail = tail.Next
	}

	return head.Next
}

// 	// 1 5 6
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	res := reverseList(head.Next)

	head.Next.Next = head
	head.Next = nil

	return res
}
