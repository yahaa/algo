package towptr

import (
	"math"
	"sort"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	if n <= 3 {
		t := 0
		for _, v := range nums {
			t += v
		}
		return t
	}
	sort.Ints(nums)

	left, right, sum, diff, res := 0, 0, 0, math.MaxInt32, 0
	for i := 0; i < n-2; i++ {
		left = i + 1
		right = n - 1
		for left < right {
			sum = nums[i] + nums[left] + nums[right]

			switch {
			case sum == target:
				return sum
			case sum < target:
				left++
			case sum > target:
				right--
			}

			if abs(target-sum) < diff {
				diff = abs(target - sum)
				res = sum
			}
		}
	}
	return res

}
