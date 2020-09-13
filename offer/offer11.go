package offer

// todo
func minArray(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	res := numbers[0]

	for i := 1; i < len(numbers); i++ {
		if numbers[i] < numbers[i-1] {
			return numbers[i]
		}
	}

	return res
}
