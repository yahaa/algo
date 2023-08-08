package offer

import "container/list"

type MinStack struct {
	s1 *list.List
	s2 *list.List
}

/** initialize your data structure here. */
func Constructor2() MinStack {
	return MinStack{
		s1: list.New(),
		s2: list.New(),
	}
}

func (t *MinStack) Push(x int) {
	t.s1.PushBack(x)

	if t.s2.Len() == 0 || t.s2.Back().Value.(int) >= x {
		t.s2.PushBack(x)
	}
}

func (t *MinStack) Pop() {
	p := t.s1.Back()

	if p == nil {
		return
	}

	t.s1.Remove(p)

	if p.Value.(int) == t.s2.Back().Value.(int) {
		t.s2.Remove(t.s2.Back())
	}

}

func (t *MinStack) Top() int {
	p := t.s1.Back()

	if p == nil {
		return 0
	}

	return p.Value.(int)
}

func (t *MinStack) Min() int {
	p := t.s2.Back()

	if p == nil {
		return 0
	}

	return p.Value.(int)

}
