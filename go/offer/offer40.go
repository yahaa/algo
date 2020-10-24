package offer

import "container/heap"

type descArr []int

func (d descArr) Len() int {
	return len(d)
}

func (d descArr) Less(i, j int) bool {
	return d[i] > d[j]
}

func (d descArr) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d *descArr) Push(v interface{}) {
	*d = append(*d, v.(int))
}

func (d *descArr) Pop() interface{} {
	old := *d
	n := len(old)

	v := old[n-1]
	*d = old[0 : n-1]

	return v
}

func getLeastNumbers(arr []int, k int) []int {
	if k <= 0 {
		return nil
	}

	h := make(descArr, 0)
	heap.Init(&h)

	for _, n := range arr {
		if h.Len() < k || n < h[0] {
			heap.Push(&h, n)
		}

		if h.Len() > k {
			heap.Pop(&h)
		}
	}

	res := make([]int, 0)
	for h.Len() > 0 {
		res = append(res, heap.Pop(&h).(int))
	}
	return res
}
