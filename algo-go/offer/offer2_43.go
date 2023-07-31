package offer

import (
	"container/list"
)

type CBTInserter struct {
	queue *list.List
	root  *TreeNode
}

func ConstructorX(root *TreeNode) CBTInserter {
	t := CBTInserter{
		queue: list.New(),
		root:  root,
	}

	t.queue.PushBack(root)

	for t.queue.Len() > 0 {
		f := t.queue.Front().Value.(*TreeNode)
		if f.Left == nil || f.Right == nil {
			break
		}

		t.queue.PushBack(f.Left)
		t.queue.PushBack(f.Right)

		t.queue.Remove(t.queue.Front())
	}

	return t
}

func (this *CBTInserter) Insert(v int) int {
	f := this.queue.Front().Value.(*TreeNode)

	if f.Left == nil {
		f.Left = &TreeNode{Val: v}
		return f.Val
	}

	f.Right = &TreeNode{Val: v}

	this.queue.PushBack(f.Left)
	this.queue.PushBack(f.Right)

	this.queue.Remove(this.queue.Front())
	return f.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}
