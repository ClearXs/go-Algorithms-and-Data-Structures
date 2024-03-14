package graph

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestBuildDigraph(t *testing.T) {
	dg := getDigraph()
	assert.Equal(t, 13, dg.vertexes)
	assert.Equal(t, 22, dg.edges)
	assert.Equal(t, []int{3, 0}, dg.Adj(2))
}

func TestReverse(t *testing.T) {
	dg := getDigraph()
	newDg := dg.Reverse()
	assert.Equal(t, 13, newDg.vertexes)
}

func TestBuildDigraphByMap(t *testing.T) {
	p := []Point{
		{v: 0, w: 1},
		{v: 1, w: 2},
		{v: 1, w: 3},
	}
	dg := BuildDigraphByPoint(p)
	s := dg.DFS(0)
	assert.Equal(t, 4, s.count)
	assert.Equal(t, 0, s.edgeTo[1])
}

func TestDFSCycleDetection(t *testing.T) {
	p := []Point{
		{v: 0, w: 5},
		{v: 5, w: 4},
		{v: 4, w: 3},
		{v: 3, w: 5},
	}
	dg := BuildDigraphByPoint(p)
	s := dg.DFS(0)
	assert.Equal(t, true, s.hasCycle)
}

func TestDFS(t *testing.T) {
	p := []Point{
		{v: 0, w: 5},
		{v: 5, w: 4},
		{v: 4, w: 3},
	}
	dg := BuildDigraphByPoint(p)
	bfs := dg.BFS(0)
	edgeTo := bfs.edgeTo
	assert.Equal(t, []int{0, 0, 0, 4, 5, 0}, edgeTo)
}

func TestBFS(t *testing.T) {
	p := []Point{
		{v: 0, w: 5},
		{v: 5, w: 4},
		{v: 4, w: 3},
	}
	dg := BuildDigraphByPoint(p)
	bfs := dg.BFS(0)
	edgeTo := bfs.edgeTo
	assert.Equal(t, []int{0, 0, 0, 4, 5, 0}, edgeTo)
}

func TestBFSCycleDetection(t *testing.T) {
	p := []Point{
		{v: 0, w: 5},
		{v: 5, w: 4},
		{v: 4, w: 3},
		{v: 3, w: 5},
	}
	dg := BuildDigraphByPoint(p)
	s := dg.BFS(0)
	assert.Equal(t, true, s.hasCycle)
}

func getDigraph() *Digraph {
	f, err := os.Open("digraph")
	if err != nil {
		return nil
	}
	defer f.Close()
	r := bufio.NewReader(f)
	graph, err := BuildDigraph(r)
	if err != nil {
		return nil
	}
	return graph
}
