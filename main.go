package main

import (
	"container/heap"
	"fmt"
)

type heapSort []int

func (h heapSort) Len() int {
	return len(h)
}

func (h heapSort) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h heapSort) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *heapSort) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *heapSort) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	h := &heapSort{1, 3, 2, 4, 5}
	heap.Init(h)

	heap.Push(h, -10)

	fmt.Printf("%v\n", (*h)[1])

	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
