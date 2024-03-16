package leetcode

import (
	"data-structure/graph"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 1971. 寻找图中是否存在路径 https://leetcode.cn/problems/find-if-path-exists-in-graph/description/

// 有一个具有 n 个顶点的 双向 图，其中每个顶点标记从 0 到 n - 1（包含 0 和 n - 1）。图中的边用一个二维整数数组 edges 表示，其中 edges[i] = [ui, vi] 表示顶点 ui 和顶点 vi 之间的双向边。 每个顶点对由 最多一条 边连接，并且没有顶点存在与自身相连的边。
//
// 请你确定是否存在从顶点 source 开始，到顶点 destination 结束的 有效路径 。
//
// 给你数组 edges 和整数 n、source 和 destination，如果从 source 到 destination 存在 有效路径 ，则返回 true，否则返回 false 。

func validPath(n int, edges [][]int, source int, destination int) bool {
	g := graph.BuildGraph(n, edges)
	s := g.DFS(source)
	return s.HasPathTo(destination)
}

// 其他解法1
func o1ValidPath(n int, edges [][]int, source int, destination int) bool {
	disjoint := graph.BuildDisjointSet(n, false)
	for _, edge := range edges {
		v := edge[0]
		w := edge[1]
		disjoint.Union(v, w)
	}
	return disjoint.IsSame(source, destination)
}

func TestValidPath(t *testing.T) {
	i1 := [][]int{{0, 1}, {1, 2}, {2, 0}}
	r1 := validPath(3, i1, 0, 2)
	ro1 := o1ValidPath(3, i1, 0, 2)
	assert.Equal(t, true, r1)
	assert.Equal(t, true, ro1)

	i2 := [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}
	r2 := validPath(6, i2, 0, 5)
	ro2 := o1ValidPath(6, i2, 0, 5)
	assert.Equal(t, false, r2)
	assert.Equal(t, false, ro2)

	i3 := [][]int{
		{4, 3},
		{1, 4},
		{4, 8},
		{1, 7},
		{6, 4},
		{4, 2},
		{7, 4},
		{4, 0},
		{0, 9},
		{5, 4},
	}
	r3 := validPath(10, i3, 5, 9)
	ro3 := o1ValidPath(10, i3, 5, 9)
	assert.Equal(t, true, r3)
	assert.Equal(t, true, ro3)
}
