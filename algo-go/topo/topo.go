package topo

import (
	"container/list"
)

// canFinish leetcode 207 拓扑排序
func canFinish(numCourses int, prerequisites [][]int) bool {
	g := make(map[int][]int)
	in := make(map[int]int)
	queue := list.New()
	result := make([]int, 0)

	for _, item := range prerequisites {
		g[item[1]] = append(g[item[1]], item[0])
		in[item[0]]++
	}

	for i := 0; i < numCourses; i++ {
		if in[i] == 0 {
			queue.PushFront(i)
		}
	}

	for queue.Len() > 0 {
		u := queue.Front().Value.(int)
		queue.Remove(queue.Front())
		result = append(result, u)

		for _, v := range g[u] {
			in[v]--

			if in[v] == 0 {
				queue.PushFront(v)
			}
		}
	}

	return len(result) == numCourses
}
