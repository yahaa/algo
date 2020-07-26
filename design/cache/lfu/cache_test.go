package lfu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	ca := Constructor(2)

	ca.Put(1, 1)
	ca.Put(2, 2)
	assert.Equal(t, 1, ca.Get(1))

	ca.Put(3, 3)

	assert.Equal(t, -1, ca.Get(2))
	assert.Equal(t, 3, ca.Get(3))
	ca.Put(4, 4)

	assert.Equal(t, -1, ca.Get(1))
	assert.Equal(t, 3, ca.Get(3))
	assert.Equal(t, 4, ca.Get(4))
}
