package sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	cases := []struct {
		a    []int
		want []int
	}{{[]int{1, 2, 4, 3}, []int{1, 2, 3, 4}}}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			MergeSort(c.a)
			for i := 0; i < len(c.want); i++ {
				if c.want[i] != c.a[i] {
					t.Error("sort algorithem error")
				}
			}
		})
	}
}
