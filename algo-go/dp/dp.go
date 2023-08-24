package dp

import (
	"fmt"
	"sort"
	"strconv"
)

func findMaxValue(nums []int, n int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	fmt.Printf("%v\n", nums)

	s := fmt.Sprintf("%v", n)
	tmp := make([]int, 0)

	for i := 0; i < len(s); i++ {
		t, _ := strconv.Atoi(s[i : i+1])
		for j := 0; j < len(nums); j++ {
			if nums[j] == t {
				tmp = append(tmp, nums[j])
				break
			} else if nums[j] < t {
				tmp = append(tmp, nums[j])
				break
			}
		}
	}

	fmt.Printf("%v\n", tmp)

	if len(tmp) != len(s) {
		tmp = make([]int, 0)

		for i := 0; i < len(s)-1; i++ {
			tmp = append(tmp, nums[0])
		}
	}

	var ans string

	for i := 0; i < len(tmp); i++ {
		ans += fmt.Sprintf("%v", tmp[i])
	}

	v, _ := strconv.Atoi(ans)
	return v
}

// maxSubArray leetcode 53 最大子数组的和
func maxSubArray(nums []int) int {
	// 定义 dp[i] 表示 以 nums[i] 结尾的连续子数组最大和
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

// lengthOfLIS leetcode 300 最长上升子序列
func lengthOfLIS(nums []int) int {
	// 定义 dp[i] 为 以 nums[i] 结尾的最长上升子序列的长度
	// dp[i]=max(dp[0].j.dp[i-1]) + 1 (当 nums[j] < nums[i])
	// dp[0]=1

	dp := make([]int, len(nums))

	dp[0] = 1

	ans := dp[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}

		ans = max(ans, dp[i])
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
