package offer

import (
	"container/list"
)

func validateStackSequences(pushed []int, popped []int) bool {
	stack := list.New()

	for _, p := range pushed {
		stack.PushFront(p)

		for len(popped) > 0 && stack.Len() > 0 && stack.Front().Value.(int) == popped[0] {
			f := stack.Front()
			stack.Remove(f)
			popped = popped[1:]
		}
	}

	return stack.Len() == 0
}
