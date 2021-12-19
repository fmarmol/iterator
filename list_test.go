package iterator

import (
	"testing"

	"gotest.tools/assert"
)

func TestList(t *testing.T) {
	l := NewList(1, 2, 3, 4)

	count := 0
	for range l.Iter() {
		count++
	}
	assert.Equal(t, 4, count)
}
