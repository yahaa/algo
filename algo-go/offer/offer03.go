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
