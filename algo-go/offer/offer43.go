package offer

// leetcode 43
func countDigitOne(n int) int {
	// "1"出现的次数 = sum ("1"在各个计数位上出现的次数)
	// 从个位开始向最高位统计
	// 3101592
	// 将数字拆分为[a...][cur][b...]
	// cur 为当前位
	var base, res int64 = 1, 0
	n64 := int64(n)
	for base <= n64 {
		// 计算 a..., cur, b...
		a := n64 / base / 10
		cur := (n64 / base) % 10
		b := n64 % base
		// 将当前位设为1，考察其他部分的变化范围
		if cur > 1 {
			// 一、cur > 1，
			//          [3101 ] 5 [92]
			// 变化范围：[0-3101] 1 [0-99]
			// 总个数：   (a+1)  *  base
			res += (a + 1) * base
		} else if cur == 1 {
			// 二、cur == 1，
			//             [310] 1 [592]
			// 1、变化范围 [0-309] 1 [0-999]
			//              a    *  base
			// 2、变化范围 [310]   1 [0-592]
			//               1   *   (b+1)
			// 总个数：a *base + (b + 1)
			res += a*base + b + 1

		} else {
			// 三、cur < 1，
			//           [31] 0 [1592]
			// 变化范围 [0-30] 1 [0-9999]
			// 总个数    a     *   base
			res += a * base
		}
		// 统计更高一位
		base *= 10
	}
	return int(res)
}
