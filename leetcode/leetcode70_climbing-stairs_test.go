package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 70. 爬楼梯 https://leetcode.cn/problems/climbing-stairs/description/

// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

func climbStairs(n int) int {
	stairs, dp0, dp1 := 0, 0, 1
	for i := 0; i < n; i++ {
		stairs = dp0 + dp1
		dp0 = dp1
		dp1 = stairs
	}
	return stairs
}

func TestClimbStairs(t *testing.T) {
	r1 := climbStairs(2)
	assert.Equal(t, 2, r1)

	r2 := climbStairs(3)
	assert.Equal(t, 3, r2)
}
