package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

// 53. 最大子数组和 https://leetcode.cn/problems/maximum-subarray/description/

// 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
// 子数组
// 是数组中的一个连续部分。

func maxSubArray(nums []int) int {
	result := math.MinInt
	count := 0
	for _, n := range nums {
		count += n
		if count > result {
			result = count
		}
		if count < 0 {
			count = 0
		}
	}

	return result
}

func TestMaxSubArray(t *testing.T) {
	r1 := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	assert.Equal(t, 6, r1)

	r2 := maxSubArray([]int{1})
	assert.Equal(t, 1, r2)

	r3 := maxSubArray([]int{5, 4, -1, 7, 8})
	assert.Equal(t, 23, r3)
}
