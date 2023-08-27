package sort

import (
	"fmt"
	"sort"
	"strings"
)

// MergeSort 归并排序算法
func MergeSort(a []int) {
	tmp := make([]int, len(a))
	mergeSort(a, tmp, 0, len(a)-1)
}

func mergeSort(a, tmp []int, left, right int) int {
	if left >= right {
		return 0
	}

	mid := (left + right) / 2

	count := mergeSort(a, tmp, left, mid) + mergeSort(a, tmp, mid+1, right)

	// do merge

	i, j, t := left, mid+1, 0

	for i <= mid && j <= right {
		if a[i] <= a[j] {
			tmp[t] = a[i]
			i++
			count += j - (mid + 1)
		} else {
			tmp[t] = a[j]
			j++
		}
		t++
	}

	for i <= mid {
		tmp[t] = a[i]
		i++
		t++

		count += j - (mid + 1)
	}

	for j <= right {
		tmp[t] = a[j]
		j++
		t++
	}

	t = 0
	for left <= right {
		a[left] = tmp[t]
		left++
		t++
	}

	return count
}

// largestNumber leetcode 179 最大的数，排序基本应用
func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		a := fmt.Sprintf("%d%d", nums[i], nums[j])
		b := fmt.Sprintf("%d%d", nums[j], nums[i])

		return a > b
	})

	var ans string
	for i := 0; i < len(nums); i++ {
		ans += fmt.Sprintf("%d", nums[i])
	}

	ans = strings.TrimLeft(ans, "0")

	if len(ans) == 0 {
		return "0"
	}

	return ans
}
