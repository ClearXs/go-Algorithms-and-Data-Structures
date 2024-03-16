package leetcode

import (
	"data-structure/graph"
	"github.com/stretchr/testify/assert"
	"testing"
)

func findRedundantConnection(edges [][]int) []int {
	disjoint := graph.BuildDisjointSet(len(edges), true)
	for _, edge := range edges {
		v := edge[0]
		w := edge[1]
		if disjoint.IsSame(v, w) {
			return edge
		}
		disjoint.Union(v, w)
	}
	return []int{}
}

func TestFindRedundantConnection(t *testing.T) {
	i1 := [][]int{{1, 2}, {1, 3}, {2, 3}}
	r1 := findRedundantConnection(i1)
	assert.Equal(t, []int{2, 3}, r1)

	i2 := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}
	r2 := findRedundantConnection(i2)
	assert.Equal(t, []int{1, 4}, r2)

	i3 := [][]int{{1, 4}, {3, 4}, {1, 3}, {1, 2}, {4, 5}}
	r3 := findRedundantConnection(i3)
	assert.Equal(t, []int{1, 3}, r3)
}
