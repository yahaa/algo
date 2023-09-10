package tree

import "container/list"

// Node node int value
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// IsUnivalTree leetcode 965
func IsUnivalTree(root *Node) bool {
	if root == nil {
		return true
	}

	var dfs func(r *Node, pVal int) bool

	dfs = func(r *Node, pVal int) bool {
		if r == nil {
			return true
		}
		return r.Val == pVal && dfs(r.Left, r.Val) && dfs(r.Right, r.Val)
	}

	return dfs(root, root.Val)
}

// longestConsecutive 298. 二叉树最长连续序列
func longestConsecutive(root *TreeNode) int {
	ans := 0

	var dfs func(root *TreeNode, cur int)

	dfs = func(root *TreeNode, cur int) {
		if root == nil {
			return
		}

		if root.Left != nil {
			if root.Left.Val == root.Val+1 {
				dfs(root.Left, cur+1)
			} else {
				dfs(root.Left, 1)
			}
		}

		if root.Right != nil {
			if root.Right.Val == root.Val+1 {
				dfs(root.Right, cur+1)
			} else {
				dfs(root.Right, 1)
			}
		}

		ans = max(ans, cur)
	}

	dfs(root, 1)

	return ans
}


// leetcode 589
type NNode struct {
	Val      int
	Children []*NNode
}

// preorder leetcode 589
func preorder(root *NNode) []int {
	res := make([]int, 0)
	var dfs func(root *NNode)

	dfs = func(root *NNode) {
		if root == nil {
			return
		}

		res = append(res, root.Val)

		for _, c := range root.Children {
			dfs(c)
		}
	}

	dfs(root)
	return res
}

// levelOrder 429
func levelOrder(root *NNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return nil
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		levelLen := queue.Len()
		levelVals := make([]int, 0)
		for levelLen > 0 {
			top := queue.Remove(queue.Front()).(*NNode)
			for _, n := range top.Children {
				queue.PushBack(n)
			}
			levelVals = append(levelVals, top.Val)
			levelLen--
		}

		result = append(result, levelVals)
	}
	return result
}

// leetcode 590
func postorder(root *NNode) []int {
	var dfs func(root *NNode)
	var res []int

	dfs = func(root *NNode) {
		if root == nil {
			return
		}

		for _, n := range root.Children {
			dfs(n)
		}
		res = append(res, root.Val)
	}

	dfs(root)
	return res
}
