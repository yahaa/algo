package buffer

// MyCircularQueue leetcode 622. 设计循环队列
type MyCircularQueue struct {
	buf   []int
	size  int
	count int
	head  int
	tail  int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		buf:   make([]int, k),
		size:  k,
		count: 0,
		head:  0,
		tail:  0,
	}

}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.count == this.size {
		return false
	}

	this.buf[this.tail] = value
	this.tail = (this.tail + 1) % this.size
	this.count++

	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.count == 0 {
		return false
	}

	this.head = (this.head + 1) % this.size
	this.count--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.count == 0 {
		return -1
	}

	return this.buf[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.count == 0 {
		return -1
	}

	index := this.tail - 1
	if index == -1 {
		index = this.size - 1
	}

	return this.buf[index]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.count == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.count == this.size
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

type ringBuffer struct {
	buf   []int
	size  int
	head  int
	tail  int
	count int
}

func NewRingBuffer(size int) *ringBuffer {
	return &ringBuffer{
		buf:   make([]int, size),
		size:  size,
		count: 0,
		head:  0,
		tail:  0,
	}
}

func (b *ringBuffer) Push(v int) {
	b.buf[b.tail] = v

	b.tail = (b.tail + 1) % b.size
	b.count++

	if b.count > b.size {
		b.count--
		b.head = (b.head + 1) % b.size
	}
}

func (b *ringBuffer) Pop() int {
	if b.count == 0 {
		return -1
	}

	res := b.buf[b.head]

	b.count--

	b.head = (b.head + 1) % b.size

	return res
}

func (b *ringBuffer) PopAll() []int {
	res := make([]int, 0)
	for v := b.Pop(); v != -1; {
		res = append(res, v)
	}
	return res
}
