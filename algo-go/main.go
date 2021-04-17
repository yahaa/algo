package main

import (
	"github.com/yahaa/leetcode/metrics"
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
	//h := &heapSort{1, 3, 2, 4, 5}
	//heap.Init(h)
	//
	//heap.Push(h, -10)
	//
	//fmt.Printf("%v\n", (*h)[1])
	//
	//for h.Len() > 0 {
	//	fmt.Printf("%d ", heap.Pop(h))
	//}

	tree := metrics.NewTree(3)

	cases := []struct {
		path string
		want string
	}{
		{"/test/0", "/test/0"},
		{"/test/1", "/test/1"},
		{"/test/2", "/test/2"},
		{"/test/3", "/test/3"},
		//{"/test/4", "/test/4"},
		//{"/test/5", "/test/5"},
		//{"/test/6", "/test/6"},
		//{"/test/7", "/test/7"},
		//{"/test/8", "/test/8"},
		//{"/test/9", "/test/9"},
		//{"/test/10", "/test/10"},
		//{"/test/11", "/test/11"},
		//{"/test/12", "/test/12"},
	}

	for _, c := range cases {
		tree.Query(c.path)
	}
}
