package offer

// offer 64 利用 && 短路特点求解
func sumNums(n int) int {
	sum := n
	_ = (n > 0) && func() bool {
		sum += sumNums(n - 1)
		return sum > 0
	}()
	return sum
}
