package offer

func exist(board [][]byte, word string) bool {
	n, m := len(board), len(board[0])
	dir := [][2]int{
		{0, -1},
		{-1, 0},
		{0, 1},
		{1, 0},
	}

	var dfs func(i, j, k int) bool

	dfs = func(i, j, k int) bool {
		if k == len(word)-1 {
			return true
		}

		if i < 0 || i >= n || j < 0 || j >= m || board[i][j] != word[k] {
			return false
		}

		board[i][j] = ' '

		res := false
		for _, d := range dir {
			res = dfs(i+d[0], j+d[1], k+1) || res
		}

		board[i][j] = word[k]

		return res
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}
