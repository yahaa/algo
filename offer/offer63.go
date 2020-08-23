package offer

import "math"

func maxProfit(prices []int) int {
	maxValue, minPrice := 0, math.MaxInt32

	for i := 0; i < len(prices); i++ {
		minPrice = min(prices[i], minPrice)
		maxValue = max(maxValue, prices[i]-minPrice)
	}
	return maxValue
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
