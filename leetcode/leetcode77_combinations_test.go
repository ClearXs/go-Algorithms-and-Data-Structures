package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 77. 组合 https://leetcode.cn/problems/combinations/description/
// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
//
// 你可以按 任何顺序 返回答案。

func combine(n int, k int) [][]int {
	// 组合问题转换为多支树遍历求解，其效率非常低，只能通过剪支来优化算法
	result := make([][]int, 0)
	result = backTracking(n, k, result, make([]int, 0), 1)
	return result
}

func backTracking(n int, k int, result [][]int, temp []int, next int) [][]int {
	if len(temp) == k {
		// 为什么是append([]int{}, temp...)
		// 因为go的函数调用传递的是变量的副本
		result = append(result, append([]int{}, temp...))
		return result
	}
	// n-(k-len(temp))+1 剪枝优化
	// 其中 k - len(temp) = 还未组合数量
	// 所以 n-(k-len(temp)) + 1= 还能够组合的最大值，
	// 比如说 k = 4, n = 4，len(temp) = 0 此时只有第一个枝叶能够组合
	// k = 4, n = 4，len(temp) = 1 此时i = 1，故也只有一个枝叶能组合
	for i := next; i <= n-(k-len(temp))+1; i++ {
		temp = append(temp, i)
		result = backTracking(n, k, result, temp, i+1)
		temp = temp[:len(temp)-1]
	}
	return result
}

func TestCombine(t *testing.T) {
	r1 := combine(1, 1)
	assert.Equal(t, [][]int{{1}}, r1)

	r2 := combine(4, 2)
	assert.Equal(t,
		[][]int{
			{1, 2},
			{1, 3},
			{1, 4},
			{2, 3},
			{2, 4},
			{3, 4},
		},
		r2)
}
