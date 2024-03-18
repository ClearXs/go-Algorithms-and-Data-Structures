package sort3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	assert.Equal(t, EXPECT_N1, InsertionSort(N1))
	assert.Equal(t, EXPECT_N2, InsertionSort(N2))
	assert.Equal(t, EXPECT_N3, InsertionSort(N3))
	assert.Equal(t, EXPECT_N4, InsertionSort(N4))
}
