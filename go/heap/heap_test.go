package heap

import (
	"fmt"
	"testing"
)

func Test_frequencySort(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{{
		"loveleetcode", "eeeeoollvtdc",
	}}

	for i, c := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			res := frequencySort(c.input)

			for res != c.want {
				t.Errorf("res= %v, want=%v", res, c.want)
			}
		})
	}
}
