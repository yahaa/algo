package search

import (
	"reflect"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		n    int
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		// {
		// 	name: "case1",
		// 	args: args{
		// 		n:    23121,
		// 		nums: []int{2, 4, 9},
		// 	},
		// 	want: []int{2, 2, 9, 9, 9},
		// },
		// {
		// 	name: "case2",
		// 	args: args{
		// 		n:    2533,
		// 		nums: []int{1, 2, 9, 4},
		// 	},
		// 	want: []int{2, 4, 9, 9},
		// },
		{
			name: "case3",
			args: args{
				n:    2543,
				nums: []int{1, 2, 5, 4},
			},
			want: []int{2, 5, 4, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.n, tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
