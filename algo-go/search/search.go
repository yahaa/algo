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

func search2(nums []int, target int) int {
	var dfs func(l, r int)
	ans := -1
	dfs = func(l, r int) {
		if l >= r {
			return
		}

		m := l + (r-l)/2
		if nums[m] == target {
			ans = m
			return
		}

		if nums[m] < target {
			l = m + 1
		} else {
			r = m
		}

		dfs(l, r)
	}

	dfs(0, len(nums))

	return ans
}
