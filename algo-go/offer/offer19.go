package offer

import "regexp"

func isMatch(s string, p string) bool {
	reg := regexp.MustCompile(p)

	return reg.Match([]byte(s))
}
