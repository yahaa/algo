package offer

func spiralOrder(matrix [][]int) []int {
	res := make([]int, 0)
	if len(matrix) == 0 {
		return res
	}

	left, up, right, down := 0, 0, len(matrix[0])-1, len(matrix)-1

	i, j := 0, 0

	for ok(left, right, up, down) {
		for j = left; j <= right && ok(left, right, up, down); j++ {
			res = append(res, matrix[i][j])
		}

		j--
		up++

		for i = up; i <= down && ok(left, right, up, down); i++ {
			res = append(res, matrix[i][j])
		}

		i--
		right--

		for j = right; j >= left && ok(left, right, up, down); j-- {
			res = append(res, matrix[i][j])
		}

		j++
		down--

		for i = down; i >= up && ok(left, right, up, down); i-- {
			res = append(res, matrix[i][j])
		}

		i++
		left++
	}

	return res
}

func ok(left, right, up, down int) bool {
	return left <= right && up <= down
}
