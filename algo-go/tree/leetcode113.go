package tree

func pathSum(root *TreeNode, sum int) [][]int {
	var res = make([][]int, 0)
	var dfs func(root *TreeNode, path []int, sum int)

	dfs = func(root *TreeNode, path []int, sum int) {
		if root == nil {
			return
		}

		path = append(path, root.Val)

		if root.Left == nil && root.Right == nil && sum == root.Val {
			cp := append(make([]int, 0), path...)
			res = append(res, cp)
		}

		dfs(root.Left, path, sum-root.Val)
		dfs(root.Right, path, sum-root.Val)
	}

	dfs(root, []int{}, sum)

	return res
}
