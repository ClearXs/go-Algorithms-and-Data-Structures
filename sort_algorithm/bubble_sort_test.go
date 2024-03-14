package sort_algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	assert.Equal(t, EXPECT_N1, bubbleSort(N1))
	assert.Equal(t, EXPECT_N2, bubbleSort(N2))
	assert.Equal(t, EXPECT_N3, bubbleSort(N3))
	assert.Equal(t, EXPECT_N4, bubbleSort(N4))

}

func bubbleSort(nums []int) []int {
	capacity := len(nums)
	if capacity <= 0 {
		return nums
	}
	for i, _ := range nums {
		j := 0
		// 内存循环终止条件为什么是capacity-i-1
		// 1.每次冒泡的结果使得某个元素在数组最后一个 所以-i
		// 2.-1目的保证数组不越位
		for j < capacity-i-1 {
			if nums[j] > nums[j+1] {
				swap := nums[j+1]
				nums[j+1] = nums[j]
				nums[j] = swap
			}
			j++
		}
	}
	return nums
}
