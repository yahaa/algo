package offer

import (
	"fmt"
	"sort"
	"strings"
)

// StringInt ccc
type StringInt []int

func (m StringInt) Len() int {
	return len(m)
}

func (m StringInt) Less(i, j int) bool {
	s1 := fmt.Sprintf("%d%d", m[i], m[j])
	s2 := fmt.Sprintf("%d%d", m[j], m[i])

	return strings.Compare(s1, s2) <= 0
}

func (m StringInt) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func minNumber(nums []int) string {
	sort.Sort(StringInt(nums))

	var res string
	for _, item := range nums {
		res += fmt.Sprintf("%d", item)
	}

	return res
}
