package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 62. 不同路径 https://leetcode.cn/problems/unique-paths/description/

// 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
//
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
//
// 问总共有多少条不同的路径？

func uniquePaths(m int, n int) int {
	dp := make([][]int, 0)
	// initial
	for i := 0; i < m; i++ {
		dpM := make([]int, 0)
		dpM = append(dpM, 1)
		dp = append(dp, dpM)
	}

	for i := 1; i < n; i++ {
		dpN := dp[0]
		dpN = append(dpN, 1)
		dp[0] = dpN
	}

	// traversal
	for i := 1; i < m; i++ {
		dpM := dp[i]
		for j := 1; j < n; j++ {
			path := dp[i-1][j] + dp[i][j-1]
			dpM = append(dpM, path)
			dp[i] = dpM
		}
	}
	return dp[m-1][n-1]
}

func TestUniquePaths(t *testing.T) {
	a := assert.New(t)
	r1 := uniquePaths(3, 7)
	a.Equal(28, r1)

	r2 := uniquePaths(3, 2)
	a.Equal(3, r2)

	r3 := uniquePaths(7, 3)
	a.Equal(28, r3)

	r4 := uniquePaths(3, 3)
	a.Equal(6, r4)
}
