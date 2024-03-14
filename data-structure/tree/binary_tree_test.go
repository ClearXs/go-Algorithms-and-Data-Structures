package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildBinaryTree(t *testing.T) {
	n1 := []int{3, 9, 20, 0, 0, 15, 7}
	t1 := BuildBinaryTree(n1)
	assert.Equal(t, 3, t1.Val)
	n2 := []int{0}
	t2 := BuildBinaryTree(n2)
	assert.Empty(t, t2)

	n3 := []int{1}
	t3 := BuildBinaryTree(n3)
	assert.Equal(t, 1, t3.Val)
}
