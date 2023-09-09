package dfs

type Node struct {
	Val       int
	Neighbors []*Node
}

// letterCombinations 17. 电话号码的字母组合
func letterCombinations(digits string) []string {
	mp := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	var ans []string

	if len(digits) == 0 {
		return ans
	}

	var dfs func(cur string, n, i int)

	dfs = func(cur string, n, i int) {
		if i >= n {
			ans = append(ans, cur)
			return
		}

		ss := mp[digits[i]]

		for k := 0; k < len(ss); k++ {
			dfs(cur+ss[k], n, i+1)
		}
	}

	dfs("", len(digits), 0)

	return ans
}

// leetcode 133
func cloneGraph(node *Node) *Node {
	var visit func(node *Node, mp map[int]*Node) *Node

	visit = func(node *Node, mp map[int]*Node) *Node {
		if node == nil {
			return nil
		}

		if _, ok := mp[node.Val]; !ok {
			n := &Node{Val: node.Val}
			mp[node.Val] = n
			for _, ne := range node.Neighbors {
				n.Neighbors = append(n.Neighbors, visit(ne, mp))
			}
		}

		return mp[node.Val]
	}

	return visit(node, make(map[int]*Node))
}
