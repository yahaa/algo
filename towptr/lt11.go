package towptr

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	if i >= j {
		return 0
	}
	var ans int
	for i < j {
		ans = max(ans, (j-i)*min(height[i], height[j]))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
