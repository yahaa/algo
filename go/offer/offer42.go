package offer

func maxSubArray(nums []int) int {
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] = max(nums[i], nums[i-1]+nums[i])
		ans = max(ans, nums[i])
	}

	return ans
}
