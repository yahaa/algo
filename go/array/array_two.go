package array

// NumRookCaptures leetcode 999
func NumRookCaptures(board [][]byte) int {
	var (
		dir = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
		res int
	)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'R' {
				for _, item := range dir {
					res += find(item[0], item[1], i, j, board)
				}
			}
		}
	}
	return res
}

func find(di, dj, i, j int, board [][]byte) int {
	if i >= len(board) || i < 0 || j >= len(board) || j < 0 {
		return 0
	}

	if board[i][j] == 'B' {
		return 0
	}

	if board[i][j] == 'p' {
		return 1
	}

	return find(di, dj, i+di, j+dj, board)
}
