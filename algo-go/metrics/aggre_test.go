package metrics

import (
	"testing"
	"time"
)

func Test_Tree(t *testing.T) {
	tree := NewAggreTree(3)

	cases := []struct {
		path string
		want string
	}{
		{"/test/0", "/test/0"},
		{"/test/1", "/test/1"},
		{"/test/2", "/test/2"},
		{"/test/3", "/test/*"},
		{"/test/4", "/test/*"},
		{"/test/5/a", "/test/*/a"},
		{"/test/6/b", "/test/*/b"},
		{"/test/7/c", "/test/*/c"},
		{"/test/8/d", "/test/*/*"},
		{"/test/8/d/a", "/test/*/*/a"},
		{"/test/8/d/b", "/test/*/*/b"},
		{"/test/8/d/c", "/test/*/*/c"},
		{"/test/8/d/e", "/test/*/*/*"},
		{"/test/9", "/test/*"},
		{"/test/10", "/test/*"},
		{"/test/11", "/test/*"},
		{"/test/12", "/test/*"},
	}

	for _, c := range cases {
		time.Sleep(time.Millisecond * 10)
		t.Run("", func(t *testing.T) {
			got := tree.Query(c.path)
			if got != c.want {
				t.Errorf("want: %v, got: %v", c.want, got)
			}
		})
	}
}
