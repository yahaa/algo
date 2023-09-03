package graph

// networkDelayTime leetcode 743 dijkstra 求最短路朴素实现
func networkDelayTime(times [][]int, n int, k int) int {
	inf := 99999999

	g := make([][]int, n)

	for i := 0; i < n; i++ {
		g[i] = make([]int, n)
		for j := 0; j < n; j++ {
			g[i][j] = inf
		}
	}

	for _, t := range times {
		u, v, w := t[0]-1, t[1]-1, t[2]
		g[u][v] = w
	}

	dist := make([]int, n)

	for i := 0; i < n; i++ {
		dist[i] = inf
	}

	dist[k-1] = 0

	used := make([]bool, n)

	// 执行松弛操作次数, 最多需要执行 n-1 次即可
	for i := 0; i < n-1; i++ {
		u := -1
		for j := 0; j < n; j++ {
			if !used[j] && (u == -1 || dist[j] < dist[u]) {
				u = j
			}
		}

		used[u] = true
		for v, w := range g[u] {
			dist[v] = min(dist[v], dist[u]+w)
		}
	}

	ans := dist[0]

	for i := 0; i < n; i++ {
		if dist[i] == inf {
			return -1
		}
		ans = max(ans, dist[i])
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
