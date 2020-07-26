package lc211

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	tree := Constructor()
	tree.AddWord("bad")
	tree.AddWord("dad")
	tree.AddWord("mad")

	assert.Equal(t, false, tree.Search("pad"))
	assert.Equal(t, true, tree.Search("bad"))
	assert.Equal(t, true, tree.Search(".ad"))
	assert.Equal(t, true, tree.Search("b.."))
}
