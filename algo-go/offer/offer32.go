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

func levelOrder4(root *TreeNode) [][]int {
	res := make([][]int, 0)

	queue := list.New()

	if root == nil {
		return res
	}

	flag := 0
	queue.PushBack(root)

	for queue.Len() > 0 {
		size := queue.Len()
		level := make([]int, 0)
		for size > 0 {
			if flag%2 == 0 {
				h := queue.Front()
				queue.Remove(h)

				n := h.Value.(*TreeNode)
				level = append(level, n.Val)

				if n.Left != nil {
					queue.PushBack(n.Left)
				}

				if n.Right != nil {
					queue.PushBack(n.Right)
				}
			} else {
				h := queue.Back()
				queue.Remove(h)

				n := h.Value.(*TreeNode)
				level = append(level, n.Val)

				if n.Right != nil {
					queue.PushFront(n.Right)
				}

				if n.Left != nil {
					queue.PushFront(n.Left)
				}
			}
			size--
		}
		flag++
		res = append(res, level)
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
