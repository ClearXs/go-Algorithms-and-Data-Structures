package graph

import (
	"bufio"
	"container/list"
	"io"
	"strconv"
	"strings"
)

const Empty = 0

type Graph struct {
	// 边数量
	edges int
	// 顶点数量
	vertexes int
	// 邻接表
	adj [][]int
}

type Search struct {
	// 标记给定起点开始
	marked []bool
	// 记录搜索到的顶点数量
	count int
	// 起点
	s int
	// 记录从某个顶点被哪些顶点连接
	edgeTo []int
}

// Graph 中添加一条边
func (g *Graph) addEdge(v, w int) {
	g.ifNilElseCreateNewVertex(v)
	g.ifNilElseCreateNewVertex(w)
	g.adj[v] = append(g.adj[v], w)
	g.adj[w] = append(g.adj[w], v)
	g.edges++
}

func (g *Graph) ifNilElseCreateNewVertex(v int) {
	if g.adj[v] == nil {
		// 创建新的顶点
		g.adj[v] = make([]int, 0)
	}
}

// Graph 中添加一条边
func (g *Graph) addStringEdge(v, w string) {
	vInt, _ := strconv.Atoi(v)
	wInt, _ := strconv.Atoi(w)
	g.addEdge(vInt, wInt)
}

// GetVertexes 获取图中顶点的数量
func (g *Graph) GetVertexes() int {
	return g.vertexes
}

// GetEdges 获取边的数量
func (g *Graph) GetEdges() int {
	return g.edges
}

// Adj 获取指定顶点的邻接表
func (g *Graph) Adj(v int) []int {
	return g.adj[v]
}

// DFS 深度优先遍历（搜索所有顶点和边）
// v 给定的起点
func (g *Graph) DFS(v int) Search {
	search := g.makeSearch(v)
	return g.dfsSearch(search, v)
}

func (g *Graph) dfsSearch(search Search, v int) Search {
	search.marked[v] = true
	search.count++
	for _, w := range g.Adj(v) {
		// 没有被搜索到
		if !search.marked[w] {
			// 记录当前顶点到起点的路径
			search.edgeTo[w] = v
			g.dfsSearch(search, w)
		}
	}
	return search
}

// BFS 广度优先搜索，找寻最短路径
// v 起点值
func (g *Graph) BFS(v int) Search {
	search := g.makeSearch(v)
	// 记录被标记，但还未检查领接表的队列
	recordMarkedAdjUncheckedQueue := list.New()
	recordMarkedAdjUncheckedQueue.PushBack(v)
	search.marked[v] = true
	for recordMarkedAdjUncheckedQueue.Len() != 0 {
		removeOfCheckedVertexEle := recordMarkedAdjUncheckedQueue.Front()
		recordMarkedAdjUncheckedQueue.Remove(removeOfCheckedVertexEle)
		checkedVertex := removeOfCheckedVertexEle.Value.(int)
		for _, ele := range g.Adj(checkedVertex) {
			if !search.marked[ele] {
				search.edgeTo[ele] = checkedVertex
				search.marked[ele] = true
				recordMarkedAdjUncheckedQueue.PushBack(ele)
			}
		}
	}
	return search
}

// makeSearch 创建一个Search实例
func (g *Graph) makeSearch(v int) Search {
	return Search{
		marked: make([]bool, g.GetVertexes()),
		count:  0,
		s:      v,
		edgeTo: make([]int, g.GetVertexes()),
	}
}

// HasPathTo 判断给定的起点是否有其路径
func (search *Search) HasPathTo(v int) bool {
	return search.marked[v]
}

// PathTo 获取给定起点的路径
func (search *Search) PathTo(v int) *list.List {
	if !search.HasPathTo(v) {
		return nil
	}
	path := list.New()
	for w := v; w != search.s; w = search.edgeTo[w] {
		if w == Empty {
			break
		}
		path.PushBack(w)
	}
	path.PushBack(search.s)
	return path
}

// Build 根据输入流，获取图的实例
// 输入流的格式应如下：
// 0 5
// 1 2
// 2 3
// 3 1
func Build(r io.Reader) (*Graph, error) {
	br := bufio.NewReader(r)
	edges := make([][]int, 0)
	for {
		line, err := br.ReadString('\n')
		if err != nil && err != io.EOF {
			return nil, err
		}
		// End Of and line 为空
		if line == "" && err == io.EOF {
			break
		}
		newLine := strings.Replace(line, "\n", "", 1)
		adjPoint := strings.Split(newLine, " ")
		v, err := strconv.Atoi(adjPoint[0])
		if err != nil {
			return nil, err
		}
		w, err := strconv.Atoi(adjPoint[1])
		if err != nil {
			return nil, err
		}
		edges = append(edges, []int{v, w})
	}
	g := BuildByEdges(edges)
	return g, nil
}

// BuildByEdges from edges two dimension for int type build graph
func BuildByEdges(edges [][]int) *Graph {
	return BuildGraph(len(edges), edges)
}

// BuildGraph Based on the number of vertex and relationship between vertices v and w using a two-dimensional int type construct graph
func BuildGraph(vertex int, edges [][]int) *Graph {
	g := Graph{
		edges:    0,
		vertexes: vertex,
		adj:      make([][]int, vertex),
	}
	for _, edge := range edges {
		v := edge[0]
		w := edge[1]
		g.addEdge(v, w)
	}
	return &g
}
