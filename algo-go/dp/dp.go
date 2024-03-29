package dp

import (
	"container/list"
	"fmt"
	"sort"
	"strconv"
)

// maxSubarraySumCircular leetcode 918 环形数组最大子数组和
func maxSubarraySumCircular(nums []int) int {
	// 设 dpmax[i] 表示以 nums[i] 结尾的子数组的最大和
	// 则 dpmax[i] = max(dpmax[i-1]+nusm[i],nums[i])
	// maxAns=max(maxAns,dpmax[i])

	// 设 dpmin[i] 表示以 nums[i] 结尾的子数组的最小和
	// 则 dpmin[i] = min(dpmin[i-1]+nums[i],nums[i])
	// minAns=min(minAns,dpmin[i])

	// res=max(maxAns,sum-minAns)

	dpMax, dpMin, sum := nums[0], nums[0], nums[0]
	maxAns, minAns := dpMax, dpMin

	for i := 1; i < len(nums); i++ {
		dpMax = max(dpMax+nums[i], nums[i])
		dpMin = min(dpMin+nums[i], nums[i])

		maxAns = max(maxAns, dpMax)
		minAns = min(minAns, dpMin)

		sum += nums[i]
	}

	if maxAns < 0 {
		return maxAns
	}

	return max(maxAns, sum-minAns)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

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

// longestCommonSubsequence leetcode 1143 最长公共子序列
func longestCommonSubsequence(text1 string, text2 string) int {
	// 定义 dp[i][j] 为一 i,j 结尾的字符串最长的公共子序列
	// 当 text1[i] == text2[j] dp[i][j]=dp[i-1][j-1]+1
	// 当 text1[i] != text2[j] dp[i][j]=max(dp[i-1][j],dp[i][j-1])
	n, m := len(text1), len(text2)
	dp := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[n][m]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

type index struct {
	i, j int
}

func findLongestFromIndex(matrix [][]int, i, j int) int {
	n, m := len(matrix), len(matrix[0])
	visit, res := make([][]int, n), 0
	for i := 0; i < n; i++ {
		visit[i] = make([]int, m)
	}

	dir := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	_ = dir

	queue := list.New()

	queue.PushBack(index{i, j})

	return res

}

func longestIncreasingPath(matrix [][]int) int {
	return 0
}

// rob2 213. 打家劫舍 II
// 思路: 假设数组 nums 的长度为 n
// 1. 如果不偷窃最后一间房屋，则偷窃房屋的下标范围是 [0,n−2]
// 2. 如果不偷窃第一间房屋，则偷窃房屋的下标范围是 [1,n−1]
func rob2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	if n == 1 {
		return nums[0]
	}

	if n == 2 {
		return max(nums[0], nums[1])
	}

	return max(rob(nums[:n-1]), rob(nums[1:]))
}

// rob 198. 打家劫舍
func rob(nums []int) int {
	// dp[i] = max(dp[i-1],dp[i-2]+nums[i])

	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))

	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[len(nums)-1]
}

// minCostClimbingStairs leetcode 746. 使用最小花费爬楼梯
func minCostClimbingStairs(cost []int) int {
	// 1. 设置 dp[i] 为爬到第 i 个台阶需要花费的最小体力
	// 2. dp[i]=min(dp[i-1]+cost[i-1],dp[i-2]+cost[i-2])
	// 3. dp[0]=0 dp[1]=0
	// 4. i>=2 开始便利

	dp := make([]int, len(cost)+1)

	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[len(cost)]
}

// longestPalindrome leetcode 5 最长回文子串 dp 实现
func longestPalindrome(s string) string {
	// dp[i][j] 表示 s[i...j] 是否为回文串
	// dp[i][i] 表示所有长度为 1 的子串
	if len(s) < 2 {
		return s
	}

	dp := make([][]bool, len(s))

	maxLen, start := 1, 0

	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}

	// 遍历所有长度
	for l := 2; l <= len(s); l++ {
		for i := 0; i < len(s); i++ {
			j := l + i - 1

			if j >= len(s) {
				break
			}

			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				start = i
			}
		}
	}
	return s[start : start+maxLen]
}
