package sort3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	assert.Equal(t, EXPECT_N1, SelectionSort(N1))
	assert.Equal(t, EXPECT_N2, SelectionSort(N2))
	assert.Equal(t, EXPECT_N3, SelectionSort(N3))
	assert.Equal(t, EXPECT_N4, SelectionSort(N4))
}
