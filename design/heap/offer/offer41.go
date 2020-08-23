package offer

import (
	"container/heap"
)

type MaxIntHeap []int

func (h MaxIntHeap) Len() int {
	return len(h)
}

func (h MaxIntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MaxIntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h *MaxIntHeap) Push(v interface{}) {
	value := v.(int)
	*h = append(*h, value)
}

func (h *MaxIntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[0 : n-1]
	return v

}

type MinIntHeap []int

func (h MinIntHeap) Len() int {
	return len(h)
}

func (h MinIntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MinIntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h *MinIntHeap) Push(v interface{}) {
	value := v.(int)
	*h = append(*h, value)
}

func (h *MinIntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[0 : n-1]
	return v
}

type MedianFinder struct {
	maxIntHeap *MaxIntHeap
	minIntHeap *MinIntHeap
	count      int
}

func Constructor() MedianFinder {
	// maxIntHeap ... minIntHeap
	maxIntHeap := &MaxIntHeap{}
	minIntHeap := &MinIntHeap{}

	heap.Init(minIntHeap)
	heap.Init(maxIntHeap)

	return MedianFinder{maxIntHeap, minIntHeap, 0}
}

func (f *MedianFinder) AddNum(num int) {
	f.count++
	if f.count%2 == 1 {
		heap.Push(f.maxIntHeap, num)

		heap.Push(f.minIntHeap, heap.Pop(f.maxIntHeap))
	} else {
		heap.Push(f.minIntHeap, num)

		heap.Push(f.maxIntHeap, heap.Pop(f.minIntHeap))
	}

}

func (f *MedianFinder) FindMedian() float64 {
	if f.count == 0 {
		return 0
	}

	min := *f.minIntHeap
	if f.count%2 == 1 {
		return float64(min[0])
	}

	max := *f.maxIntHeap

	return float64(max[0]+min[0]) / 2
}
