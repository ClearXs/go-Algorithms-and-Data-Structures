// 两数之和，采用双指针
package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSun(t *testing.T) {
	assert.Equal(t, []int{1, 2}, twoSum([]int{2, 7, 11, 15}, 9))
	assert.Equal(t, []int{1, 3}, twoSum([]int{2, 3, 4}, 6))
	assert.Equal(t, []int{1, 2}, twoSum([]int{-1, 0}, -1))
}

func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else if sum > target {
			right--
		}
	}
	return []int{}
}
