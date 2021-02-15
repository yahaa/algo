package lt760

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashmap(t *testing.T) {
	mp := Constructor()
	mp.Put(1, 1)
	mp.Put(2, 2)

	assert.Equal(t, 1, mp.Get(1))
	assert.Equal(t, -1, mp.Get(3))

	mp.Put(2, 1)

	assert.Equal(t, 1, mp.Get(2))
	mp.Remove(2)

	assert.Equal(t, -1, mp.Get(2))

}
