package recursive

import (
	"fmt"
	"sort"
	"unicode"
)

// Subsets1 return subset of nums slice
// 方法 1 比较蠢没有方法论
func Subsets1(nums []int) [][]int {
	var (
		ans = make([][]int, 0)
		cur = make([]int, 0)
		dfs func(n int, cur []int, nums []int)
	)

	cur = make([]int, 0)
	dfs = func(n int, cur []int, nums []int) {
		if n > len(nums) {
			return
		}

		if n == 0 {
			ans = append(ans, []int{})
			return
		}

		if len(cur) == n {
			t := make([]int, n)
			copy(t, cur)
			ans = append(ans, t)
			return
		}

		for _, item := range nums {
			var max int
			if len(cur) == 0 {
				max = item
			} else {
				max = cur[len(cur)-1]
			}
			if !Found(item, cur) && item >= max && len(cur) < n {
				cur = append(cur, item)
				dfs(n, cur, nums)
				cur = Del(cur, item)
			}
		}
	}

	sort.Ints(nums)
	for i := 0; i <= len(nums); i++ {
		dfs(i, cur, nums)
	}
	return ans
}

// Subsets2 子集树解法
func Subsets2(nums []int) [][]int {
	/*** 子集树思路
		[1,2,3]
			                    ([])
		           ([1])                        ([])
	      ([1,2])         ([1])           ([2])       ([])
	 ([1,2,3]) ([1,2]) ([1,3])([1])    ([2,3])([2]) ([3]) ([])
		***/
	var (
		res = make([][]int, 0)
		dfs func(nums, cur []int, index int)
	)

	dfs = func(nums, cur []int, index int) {
		if index == len(nums) {
			t := make([]int, len(cur))
			copy(t, cur)
			res = append(res, t)
			return
		}

		dfs(nums, cur, index+1) // 不选

		cur = append(cur, nums[index]) // 选，需要加入当前状态
		dfs(nums, cur, index+1)
	}

	dfs(nums, []int{}, 0)
	return res
}

// Subsets3 站在子集的角度去考虑
func Subsets3(nums []int) [][]int {
	var (
		res = make([][]int, 0)
		dfs func(nums, cur []int, index int)
	)

	/***
	子集长度为 0 有几个
	子集长度为 1 有几个
	子集长度为 2 有几个
	...
	子集长度为 n 有几个
	***/

	dfs = func(nums, cur []int, index int) {
		t := make([]int, len(cur))
		copy(t, cur)
		res = append(res, t)

		for i := index; i < len(nums); i++ {
			cur = append(cur, nums[i])
			// 下一次的 index 是 i+1 而不是 index+1
			dfs(nums, cur, i+1)
			cur = cur[0 : len(cur)-1]
		}
	}
	dfs(nums, []int{}, 0)
	return res
}

// Found find whether target in nums slice
func Found(target int, nums []int) bool {
	for _, v := range nums {
		if v == target {
			return true
		}
	}
	return false
}

// Del delete item for slice
func Del(nums []int, target int) []int {
	for i, v := range nums {
		if v == target {
			nums = append(nums[0:i], nums[i+1:]...)
			break
		}
	}
	return nums
}

// Permutation 全排列,排列树
func Permutation(nums []int) [][]int {
	var (
		res = make([][]int, 0)
		dfs func(nums []int, n int)
	)

	dfs = func(nums []int, n int) {
		if n == len(nums) {
			t := make([]int, n)
			copy(t, nums)
			res = append(res, t)
			return
		}

		for i := n; i < len(nums); i++ {
			nums[n], nums[i] = nums[i], nums[n]
			dfs(nums, n+1)
			nums[n], nums[i] = nums[i], nums[n]
		}
	}

	dfs(nums, 0)
	return res
}

// SubsetsWithDup 包含重复数字的子集
func SubsetsWithDup(nums []int) [][]int {
	var (
		res = make([][]int, 0)
		dfs func(nums, cur []int, index int, pre bool)
	)

	dfs = func(nums, cur []int, index int, pre bool) {
		if index == len(nums) {
			t := make([]int, len(cur))
			copy(t, cur)
			res = append(res, t)
			return
		}

		dfs(nums, cur, index+1, false)
		if index > 0 && nums[index] != nums[index-1] || pre {
			cur = append(cur, nums[index])
			dfs(nums, cur, index+1, true)
		}
	}
	dfs(nums, []int{}, 0, true)
	return res
}

