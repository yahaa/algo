package offer

import (
	"container/heap"
)

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
	k    int
	heap minHeap
}

func Constructor1(k int, nums []int) KthLargest {
	h := minHeap(nums)

	heap.Init(&h)

	for h.Len() > k {
		heap.Pop(&h)
	}

	return KthLargest{
		k:    k,
		heap: h,
	}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(&this.heap, val)

	if this.heap.Len() > this.k {
		heap.Pop(&this.heap)
	}

	return this.heap[0]
}
