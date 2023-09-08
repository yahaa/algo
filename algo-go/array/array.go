package array

import (
	"sort"
)

type Pairs struct {
	index int
	value int
}

// singleNumber260 leetcode 260. 只出现一次的数字 III(有两个数组出现一次，其余都出现两次)
func singleNumber260(nums []int) []int {
	// 1 0000 0001
	// 2 0000 0010
	//  3 0000 0011
	//  5 0000 0101

	// 3^5 0000 0110

	// 3和5 第一个不相同的位置 0000 0010

	// 根据第一个不相同的位置分组

	s := 0

	for _, n := range nums {
		s ^= n
	}

	d := 1
	for i := 0; i < 64; i++ {
		if s&d > 0 {
			break
		}
		d <<= 1
	}

	a, b := 0, 0
	for _, n := range nums {
		if n&d > 0 {
			a ^= n
		} else {
			b ^= n
		}
	}

	return []int{a, b}
}

// singleNumber 137. 只出现一次的数字 II(一个数字出现一次，其他数字出现 3 次)
func singleNumber(nums []int) int {
	// 2 0000 0010
	// 2 0000 0010
	// 2 0000 0010
	// 3 0000 0011
	// c 0000 0011

	// 1 0000 0001
	// 2 0000 0010
	// 4 0000 0100
	counts := make([]int, 64)

	for i := 0; i <= 64; i++ {
		for _, n := range nums {
			if (n >> i & 1) > 0 {
				counts[i]++
			}
		}

		counts[i] = counts[i] % 3
	}

	ans := 0
	base := 1
	for i := 0; i <= 64; i++ {
		if counts[i] > 0 {
			ans |= base
		}
		base <<= 1
	}

	return ans
}

func twoSum(nums []int, target int) []int {
	pairs := make([]Pairs, len(nums))
	for i := 0; i < len(nums); i++ {
		pairs[i] = Pairs{i, nums[i]}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].value < pairs[j].value
	})

	for i := 0; i < len(pairs); i++ {
		tmp := target - pairs[i].value

		k := sort.Search(len(pairs), func(j int) bool {
			return pairs[j].value >= tmp
		})

		if k < len(pairs) && pairs[k].value == tmp && pairs[k].index != pairs[i].index {
			return []int{pairs[i].index, pairs[k].index}
		}
	}

	return []int{}
}

// SortArrayByParityII leetcode 922
func SortArrayByParityII(A []int) []int {
	n := len(A)
	for i := 0; i < n; i++ {
		k, j := i%2, i+1
		for j < n && A[i]%2 != k {
			A[i], A[j] = A[j], A[i]
			j++
		}
	}
	return A
}

// Min 求最小值
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Abs 求绝对值
func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// MinimumAbsDifference leetcode 1200
func MinimumAbsDifference(arr []int) [][]int {
	min, res, n := 99999999, make([][]int, 0), len(arr)
	sort.Ints(arr)
	for i := 0; i < n-1; i++ {
		if arr[i+1]-arr[i] < min {
			min = arr[i+1] - arr[i]
			res = res[0:0]
			res = append(res, []int{arr[i], arr[i+1]})
		} else if arr[i+1]-arr[i] == min {
			res = append(res, []int{arr[i], arr[i+1]})
		}
	}
	return res
}

// Fib leetcode 509
func Fib(n int) int {
	f := [31]int{0, 1}
	for i := 2; i < 31; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

// MinCostToMoveChips leetcode 1217
func MinCostToMoveChips(chips []int) int {
	even, old := 0, 0
	for _, item := range chips {
		if item%2 == 0 {
			even++
		} else {
			old++
		}
	}
	if even < old {
		return even
	}
	return old
}

// UniqueOccurrences leetcode 1207
func UniqueOccurrences(arr []int) bool {
	countMap := make(map[int]int)
	for _, item := range arr {
		countMap[item]++
	}

	valueMap := make(map[int]bool)

	for _, v := range countMap {
		if _, ok := valueMap[v]; ok {
			return false
		}
		valueMap[v] = true
	}
	return true
}

// SmallestRangeI leetcode 908
func SmallestRangeI(arr []int, k int) int {
	if len(arr) <= 0 {
		return 0
	}
	maxVal, minVal := arr[0], arr[0]

	for _, item := range arr {
		maxVal = Max(maxVal, item)
		minVal = Min(minVal, item)
	}
	return maxVal - minVal - 2*k
}

// Max 求最大值
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ProjectionArea 883
func ProjectionArea(grid [][]int) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		front, left := 0, 0
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] > 0 {
				res++
			}
			front = Max(front, grid[i][j])
			left = Max(left, grid[j][i])
		}
		res += front + left
	}
	return res
}

// merge leetcode 88
func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for k >= 0 {
		if i >= 0 && j >= 0 {
			if nums1[i] < nums2[j] {
				nums1[k] = nums2[j]
				j--
			} else {
				nums1[k] = nums1[i]
				i--
			}
		} else if i >= 0 {
			nums1[k] = nums1[i]
			i--
		} else if j >= 0 {
			nums1[k] = nums2[j]
			j--
		}

		k--
	}
}
