package metrics

import (
	"testing"
	"time"
)

func Test_Tree(t *testing.T) {
	tree := NewTree(10)

	cases := []struct {
		path string
		want string
	}{
		{"/test/debug", "/test/debug"},
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
		time.Sleep(time.Second)
		t.Run("", func(t *testing.T) {
			got := tree.Query(c.path)
			if got != c.want {
				t.Errorf("want: %v, got: %v", c.want, got)
			}

			//d, _ := json.Marshal(tree.root)
			//t.Logf("%v", debug)
		})
	}
}
