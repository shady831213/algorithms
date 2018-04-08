package graph

import (
	"container/list"
)

const (
	WHITE = 0
	GRAY  = 1
	BLACK = 2
)

type Edge struct {
	Start, End interface{}
}

type Graph interface {
	AddVertex(interface{})
	AddEdge(Edge)
	AddEdgeBi(Edge)
	AllVertices() []interface{}
	AllEdges() []Edge
	AllConnectedVertices(interface{}) []interface{}
	Transpose()Graph
	GetGraph() (interface{})
}

type linkedMap struct {
	keyL *list.List
	m    map[interface{}]interface{}
}

func (m *linkedMap) init() *linkedMap {
	m.keyL = list.New()
	m.m = make(map[interface{}]interface{})
	return m
}

func (m *linkedMap) exist(key interface{}) bool {
	_, ok := m.m[key]
	return ok
}

func (m *linkedMap) add(key, value interface{}) {
	if !m.exist(key) {
		e := m.keyL.PushBack(key)
		m.m[key] = []interface{}{e, value}
	} else {
		m.m[key].([]interface{})[1] = value
	}
}

func (m *linkedMap) get(key interface{}) interface{} {
	if m.exist(key) {
		return m.m[key].([]interface{})[1]
	} else {
		return nil
	}
}

func (m *linkedMap) delete(key interface{}) {
	if m.exist(key) {
		i := m.m[key].([]interface{})
		m.keyL.Remove(i[0].(*list.Element))
		delete(m.m, key)
	}
}

func (m *linkedMap) frontKey() interface{} {
	if m.keyL.Len() == 0 {
		return nil
	}
	return m.keyL.Front().Value
}

func (m *linkedMap) backKey() interface{} {
	if m.keyL.Len() == 0 {
		return nil
	}
	return m.keyL.Back().Value
}

func (m *linkedMap) nextKey(key interface{}) interface{} {
	if e := m.m[key].([]interface{})[0].(*list.Element).Next(); e != nil {
		return e.Value
	}
	return nil
}

type AdjacencyMatrix struct {
	matrix linkedMap
	Graph
}

func (g *AdjacencyMatrix) Init() *AdjacencyMatrix {
	g.matrix = *new(linkedMap).init()

	return g
}

func (g *AdjacencyMatrix) AddVertex(vertex interface{}) {
	if !g.matrix.exist(vertex) {
		g.matrix.add(vertex,new(linkedMap).init())
	}
}

func (g *AdjacencyMatrix) AddEdge(e Edge) {
	g.AddVertex(e.Start)
	g.AddVertex(e.End)
	g.matrix.get(e.Start).(*linkedMap).add(e.End, true)
}

func (g *AdjacencyMatrix) AddEdgeBi(e Edge) {
	g.AddVertex(e.Start)
	g.AddVertex(e.End)
	g.matrix.get(e.Start).(*linkedMap).add(e.End, true)
	g.matrix.get(e.End).(*linkedMap).add(e.Start, true)
}

func (g *AdjacencyMatrix) AllVertices() []interface{} {
	keys := make([]interface{}, 0, 0)
	for v := g.matrix.frontKey(); v != nil; v = g.matrix.nextKey(v) {
		keys = append(keys, v)
	}
	return keys
}

func (g *AdjacencyMatrix) AllEdges() []Edge {
	edges := make([]Edge, 0, 0)
	for start := g.matrix.frontKey(); start != nil; start = g.matrix.nextKey(start) {
		for end := g.matrix.get(start).(*linkedMap).frontKey(); end != nil; end = g.matrix.get(start).(*linkedMap).nextKey(end) {
			edges = append(edges, Edge{start, end})
		}
	}
	return edges
}

func (g *AdjacencyMatrix) AllConnectedVertices(v interface{}) []interface{} {
	keys := make([]interface{}, 0, 0)
	if g.matrix.exist(v){
		for key := g.matrix.get(v).(*linkedMap).frontKey(); key != nil; key = g.matrix.get(v).(*linkedMap).nextKey(key) {
			keys = append(keys, key)
		}
	}
	return keys
}

