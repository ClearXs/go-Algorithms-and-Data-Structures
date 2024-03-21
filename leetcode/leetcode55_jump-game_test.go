package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

// 55. 跳跃游戏 https://leetcode.cn/problems/jump-game/description/

// 给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
// 判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。

func canJump(nums []int) bool {
	if len(nums) <= 0 {
		return true
	}
	cover := 0
	for i := 0; i <= cover; i++ {
		// why i + nums[i], because first nums[i] is jump to range, and now
		// calculate current index add nums[i] for be able to cover distance in current cover
		// seek maximum
		cover = int(math.Max(float64(i+nums[i]), float64(cover)))
		if cover >= len(nums)-1 {
			return true
		}
	}
	return false
}

func TestCanJump(t *testing.T) {
	r1 := canJump([]int{2, 3, 1, 1, 4})
	assert.Equal(t, true, r1)

	r2 := canJump([]int{3, 2, 1, 0, 4})
	assert.Equal(t, false, r2)
}
