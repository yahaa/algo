package offer

import (
	"fmt"
	"strconv"
)

func translateNum(num int) int {
	s := fmt.Sprintf("%d", num)

	if len(s) < 2 {
		return 1
	}

	dp := make([]int, len(s)+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i <= len(s); i++ {
		t := s[i-2 : i]

		n, _ := strconv.Atoi(t)
		if n >= 10 && n <= 25 {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}

	return dp[len(s)]
}
