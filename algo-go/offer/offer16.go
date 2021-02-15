package offer

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	k := 1
	if n < 0 {
		k = -1
		n = -n
	}

	res := 1.0
	for n > 0 {
		if (n & 1) > 0 {
			res = res * x
		}
		x = x * x

		n >>= 1
	}

	if k == 1 {
		return res
	}

	return 1 / res
}
