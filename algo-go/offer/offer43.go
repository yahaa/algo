package offer

// leetcode 43
func countDigitOne(n int) int {
	// n 划分为三个部位 3101592
	// cur 为 base 位的值
	// a 为 base 位左边的值
	// b 为 base 位右边剩下的值
	// [a][cur][b]

	// 如果 cur > 1
	// a 的范围为 [0-a]，总共 a+1 个
	// b 的范围为 [0-base]，总共 base 个
	// 总的个数为 (a+1)*base

	// 如果 cur == 1
	// a 的范围为 [0-a],总共 a+1 个
	// 当 a 的范围取最大值 "a" 时， b 的范围为 [0-b] 总共 b+1 个
	// 当 a 的范围为 [0-a) 时，b 的范围为 [0-base_1] 总共 base 个
	// 总个数为 b+1+a*base

	// 如果 cur < 1
	// a 的范围为 [0-a), 总共 a 个，为什么 a 的范围不能是 [0-a]?
	//    因为我们求的是 base 位为 1 ,
	//    当我们把 base 位设置为 1 后 a 的范围是 [0-a],那么当前数字将会 > n
	// 当 a 的范围为 [0-a) 时，b 的范围为 [0-base_1] 总共 base 个
	// 总共 (a)*(base) 个

	base := 1
	res := 0
	for base <= n {
		a := n / base / 10
		cur := (n / base) % 10
		b := n % base

		if cur > 1 {
			res += (a + 1) * base
		} else if cur == 1 {
			res += a*base + b + 1
		} else {
			res += a * base
		}

		base *= 10
	}

	return res
}
