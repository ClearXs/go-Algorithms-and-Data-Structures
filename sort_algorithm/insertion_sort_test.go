package sort_algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	assert.Equal(t, EXPECT_N1, insertionSort(N1))
	assert.Equal(t, EXPECT_N2, insertionSort(N2))
	assert.Equal(t, EXPECT_N3, insertionSort(N3))
	assert.Equal(t, EXPECT_N4, insertionSort(N4))
}

func insertionSort(nums []int) []int {
	capacity := len(nums)
	if capacity <= 1 {
		return nums
	}
	for i, _ := range nums {
		// 选择尾插入法，内循环初始化条件需要以j=i，如果是头插法，则j=0（这种方法需要包含移动的第三层循环）
		j := i
		for j > 0 {
			// 与已经排好序的数组进行比较
			if nums[j] < nums[j-1] {
				// 移动元素
				swap := nums[j]
				nums[j] = nums[j-1]
				nums[j-1] = swap
			}
			j--
		}
	}
	return nums
}
