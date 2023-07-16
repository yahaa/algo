package offer

import (
	"testing"
)

func Test_serialize(t *testing.T) {
	cases := []struct {
		Name string
		Args *TreeNode
		Want string
	}{
		{
			Name: "case1",
			Args: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
			},
			Want: "1,2,nil,nil,3,4,nil,nil,5,nil,nil",
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			got := serialize(c.Args)
			if got != c.Want {
				t.Errorf("want: %v, got: %v", c.Want, got)
			}
		})
	}
}

func Test_dserialize(t *testing.T) {
	cases := []struct {
		Name string
		Want string
		Args string
	}{
		{
			Name: "case1",
			Want: "1,2,nil,nil,3,4,nil,nil,5,nil,nil",
			Args: "1,2,nil,nil,3,4,nil,nil,5,nil,nil",
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			n := deserialize(c.Args)
			got := serialize(n)
			if got != c.Args {
				t.Errorf("want: %v, got: %v", c.Want, got)
			}
		})
	}
}
