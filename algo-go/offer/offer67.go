package offer

import (
	"math"
	"strings"
)

func getNumStr(str string) (string, int) {
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return str, 0
	}

	f := uint8(str[0])
	var k = 1
	if f == 45 || f == 43 {
		str = str[1:]
		k = 44 - int(f)
	}

	str = strings.TrimLeft(str, "0")

	if len(str) == 0 {
		return str, 0
	}

	i := 0
	for i < len(str) {
		t := uint8(str[i])
		if t < 48 || t > 57 {
			break
		}

		i++
	}

	if i == 0 {
		return str, 0
	}

	str = str[0:i]

	return str, k
}

func strToInt(str string) int {
	str, k := getNumStr(str)
	if k == 0 {
		return 0
	}

	if len(str) == 0 {
		return 0
	}

	t := uint8(str[0])

	if t == 45 || t == 43 {
		return 0
	}

	if len(str) > 10 {
		if k == 1 {
			return math.MaxInt32
		} else if k == -1 {
			return math.MinInt32
		}
	}

	sum, b := 0, 1

	for i := len(str) - 1; i >= 0; i-- {
		sum += (int(str[i]) - 48) * b

		if k == 1 {
			if sum >= math.MaxInt32 {
				return math.MaxInt32
			}
		} else if k == -1 {
			if sum == math.MaxInt32-1 {
				return math.MinInt32
			} else if sum > math.MaxInt32 {
				return math.MinInt32
			}
		}

		b *= 10
	}

	sum *= k

	if sum > math.MaxInt32 {
		return math.MaxInt32
	}

	if sum < math.MinInt32 {
		return math.MinInt32
	}

	return sum
}
