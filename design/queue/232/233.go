package lc232

import "container/list"

// MyQueue leetcode 232 用 stack 实现队列
// 实现思路: 两个栈模拟,一个栈 s1 用来实现入队，另外一个栈 s2 实现出栈
// 1. 出栈的时候，如果 s2 为空，把 s1 中所有元素都倒入 s2 中
// 2. s2 栈顶出栈
type MyQueue struct {
	s1    *list.List
	s2    *list.List
	count int
}

// Constructor
func Constructor() MyQueue {
	return MyQueue{
		s1: list.New(),
		s2: list.New(),
	}
}

// Push
// 1. x 入栈 s1
// 2. count++
func (q *MyQueue) Push(x int) {
	q.s1.PushBack(x)
	q.count++
}

// Pop
// 1. 如果 s2 为空，把 s1 中所有元素倒入 s2
// 2. s2 执行出栈
// 3. count--
func (q *MyQueue) Pop() int {
	if q.Empty() {
		return -1
	}

	if q.s2.Len() == 0 {
		for q.s1.Len() > 0 {
			e := q.s1.Back()
			q.s1.Remove(e)

			q.s2.PushBack(e.Value.(int))
		}
	}

	e := q.s2.Back()
	q.s2.Remove(e)

	q.count--

	return e.Value.(int)
}

// Peek
func (q *MyQueue) Peek() int {
	if q.Empty() {
		return -1
	}

	if q.s2.Len() == 0 {
		for q.s1.Len() > 0 {
			e := q.s1.Back()
			q.s1.Remove(e)

			q.s2.PushBack(e.Value.(int))
		}
	}

	return q.s2.Back().Value.(int)
}

// Empty
func (q *MyQueue) Empty() bool {
	return q.count == 0
}
