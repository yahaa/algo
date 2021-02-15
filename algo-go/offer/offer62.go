package offer

func lastRemaining(n int, m int) int {
	if n == 1 {
		return 0
	}

	x := lastRemaining(n-1, m)

	return (x + m) % n
}
