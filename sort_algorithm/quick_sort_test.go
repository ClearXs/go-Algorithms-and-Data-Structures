package sort_algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuickSort(t *testing.T) {
	assert.Equal(t, EXPECT_N1, quickSort(N1))
	assert.Equal(t, EXPECT_N2, quickSort(N2))
	assert.Equal(t, EXPECT_N3, quickSort(N3))
	assert.Equal(t, EXPECT_N4, quickSort(N4))
}

func quickSort(nums []int) []int {
	return decompose(nums, 0, len(nums)-1)
}

func decompose(nums []int, p int, q int) []int {
	if p >= q {
		return nums
	}
	// 需要注意，分区点不参与下次的分区，因为分区是已经确定比该值大或者小
	pivot := partition(nums, p, q)
	decompose(nums, p, pivot-1)
	decompose(nums, pivot+1, q)
	return nums
}

func partition(nums []int, p int, q int) int {
	pivot := nums[q]
	i, j := p, p
	for j < q {
		if nums[j] < pivot {
			Swap(nums, i, j)
			i++
		}
		j++
	}
	// swap for pivot
	Swap(nums, i, q)
	return i
}
