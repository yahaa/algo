package offer

func findNumberIn2DArray(matrix [][]int, target int) bool {
	n := len(matrix)

	if n <= 0 {
		return false
	}

	m := len(matrix[0])
	i, j := n-1, 0

	for i >= 0 && j < m {
		if matrix[i][j] == target {
			return true
		}

		if matrix[i][j] > target {
			i--
		} else {
			j++
		}
	}
	return false
}
