package stack

import (
	"container/heap"
	"container/list"
	"math"
)

func longestValidParentheses(s string) int {
	return 0
}

// minOperations 1598. 文件夹操作日志搜集器
func minOperations(logs []string) int {
	stack := list.New()

	for _, item := range logs {
		if item == "./" {
			continue
		}

		if item == "../" {
			if stack.Len() > 0 {
				stack.Remove(stack.Front())
			}

			continue
		}

		stack.PushFront(item)
	}

	return stack.Len()
}

type pair struct {
	value int
	index int
}

type StockSpanner struct {
	stack  *list.List
	curDay int
}

func Constructor() StockSpanner {
	stack := list.New()
	stack.PushFront(pair{math.MaxInt, -1})
	return StockSpanner{
		stack:  stack,
		curDay: -1,
	}
}

func (this *StockSpanner) Next(price int) int {
	for price >= this.stack.Front().Value.(pair).value {
		this.stack.Remove(this.stack.Front())
	}

	this.curDay++
	this.stack.PushFront(pair{price, this.curDay})
	return this.curDay - this.stack.Front().Next().Value.(pair).index
}

type MyStack struct {
	que1 *list.List
	que2 *list.List
}

func NewMyStack() MyStack {
	return MyStack{
		que1: list.New(),
		que2: list.New(),
	}
}

func (this *MyStack) Push(x int) {
	this.que1.PushBack(x)
}

func (this *MyStack) Pop() int {
	if this.que1.Len() == 0 {
		return -1
	}

	for this.que1.Len() > 1 {
		e := this.que1.Front()
		this.que1.Remove(e)
		this.que2.PushBack(e.Value.(int))
	}

	res := this.que1.Front()
	this.que1.Remove(res)

	for this.que2.Len() > 0 {
		e := this.que2.Front()
		this.que2.Remove(e)
		this.que1.PushBack(e.Value.(int))
	}

	return res.Value.(int)
}

func (this *MyStack) Top() int {
	if this.que1.Len() == 0 {
		return -1
	}

	for this.que1.Len() > 0 {
		e := this.que1.Front()
		this.que1.Remove(e)
		this.que2.PushBack(e.Value.(int))
	}

	for this.que2.Len() > 1 {
		e := this.que2.Front()
		this.que2.Remove(e)
		this.que1.PushBack(e.Value.(int))
	}

	res := this.que2.Front()
	this.que2.Remove(res)
	this.que1.PushBack(res.Value.(int))

	return res.Value.(int)
}

func (this *MyStack) Empty() bool {
	return this.que1.Len() == 0
}

// isValid 20
func isValid(s string) bool {
	stack := list.New()

	ss := []byte(s)

	for _, c := range ss {
		if c == '[' || c == '{' || c == '(' {
			stack.PushFront(c)
			continue
		}

		if stack.Len() == 0 {
			return false
		}

		t := stack.Front().Value.(byte)

		if c == ']' && t != '[' {
			return false
		}

		if c == ')' && t != '(' {
			return false
		}

		if c == '}' && t != '{' {
			return false
		}

		stack.Remove(stack.Front())
	}

	return stack.Len() != 0
}

type maxHeap []pair

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h maxHeap) Less(i, j int) bool {
	return h[i].value > h[j].value
}

func (h *maxHeap) Push(v any) {
	*h = append(*h, v.(pair))
}

func (h *maxHeap) Pop() any {
	old := *h
	n := len(old)
	tmp := old[n-1]
	*h = old[:n-1]

	return tmp
}

func maxSlidingWindow(nums []int, k int) []int {
	h := maxHeap{}
	heap.Init(&h)

	res := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		heap.Push(&h, pair{value: nums[i], index: i})
		if i+1 >= k {
			for h[0].index <= i-k {
				heap.Pop(&h)
			}

			res = append(res, h[0].value)
		}
	}

	return res
}
