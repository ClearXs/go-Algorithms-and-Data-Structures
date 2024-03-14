package graph

import (
	"bufio"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func getGraph() *Graph {
	f, err := os.Open("graph")
	if err != nil {
		return nil
	}
	defer f.Close()
	r := bufio.NewReader(f)
	graph, err := Build(r)
	if err != nil {
		panic(err)
	}
	return graph
}

func TestBuild(t *testing.T) {
	graph := getGraph()
	adj := graph.Adj(0)
	for _, ele := range adj {
		t.Log("vertex: x -> adj", 0, ele)
	}
	assertions := assert.New(t)
	assertions.Equal(6, graph.vertexes)
	assertions.Equal(6, graph.edges)
}

func TestDfs(t *testing.T) {
	graph := getGraph()
	search := graph.DFS(2)
	assert.Equal(t, search.count, 3)
}

func TestPath(t *testing.T) {
	graph := getGraph()
	sp := 1
	search := graph.DFS(sp)
	// 顶点3到起点1的路径
	v := 3
	path := search.PathTo(v)
	// 路径描述
	builder := strings.Builder{}
	builder.WriteString("vertex v = ")
	builder.WriteString(fmt.Sprintf("%d", v))
	builder.WriteString(" path to: ")
	for ele := path.Front(); ele != nil; ele = ele.Next() {
		value := ele.Value
		builder.WriteString(fmt.Sprintf("%d", value))
		builder.WriteString(" -> ")
	}
	t.Log(builder.String())
}

func TestBfs(t *testing.T) {
	graph := getGraph()
	search := graph.BFS(2)

	println(search.marked)
}