func (g *AdjacencyMatrix) Transpose() Graph {
	gt := NewAdjacencyMatrix()
	for _, e := range g.AllEdges() {
		gt.AddEdge(Edge{e.End,e.Start})
	}
	return gt
}

func (g *AdjacencyMatrix) GetGraph() interface{} {
	return g
}

func NewAdjacencyMatrix() *AdjacencyMatrix {
	return new(AdjacencyMatrix).Init()
}

type AdjacencyList struct {
	matrix linkedMap
	Graph
}

func (g *AdjacencyList) Init() *AdjacencyList {
	g.matrix = *new(linkedMap).init()
	return g
}

func (g *AdjacencyList) AddVertex(vertex interface{}) {
	if !g.matrix.exist(vertex){
		g.matrix.add(vertex,list.New())
	}
}

func (g *AdjacencyList) AddEdge(e Edge) {
	g.AddVertex(e.Start)
	g.AddVertex(e.End)
	for le := g.matrix.get(e.Start).(*list.List).Front(); le != nil; le = le.Next() {
		if le.Value == e.End {
			return
		}
	}
	g.matrix.get(e.Start).(*list.List).PushBack(e.End)
}

func (g *AdjacencyList) AddEdgeBi(e Edge) {
	g.AddEdge(e)
	g.AddEdge(Edge{e.End,e.Start})
}

func (g *AdjacencyList) AllVertices() []interface{} {
	keys := make([]interface{}, 0, 0)
	for v := g.matrix.frontKey(); v != nil; v = g.matrix.nextKey(v) {
		keys = append(keys, v)
	}
	return keys
}

func (g *AdjacencyList) AllConnectedVertices(v interface{}) []interface{} {
	value := make([]interface{}, 0, 0)
	if g.matrix.exist(v) {
		for e := g.matrix.get(v).(*list.List).Front(); e != nil; e = e.Next() {
			value = append(value, e.Value)
		}
	}
	return value
}

func (g *AdjacencyList) AllEdges() []Edge {
	edges := make([]Edge, 0, 0)
	for v := g.matrix.frontKey(); v != nil; v = g.matrix.nextKey(v) {
		for e := g.matrix.get(v).(*list.List).Front(); e != nil; e = e.Next() {
			edges = append(edges, Edge{v, e.Value})
		}
	}
	return edges
}


func (g *AdjacencyList) Transpose() Graph {
	gt := NewAdjacencyList()
	for _, e := range g.AllEdges() {
		gt.AddEdge(Edge{e.End,e.Start})
	}
	return gt
}


func (g *AdjacencyList) GetGraph() interface{} {
	return g
}

func NewAdjacencyList() *AdjacencyList {
	return new(AdjacencyList).Init()
}

func AdjacencyList2AdjacencyMatrix(l *AdjacencyList) *AdjacencyMatrix {
	m := NewAdjacencyMatrix()
	for v := l.matrix.frontKey(); v != nil; v = l.matrix.nextKey(v) {
		for e := l.matrix.get(v).(*list.List).Front(); e != nil; e = e.Next() {
			m.AddEdge(Edge{v, e.Value})
		}
	}
	return m
}

func AdjacencyMatrix2AdjacencyList(m *AdjacencyMatrix) *AdjacencyList {
	l := NewAdjacencyList()
	for start := m.matrix.frontKey(); start != nil; start = m.matrix.nextKey(start) {
		for end := m.matrix.get(start).(*linkedMap).frontKey(); end != nil; end = m.matrix.get(start).(*linkedMap).nextKey(end) {
			l.AddEdge(Edge{start, end})
		}
	}
	return l
}

func CreateGraphByType(g Graph) (newG Graph) {
	if _, isList := g.GetGraph().(*AdjacencyList); isList {
		newG = NewAdjacencyList()
	} else {
		newG = NewAdjacencyMatrix()
	}
	return
}