// LetterCombinations 字母组合
func LetterCombinations(digits string) []string {
	var (
		dfs      func(digits, cur string, index int)
		res      = make([]string, 0)
		chartMap = map[byte][]string{
			'2': {"a", "b", "c"},
			'3': {"d", "e", "f"},
			'4': {"g", "h", "i"},
			'5': {"j", "k", "l"},
			'6': {"m", "n", "o"},
			'7': {"p", "q", "r", "s"},
			'8': {"t", "u", "v"},
			'9': {"w", "x", "y", "z"},
		}
	)
	dfs = func(digits, cur string, index int) {
		if index == len(digits) {
			res = append(res, cur)
			return
		}

		num := digits[index]

		chars := chartMap[num]

		for _, s := range chars {
			dfs(digits, cur+s, index+1)
		}
	}

	if digits == "" {
		return res
	}

	dfs(digits, "", 0)
	return res
}

// GenerateParenthesis generate well-format parenthesis
func GenerateParenthesis(n int) []string {
	var (
		res = make([]string, 0)
		dfs func(left, right, index int, cur string)
	)

	dfs = func(left, right, index int, cur string) {
		if left > right {
			return
		}
		if index == n*2 {
			res = append(res, cur)
			return
		}

		if left > 0 {
			dfs(left-1, right, index+1, cur+"(")
		}
		if right > 0 {
			dfs(left, right-1, index+1, cur+")")
		}
	}

	dfs(n, n, 0, "")
	return res
}

// SolveSudoku solve sudoku by backtracking
func SolveSudoku(board [][]byte) {
	var (
		dfs         func(total int) bool
		checkRow    func(c byte, i, j int) bool
		checkColumn func(c byte, i, j int) bool
		checkBox    func(c byte, i, j int) bool
		nums        = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
		total       = 0
		rows        = len(board)
		column      = 0
	)

	if rows == 0 {
		return
	}

	column = len(board[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < column; j++ {
			if board[i][j] == '.' {
				total++
			}
		}
	}

	checkRow = func(c byte, i, j int) bool {
		for r := i + 1; r < rows; r++ {
			if board[r][j] == c {
				return false
			}
		}

		for r := i - 1; r >= 0; r-- {
			if board[r][j] == c {
				return false
			}
		}
		return true
	}

	checkColumn = func(c byte, i, j int) bool {
		for r := j + 1; r < column; r++ {
			if board[i][r] == c {
				return false
			}
		}

		for r := j - 1; r >= 0; r-- {
			if board[i][r] == c {
				return false
			}
		}
		return true
	}

	checkBox = func(c byte, i, j int) bool {
		br, bc := i/3*3, j/3*3
		for x := 0; x < 3; x++ {
			for y := 0; y < 3; y++ {
				if board[x+br][y+bc] == c {
					return false
				}
			}
		}
		return true
	}

	dfs = func(total int) bool {
		if total == 0 {
			return true
		}

		for x := 0; x < rows; x++ {
			for y := 0; y < column; y++ {
				if board[x][y] == '.' {
					for _, num := range nums {
						if checkRow(num, x, y) && checkColumn(num, x, y) && checkBox(num, x, y) {
							board[x][y] = num
							if dfs(total - 1) {
								return true
							}
						}
					}
					board[x][y] = '.'
					return false
				}
			}
		}
		return false
	}

	dfs(total)
}

// GetPermutation get n'th permutation leetcode 60
func GetPermutation(n int, k int) string {
	var (
		f    = map[int]int{1: 1}
		nums = make([]int, n)
		res  string
	)

	if n == 0 {
		return res
	}

	k--
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}

	for i := 2; i <= n; i++ {
		f[i] = f[i-1] * i
	}

	for i := n - 1; i >= 1; i-- {
		index := k / f[i]
		res = fmt.Sprintf("%s%d", res, nums[index])
		nums = append(nums[0:index], nums[index+1:]...)
		k = k % f[i]
	}

	res = fmt.Sprintf("%s%d", res, nums[0])
	return res
}

