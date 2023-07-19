package offer

import (
	"container/heap"
	"container/list"
)

// MaxQueue 队列最大值
type MaxQueue struct {
	queue     *list.List
	sortQueue *list.List
}

// Constructor bbb
func Constructor() MaxQueue {
	return MaxQueue{
		queue:     list.New(),
		sortQueue: list.New(),
	}
}

// MaxValue v
func (q *MaxQueue) MaxValue() int {
	if q.sortQueue.Len() > 0 {
		return q.sortQueue.Front().Value.(int)
	}

	return -1
}

// PushBack p
func (q *MaxQueue) PushBack(value int) {
	q.queue.PushBack(value)

	for q.sortQueue.Len() > 0 && value > q.sortQueue.Back().Value.(int) {
		b := q.sortQueue.Back()
		q.sortQueue.Remove(b)
	}

	q.sortQueue.PushBack(value)
}

// PopFront f
func (q *MaxQueue) PopFront() int {
	if q.queue.Len() == 0 {
		return -1
	}

	f := q.queue.Front()
	q.queue.Remove(f)

	if q.sortQueue.Front().Value.(int) == f.Value.(int) {
		sf := q.sortQueue.Front()
		q.sortQueue.Remove(sf)
	}

	return f.Value.(int)
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxSlidingWindow(nums []int, k int) []int {
	dqueue := list.New()
	res := make([]int, 0)

	for i := 0; i < k; i++ {
		for dqueue.Len() > 0 && dqueue.Back().Value.(int) < nums[i] {
			dqueue.Remove(dqueue.Back())
		}
		dqueue.PushBack(nums[i])
	}

	res = append(res, dqueue.Front().Value.(int))
	for i := k; i < len(nums); i++ {
		if nums[i-k] == dqueue.Front().Value.(int) {
			dqueue.Remove(dqueue.Front())
		}

		for dqueue.Len() > 0 && dqueue.Back().Value.(int) < nums[i] {
			dqueue.Remove(dqueue.Back())
		}

		dqueue.PushBack(nums[i])
		res = append(res, dqueue.Front().Value.(int))
	}
	return res
}
