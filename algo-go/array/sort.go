package array

import (
	"sort"
)

// PriArray 自定义优先级数组
type PriArray struct {
	data     []int
	priority map[int]int
}

// Len 长度
func (p PriArray) Len() int {
	return len(p.data)
}

// Less 比较大小
func (p PriArray) Less(i, j int) bool {
	const max = 9999999
	pi, pj := p.priority[p.data[i]], p.priority[p.data[j]]

	if pi == 0 {
		pi = max
	}

	if pj == 0 {
		pj = max
	}

	if pi == pj && pj == max {
		return p.data[i] < p.data[j]
	}
	return pi < pj
}

func (p PriArray) Swap(i, j int) {
	p.data[i], p.data[j] = p.data[j], p.data[i]
}

// RelativeSortArray leetcode 1122
func RelativeSortArray(arr1 []int, arr2 []int) []int {
	priMap := make(map[int]int)
	for i, item := range arr2 {
		priMap[item] = i + 1
	}

	res := PriArray{
		data:     arr1,
		priority: priMap,
	}

	sort.Sort(res)
	return res.data
}
