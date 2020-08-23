package offer

// singleNumber1 leetcode 136
func singleNumber1(nums []int) int {
	var res = 0
	for _, num := range nums {
		res = res ^ num
	}

	return res
}

// singleNumber leetcode 137
func singleNumber2(nums []int) int {
	var res int
	for i := 0; i <= 64; i++ {
		count := 0
		for _, num := range nums {
			count += num >> i & 1
		}

		// count % 3 在这题里面永远为 0 或 1
		res |= (count % 3) << i
	}
	return res
}

func singleNumber3(nums []int) []int {
	var t int
	for _, num := range nums {
		t ^= num
	}

	f := 1
	for t&f == 0 {
		f <<= 1
	}

	var a, b int
	for _, num := range nums {
		if num&f == 0 {
			a ^= num
		} else {
			b ^= num
		}
	}

	return []int{a, b}
}


