package lc225

import "container/list"

// MyStack leetcode 225 两个队列模拟 stack
type MyStack struct {
	q1    *list.List
	q2    *list.List
	count int
}

// Constructor 结构体
func Constructor() MyStack {
	s := MyStack{
		q1: list.New(),
		q2: list.New(),
	}
	return s
}

// Push 向 stack  添加元素
// 1. 向 q1 队尾插入 x
// 2. q1 队首元素出队
// 3. 2 中的元素往 q1 队尾中插入
// 4. count++
// 时间复杂度 O(1)
func (s *MyStack) Push(x int) {
	s.q1.PushBack(x)

	if s.q1.Len() > 1 {
		e := s.q1.Front()
		s.q1.Remove(e)

		s.q2.PushBack(e.Value.(int))
	}

	s.count++
}

// Pop 弹出栈顶元素
// 1. q1  出队(这个时候 q1 为空)
// 2. q2 队列元素出队，并入 q1 元素直至 q2 队列长度为 1
// 3. q1 q2 指针互换
// 4. count--
// 时间复杂度 0(n)
func (s *MyStack) Pop() int {
	if s.Empty() {
		return -1
	}

	// 只要 !s.Empty() e 就不会为 nil
	e := s.q1.Front()
	s.q1.Remove(e)

	for s.q2.Len() > 1 {
		t := s.q2.Front()
		s.q2.Remove(t)

		s.q1.PushBack(t.Value.(int))
	}

	if s.q2.Len() == 1 {
		s.q1, s.q2 = s.q2, s.q1
	}

	s.count--

	return e.Value.(int)
}

// Top 获取当前栈顶元素值
// 1. 返回 q1 队首元素值即可
func (s *MyStack) Top() int {
	if s.Empty() {
		return -1
	}

	e := s.q1.Front()
	return e.Value.(int)
}

// Empty 判断栈是否为空
func (s *MyStack) Empty() bool {
	return s.count == 0
}
