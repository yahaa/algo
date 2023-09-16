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

func nthUglyNumber(n int) int {
	h := &minHeap{1}
	heap.Init(h)

	m := make(map[int]struct{})
	factors := []int{2, 3, 5}
	for i := 1; i <= n; i++ {
		x := heap.Pop(h).(int)

		if i == n {
			return x
		}

		for _, f := range factors {
			n := x * f
			if _, ok := m[n]; !ok {
				heap.Push(h, n)
				m[n] = struct{}{}
			}
		}
	}

	return 0
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

// findKthLargest2 leetcode 215 找到数组中第 k 大的数字
func findKthLargest2(nums []int, k int) int {
	h := &minHeap{}
	heap.Init(h)

	for i := 0; i < len(nums); i++ {
		heap.Push(h, nums[i])

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return (*h)[0]
}

type item struct {
	val  int
	freq int
}

type minHeapItem []item

func (h minHeapItem) Len() int           { return len(h) }
func (h minHeapItem) Less(i, j int) bool { return h[i].freq < h[j].freq }
func (h minHeapItem) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeapItem) Push(v any) {
	*h = append(*h, v.(item))
}
func (h *minHeapItem) Pop() any {
	old := *h
	n := len(old)
	res := old[n-1]
	old = old[:n-1]
	*h = old

	return res
}

func topKFrequent2(nums []int, k int) []int {
	hashMap := make(map[int]int)

	for _, num := range nums {
		hashMap[num]++
	}

	h := &minHeapItem{}
	heap.Init(h)
	for val, freq := range hashMap {
		heap.Push(h, item{val, freq})

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := make([]int, 0)
	for h.Len() > 0 {
		tmp := heap.Pop(h).(item)
		res = append(res, tmp.val)
	}

	return res
}

// rearrangeBarcodes leetcode 1054 重排条形码，使得任意相邻条形码不相同
func rearrangeBarcodes(barcodes []int) []int {
	hashMap := make(map[int]int)

	for _, num := range barcodes {
		hashMap[num]++
	}

	h := &minHeap{}
	heap.Init(h)
	for val, freq := range hashMap {
		heap.Push(h, item{val, freq})
	}

	res := make([]int, 0)
	for h.Len() > 0 {
		t1 := heap.Pop(h).(item)

		if len(res) == 0 || res[len(res)-1] != t1.val {
			res = append(res, t1.val)

			t1.freq--

			if t1.freq > 0 {
				heap.Push(h, t1)
			}

			continue
		}

		if res[len(res)-1] == t1.val && h.Len() > 0 {
			t2 := heap.Pop(h).(item)

			heap.Push(h, t1)

			res = append(res, t2.val)

			t2.freq--

			if t2.freq > 0 {
				heap.Push(h, t2)
			}
		}
	}

	return res
}

type arrayIndex struct {
	array *[]int
	index int
}

type maxHeap2 []arrayIndex

func (h maxHeap2) Len() int      { return len(h) }
func (h maxHeap2) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h maxHeap2) Less(i, j int) bool {
	return (*h[i].array)[h[i].index] < (*h[j].array)[h[j].index]
}
func (h *maxHeap2) Push(v any) {
	*h = append(*h, v.(arrayIndex))
}
func (h *maxHeap2) Pop() any {
	old := *h
	n := len(old)
	res := old[n-1]
	old = old[:n-1]
	*h = old

	return res
}

func kthSmallest2(matrix [][]int, k int) int {
	h := &maxHeap2{}

	heap.Init(h)

	for i := 0; i < len(matrix); i++ {
		if len(matrix[i]) == 0 {
			continue
		}

		heap.Push(h, arrayIndex{&matrix[i], 0})
	}

	count := 1

	for count < k {
		top := heap.Pop(h).(arrayIndex)

		top.index++

		if top.index < len(*top.array) {
			heap.Push(h, top)
		}

		count++
	}

	top := heap.Pop(h).(arrayIndex)

	return (*top.array)[top.index]
}

type pairs struct {
	sum int
	i   int
	j   int
}

type minHeap1 []pairs

func (h minHeap1) Len() int           { return len(h) }
func (h minHeap1) Less(i, j int) bool { return h[i].sum < h[j].sum }
func (h minHeap1) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap1) Push(v any)        { *h = append(*h, v.(pairs)) }
func (h *minHeap1) Pop() any {
	old := *h
	n := len(old)
	res := old[n-1]
	old = old[:n-1]
	*h = old

	return res
}

// kSmallestPairs 373. 查找和最小的 K 对数字
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	res := make([][]int, 0)
	h := &minHeap1{}
	heap.Init(h)
	n, m := len(nums1), len(nums2)

	for i := 0; i < n; i++ {
		heap.Push(h, pairs{nums1[i] + nums2[0], i, 0})
	}

	for h.Len() > 0 && k > 0 {
		top := heap.Pop(h).(pairs)

		res = append(res, []int{nums1[top.i], nums2[top.j]})
		k--

		if top.j+1 < m {
			heap.Push(h, pairs{nums1[top.i] + nums2[top.j+1], top.i, top.j + 1})
		}
	}
	return res
}

// [1,1,2] [1,2,3]

// (0,0) (0,1) (0,2)
// (1,0) (1,1) (1,2)
// (2,0) (2,1) (2,2)

// (1,1) (1,2) (1,3)
// (1,1) (1,2) (1,3)
// (2,1) (2,2) (2,3)
