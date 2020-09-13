package offer

func numWays(n int) int {
	if n < 2 {
		return 1
	}

	m := 1000000007
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i <= n; i++ {
		dp[i] = (dp[i-1]%m + dp[i-2]%m) % m
	}

	return dp[n]
}
