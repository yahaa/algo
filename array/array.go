package array

import (
	"sort"
)

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
