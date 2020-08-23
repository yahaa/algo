package offer

func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	cur, count := nums[0], 0
	for _, n := range nums {
		if count == 0 {
			cur = n
			count++
		} else {
			if n == cur {
				count++
			} else {
				count--
			}
		}
	}

	return cur
}
