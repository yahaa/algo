package offer

import (
	"container/list"
)

func posSum(n int) int {
	if n <= 0 {
		return 0
	}

	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}

	return sum
}

func movingCount(m int, n int, k int) int {
	mp := make([][]int, 0)
	for i := 0; i < m; i++ {
		mp = append(mp, make([]int, n))
		for j := 0; j < n; j++ {
			if posSum(i)+posSum(j) > k {
				mp[i][j] = 0
			} else {
				mp[i][j] = 1
			}
		}
	}

	count := 1
	q := list.New()
	q.PushBack([2]int{0, 0})
	dir := [2][4]int{{-1, 1, 0, 0}, {0, 0, -1, 1}}
	mp[0][0] = 2

	for q.Len() > 0 {
		f := q.Front()
		i, j := f.Value.([2]int)[0], f.Value.([2]int)[1]

		q.Remove(f)

		for d := 0; d < 4; d++ {
			ti := dir[0][d] + i
			tj := dir[1][d] + j

			if ti < 0 || ti >= m || tj >= n || tj < 0 || mp[ti][tj] == 2 || mp[ti][tj] == 0 {
				continue
			}

			mp[ti][tj] = 2

			q.PushBack([2]int{ti, tj})
			count++
		}
	}

	return count
}
