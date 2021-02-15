package queue

type MyCircularQueue struct {
	size      int
	count     int
	data      []int
	headIndex int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		size:      k,
		count:     0,
		data:      make([]int, k),
		headIndex: 0,
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (cir *MyCircularQueue) EnQueue(value int) bool {
	if cir.IsFull() {
		return false
	}
	cir.data[(cir.headIndex+cir.count)%cir.size] = value
	cir.count++
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (cir *MyCircularQueue) DeQueue() bool {
	if cir.IsEmpty() {
		return false
	}
	cir.headIndex = (cir.headIndex + 1) % cir.size
	cir.count--
	return true
}

/** Get the front item from the queue. */
func (cir *MyCircularQueue) Front() int {
	if cir.IsEmpty() {
		return -1
	}

	return cir.data[cir.headIndex]
}

/** Get the last item from the queue. */
func (cir *MyCircularQueue) Rear() int {
	if cir.IsEmpty() {
		return -1
	}

	return cir.data[(cir.headIndex+cir.count-1+cir.size)%cir.size]
}

/** Checks whether the circular queue is empty or not. */
func (cir *MyCircularQueue) IsEmpty() bool {
	return cir.count == 0

}

/** Checks whether the circular queue is full or not. */
func (cir *MyCircularQueue) IsFull() bool {
	return cir.count == cir.size
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
