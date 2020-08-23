package offer

func replaceSpace(s string) string {
	res := ""
	for _, c := range s {
		if c == ' ' {
			res += "%20"
		} else {
			res += string(c)
		}
	}
	return res

}
