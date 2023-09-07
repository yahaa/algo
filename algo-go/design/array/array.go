package array

// firstMissingPositive leetcode 41 缺失的第一个正数
func firstMissingPositive(nums []int) int {
	// nums = [1, 2, 0]
	// nums = [1, -1 ,3, 4]
	// nums = [7, 8, 9, 11,12]
	// nums = [1]
	// nums = [1, 1]

	m := len(nums)
	for i := 0; i < m; i++ {
		n := nums[i]
		for n != i+1 && n-1 >= 0 && n <= m && nums[n-1] != n {
			nums[i], nums[n-1] = nums[n-1], nums[i]
			n = nums[i]
		}
	}

	for i := 0; i < m; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return m + 1
}
