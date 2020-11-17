package towptr

import (
	"fmt"
	"testing"
)

func Test_lt234(t *testing.T) {
	tests := []struct {
		Input string
		Want  string
	}{
		{Input: "hello", Want: "holle"},
	}

	for i, c := range tests {
		t.Run(fmt.Sprintf("case %d ", i), func(t *testing.T) {
			res := reverseVowels(c.Input)
			if c.Want != res {
				t.Errorf("want: %v, got: %v", c.Want, res)
			}
		})
	}
}

func Test_lt76(t *testing.T) {
	tests := []struct {
		Inputs string
		Inputt string
		Want   string
	}{
		{Inputs: "ADOBECODEBANC", Inputt: "ABC", Want: "BANC"},
		{Inputs: "AA", Inputt: "AA", Want: "AA"},
	}

	for i, c := range tests {
		t.Run(fmt.Sprintf("case %d ", i), func(t *testing.T) {
			res := minWindow(c.Inputs, c.Inputt)
			t.Logf("res: %v", res)
			if c.Want != res {
				t.Errorf("want: %v, got: %v", c.Want, res)
			}
		})
	}
}

func Test_lt3(t *testing.T) {
	tests := []struct {
		Input string
		Want  int
	}{
		{"abcabcbb", 3},
		{"bbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"tmmzuxt", 5},
		{"dvdf", 3},
	}

	for i, c := range tests {
		t.Run(fmt.Sprintf("case %d ", i), func(t *testing.T) {
			res := lengthOfLongestSubstring(c.Input)
			if c.Want != res {
				t.Errorf("want: %v, got: %v", c.Want, res)
			}
		})
	}
}

func Test_lt61(t *testing.T) {

}
