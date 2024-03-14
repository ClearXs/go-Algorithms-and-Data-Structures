package sort_algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	assert.Equal(t, EXPECT_N1, selectionSort(N1))
	assert.Equal(t, EXPECT_N2, selectionSort(N2))
	assert.Equal(t, EXPECT_N3, selectionSort(N3))
	assert.Equal(t, EXPECT_N4, selectionSort(N4))
}

func selectionSort(nums []int) []int {
	capacity := len(nums)
	if capacity <= 1 {
		return nums
	}

	for i, _ := range nums {
		minIndex := i
		innerIndex := i
		for innerIndex < capacity {
			if nums[innerIndex] < nums[minIndex] {
				minIndex = innerIndex
			}
			innerIndex++
		}
		// swap
		swap := nums[i]
		nums[i] = nums[minIndex]
		nums[minIndex] = swap
	}
	return nums
}