// Combine leetcode 77
func Combine(n int, k int) [][]int {
	var (
		res = make([][]int, 0)
		dfs func(cur []int, f, k int)
	)

	dfs = func(cur []int, f, k int) {
		if k == 0 {
			res = append(res, append(make([]int, 0), cur...))
			return
		}

		for i := f; i <= n; i++ {
			cur = append(cur, i)
			dfs(cur, i+1, k-1)
			cur = cur[0 : len(cur)-1]
		}
	}

	dfs([]int{}, 1, k)

	return res
}

// Exist leetcode probleam 79, find whether the specify word exist in board
func Exist(board [][]byte, word string) bool {
	/*
		[["A","B","C","E"],
		 ["S","F","E","S"],
		 ["A","D","E","E"]]
		"ABCESEEEFS"
	*/
	var (
		rows, colume = len(board), len(board[0])
		dx           = []int{0, -1, 0, 1}
		dy           = []int{-1, 0, 1, 0}
		dfs          func(x, y, index int) bool
	)

	dfs = func(x, y, index int) bool {

		if index >= len(word) {
			return true
		}

		if x < 0 || y < 0 || x >= rows || y >= colume {
			return false
		}
		if board[x][y] == '*' {
			return false
		}

		if board[x][y] == word[index] {
			board[x][y] = '*'
			for i := 0; i < 4; i++ {
				if dfs(x+dx[i], y+dy[i], index+1) {
					return true
				}
			}
			board[x][y] = word[index]
		}

		return false
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < colume; j++ {
			if board[i][j] == word[0] {
				if dfs(i, j, 0) {
					return true
				}
			}
		}
	}

	return false
}

// LetterCasePermutation leetcode 784
func LetterCasePermutation(s string) []string {
	var (
		res = make([]string, 0)
		dfs func(cur string, index int)
	)

	dfs = func(cur string, i int) {
		if i == len(s) {
			res = append(res, cur)
			return
		}

		dfs(cur+string(s[i]), i+1)

		if unicode.IsLower(rune(s[i])) {
			cur = cur + string(unicode.ToUpper(rune(s[i])))
			dfs(cur, i+1)
		}

		if unicode.IsUpper(rune(s[i])) {
			cur = cur + string(unicode.ToLower(rune(s[i])))
			dfs(cur, i+1)
		}
	}

	dfs("", 0)
	return res
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	}

	if root.Val > high {
		return rangeSumBST(root.Left, low, high)
	}

	return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	ans := make([][]int, 0)
	var dfs func(candidates []int, cur []int, target, index int)

	dfs = func(candidates []int, cur []int, target, index int) {
		if target < 0 {
			return
		}

		if target == 0 {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			ans = append(ans, tmp)
			return
		}

		for i := index; i < len(candidates); i++ {
			if target < candidates[i] {
				break
			}

			cur = append(cur, candidates[i])
			dfs(candidates, cur, target-candidates[i], i)
			cur = cur[:len(cur)-1]
		}
	}

	dfs(candidates, []int{}, target, 0)

	return ans
}

func combinationSum2(candidates []int, target int) [][]int {
	// 1. 排序 [1,1,2,5,6,7,10]
	// 2. 递归回溯
	//    递归出口
	//.      target < 0
	//       target == 0
	//

	sort.Ints(candidates)

	ans := make([][]int, 0)

	var dfs func(candidates []int, ans []int, target int, start int)

	dfs = func(candidates []int, cur []int, target int, start int) {
		if target < 0 {
			return
		}

		if target == 0 {
			cp := make([]int, len(cur))
			copy(cp, cur)
			ans = append(ans, cp)
			return
		}

		for i := start; i < len(candidates); i++ {
			if candidates[i] > target {
				break
			}

			if i != start && candidates[i] == candidates[i-1] {
				continue
			}

			cur = append(cur, candidates[i])
			dfs(candidates, cur, target-candidates[i], i+1)
			cur = cur[:len(cur)-1]
		}
	}

	dfs(candidates, []int{}, target, 0)
	return ans
}
