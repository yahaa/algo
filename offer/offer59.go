package offer

import "container/list"

// MaxQueue 队列最大值
type MaxQueue struct {
	queue     *list.List
	sortQueue *list.List
}

// Constructor bbb
func Constructor() MaxQueue {
	return MaxQueue{
		queue:     list.New(),
		sortQueue: list.New(),
	}
}

// MaxValue v
func (q *MaxQueue) MaxValue() int {
	if q.sortQueue.Len() > 0 {
		return q.sortQueue.Front().Value.(int)
	}

	return -1
}

// PushBack p
func (q *MaxQueue) PushBack(value int) {
	q.queue.PushBack(value)

	for q.sortQueue.Len() > 0 && value > q.sortQueue.Back().Value.(int) {
		b := q.sortQueue.Back()
		q.sortQueue.Remove(b)
	}

	q.sortQueue.PushBack(value)
}

// PopFront f
func (q *MaxQueue) PopFront() int {
	if q.queue.Len() == 0 {
		return -1
	}

	f := q.queue.Front()
	q.queue.Remove(f)

	if q.sortQueue.Front().Value.(int) == f.Value.(int) {
		sf := q.sortQueue.Front()
		q.sortQueue.Remove(sf)
	}

	return f.Value.(int)
}
