package offer

func permutation(s string) []string {
	var mp = make(map[string]bool)
	var dfs func(s []byte, n int)

	dfs = func(s []byte, n int) {
		if n == len(s) {
			mp[string(s)] = true
			return
		}

		for i := n; i < len(s); i++ {
			s[n], s[i] = s[i], s[n]
			dfs(s, n+1)
			s[n], s[i] = s[i], s[n]
		}
	}

	dfs([]byte(s), 0)
	var res = make([]string, 0)

	for k := range mp {
		res = append(res, k)
	}
	return res
}
