package offer

import (
	"container/list"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder5(root *TreeNode) []int {
	res := make([]int, 0)

	if root == nil {
		return res
	}

	que := list.New()

	que.PushBack(root)

	for que.Len() > 0 {
		count := que.Len()
		for count > 0 {
			t := que.Front().Value.(*TreeNode)
			que.Remove(que.Front())

			res = append(res, t.Val)

			if t.Left != nil {
				que.PushBack(t.Left)
			}

			if t.Right != nil {
				que.PushBack(t.Right)
			}

			count--
		}
	}

	return res
}

func levelOrder(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	queue := list.New()

	queue.PushBack(root)

	for queue.Len() > 0 {
		h := queue.Front()

		queue.Remove(h)

		head := h.Value.(*TreeNode)

		if head.Left != nil {
			queue.PushBack(head.Left)
		}

		if head.Right != nil {
			queue.PushBack(head.Right)
		}

		res = append(res, head.Val)
	}

	return res
}

func levelOrder2(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := list.New()

	queue.PushBack(root)

	for queue.Len() > 0 {
		tmp := make([]int, 0)
		levelSize := queue.Len()

		for i := 0; i < levelSize; i++ {
			h := queue.Front()

			queue.Remove(h)

			head := h.Value.(*TreeNode)

			if head.Left != nil {
				queue.PushBack(head.Left)
			}

			if head.Right != nil {
				queue.PushBack(head.Right)
			}
			tmp = append(tmp, head.Val)
		}

		res = append(res, tmp)
	}

	return res
}

func levelOrder3(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := list.New()

	queue.PushBack(root)

	for queue.Len() > 0 {
		tmp := make([]int, 0)
		levelSize := queue.Len()

		for i := 0; i < levelSize; i++ {
			h := queue.Front()

			queue.Remove(h)

			head := h.Value.(*TreeNode)

			if head.Left != nil {
				queue.PushBack(head.Left)
			}

			if head.Right != nil {
				queue.PushBack(head.Right)
			}
			tmp = append(tmp, head.Val)
		}

		if len(res)%2 == 1 {
			reverse(tmp)
		}

		res = append(res, tmp)
	}

	return res
}

func reverse(a []int) {
	i, j := 0, len(a)-1
	for i < j {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
}
