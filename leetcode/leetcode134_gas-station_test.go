package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 134. 加油站  https://leetcode.cn/problems/gas-station/description/

// 在一条环路上有 n 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
//
// 你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。
//
// 给定两个整数数组 gas 和 cost ，如果你可以按顺序绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1 。如果存在解，则 保证 它是 唯一 的。

func canCompleteCircuit(gas []int, cost []int) int {
	curSum, totalSum, startIndex := 0, 0, 0

	for index := range gas {
		curSum += gas[index] - cost[index]
		totalSum += gas[index] - cost[index]
		if curSum < 0 {
			// change to start index
			startIndex = index + 1
			// reset current sum to 0
			curSum = 0
		}
	}
	if totalSum < 0 {
		return -1
	}

	return startIndex
}

func TestCanCompleteCircuit(t *testing.T) {
	r1 := canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2})
	assert.Equal(t, 3, r1)

	r2 := canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3})
	assert.Equal(t, -1, r2)
}
