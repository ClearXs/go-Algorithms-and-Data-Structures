package graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDisjoint(t *testing.T) {
	d1 := BuildDisjointSet(3, false)

	// 1 -> 0
	d1.Union(1, 0)
	assert.Equal(t, 1, d1.father[0])

	// 1 -> 2
	d1.Union(1, 2)
	assert.Equal(t, 1, d1.father[2])

	assert.Equal(t, true, d1.IsSame(0, 1))
	assert.Equal(t, true, d1.IsSame(1, 2))

	assert.Equal(t, true, d1.IsSame(2, 0))

	d2 := BuildDisjointSet(3, false)
	d2.Union(0, 1)
	assert.Equal(t, 0, d2.father[1])
	d2.Union(1, 2)
	assert.Equal(t, 0, d2.father[2])

}
