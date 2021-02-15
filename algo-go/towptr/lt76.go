package towptr

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	need := make(map[rune]int)

	for _, c := range t {
		need[c]++
	}

	neeCount, left, right, i := len(t), 0, len(s), 0

	for j := 0; j < len(s); j++ {
		if _, ok := need[rune(s[j])]; ok {
			if need[rune(s[j])] > 0 {
				neeCount--
			}
			need[rune(s[j])]--
		}

		if neeCount == 0 {
			for {
				if v, ok := need[rune(s[i])]; ok && v == 0 {
					break
				} else if ok {
					need[rune(s[i])]++
				}
				i++
			}

			if j-i < right-left {
				left = i
				right = j
			}

			need[rune(s[i])]++
			neeCount++
			i++
		}

	}

	if right == len(s) {
		return ""
	}

	return s[left : right+1]
}
