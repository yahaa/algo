package str

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SubdomainVisits(t *testing.T) {
	test := []string{"9001 discuss.leetcode.com"}
	fmt.Println(SubdomainVisits(test))
}

func Test_AddBinary(t *testing.T) {
	cases := []struct {
		A    string
		B    string
		Want string
	}{
		{"1", "1", "10"},
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
	}

	for _, item := range cases {
		res := AddBinary(item.A, item.B)
		assert.Equal(t, item.Want, res)
	}
}

func Test_CountCharacters(t *testing.T) {
	cases := []struct {
		Words []string
		Chars string
		Want  int
	}{
		{Words: []string{"cat", "bt", "hat", "tree"}, Chars: "atach", Want: 6},
		{Words: []string{"hello", "world", "leetcode"}, Chars: "welldonehoneyr", Want: 10},
	}
	for _, item := range cases {
		res := CountCharacters(item.Words, item.Chars)
		assert.Equal(t, item.Want, res)
	}

	var tt *int
	fmt.Print(*tt)
}

func Test_LongestPalindrome(t *testing.T) {
	cases := []struct {
		Input string
		Want  int
	}{
		{"abccccdd", 7},
		{"civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth", 983},
		{"bb", 2},
	}

	for _, item := range cases {
		res := LongestPalindrome(item.Input)
		assert.Equal(t, item.Want, res)
	}
}
