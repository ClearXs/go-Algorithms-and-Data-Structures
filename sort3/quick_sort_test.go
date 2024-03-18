package sort3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuickSort(t *testing.T) {
	assert.Equal(t, EXPECT_N1, QuickSort(N1))
	assert.Equal(t, EXPECT_N2, QuickSort(N2))
	assert.Equal(t, EXPECT_N3, QuickSort(N3))
	assert.Equal(t, EXPECT_N4, QuickSort(N4))
}
