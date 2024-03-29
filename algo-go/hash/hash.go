package hash

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// smallerNumbersThanCurrent 1365. 有多少小于当前数字的数字
// 计数排序
func smallerNumbersThanCurrent(nums []int) []int {
	count := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
	}

	for i := 1; i <= 100; i++ {
		count[i] += count[i-1]
	}

	result := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			result[i] = 0
		} else {
			result[i] = count[nums[i]-1]
		}
	}

	return result
}

// distributeCandies leetcode 575. 分糖果
func distributeCandies(candyType []int) int {
	mp := map[int]struct{}{}

	for _, n := range candyType {
		mp[n] = struct{}{}
	}

	return min(len(candyType)/2, len(mp))
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

// longestConsecutive 128. 最长连续序列
func longestConsecutive(nums []int) int {
	set := make(map[int]struct{})

	for _, n := range nums {
		set[n] = struct{}{}
	}

	ans := 0

	for n := range set {
		if _, ok := set[n-1]; !ok {
			cur := n
			curLen := 1

			for {
				_, ok := set[cur+1]
				if !ok {
					break
				}

				cur++
				curLen++
			}

			ans = max(ans, curLen)
		}
	}
	return ans
}

// romanToInt leetcode 13. 罗马数字转整数
func romanToInt(s string) int {
	m := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	res := 0

	n := len(s)
	for i := 0; i < n; i++ {
		var k1 string
		var k2 string
		k1 = s[i : i+1]
		if i+2 <= n {
			k2 = s[i+1 : i+2]
		}

		if m[k1] < m[k2] {
			res += m[k2] - m[k1]
			i++
		} else {
			res += m[k1]
		}
	}
	return res
}
