package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SortArrayByParityII(t *testing.T) {
	tests := []struct {
		Input []int
		Want  []int
	}{
		{Input: []int{4, 2, 5, 7}, Want: []int{4, 5, 2, 7}},
		{Input: []int{4, 5, 2, 7}, Want: []int{4, 5, 2, 7}},
	}

	for _, item := range tests {
		res := SortArrayByParityII(item.Input)
		assert.Equal(t, item.Want, res)
	}
}

func Test_Fib(t *testing.T) {
	tests := []struct {
		Input int
		Want  int
	}{
		{Input: 0, Want: 0},
		{Input: 1, Want: 1},
		{Input: 2, Want: 1},
		{Input: 3, Want: 2},
		{Input: 4, Want: 3},
		{Input: 5, Want: 5},
	}
	for _, item := range tests {
		res := Fib(item.Input)
		assert.Equal(t, item.Want, res)
	}
}

func Test_MinimumAbsDifference(t *testing.T) {
	test := []int{4, 2, 1, 3}
	t.Log(MinimumAbsDifference(test))
}

func Test_NumRookCaptures(t *testing.T) {
	cases := []struct {
		Input [][]byte
		Want  int
	}{
		{
			Input: [][]byte{
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', 'p', '.', '.', '.', '.'},
				{'.', '.', '.', 'R', '.', '.', '.', 'p'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', 'p', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
			},
			Want: 3,
		},
		{
			Input: [][]byte{
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', 'p', 'p', 'p', 'p', 'p', '.', '.'},
				{'.', 'p', 'p', 'B', 'p', 'p', '.', '.'},
				{'.', 'p', 'B', 'R', 'B', 'p', '.', '.'},
				{'.', 'p', 'p', 'B', 'p', 'p', '.', '.'},
				{'.', 'p', 'p', 'p', 'p', 'p', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
			},
			Want: 0,
		},
		{
			Input: [][]byte{
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', 'p', '.', '.', '.', '.'},
				{'.', '.', '.', 'p', '.', '.', '.', '.'},
				{'p', 'p', '.', 'R', '.', 'p', 'B', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', 'B', '.', '.', '.', '.'},
				{'.', '.', '.', 'p', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.'},
			},
			Want: 3,
		},
	}

	for _, test := range cases {
		res := NumRookCaptures(test.Input)
		assert.Equal(t, test.Want, res)
	}
}

func Test_ProjectionArea(t *testing.T) {
	cases := []struct {
		Array [][]int
		Want  int
	}{
		{[][]int{{1, 2}, {3, 4}}, 17},
		{[][]int{{2}}, 5},
		{[][]int{{1, 0}, {0, 2}}, 8},
	}

	for _, test := range cases {
		res := ProjectionArea(test.Array)
		assert.Equal(t, test.Want, res)
	}

}
