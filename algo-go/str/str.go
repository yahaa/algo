package str

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// SubdomainVisits leetcode 811
func SubdomainVisits(cpdomains []string) []string {
	visitMap := make(map[string]int)

	for _, item := range cpdomains {
		ts := strings.Split(item, " ")
		if len(ts) != 2 {
			continue
		}

		visit, err := strconv.Atoi(ts[0])
		if err != nil {
			continue
		}

		d := ts[1]
		for len(d) > 0 {
			visitMap[d] += visit
			idx := strings.Index(d, ".")
			if idx >= 0 {
				d = d[idx+1:]
			} else {
				break
			}
		}
	}

	res := make([]string, 0)
	for k, v := range visitMap {
		res = append(res, fmt.Sprintf("%d %s", v, k))
	}
	return res
}

// DayOfTheWeek leetcode 1185
func DayOfTheWeek(day int, month int, year int) string {
	w := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	dstr := fmt.Sprintf("%d-%02d-%02d", year, month, day)
	d, err := time.Parse("2006-01-02", dstr)
	if err != nil {
		return ""
	}

	return w[d.Weekday()]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// CommonChars leetcode 1002
func CommonChars(arr []string) []string {
	countMap := make(map[rune]int)
	res := make([]string, 0)

	for i, item := range arr {
		if i == 0 {
			for _, c := range item {
				countMap[c]++
			}
		} else {
			itemMap := make(map[rune]int)
			for _, c := range item {
				itemMap[c]++
			}

			for k, v1 := range countMap {
				if v2, ok := itemMap[k]; ok {
					countMap[k] = min(v1, v2)
				} else {
					delete(countMap, k)
				}
			}
		}

	}

	for k, v := range countMap {
		for v > 0 {
			res = append(res, string(k))
			v--
		}
	}

	return res
}

// AddBinary 67
func AddBinary(a string, b string) string {
	an, bn, c, res := len(a), len(b), '0', make([]rune, 0)
	maxLen := Max(an, bn)
	for i := 1; i <= maxLen; i++ {
		ai := an - i
		bi := bn - i
		if ai >= 0 && bi >= 0 {
			ac := a[ai]
			bc := b[bi]
			if c == '1' {
				if ac == bc {
					res = append(res, '1')
					if ac == '1' {
						c = '1'
					} else {
						c = '0'
					}
				} else {
					res = append(res, '0')
					c = '1'
				}
			} else {
				if ac == bc {
					res = append(res, '0')
					if ac == '1' {
						c = '1'
					}
				} else {
					res = append(res, '1')
				}
			}

		} else if ai >= 0 {
			ac := a[ai]
			if ac == '0' {
				if c == '0' {
					res = append(res, '0')
				} else {
					res = append(res, '1')
					c = '0'
				}
			} else {
				if c == '0' {
					res = append(res, '1')
					c = '0'
				} else {
					res = append(res, '0')
					c = '1'
				}
			}
		} else if bi >= 0 {
			bc := b[bi]
			if bc == '0' {
				if c == '0' {
					res = append(res, '0')
				} else {
					res = append(res, '1')
					c = '0'
				}
			} else {
				if c == '0' {
					res = append(res, '1')
					c = '0'
				} else {
					res = append(res, '0')
					c = '1'
				}
			}
		}
	}

	if c == '1' {
		res = append(res, '1')
	}
	return reverseString(res)
}

func reverseString(s []rune) string {
	for from, to := 0, len(s)-1; from < to; from, to = from+1, to-1 {
		s[from], s[to] = s[to], s[from]
	}
	return string(s)
}

// CountCharacters 1160
func CountCharacters(words []string, chars string) int {
	charsMap := make(map[rune]int)
	res := 0
	for _, c := range chars {
		charsMap[c]++
	}

	for _, w := range words {
		wMap := make(map[rune]int)
		isOk := true
		for _, c := range w {
			if v, ok := wMap[c]; ok {
				if v > 0 {
					wMap[c]--
				} else {
					isOk = false
					break
				}
			} else {
				if v, ok := charsMap[c]; ok {
					wMap[c] = v - 1
				} else {
					isOk = false
					break
				}
			}
		}
		if isOk {
			res += len(w)
		}
	}
	return res
}

// Max 最大值
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// LongestPalindrome 409
func LongestPalindrome(s string) int {
	chartMap, res := make(map[rune]int), 0
	for _, c := range s {
		chartMap[c]++
	}
	maxOld := 0
	for _, v := range chartMap {
		if v%2 == 1 && v > maxOld {
			maxOld = v
		}

		if v%2 == 0 {
			res += v
		} else {
			res += v - 1
		}
	}

	if maxOld > 0 {
		res++
	}

	return res
}
