package search

// search leetcode 704 二分搜索
func search(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		m := (left + right) / 2
		if nums[m] == target {
			return m
		}

		if nums[m] < target {
			left = m + 1
		} else {
			right = m
		}
	}

	return -1
}
