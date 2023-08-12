package heap

import (
	"container/heap"
	"container/list"
	"sort"
)

type pointMaxHeap [][]int

func (h pointMaxHeap) Len() int {
	return len(h)
}

func (h pointMaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h pointMaxHeap) Less(i, j int) bool {
	a := h[i][0]*h[i][0] + h[i][1]*h[i][1]
	b := h[j][0]*h[j][0] + h[j][1]*h[j][1]
	return a > b
}

func (h *pointMaxHeap) Push(v interface{}) {
	*h = append(*h, v.([]int))
}

func (h *pointMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
	h := pointMaxHeap(points)
	heap.Init(&h)

	for h.Len() > k {
		heap.Pop(&h)
	}

	return h
}

type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

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

func (h minHeap) Top() interface{} {
	return h[0]
}

type maxHeap []int

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h *maxHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h maxHeap) Top() interface{} {
	return h[0]
}

type MedianFinder struct {
	min *minHeap
	max *maxHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	m := MedianFinder{
		min: &minHeap{},
		max: &maxHeap{},
	}

	heap.Init(m.min)
	heap.Init(m.max)
	return m
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.min, num)
	heap.Push(this.max, heap.Pop(this.min))

	if this.max.Len() > this.min.Len() {
		heap.Push(this.min, heap.Pop(this.max))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if (this.max.Len()+this.min.Len())%2 == 1 {
		return float64(this.min.Top().(int))
	}
	return float64(this.min.Top().(int)+this.max.Top().(int)) / 2
}

func findKthLargest(nums []int, k int) int {
	h := &minHeap{}
	heap.Init(h)

	for _, n := range nums {
		heap.Push(h, n)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := heap.Pop(h)
	return res.(int)
}

func maxSlidingWindow(nums []int, k int) []int {
	dqueue, res := list.New(), make([]int, 0)

	for i, n := range nums {
		if i >= k {
			if nums[i-k] == dqueue.Back().Value.(int) {
				dqueue.Remove(dqueue.Back())
			}
		}

		for dqueue.Len() > 0 && dqueue.Front().Value.(int) < n {
			dqueue.Remove(dqueue.Front())
		}
		dqueue.PushFront(n)

		if i+1 >= k {
			res = append(res, dqueue.Back().Value.(int))
		}
	}

	return res
}

func lastStoneWeight(stones []int) int {
	h := maxHeap(stones)
	heap.Init(&h)

	for h.Len() > 1 {
		rest := heap.Pop(&h).(int) - heap.Pop(&h).(int)
		heap.Push(&h, rest)
	}

	return heap.Pop(&h).(int)
}

// KthLargest b
type KthLargest struct {
	h *minHeap
	k int
}

func Constructor2(k int, nums []int) KthLargest {
	h := minHeap(nums)

	heap.Init(&h)

	for h.Len() > k {
		heap.Pop(&h)
	}

	return KthLargest{h: &h, k: k}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.h, val)

	if this.h.Len() > this.k {
		heap.Pop(this.h)
	}

	return this.h.Top().(int)
}

func stringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func frequencySort(s string) string {
	mp := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		mp[s[i]]++
	}

	bs := []byte(s)
	sort.Slice(bs, func(i, j int) bool {
		if mp[bs[i]] == mp[bs[j]] {
			return bs[i] > bs[j]
		}

		return mp[bs[i]] > mp[bs[j]]
	})

	return string(bs)
}

type numCount struct {
	v     int
	count int
}

type numCountMinHeap []numCount

func (h numCountMinHeap) Len() int {
	return len(h)
}

func (h numCountMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h numCountMinHeap) Less(i, j int) bool {

	return h[i].count < h[j].count
}

func (h *numCountMinHeap) Push(v interface{}) {
	*h = append(*h, v.(numCount))
}

func (h *numCountMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	mp := make(map[int]int)
	for _, n := range nums {
		mp[n]++
	}

	h := &numCountMinHeap{}
	heap.Init(h)

	for num, count := range mp {
		heap.Push(h, numCount{v: num, count: count})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := make([]int, h.Len())

	for i, nc := range *h {
		res[i] = nc.v
	}

	return res
}

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	left, right := matrix[0][0], matrix[n-1][n-1]

	for left < right {
		mid := left + (right-left)/2

		count := getSmallEqualNum(matrix, mid)

		if count >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func getSmallEqualNum(matrix [][]int, target int) int {
	i, j, count := len(matrix)-1, 0, 0

	for i >= 0 && j < len(matrix) {
		if matrix[i][j] <= target {
			count += i + 1
			j++
		} else {
			i--
		}
	}
	return count
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type minHeapNode []*ListNode

func (h minHeapNode) Len() int           { return len(h) }
func (h minHeapNode) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h minHeapNode) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeapNode) Push(v any) {
	*h = append(*h, v.(*ListNode))
}

func (h *minHeapNode) Pop() any {
	old := *h
	n := len(old)
	res := old[n-1]
	old = old[:n-1]

	*h = old
	return res
}

func mergeKLists(lists []*ListNode) *ListNode {
	head := &ListNode{}
	p := head

	h := &minHeapNode{}
	heap.Init(h)

	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(h, lists[i])
		}
	}

	for h.Len() > 0 {
		tmp := heap.Pop(h).(*ListNode)
		p.Next = tmp
		p = p.Next

		if tmp.Next != nil {
			heap.Push(h, tmp.Next)
		}
	}

	return head.Next
}
