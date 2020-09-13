package bst

import "fmt"

// TreeNode node
type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
	Father *TreeNode
}

// BSTree 二叉搜索树
type BSTree struct {
	root *TreeNode
}

func (bst *BSTree) String() string {
	var inOrder func(root *TreeNode)
	var res string

	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inOrder(root.Left)
		res = fmt.Sprintf("%s %d", res, root.Val)

		inOrder(root.Right)
	}

	inOrder(bst.root)
	return res
}

// Insert insert one node into the BST
func (bst *BSTree) Insert(val int) {
	t := &TreeNode{Val: val}

	if bst.root == nil {
		bst.root = t
		return
	}

	_, p := find(bst.root, val)
	if p == nil {
		return
	}

	t.Father = p

	if p.Val > val {
		p.Left = t
	} else {
		p.Right = t
	}
}

// Delete del
func (bst *BSTree) Delete(val int) {
	t, _ := find(bst.root, val)
	if t == nil {
		return
	}

	// 1 t if leaf node
	if t.Left == nil && t.Right == nil {
		// 只有 root 节点的 father 为空
		if t.Father == nil {
			bst.root = nil
		} else if t.Father.Left == t {
			t.Father.Left = nil
		} else {
			t.Father.Right = nil
		}

		return
	}

	// 2 t have left and right node
	if t.Left != nil && t.Right != nil {
		leftMax := t.Left

		for leftMax.Right != nil {
			leftMax = leftMax.Right
		}

		if t.Left == leftMax {
			leftMax.Right = t.Right
			t.Right = nil
			t.Left = nil

			if t.Father == nil {
				bst.root = leftMax
			} else if t.Father.Left == t {
				t.Father.Left = leftMax
			} else if t.Father.Right == t {
				t.Father.Right = leftMax
			}

			leftMax.Father = t.Father

			// 使得 t 指向为空被 gc
			t.Father = nil

		} else {
			t.Right, leftMax.Right = leftMax.Right, t.Right
			leftMax.Left, t.Left = t.Left, leftMax.Left

			if t.Father == nil {
				bst.root = leftMax
			} else if t.Father.Left == t {
				t.Father.Left = leftMax
			} else if t.Father.Right == t {
				t.Father.Right = leftMax
			}

			if leftMax.Father.Left == leftMax {
				leftMax.Father.Left = t
			} else if leftMax.Father.Right == leftMax {
				leftMax.Father.Right = t
			}

			if t.Father == nil {
				leftMax.Father = nil
			} else {
				t.Father, leftMax.Father = leftMax.Father, t.Father
			}
		}

		if leftMax.Left != nil {
			leftMax.Left.Father = leftMax
		}

		if leftMax.Right != nil {
			leftMax.Right.Father = leftMax
		}

		if t.Left != nil {
			t.Left.Father = t
		}

		if t.Right != nil {
			t.Right.Father = t
		}
	}

	// 3
	if t.Left != nil && t.Right == nil {
		if t.Father == nil {
			bst.root = t.Left
		} else if t.Father.Left == t {
			t.Father.Left = t.Left
		} else if t.Father.Right == t {
			t.Father.Right = t.Left
		}

		t.Left.Father = t.Father

		t.Left = nil
		t.Father = nil
		return
	}

	if t.Right != nil && t.Left == nil {
		if t.Father == nil {
			bst.root = t.Right
		} else if t.Father.Left == t {
			t.Father.Left = t.Right
		} else if t.Father.Right == t {
			t.Father.Right = t.Right
		}

		t.Right.Father = t.Father

		t.Right = nil
		t.Father = nil
		return
	}
}

// Find find
func (bst *BSTree) Find(val int) bool {
	if bst.root == nil {
		return false
	}

	t, _ := find(bst.root, val)
	return t != nil
}

// find return target
func find(r *TreeNode, val int) (t, p *TreeNode) {
	var f *TreeNode
	for r != nil {
		f = r

		if r.Val == val {
			return r, r.Father
		}
		if r.Val > val {
			r = r.Left
		} else {
			r = r.Right
		}
	}

	return r, f
}

// NewBSTree new tree node
func NewBSTree() *BSTree {
	return &BSTree{}
}
