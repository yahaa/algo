package bst

import (
	"fmt"
	"testing"
)

func TestNewBSTree(t *testing.T) {
	bst := NewBSTree()

	bst.Insert(1)
	bst.Insert(2)
	bst.Insert(30)
	bst.Insert(3)
	bst.Insert(0)

	tests := []struct {
		input int
		want  bool
	}{
		{1, true},
		{2, true},
		{3, true},
		{0, true},
		{100, false},
		{30, true},
	}

	for i, c := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			res := bst.Find(c.input)
			if res != c.want {
				t.Errorf("want: %v, got: %v\n", c.want, res)
			}
		})
	}
}

func TestNewBSTreeDelete(t *testing.T) {
	bst := NewBSTree()

	bst.Insert(30)
	bst.Insert(5)
	bst.Insert(2)

	bst.Insert(40)

	bst.Insert(80)
	bst.Insert(60)
	bst.Insert(85)

	bst.Insert(35)
	bst.Insert(37)
	bst.Insert(36)

	bst.Insert(32)
	bst.Insert(31)
	bst.Insert(33)

	fmt.Printf("%v\n", bst)
	tests := []struct {
		input int
		want  bool
	}{
		{30, false},
		{5, false},
		{33, false},
		{31, false},
		{30, false},
		{40, false},
		{41, false},
		{40, false},
	}

	for i, c := range tests {
		t.Run(fmt.Sprintf("case %d delete %d", i, c.input), func(t *testing.T) {
			fmt.Printf("start: %v\n", bst)

			bst.Delete(c.input)

			fmt.Printf("ended: %v\n", bst)

			res := bst.Find(c.input)
			if res != c.want {
				t.Errorf("want: %v, got: %v\n", c.want, res)
			}
		})
	}

	bst.Insert(40)

	tests = []struct {
		input int
		want  bool
	}{
		{40, true},
		{42, false},
		{80, true},
	}

	for i, c := range tests {
		t.Run(fmt.Sprintf("case %d find %d", i, c.input), func(t *testing.T) {
			res := bst.Find(c.input)
			if res != c.want {
				t.Errorf("want: %v, got: %v\n", c.want, res)
			}
		})
	}

	it := NewIterator(bst)

	fmt.Printf("start: %v\n", bst)

	for it.HasNext() {
		fmt.Printf("%d ", it.Next())
	}
	fmt.Printf("\n")

}
