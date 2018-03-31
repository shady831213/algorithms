package graph

import (
	"container/list"
)

type Edge struct {
	Start, End interface{}
}

type Graph interface {
	AddVertex(interface{})
	AddEdge(Edge)
	AddEdgeBi(Edge)
	AllVertexes() []interface{}
	AllEdges() []Edge
	GetGraph() (interface{})
}

type AdjacencyMatrix struct {
	Matrix map[interface{}]map[interface{}]bool
	Graph
}

func (g *AdjacencyMatrix) Init() *AdjacencyMatrix {
	g.Matrix = make(map[interface{}]map[interface{}]bool)
	return g
}

func (g *AdjacencyMatrix) AddVertex(vertex interface{}) {
	if _, ok := g.Matrix[vertex]; !ok {
		g.Matrix[vertex] = make(map[interface{}]bool)
	}
}

func (g *AdjacencyMatrix) AddEdge(e Edge) {
	g.AddVertex(e.Start)
	g.AddVertex(e.End)
	g.Matrix[e.Start][e.End] = true
}

func (g *AdjacencyMatrix) AddEdgeBi(e Edge) {
	g.AddVertex(e.Start)
	g.AddVertex(e.End)
	g.Matrix[e.Start][e.End] = true
	g.Matrix[e.End][e.Start] = true
}

func (g *AdjacencyMatrix) AllVertexes() []interface{} {
	keys := make([]interface{}, 0, 0)
	for v := range g.Matrix {
		keys = append(keys, v)
	}
	return keys
}

func (g *AdjacencyMatrix) AllEdges() []Edge {
	edges := make([]Edge, 0, 0)
	for start := range g.Matrix {
		for end := range g.Matrix[start] {
			edges = append(edges, Edge{start, end})
		}
	}
	return edges
}

func (g *AdjacencyMatrix) GetGraph() interface{} {
	return g
}

func NewAdjacencyMatrix() *AdjacencyMatrix {
	return new(AdjacencyMatrix).Init()
}

type AdjacencyList struct {
	List map[interface{}]*list.List
	Graph
}

func (g *AdjacencyList) Init() *AdjacencyList {
	g.List = make(map[interface{}]*list.List)
	return g
}

func (g *AdjacencyList) AddVertex(vertex interface{}) {
	if _, ok := g.List[vertex]; !ok {
		g.List[vertex] = list.New()
	}
}

func (g *AdjacencyList) AddEdge(e Edge) {
	g.AddVertex(e.Start)
	g.AddVertex(e.End)
	g.List[e.Start].PushBack(e.End)
}

func (g *AdjacencyList) AddEdgeBi(e Edge) {
	g.AddVertex(e.Start)
	g.AddVertex(e.End)
	g.List[e.Start].PushBack(e.End)
	g.List[e.End].PushBack(e.Start)
}


func (g *AdjacencyList) AllVertexes() []interface{} {
	keys := make([]interface{}, 0, 0)
	for v := range g.List {
		keys = append(keys, v)
	}
	return keys
}

func (g *AdjacencyList) GetGraph() interface{} {
	return g
}

func (g *AdjacencyList) AllEdges() []Edge {
	edges := make([]Edge, 0, 0)
	for v := range g.List {
		for e := g.List[v].Front(); e != nil; e = e.Next() {
			edges = append(edges, Edge{v, e.Value})
		}
	}
	return edges
}

func NewAdjacencyList() *AdjacencyList {
	return new(AdjacencyList).Init()
}

func AdjacencyList2AdjacencyMatrix(l *AdjacencyList) *AdjacencyMatrix {
	m := NewAdjacencyMatrix()
	for v := range l.List {
		for e := l.List[v].Front(); e != nil; e = e.Next() {
			m.AddEdge(Edge{v, e.Value})
		}
	}
	return m
}

func AdjacencyMatrix2AdjacencyList(m *AdjacencyMatrix) *AdjacencyList {
	l := NewAdjacencyList()
	for start := range m.Matrix {
		for end := range m.Matrix[start] {
			l.AddEdge(Edge{start, end})
		}
	}
	return l
}
