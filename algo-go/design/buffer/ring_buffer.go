package buffer

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
