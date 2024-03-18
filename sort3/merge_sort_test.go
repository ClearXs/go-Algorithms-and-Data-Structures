package sort3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeSort(t *testing.T) {
	// 写了很久，主要调整参数、调整逻辑，如子数组大小
	assert.Equal(t, EXPECT_N1, MergeSort(N1))
	assert.Equal(t, EXPECT_N2, MergeSort(N2))
	assert.Equal(t, EXPECT_N3, MergeSort(N3))
	assert.Equal(t, EXPECT_N4, MergeSort(N4))
}
