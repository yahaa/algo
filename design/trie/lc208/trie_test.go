package lc208

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	tree := Constructor()
	tree.Insert("apple")
	assert.Equal(t, true, tree.Search("apple"))

	assert.Equal(t, true, tree.StartsWith("app"))

	assert.Equal(t, false, tree.Search("app"))
	tree.Insert("app")
	assert.Equal(t, true, tree.Search("app"))

}
