package offer

import "container/heap"

type ListHeap []*ListNode

func (l ListHeap) Len() int {
	return len(l)
}

func (l ListHeap) Less(i, j int) bool {
	return l[i].Val < l[j].Val
}

func (l ListHeap) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l *ListHeap) Push(v any) {
	*l = append(*l, v.(*ListNode))
}

func (l *ListHeap) Pop() any {
	old := *l
	n := len(old)
	*l = old[:n-1]
	return old[n-1]
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := &ListHeap{}
	heap.Init(h)

	for _, l := range lists {
		if l != nil {
			heap.Push(h, l)
		}
	}

	head := &ListNode{}
	cur := head

	for h.Len() > 0 {
		pop := heap.Pop(h).(*ListNode)

		cur.Next = pop
		cur = pop

		pop = pop.Next

		if pop != nil {
			heap.Push(h, pop)
		}
	}

	return head.Next
}
