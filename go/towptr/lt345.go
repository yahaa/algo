package towptr

func reverseVowels(s string) string {
	if len(s) == 0 {
		return s
	}

	bs := []byte(s)
	mp := map[byte]bool{
		byte('a'): true,
		byte('e'): true,
		byte('i'): true,
		byte('o'): true,
		byte('u'): true,
	}
	i, j := 0, len(bs)-1

	for i < j {
		for !mp[bs[i]] && i < j {
			i++
		}

		for !mp[bs[j]] && j > i {
			j--
		}

		bs[i], bs[j] = bs[j], bs[i]
		i++
		j--
	}

	return string(bs)
}
