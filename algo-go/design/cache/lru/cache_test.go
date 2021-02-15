package lru

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

	ca.Put(4, 4)

	assert.Equal(t, -1, ca.Get(1))

	cb := Constructor(1)

	cb.Put(1, 1)

	assert.Equal(t, 1, cb.Get(1))

	cb.Put(2, 2)

	assert.Equal(t, 2, cb.Get(2))
	assert.Equal(t, -1, cb.Get(1))

	assert.Equal(t, -1, cb.Get(3))

	cc := Constructor(0)

	cc.Put(1, 1)

	assert.Equal(t, -1, cc.Get(1))

	cd := Constructor(2)

	cd.Put(2, 1)
	cd.Put(1, 1)
	cd.Put(2, 3)
	cd.Put(4, 1)

	assert.Equal(t, -1, cd.Get(1))
	assert.Equal(t, 3, cd.Get(2))

}
