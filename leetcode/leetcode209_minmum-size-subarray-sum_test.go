package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

// 209. 长度最小的子数组 https://leetcode.cn/problems/minimum-size-subarray-sum/

// 给定一个含有 n 个正整数的数组和一个正整数 target 。
//
// 找出该数组中满足其总和大于等于 target 的长度最小的 连续
// 子数组
// [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。

func minSubArrayLen(target int, nums []int) int {
	leftIndex, rightIndex, miniLength, count := 0, 0, math.MaxInt, 0
	for _, num := range nums {
		count += num
		if count >= target {
			for {
				if count < target {
					break
				}
				length := rightIndex - leftIndex + 1
				if miniLength > length {
					miniLength = length
				}
				count -= nums[leftIndex]
				if leftIndex == rightIndex {
					break
				}
				leftIndex += 1
			}
		}
		rightIndex += 1
	}
	if miniLength == math.MaxInt {
		return 0
	}
	return miniLength
}

func TestMinSubArrayLen(t *testing.T) {
	r1 := minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})
	assert.Equal(t, 2, r1)

	r2 := minSubArrayLen(4, []int{1, 4, 4})
	assert.Equal(t, 1, r2)

	r3 := minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1})
	assert.Equal(t, 0, r3)

	r4 := minSubArrayLen(11, []int{1, 2, 3, 4, 5})
	assert.Equal(t, 3, r4)
}
