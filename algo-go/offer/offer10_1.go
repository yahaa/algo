package offer

func fib(n int) int {
	if n < 2 {
		return n
	}

	m := 1000000007
	f := make([]int, n+1)
	f[0], f[1] = 0, 1
	for i := 2; i <= n; i++ {
		f[i] = (f[i-1]%m + f[i-2]%m) % m
	}

	return f[n]
}
