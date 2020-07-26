package stack

type node struct {
	value int
	next  *node
}

// 单链表实现栈
type stack struct {
	head *node
}

func (s *stack) Push(n node) {
	n.next = s.head
	s.head = &n
}

func (s *stack) Pop() *node {
	if s.head == nil {
		return nil
	}

	head := s.head
	s.head = s.head.next

	return head
}

func (s *stack) Top() *node {
	return s.head
}

// MinStack 提供返回最小值的栈
type MinStack struct {
	dataStack stack
	minStack  stack
}

// Constructor 构造器
func Constructor() MinStack {
	return MinStack{}
}

func (m *MinStack) Push(x int) {
	n := node{value: x}

	m.dataStack.Push(n)

	min := m.minStack.Top()

	if min == nil || n.value <= min.value {
		m.minStack.Push(n)
	}
}

func (m *MinStack) Pop() {
	t := m.dataStack.Pop()

	if t == nil {
		return
	}

	if m.minStack.Top().value == t.value {
		m.minStack.Pop()
	}
}

func (m *MinStack) Top() int {
	n := m.dataStack.Top()

	if n == nil {
		return 0
	}

	return n.value
}

func (m *MinStack) GetMin() int {
	n := m.minStack.Top()

	if n == nil {
		return 0
	}

	return n.value
}
