package lc225

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test225(t *testing.T) {
	stack := Constructor()

	stack.Push(1)
	stack.Push(2)

	assert.Equal(t, 2, stack.Top())
	assert.Equal(t, 2, stack.Pop())

	assert.Equal(t, false, stack.Empty())

	stack = Constructor()

	stack.Push(1)
	stack.Push(2)

	assert.Equal(t, 2, stack.Pop())

	assert.Equal(t, 1, stack.Top())

	assert.Equal(t, false, stack.Empty())

}
