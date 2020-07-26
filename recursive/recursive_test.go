package recursive

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Del(t *testing.T) {
	tests := []struct {
		Nums   []int
		Target int

		Want []int
	}{
		{[]int{1, 2, 3, 4}, 3, []int{1, 2, 4}},
		{[]int{1, 2, 3, 4}, 6, []int{1, 2, 3, 4}},
		{[]int{1}, 1, []int{}},
		{[]int{}, 3, []int{}},
		{[]int{1, 2, 3, 4}, 4, []int{1, 2, 3}},
	}

	for _, test := range tests {
		res := Del(test.Nums, test.Target)

		assert.Equal(t, test.Want, res)
	}
}

func Test_Subsets(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	res := Subsets1(nums)
	t.Logf("----------->%#v\n", res)
	t.Helper()
}

func TestLetterCasePermutation(t *testing.T) {
	input := "12ab"
	res := LetterCasePermutation(input)
	t.Log(res)
}
