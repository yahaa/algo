package offer

func minWindow(s string, t string) string {
	minLen := math.MaxInt32
	var left, right int
	var i, j int

	counts, dyCounts := make(map[rune]int), make(map[rune]int)
	for _, c := range t {
		counts[c]++
	}

	check := func() bool {
		for k, v := range counts {
			if dyCounts[k] < v {
				return false
			}
		}
		return true
	}

	for i <= j && j < len(s) {
		if check() {
			if (j - i) < minLen {
				minLen = j - i
				right = j
				left = i
			}
			if counts[rune(s[i])] > 0 {
				dyCounts[rune(s[i])]--
			}
			i++
		} else {
			if counts[rune(s[j])] > 0 {
				dyCounts[rune(s[j])]++
			}
			j++
		}
	}

	for i < len(s) {
		if check() {
			if (j - i) < minLen {
				minLen = j - i
				right = j
				left = i
			}
			if counts[rune(s[i])] > 0 {
				dyCounts[rune(s[i])]--
			}
		}
		i++
	}

	return s[left:right]
}
