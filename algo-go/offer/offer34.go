package offer

// pathSum leetcode 113
func pathSum(root *TreeNode, sum int) [][]int {
	res := make([][]int, 0)
	var dfs func(root *TreeNode, p []int, r int)

	dfs = func(root *TreeNode, p []int, r int) {
		if root == nil {
			return
		}

		p = append(p, root.Val)
		r -= root.Val

		if root.Left == nil && root.Right == nil && r == 0 {
			copy := append(make([]int, 0, len(p)), p...)
			res = append(res, copy)
		}

		dfs(root.Left, p, r)
		dfs(root.Right, p, r)
	}

	dfs(root, []int{}, sum)

	return res
}
