package offer

func findRepeatNumber(nums []int) int {
	v := make(map[int]int)
	res := 0
	for _, num := range nums {
		if _, ok := v[num]; ok {
			res = num
			break
		}
	}
	return res
}

// findRepeatNumber2 剑指 Offer 03. 数组中重复的数字
func findRepeatNumber2(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}

			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}

	return 0
}
