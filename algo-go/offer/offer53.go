package offer

// missingNumber offer 53 0～n-1 中缺失的数字
func missingNumber(nums []int) int {
	i := 0
	for i < len(nums) {
		if nums[i] >= len(nums) {
			return i
		}

		if nums[i] != i {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		} else {
			i++
		}
	}

	return i
}
