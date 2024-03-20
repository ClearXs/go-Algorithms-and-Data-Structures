package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 376. 摆动序列 https://leetcode.cn/problems/wiggle-subsequence/description/

// 如果连续数字之间的差严格地在正数和负数之间交替，则数字序列称为 摆动序列 。第一个差（如果存在的话）可能是正数或负数。仅有一个元素或者含两个不等元素的序列也视作摆动序列。
//
// 例如， [1, 7, 4, 9, 2, 5] 是一个 摆动序列 ，因为差值 (6, -3, 5, -7, 3) 是正负交替出现的。
//
// 相反，[1, 4, 7, 2, 5] 和 [1, 7, 4, 5, 5] 不是摆动序列，第一个序列是因为它的前两个差值都是正数，第二个序列是因为它的最后一个差值为零。
// 子序列 可以通过从原始序列中删除一些（也可以不删除）元素来获得，剩下的元素保持其原始顺序。
//
// 给你一个整数数组 nums ，返回 nums 中作为 摆动序列 的 最长子序列的长度 。

func wiggleMaxLength(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	// 考虑单调平坡情况
	// 考虑只有两个元素情况，设置初值在数组前加上和数组第一个元素相等元素，此时leftDiff = 0
	leftDiff := 0
	// 如果存在一直单调那么最后就只会剩下两个元素，此时摆动长度为2，故初始1
	wiggleLength := 1
	for i := 1; i < len(nums); i++ {
		rightDiff := nums[i] - nums[i-1]
		// greed algorithm judge part extreme
		if (leftDiff <= 0 && rightDiff > 0) || (leftDiff >= 0 && rightDiff < 0) {
			wiggleLength += 1
			leftDiff = rightDiff
		}
	}
	return wiggleLength
}

func TestWiggleMaxLength(t *testing.T) {
	r1 := wiggleMaxLength([]int{1, 7, 4, 9, 2, 5})
	assert.Equal(t, 6, r1)

	r2 := wiggleMaxLength([]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8})
	assert.Equal(t, 7, r2)

	r3 := wiggleMaxLength([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	assert.Equal(t, 2, r3)
}
