package link

import (
	"fmt"
	"testing"
)

func Test_orderInsert(t *testing.T) {

	tests := []struct {
		head *ListNode
		node *ListNode
		want *ListNode
	}{
		{
			head: &ListNode{
				Val:  1,
				Next: &ListNode{Val: 5},
			},
			node: &ListNode{Val: 3},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
		{
			head: &ListNode{
				Val:  1,
				Next: &ListNode{Val: 5},
			},
			node: &ListNode{Val: 1},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},

		{
			head: &ListNode{
				Val:  1,
				Next: &ListNode{Val: 1},
			},
			node: &ListNode{Val: 3},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		},

		{
			head: &ListNode{
				Val:  1,
				Next: &ListNode{Val: 2},
			},
			node: &ListNode{Val: 3},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		},

		{
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 2,
					}},
			},
			node: &ListNode{Val: 3},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  2,
						Next: &ListNode{Val: 3},
					},
				},
			},
		},
	}

	for i, c := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			head := orderInsert(c.node, c.head)

			for head != nil && c.want != nil {
				if head.Val != c.want.Val {
					t.Errorf("head.Val= %v, Want.Val=%v", head.Val, c.want.Val)
					break
				}
				head = head.Next
				c.want = c.want.Next
			}
		})
	}

}
