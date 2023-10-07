package stack

import (
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
