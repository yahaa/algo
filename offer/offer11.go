package offer

// todo
func minArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			return nums[i]
		}
	}

	return res
}

// offer 11
func minArray2(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] < nums[right] {
			right = mid
		} else if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right--
		}
	}

	return nums[left]
}
