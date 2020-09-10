package offer

import "strings"

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	i, j, res := 0, 0, 1

	for i <= j && j < len(s) {
		if i == j {
			j++
			continue
		}

		sub := s[i:j]
		if strings.Contains(sub, string(s[j])) {
			i++
		} else {
			j++
		}

		res = max(res, j-i)
	}

	return res
}

func lengthOfLongestSubstring2(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	i, res := 0, 0

	mp := make(map[byte]int)

	for j := 0; j < len(s); j++ {
		n, ok := mp[s[j]]
		if ok {
			i = max(n, i)
		}
		mp[s[j]] = j + 1

		res = max(res, j-i+1)
	}

	return res
}
