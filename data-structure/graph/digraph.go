package graph

import (
	"bufio"
	"data-structure/queue"
	"io"
	"math"
	"strconv"
	"strings"
)

// 定义图的指向结构
type Point struct {
	// 起点
	v int
	// 终点
	w int
}

type Digraph struct {
	vertexes int     // 顶点数量
	edges    int     // 边数量
	adj      [][]int // 邻接表
}

type DigSearch struct {
	Search
	// 判断有向图中是否有环
	hasCycle bool
}

// addEdge 有向图添加边，其方向为 v -> w
func (dg *Digraph) addEdge(v int, w int) {
	// 扩容
	l := len(dg.adj)
	if l <= v || l <= w {
		expand := int(math.Max(float64(v-l+1), float64(w-l+1)))
		for i := 0; i < expand; i++ {
			dg.adj = append(dg.adj, make([]int, 0))
		}
	}
	dg.adj[v] = append(dg.adj[v], w)
	dg.vertexes = len(dg.adj)
	dg.edges++
}

// Adj 查找某个顶点的邻接矩阵
func (dg *Digraph) Adj(v int) []int {
	return dg.adj[v]
}

// Reverse 当前有向图取反向
func (dg *Digraph) Reverse() *Digraph {
	newDg := Digraph{vertexes: 0, edges: 0, adj: make([][]int, 0)}
	for v := 0; v < dg.vertexes; v++ {
		for _, w := range dg.Adj(v) {
			newDg.addEdge(w, v)
		}
	}
	return &newDg
}

// DFS 有向图的深度优先搜索，逻辑和无向图一致
func (dg *Digraph) DFS(v int) *DigSearch {
	search := DigSearch{
		hasCycle: false,
	}
	search.marked = make([]bool, dg.vertexes)
	search.count = 0
	search.s = v
	search.edgeTo = make([]int, dg.vertexes)
	dg.digraphDFS(&search, v)
	return &search
}

func (dg *Digraph) digraphDFS(search *DigSearch, v int) {
	search.marked[v] = true
	search.count = search.count + 1
	for _, w := range dg.Adj(v) {
		if !search.marked[w] {
			dg.digraphDFS(search, w)
			search.edgeTo[w] = v
		} else if search.marked[w] {
			search.hasCycle = true
		}
	}
}

// BFS 有向图的广度遍历
func (dg *Digraph) BFS(v int) *DigSearch {
	search := DigSearch{
		hasCycle: false,
	}
	search.marked = make([]bool, dg.vertexes)
	search.count = 0
	search.s = v
	search.edgeTo = make([]int, dg.vertexes)
	dg.digraphBFS(&search, v)
	return &search
}

func (dg *Digraph) digraphBFS(search *DigSearch, v int) {
	search.marked[v] = true
	search.count += 1
	// 借助队列数据结构，记录需要遍历的顶点
	traverseVertexes := queue.NewQueue[int]()
	for _, w := range dg.adj[v] {
		if !search.marked[w] {
			search.edgeTo[w] = v
			traverseVertexes.Push(&w)
			// indicate w vertex point to traversed vertex
		} else {
			search.hasCycle = true
		}
	}
	for {
		nextVertex := traverseVertexes.Pop()
		if nextVertex != nil {
			dg.digraphDFS(search, *nextVertex)
		} else {
			break
		}
	}
}

// BuildDigraphByPoint 根据指向结构构建有向图
func BuildDigraphByPoint(g []Point) *Digraph {
	dg := Digraph{vertexes: 0, edges: 0, adj: make([][]int, 0)}
	for i := range g {
		p := g[i]
		dg.addEdge(p.v, p.w)
	}
	return &dg
}

// BuildDigraph 根据文件内容构建有向图
func BuildDigraph(in io.Reader) (*Digraph, error) {
	br := bufio.NewReader(in)
	dg := Digraph{vertexes: 0, edges: 0, adj: make([][]int, 0)}
	for {
		line, err := br.ReadString('\n')
		if err != nil && err != io.EOF {
			return nil, err
		}
		newLine := strings.Replace(line, "\n", "", 1)
		adjPoint := strings.Split(newLine, " ")
		v, _ := strconv.Atoi(adjPoint[0])
		w, _ := strconv.Atoi(adjPoint[1])
		dg.addEdge(v, w)
		if err == io.EOF {
			break
		}
	}
	return &dg, nil
}
