package offer

import (
	"fmt"
	"math"
	"testing"
)

func TestStrToInt(t *testing.T) {
	cases := []struct {
		Input string
		Want  int64
	}{
		{"123", 123},
		{"+123", 123},
		{"-123", -123},
		{"-a123", 0},
		{"    -42", -42},
		{"4193 with words", 4193},
		{"words and 987", 0},
		{"-91283472332", -2147483648},
		{"+-2", 0},
		{"9223372036854775808", 2147483647},
		{"-91283472332", -2147483648},
		{"-2147483647", -2147483647},
		{"0-1", 0},
		{"10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000522545459", 2147483647},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			res := strToInt(c.Input)
			if int64(res) != c.Want {
				t.Errorf("not equal res %v ,want %v", res, c.Want)
			}
		})
	}
}

func TestSingleNumber(t *testing.T) {
	//cases := []struct {
	//	Input []int
	//	Want  int
	//}{
	//	{[]int{1, 2, 3, 1, 1, 2, 2}, 3},
	//	{[]int{3, 4, 3, 3}, 4},
	//	{[]int{9, 1, 7, 9, 7, 9, 7}, 1},
	//}
	//
	//for i, c := range cases {
	//	t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
	//		res := singleNumber1(c.Input)
	//		if res != c.Want {
	//			t.Errorf("not equal res %v ,want %v", res, c.Want)
	//		}
	//	})
	//}
}

func TestOffer16(t *testing.T) {
	cases := []struct {
		a    float64
		b    int
		want float64
	}{{2.0000, -2, 0.25}}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			res := myPow(c.a, c.b)

			if math.Abs(res-c.want) > 0.0000001 {
				t.Errorf("got %v,want %v", res, c.want)
			}
		})
	}
}

func TestOffer31(t *testing.T) {
	cases := []struct {
		a    []int
		b    []int
		want bool
	}{{[]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}, false}}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			res := validateStackSequences(c.a, c.b)

			if res != c.want {
				t.Errorf("got %v,want %v", res, c.want)
			}
		})
	}
}

func TestOffer45(t *testing.T) {
	cases := []struct {
		a    []int
		want string
	}{{[]int{10, 2}, "102"}}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			res := minNumber(c.a)

			if res != c.want {
				t.Errorf("got %v,want %v", res, c.want)
			}
		})
	}
}

func TestOffer13(t *testing.T) {

	tests := []struct {
		m, n, k int
		want    int
	}{
		{2, 3, 1, 3},
		{3, 2, 17, 6},
	}

	for i, c := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			res := movingCount(c.m, c.n, c.k)
			if res != c.want {
				t.Errorf("got %v, want %v", res, c.want)
			}
		})
	}
}

func TestOffer48(t *testing.T) {

	tests := []struct {
		s    string
		want int
	}{
		{"abc", 3},
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"abba", 2},
		{"pwwkew", 3},
	}

	for i, c := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			res := lengthOfLongestSubstring(c.s)
			if res != c.want {
				t.Errorf("got %v, want %v", res, c.want)
			}
		})
	}

	for i, c := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			res := lengthOfLongestSubstring2(c.s)
			if res != c.want {
				t.Errorf("got %v, want %v", res, c.want)
			}
		})
	}
}

func TestOffer11(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 4, 5, 1, 2}, 1},
	}

	for i, c := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			res := minArray2(c.nums)
			if res != c.want {
				t.Errorf("got %v, want %v", res, c.want)
			}
		})
	}
}
func TestOffer10(t *testing.T) {
	tests := []struct {
		num  int
		want int
	}{
		{2, 2},
		{7, 21},
		{0, 1},
	}

	for i, c := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			res := numWays(c.num)
			if res != c.want {
				t.Errorf("got %v, want %v", res, c.want)
			}
		})
	}
}
