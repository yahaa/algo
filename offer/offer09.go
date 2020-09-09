package offer

import "container/list"

type CQueue struct {
	stack1 list.List
	stack2 list.List
}

func (this *CQueue) AppendTail(value int) {
	this.stack1.PushFront(value)
}

func (this *CQueue) DeleteHead() int {
	if this.stack2.Len() == 0 {
		for this.stack1.Len() > 0 {
			f := this.stack1.Front()
			this.stack1.Remove(f)
			this.stack2.PushFront(f.Value.(int))
		}
	}

	if this.stack2.Len() > 0 {
		f := this.stack2.Front()
		this.stack2.Remove(f)
		return f.Value.(int)
	}
	return -1
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
