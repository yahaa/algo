package main

import (
	"fmt"
	"sort"
)

func solve(nums []int, n int) {
	// 2 4 9
	// 23121
	nn := make([]int, 0)

	tt := n

	for n > 0 {
		t := n % 10
		nn = append(nn, t)
		n = n / 10
	}

	sort.Ints(nums)
	fmt.Printf("nums: %v n: %v\n", nums, tt)

	res := make([]int, 0)
	existLess := false
	skip := false

	for i := len(nn) - 1; i >= 0; i-- {
		tmp := nn[i]
		if existLess || skip {
			res = append(res, nums[len(nums)-1])
		} else {
			skipCount := 0
			for j := len(nums) - 1; j >= 0; j-- {
				if nums[j] > tmp {
					skipCount++
					continue
				} else if nums[j] == tmp {
					res = append(res, nums[j])
				} else {
					res = append(res, nums[j])
					existLess = true
				}
			}

			if skipCount == len(nums) {
				skip = true
			}
		}
	}

	fmt.Printf("%v", res)

}

// 给定一个数 n,如 23121;
// 给定一组数字A如{2,4,9},求由 A 中元素组成的、小于 n 的最大数,如小于 23121 的最大数为 22999。
// 请实现这个函数 maxLessEqual
// 例如:
// f(23121, {2,4,9}) = 22999
// f(23121, {5,4,9}) = 9999
// f(24121, {2,4,9}) = 22999
// f(24244, {2,4,9}) = 24242
// f(22222, {2,4,9}) = 9999
func maxLessEqual(nums []int, target int) int {
}

func repeatAppend(s string, c byte, count int) string {
	for i := 0; i < count; i++ {
		s += string(c)
	}
	return s
}

func main() {
	//solve([]int{2, 4, 9}, 23121)
	//solve([]int{5, 4, 9}, 23121)
	solve([]int{2, 4, 9}, 24121)
}
