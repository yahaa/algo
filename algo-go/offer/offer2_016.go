package offer

import "math"

func lengthOfLongestSubstring_2(s string) int {
	i, j, ans, count := 0, 0, 0, make(map[byte]int)

	for i <= j && j < len(s) {
		if _, ok := count[s[j]]; !ok {
			count[s[j]]++
			j++
		} else {
			delete(count, s[i])
			i++
		}

		if j-i > ans {
			ans = j - i
		}
	}

	return ans
}
