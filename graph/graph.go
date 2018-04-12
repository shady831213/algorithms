package graph

import (
	"container/list"
	"sync"
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
	IterConnectedVertices(interface{}) <-chan interface{}
	Transpose() graph
	GetGraph() interface{}
}

type chanWithStatus struct {
	ch    chan interface{}
	close bool
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
	}
	return nil
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

type adjacencyMatrix struct {
	matrix linkedMap
	sync.RWMutex
	chs map[interface{}]*chanWithStatus
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

func (g *adjacencyMatrix) IterConnectedVertices(v interface{}) <-chan interface{} {
	if _, ok := g.chs[v]; !ok {
		g.chs[v] = &chanWithStatus{make(chan interface{}), false}
		go func() {
			if g.matrix.exist(v) {
				for key := g.matrix.get(v).(*linkedMap).frontKey(); key != nil; key = g.matrix.get(v).(*linkedMap).nextKey(key) {
					g.chs[v].ch <- key
				}
			}
			g.Lock()
			close(g.chs[v].ch)
			g.chs[v].close = true
			g.Unlock()
		}()
	} else if g.chs[v].close {
		lastChan := g.chs[v].ch
		delete(g.chs, v)
		return lastChan
	}
	return g.chs[v].ch
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
	sync.RWMutex
	chs map[interface{}]*chanWithStatus
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

func (g *adjacencyList) IterConnectedVertices(v interface{}) <-chan interface{} {
	if _, ok := g.chs[v]; !ok {
		g.chs[v] = &chanWithStatus{make(chan interface{}), false}
		go func() {
			if g.matrix.exist(v) {
				for e := g.matrix.get(v).(*list.List).Front(); e != nil; e = e.Next() {
					g.chs[v].ch <- e.Value
				}
			}
			g.Lock()
			close(g.chs[v].ch)
			g.chs[v].close = true
			g.Unlock()
		}()
	} else if g.chs[v].close {
		lastChan := g.chs[v].ch
		delete(g.chs, v)
		return lastChan
	}

	return g.chs[v].ch
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
