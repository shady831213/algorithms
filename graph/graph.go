package graph

import (
	"container/list"
)

const (
	white = 0
	gray  = 1
	black = 2
)

type edge struct {
	Start, End interface{}
}

type graph interface {
	AddVertex(interface{})
	AddEdge(edge)
	AddEdgeBi(edge)
	AllVertices() []interface{}
	AllEdges() []edge
	AllConnectedVertices(interface{}) []interface{}
	IterConnectedVertices(interface{}) iterator
	Transpose() graph
	GetGraph() interface{}
}

type chanWithStatus struct {
	ch    chan interface{}
	close bool
}

type adjacencyMatrix struct {
	matrix linkedMap
	chs    map[interface{}]*chanWithStatus
	graph
}

func (g *adjacencyMatrix) init() *adjacencyMatrix {
	g.matrix = *new(linkedMap).init()
	g.chs = make(map[interface{}]*chanWithStatus)
	return g
}

func (g *adjacencyMatrix) AddVertex(vertex interface{}) {
	if !g.matrix.exist(vertex) {
		g.matrix.add(vertex, new(linkedMap).init())
	}
}

func (g *adjacencyMatrix) AddEdge(e edge) {
	g.AddVertex(e.Start)
	g.AddVertex(e.End)
	g.matrix.get(e.Start).(*linkedMap).add(e.End, true)
}

func (g *adjacencyMatrix) AddEdgeBi(e edge) {
	g.AddVertex(e.Start)
	g.AddVertex(e.End)
	g.matrix.get(e.Start).(*linkedMap).add(e.End, true)
	g.matrix.get(e.End).(*linkedMap).add(e.Start, true)
}

func (g *adjacencyMatrix) AllVertices() []interface{} {
	keys := make([]interface{}, 0, 0)
	for v := g.matrix.frontKey(); v != nil; v = g.matrix.nextKey(v) {
		keys = append(keys, v)
	}
	return keys
}

func (g *adjacencyMatrix) AllEdges() []edge {
	edges := make([]edge, 0, 0)
	for start := g.matrix.frontKey(); start != nil; start = g.matrix.nextKey(start) {
		for end := g.matrix.get(start).(*linkedMap).frontKey(); end != nil; end = g.matrix.get(start).(*linkedMap).nextKey(end) {
			edges = append(edges, edge{start, end})
		}
	}
	return edges
}

func (g *adjacencyMatrix) AllConnectedVertices(v interface{}) []interface{} {
	keys := make([]interface{}, 0, 0)
	if g.matrix.exist(v) {
		for key := g.matrix.get(v).(*linkedMap).frontKey(); key != nil; key = g.matrix.get(v).(*linkedMap).nextKey(key) {
			keys = append(keys, key)
		}
	}
	return keys
}

func (g *adjacencyMatrix) IterConnectedVertices(v interface{}) iterator {
	if g.matrix.exist(v) {
		return newLinkedMapIterator(g.matrix.get(v).(*linkedMap))
	}
	return nil
}

func (g *adjacencyMatrix) Transpose() graph {
	gt := newAdjacencyMatrix()
	for _, e := range g.AllEdges() {
		gt.AddEdge(edge{e.End, e.Start})
	}
	return gt
}

func (g *adjacencyMatrix) GetGraph() interface{} {
	return g
}

func newAdjacencyMatrix() *adjacencyMatrix {
	return new(adjacencyMatrix).init()
}

type adjacencyList struct {
	matrix linkedMap
	chs    map[interface{}]*chanWithStatus
	graph
}

func (g *adjacencyList) init() *adjacencyList {
	g.matrix = *new(linkedMap).init()
	g.chs = make(map[interface{}]*chanWithStatus)
	return g
}

func (g *adjacencyList) AddVertex(vertex interface{}) {
	if !g.matrix.exist(vertex) {
		g.matrix.add(vertex, list.New())
	}
}

func (g *adjacencyList) AddEdge(e edge) {
	g.AddVertex(e.Start)
	g.AddVertex(e.End)
	for le := g.matrix.get(e.Start).(*list.List).Front(); le != nil; le = le.Next() {
		if le.Value == e.End {
			return
		}
	}
	g.matrix.get(e.Start).(*list.List).PushBack(e.End)
}

func (g *adjacencyList) AddEdgeBi(e edge) {
	g.AddEdge(e)
	g.AddEdge(edge{e.End, e.Start})
}

func (g *adjacencyList) AllVertices() []interface{} {
	keys := make([]interface{}, 0, 0)
	for v := g.matrix.frontKey(); v != nil; v = g.matrix.nextKey(v) {
		keys = append(keys, v)
	}
	return keys
}

func (g *adjacencyList) AllConnectedVertices(v interface{}) []interface{} {
	value := make([]interface{}, 0, 0)
	if g.matrix.exist(v) {
		for e := g.matrix.get(v).(*list.List).Front(); e != nil; e = e.Next() {
			value = append(value, e.Value)
		}
	}
	return value
}

func (g *adjacencyList) IterConnectedVertices(v interface{}) iterator {
	if g.matrix.exist(v) {
		return newListIterator(g.matrix.get(v).(*list.List))
	}
	return nil
}

func (g *adjacencyList) AllEdges() []edge {
	edges := make([]edge, 0, 0)
	for v := g.matrix.frontKey(); v != nil; v = g.matrix.nextKey(v) {
		for e := g.matrix.get(v).(*list.List).Front(); e != nil; e = e.Next() {
			edges = append(edges, edge{v, e.Value})
		}
	}
	return edges
}

func (g *adjacencyList) Transpose() graph {
	gt := newAdjacencyList()
	for _, e := range g.AllEdges() {
		gt.AddEdge(edge{e.End, e.Start})
	}
	return gt
}

func (g *adjacencyList) GetGraph() interface{} {
	return g
}

func newAdjacencyList() *adjacencyList {
	return new(adjacencyList).init()
}

func adjacencyList2AdjacencyMatrix(l *adjacencyList) *adjacencyMatrix {
	m := newAdjacencyMatrix()
	for v := l.matrix.frontKey(); v != nil; v = l.matrix.nextKey(v) {
		for e := l.matrix.get(v).(*list.List).Front(); e != nil; e = e.Next() {
			m.AddEdge(edge{v, e.Value})
		}
	}
	return m
}

func adjacencyMatrix2AdjacencyList(m *adjacencyMatrix) *adjacencyList {
	l := newAdjacencyList()
	for start := m.matrix.frontKey(); start != nil; start = m.matrix.nextKey(start) {
		for end := m.matrix.get(start).(*linkedMap).frontKey(); end != nil; end = m.matrix.get(start).(*linkedMap).nextKey(end) {
			l.AddEdge(edge{start, end})
		}
	}
	return l
}

func createGraphByType(g graph) (newG graph) {
	if _, isList := g.GetGraph().(*adjacencyList); isList {
		newG = newAdjacencyList()
	} else {
		newG = newAdjacencyMatrix()
	}
	return
}
