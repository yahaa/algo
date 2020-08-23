package lc1472

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_link(t *testing.T) {
	b := Constructor("leetcode.com")
	b.Visit("google.com")
	b.Visit("facebook.com")
	b.Visit("youtube.com")

	assert.Equal(t, "facebook.com", b.Back(1))

	assert.Equal(t, "google.com", b.Back(1))

	assert.Equal(t, "facebook.com", b.Forward(1))

	b.Visit("linkedin.com")

	assert.Equal(t, "linkedin.com", b.Forward(2))
	assert.Equal(t, "google.com", b.Back(2))
	assert.Equal(t, "leetcode.com", b.Back(7))

}
