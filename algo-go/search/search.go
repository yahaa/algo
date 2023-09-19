package search

import (
	"math"
)

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

// leetcode 704 二分查找递归实现
func search2(nums []int, target int) int {
	var dfs func(l, r int) int
	dfs = func(l, r int) int {
		if l >= r {
			return -1
		}

		m := l + (r-l)/2
		if nums[m] == target {
			return m
		}

		if nums[m] < target {
			return dfs(m+1, r)
		}

		return dfs(l, m)
	}

	return dfs(0, len(nums))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// minSubArrayLen leetcode 209 最长子数组滑动窗口实现
func minSubArrayLen(target int, nums []int) int {
	l, r, tmp, ans := 0, 0, 0, math.MaxInt64
	for r < len(nums) {
		tmp += nums[r]
		r++
		for tmp >= target {
			ans = min(ans, r-l)
			tmp -= nums[l]
			l++
		}
	}

	if ans == math.MaxInt64 {
		return 0
	}

	return ans
}

// searchRange leetcode 34 在有序数组中查找 target 出现的第一次和最后一次
func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)

	ans := -1
	for l < r {
		m := l + (r-l)/2
		if nums[m] == target {
			ans = m
			break
		}

		if nums[m] > target {
			r = m
		} else {
			l = m + 1
		}
	}

	if ans == -1 {
		return []int{-1, -1}
	}

	l = ans
	for l-1 >= 0 && nums[l-1] == target {
		l--
	}

	r = ans
	for r < len(nums)-1 && nums[r+1] == target {
		r++
	}

	return []int{l, r}
}

// bSearch 自定义二分查找
// 找到返回下标，找不到返回 -1
func bSearch(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := l + (r-l)/2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			l = mid + 1
		}

		if nums[mid] > target {
			r = mid
		}
	}

	return -1
}

// search3 33. 搜索旋转排序数组
// 通过 findMin 找到旋转数组中最小值
// 把原始数组分为两个有序数组
// 用二分在两个有序数组中二分查找
func search3(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	b := findMin(nums)

	f := nums[0:b]
	e := nums[b:]

	i := bSearch(f, target)

	if i != -1 {
		return i
	}

	i = bSearch(e, target)

	if i != -1 {
		return len(f) + i
	}

	return -1
}

// findMin 剑指 Offer 11. 旋转数组的最小数字思想
// 返回最小数字的下标
func findMin(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		mid := l + (r-l)/2

		if nums[mid] < nums[r] {
			r = mid
		} else if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r--
		}
	}

	return l
}
