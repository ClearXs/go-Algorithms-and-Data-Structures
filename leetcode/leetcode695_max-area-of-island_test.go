package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 695. 岛屿的最大面积 https://leetcode.cn/problems/max-area-of-island/
// 给你一个大小为 m x n 的二进制矩阵 grid 。
//
// 岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在 水平或者竖直的四个方向上 相邻。你可以假设 grid 的四个边缘都被 0（代表水）包围着。
//
// 岛屿的面积是岛上值为 1 的单元格的数目。
//
// 计算并返回 grid 中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	traversal := make([][]byte, len(grid))

	// 赋予初始值
	for r := range grid {
		traversal[r] = make([]byte, len(grid[r]))
	}

	for r := range grid {
		for c := range grid[r] {
			isIsland := grid[r][c]
			if 1 == isIsland && traversal[r][c] != byte(1) {
				temp := 0
				AreaGridDFS(grid, r, c, traversal, func(r int, c int) {
					temp += grid[r][c]
				})

				if temp > maxArea {
					maxArea = temp
				}
			}
		}
	}
	return maxArea
}

func AreaGridDFS(grid [][]int, r int, c int, traversal [][]byte, accept func(r int, c int)) {
	// 递归终止条件：越界与zero
	if areaOverBoundary(grid, r, c) {
		return
	}
	if grid[r][c] == 0 {
		return
	}

	// 不能让它一直兜圈子
	if traversal[r][c] == byte(1) {
		return
	}

	traversal[r][c] = byte(1)

	if accept != nil {
		accept(r, c)
	}

	AreaGridDFS(grid, r, c-1, traversal, accept) // 上
	AreaGridDFS(grid, r-1, c, traversal, accept) // 左
	AreaGridDFS(grid, r+1, c, traversal, accept) // 右
	AreaGridDFS(grid, r, c+1, traversal, accept) // 下
}

// 判断给定的横纵坐标是否在grid越界
func areaOverBoundary(grid [][]int, r int, c int) bool {
	return (r < 0 || c < 0) ||
		r >= len(grid) || c >= len(grid[r])
}

func TestMaxAreaOfIsland(t *testing.T) {
	grid := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}

	maxArea := maxAreaOfIsland(grid)
	assert.Equal(t, 6, maxArea)
}
