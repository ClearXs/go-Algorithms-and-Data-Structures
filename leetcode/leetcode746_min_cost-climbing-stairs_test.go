package leetcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

// 746. 使用最小花费爬楼梯 https://leetcode.cn/problems/min-cost-climbing-stairs/description/

// 给你一个整数数组 cost ，其中 cost[i] 是从楼梯第 i 个台阶向上爬需要支付的费用。一旦你支付此费用，即可选择向上爬一个或者两个台阶。
//
// 你可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯。
//
// 请你计算并返回达到楼梯顶部的最低花费。

func minCostClimbingStairs(cost []int) int {
	dp0, dp1, minimal := 0, 0, 0
	for i := 2; i <= len(cost); i++ {
		minimal = int(math.Min(float64(dp1+cost[i-1]), float64(dp0+cost[i-2])))
		dp0 = dp1
		dp1 = minimal
	}
	return minimal
}

func TestMinCostClimbingStairs(t *testing.T) {

	r1 := minCostClimbingStairs([]int{10, 15, 20})
	assert.Equal(t, 15, r1)

	r2 := minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})
	assert.Equal(t, 6, r2)
}
