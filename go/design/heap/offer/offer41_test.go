package offer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_offer(t *testing.T) {
	fd := Constructor()

	fd.AddNum(2)
	assert.Equal(t, float64(2), fd.FindMedian())
	fd.AddNum(3)
	fd.AddNum(4)
	assert.Equal(t, float64(3), fd.FindMedian())

	fd.AddNum(5)
	fd.AddNum(6)
	fd.AddNum(7)
	fd.AddNum(8)

}
