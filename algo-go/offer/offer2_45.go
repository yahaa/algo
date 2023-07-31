package offer

func findBottomLeftValue2(root *TreeNode) int {
	if root == nil {
		return -1
	}

	maxLevel := -1
	maxLeft := root.Val

	var dfs func(root *TreeNode, level int)

	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if level >= maxLevel {
			maxLevel = level
			if root.Left != nil {
				maxLeft = root.Left.Val
			} else if root.Right != nil {
				maxLeft = root.Right.Val
			}
		}

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)

	return maxLeft
}

func findBottomLeftValue(root *TreeNode) int {
	maxLevel := 1
	firsts := make([]int, 0)

	var dfs func(root *TreeNode, level int)

	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if len(firsts) < level {
			firsts = append(firsts, root.Val)
		}

		maxLevel = max(maxLevel, level)

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 1)

	return firsts[maxLevel-1]
}
