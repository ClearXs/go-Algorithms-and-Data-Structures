package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 岛屿数量 https://leetcode.cn/problems/number-of-islands/description/
// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

// 这是一道grid的题目，可以看作为是图，适用于图的遍历方法

func numIslands(grid [][]byte) int {
	count := 0
	traversal := make([][]byte, len(grid))

	// 赋予初始值
	for r := range grid {
		traversal[r] = make([]byte, len(grid[r]))
	}

	for r := range grid {
		for c := range grid[r] {
			isIsland := grid[r][c]
			if byte(1) == isIsland && traversal[r][c] != byte(1) {
				GridDFS(grid, r, c, traversal, nil)
				count += 1
			}
		}
	}
	return count
}

func GridDFS(grid [][]byte, r int, c int, traversal [][]byte, accept func(r int, c int)) {
	// 递归终止条件：越界与zero
	if overBoundary(grid, r, c) {
		return
	}
	if grid[r][c] == byte(0) {
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

	GridDFS(grid, r, c-1, traversal, accept) // 上
	GridDFS(grid, r-1, c, traversal, accept) // 左
	GridDFS(grid, r+1, c, traversal, accept) // 右
	GridDFS(grid, r, c+1, traversal, accept) // 下
}

// 判断给定的横纵坐标是否在grid越界
func overBoundary(grid [][]byte, r int, c int) bool {
	return (r < 0 || c < 0) ||
		r >= len(grid) || c >= len(grid[r])
}

func TestNumIslands(t *testing.T) {

	grid1 := makeGrid([][]string{
		{"1", "1", "1", "1", "0"},
		{"1", "1", "0", "1", "0"},
		{"1", "1", "0", "0", "0"},
		{"0", "0", "0", "0", "0"},
	})

	islands1 := numIslands(grid1)
	assert.Equal(t, 1, islands1)

	grid2 := makeGrid([][]string{
		{"1", "1", "0", "0", "0"},
		{"1", "1", "0", "0", "0"},
		{"0", "0", "1", "0", "0"},
		{"0", "0", "0", "1", "1"},
	})
	islands2 := numIslands(grid2)
	assert.Equal(t, 3, islands2)

	grid3 := makeGrid([][]string{{"1", "1", "1", "1", "0"}, {"1", "1", "0", "1", "0"}, {"1", "1", "0", "0", "0"}, {"0", "0", "0", "0", "0"}})
	islands3 := numIslands(grid3)
	assert.Equal(t, 1, islands3)
}

func makeGrid(input [][]string) [][]byte {
	grid := make([][]byte, len(input))
	for r := range input {
		if grid[r] == nil {
			grid[r] = make([]byte, len(input[r]))
		}

		for c := range input[r] {
			val := input[r][c]

			if val == "1" {
				grid[r][c] = byte(1)
			} else if val == "0" {
				grid[r][c] = byte(0)
			}
		}
	}
	return grid
}
