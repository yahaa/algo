package offer

func findAnagrams(s string, p string) []int {
	sl, pl, ans := len(s), len(p), make([]int, 0)
	if sl < pl {
		return ans
	}

	var sc, pc [26]int
	for i := 0; i < pl; i++ {
		sc[s[i]-'a']++
		pc[p[i]-'a']++
	}

	if sc == pc {
		ans = append(ans, 0)
	}

	for i, c := range s[:sl-pl] {
		sc[c-'a']--
		sc[s[i+pl]-'a']++

		if sc == pc {
			ans = append(ans, i+1)
		}
	}

	return ans
}
