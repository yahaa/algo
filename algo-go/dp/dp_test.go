package dp

import (
	"reflect"
	"testing"
)

func Test_findMaxValue(t *testing.T) {
	{
		got := findMaxValue([]int{3, 5}, 135600)
		want := 55555
		if !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v got: %v", want, got)

		}
	}

	{

		got := findMaxValue([]int{1, 3, 5}, 135600)
		want := 135555
		if !reflect.DeepEqual(want, got) {
			t.Errorf("want: %v got: %v", want, got)

		}
	}

}
