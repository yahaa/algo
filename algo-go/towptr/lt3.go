package towptr

func lengthOfLongestSubstring(s string) int {
	res, mp, i, j := 0, make(map[rune]int), 0, 0

	for j = 0; j < len(s); j++ {
		if mp[rune(s[j])] > 0 {
			mp[rune(s[j])]++
			for mp[rune(s[j])] > 1 {
				mp[rune(s[i])]--
				i++
			}
		} else {
			mp[rune(s[j])] = 1
		}
		res = max(j-i+1, res)
	}

	//res = max(j-i, res)
	return res
}
