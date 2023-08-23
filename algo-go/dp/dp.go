package dp

// maxSubArray leetcode 53 最大子数组的和
func maxSubArray(nums []int) int {
	// 定义 dp[i] 表示 以 i 结尾的连续子数组最大和
	// dp[0]=nums[0]
	// dp[i]=max(dp[i-1]+nums[i],nums[i])

	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))

	dp[0] = nums[0]
	ans := dp[0]

	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}

		ans = max(ans, dp[i])

	}

	return ans
}

// maxSubArray leetcode 53 最大子数组的和, 优化空间
func maxSubArray2(nums []int) int {
	// 定义 dp[i] 表示 以 i 结尾的连续子数组最大和
	// dp[0]=nums[0]
	// dp[i]=max(dp[i-1]+nums[i],nums[i])

	if len(nums) == 0 {
		return 0
	}

	dp := nums[0]
	ans := dp

	for i := 1; i < len(nums); i++ {
		if dp > 0 {
			dp = dp + nums[i]
		} else {
			dp = nums[i]
		}

		ans = max(ans, dp)
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
