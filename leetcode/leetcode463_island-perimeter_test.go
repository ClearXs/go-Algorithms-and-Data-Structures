package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 463. 岛屿的周长 https://leetcode.cn/problems/island-perimeter/description/

// 网格中的格子 水平和垂直 方向相连（对角线方向不相连）。整个网格被水完全包围，但其中恰好有一个岛屿（或者说，一个或多个表示陆地的格子相连组成的岛屿）。
//
// 岛屿中没有“湖”（“湖” 指水域在岛屿内部且不和岛屿周围的水相连）。格子是边长为 1 的正方形。网格为长方形，且宽度和高度均不超过 100 。计算这个岛屿的周长。

func islandPerimeter(grid [][]int) int {
	maxPerimeter := 0
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
				PerimeterGridDFS(grid, r, c, traversal, &temp)
				if temp > maxPerimeter {
					maxPerimeter = temp
				}
			}
		}
	}
	return maxPerimeter
}

func PerimeterGridDFS(grid [][]int, r int, c int, traversal [][]byte, perimeter *int) {
	// 当向上左下右遍历时，如果发生终止，则返回1（如越界、水即无法连接的），否则则可以视为该条边存在连接，则返回0

	// 递归终止条件：越界与zero
	if perimeterOverBoundary(grid, r, c) {
		*perimeter += 1
		return
	}

	if grid[r][c] == 0 {
		*perimeter += 1
		return
	}

	// 不能让它一直兜圈子
	if traversal[r][c] == byte(1) {
		return
	}

	traversal[r][c] = byte(1)

	PerimeterGridDFS(grid, r, c-1, traversal, perimeter) // 左
	PerimeterGridDFS(grid, r, c+1, traversal, perimeter) // 右
	PerimeterGridDFS(grid, r-1, c, traversal, perimeter) // 上
	PerimeterGridDFS(grid, r+1, c, traversal, perimeter) // 下
}

// 判断给定的横纵坐标是否在grid越界
func perimeterOverBoundary(grid [][]int, r int, c int) bool {
	return (r < 0 || c < 0) ||
		r >= len(grid) || c >= len(grid[r])
}

func TestIslandPerimeter(t *testing.T) {
	grid1 := [][]int{{1}}
	perimeter1 := islandPerimeter(grid1)
	assert.Equal(t, 4, perimeter1)

	grid2 := [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}
	perimeter2 := islandPerimeter(grid2)
	assert.Equal(t, 16, perimeter2)
}
