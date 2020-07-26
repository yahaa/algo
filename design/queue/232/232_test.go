package lc232

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test232(t *testing.T) {
	q := Constructor()

	q.Push(1)
	q.Push(2)

	assert.Equal(t, 1, q.Peek())

	assert.Equal(t, 1, q.Pop())

	assert.Equal(t, false, q.Empty())

}
