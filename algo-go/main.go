package main

import (
	"github.com/yahaa/leetcode/metrics"
)

func main() {
	tree := metrics.NewAggreTree(3)

	cases := []struct {
		path string
		want string
	}{
		{"/test/0", "/test/0"},
		{"/test/1", "/test/1"},
		{"/test/2", "/test/2"},
		{"/test/3", "/test/3"},
		{"/test/4", "/test/4"},
		{"/test/5", "/test/5"},
		{"/test/6", "/test/6"},
		{"/test/7", "/test/7"},
		{"/test/8", "/test/8"},
		{"/test/9", "/test/9"},
		{"/test/10", "/test/10"},
		{"/test/11", "/test/11"},
		{"/test/12", "/test/12"},
	}

	for _, c := range cases {
		tree.Query(c.path)
	}
}
