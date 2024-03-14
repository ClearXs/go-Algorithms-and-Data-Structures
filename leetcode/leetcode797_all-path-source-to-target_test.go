package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAllPathsSourceTarget(t *testing.T) {
	adj1 := [][]int{{1, 2}, {3}, {3}, {}}
	paths1 := allPathsSourceTarget(adj1)
	assert.Equal(t, [][]int{{0, 1, 3}, {0, 2, 3}}, paths1)

	adj2 := [][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}
	paths2 := allPathsSourceTarget(adj2)
	assert.Equal(t, [][]int{{0, 4}, {0, 3, 4}, {0, 1, 3, 4}, {0, 1, 2, 3, 4}, {0, 1, 4}}, paths2)

	adj3 := [][]int{{3, 1}, {4, 6, 7, 2, 5}, {4, 6, 3}, {6, 4}, {7, 6, 5}, {6}, {7}, {}}

	paths3 := allPathsSourceTarget(adj3)
	println(paths3)

}

func allPathsSourceTarget(graph [][]int) [][]int {
	paths := make([][]int, 0)
	path := make([]int, 0)
	return traverse(graph, 0, paths, path)
}

// traverse 基于dfs 该方法能够从给定顶点v寻找该点的所有到达终点的数据
func traverse(adj [][]int, v int, paths [][]int, path []int) [][]int {
	path = append(path, v)
	gLength := len(adj)
	if v == gLength-1 {
		paths = append(paths, append([]int{}, path...))
	}

	for _, w := range adj[v] {
		paths = traverse(adj, w, paths, path)
	}

	return paths
}
