package offer

import (
	"fmt"
	"strconv"
	"strings"
)

func serialize(root *TreeNode) string {
	var dfs func(root *TreeNode)
	var res []string
	dfs = func(root *TreeNode) {
		if root == nil {
			res = append(res, "nil")
			return
		}

		res = append(res, fmt.Sprintf("%v", root.Val))

		dfs(root.Left)
		dfs(root.Right)
	}

	dfs(root)

	return strings.Join(res, ",")
}

func deserialize(s string) *TreeNode {
	ss := strings.Split(s, ",")
	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if len(ss) == 0 {
			return nil
		}

		if ss[0] == "nil" {
			ss = ss[1:]
			return nil
		}
		v, _ := strconv.Atoi(ss[0])
		ss = ss[1:]

		root := &TreeNode{Val: v}
		root.Left = dfs()
		root.Right = dfs()
		return root
	}

	return dfs()
}
