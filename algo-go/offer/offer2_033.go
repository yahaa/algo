package offer

import "sort"

func sortString(s string) string {
	slice := []rune(s)

	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})

	return string(slice)
}

func groupAnagrams(strs []string) [][]string {
	group := make(map[string][]string)

	for _, item := range strs {
		ss := sortString(item)
		if len(group[ss]) > 0 {
			group[ss] = append(group[ss], item)
		} else {
			group[ss] = []string{item}
		}
	}

	ans := make([][]string, 0)

	for _, v := range group {
		ans = append(ans, v)
	}
	return ans
}
