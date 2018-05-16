package graph

import (
	"fmt"
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
	CheckVertex(interface{}) bool
	DeleteVertex(interface{})
	AddEdge(edge)
	CheckEdge(edge) bool
	DeleteEdge(edge)
	AddEdgeBi(edge)
	DeleteEdgeBi(edge)
	AllVertices() []interface{}
	AllEdges() []edge
	AllConnectedVertices(interface{}) []interface{}
	IterConnectedVertices(interface{}) iterator
	Transpose() graph
}

type adjacencyMatrix struct {
	matrix linkedMap
	graph
}

func (g *adjacencyMatrix) init() *adjacencyMatrix {
	g.matrix = *new(linkedMap).init()
	return g
}

func (g *adjacencyMatrix) AddVertex(vertex interface{}) {
	if !g.matrix.exist(vertex) {
		g.matrix.add(vertex, new(linkedMap).init())
	}
}

func (g *adjacencyMatrix) CheckVertex(vertex interface{}) bool {
	return g.matrix.exist(vertex)
}

func (g *adjacencyMatrix) DeleteVertex(vertex interface{}) {
	g.matrix.delete(vertex)
	for v := g.matrix.frontKey(); v != nil; v = g.matrix.nextKey(v) {
		g.DeleteEdge(edge{v, vertex})
	}
}

func (g *adjacencyMatrix) AddEdge(e edge) {
	g.AddVertex(e.Start)
	g.AddVertex(e.End)
	g.matrix.get(e.Start).(*linkedMap).add(e.End, true)
}

func (g *adjacencyMatrix) CheckEdge(e edge) bool {
	if !g.CheckVertex(e.Start) {
		return false
	}
	return g.matrix.get(e.Start).(*linkedMap).exist(e.End)
}

func (g *adjacencyMatrix) DeleteEdge(e edge) {
	if g.matrix.exist(e.Start) {
		g.matrix.get(e.Start).(*linkedMap).delete(e.End)
	}
}

func (g *adjacencyMatrix) AddEdgeBi(e edge) {
	g.AddVertex(e.Start)
	g.AddVertex(e.End)
	g.matrix.get(e.Start).(*linkedMap).add(e.End, true)
	g.matrix.get(e.End).(*linkedMap).add(e.Start, true)
}

func (g *adjacencyMatrix) DeleteEdgeBi(e edge) {
	g.DeleteEdge(e)
	g.DeleteEdge(edge{e.End, e.Start})
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
	gt := newGraph()
	for _, e := range g.AllEdges() {
		gt.AddEdge(edge{e.End, e.Start})
	}
	return gt
}

func newGraph() graph {
	return new(adjacencyMatrix).init()
}

type weightedGraph interface {
	graph
	Weight(edge) int
	AddEdgeWithWeight(edge, int)
	AddEdgeWithWeightBi(edge, int)
	TotalWeight() int
}

type adjacencyMatrixWithWeight struct {
	adjacencyMatrix
	weights map[edge]int
	tw      int
}

func (g *adjacencyMatrixWithWeight) Init() *adjacencyMatrixWithWeight {
	g.adjacencyMatrix.init()
	g.weights = make(map[edge]int)
	g.tw = 0
	return g
}

func (g *adjacencyMatrixWithWeight) AddEdgeWithWeight(e edge, w int) {
	g.adjacencyMatrix.AddEdge(e)
	if _, ok := g.weights[e]; ok {
		g.tw = g.tw - g.weights[e] + w
	} else {
		g.tw = g.tw + w
	}
	g.weights[e] = w
}

func (g *adjacencyMatrixWithWeight) DeleteEdge(e edge) {
	g.adjacencyMatrix.DeleteEdge(e)
	if w, ok := g.weights[e]; ok {
		g.tw -= w
		delete(g.weights, e)
	}
}

func (g *adjacencyMatrixWithWeight) AddEdgeWithWeightBi(e edge, w int) {
	g.AddEdgeWithWeight(e, w)
	g.AddEdgeWithWeight(edge{e.End, e.Start}, w)
}

func (g *adjacencyMatrixWithWeight) DeleteEdgeBi(e edge) {
	g.DeleteEdge(e)
	g.DeleteEdge(edge{e.End, e.Start})
}

func (g *adjacencyMatrixWithWeight) Weight(e edge) int {
	if value, ok := g.weights[e]; ok {
		return value
	}
	return -1
}

func (g *adjacencyMatrixWithWeight) TotalWeight() int {
	return g.tw
}

func newWeightedGraph() weightedGraph {
	return new(adjacencyMatrixWithWeight).Init()
}

type flowGraph interface {
	graph
	Cap(edge) int
	Flow(edge) int
	AddEdgeWithCap(edge, int)
	AddEdgeWithFlow(edge, int)
}

type adjacencyMatrixWithFlow struct {
	adjacencyMatrix
	cap  map[edge]int
	flow map[edge]int
}

func (g *adjacencyMatrixWithFlow) init() *adjacencyMatrixWithFlow {
	g.adjacencyMatrix.init()
	g.cap = make(map[edge]int)
	g.flow = make(map[edge]int)
	return g
}

func (g *adjacencyMatrixWithFlow) AddEdgeWithCap(e edge, c int) {
	g.adjacencyMatrix.AddEdge(e)
	g.cap[e] = c
}

func (g *adjacencyMatrixWithFlow) AddEdgeWithFlow(e edge, f int) {
	g.adjacencyMatrix.AddEdge(e)
	g.flow[e] = f
}

func (g *adjacencyMatrixWithFlow) Cap(e edge) int {
	if _, ok := g.cap[e]; !ok {
		return 0
	}
	return g.cap[e]
}

func (g *adjacencyMatrixWithFlow) Flow(e edge) int {
	if _, ok := g.flow[e]; !ok {
		return 0
	}
	return g.flow[e]
}

func (g *adjacencyMatrixWithFlow) DeleteEdge(e edge) {
	g.adjacencyMatrix.DeleteEdge(e)
	delete(g.cap, e)
	delete(g.flow, e)
}

func (g *adjacencyMatrixWithFlow) AddEdgeBi(e edge) {
	panic(fmt.Sprintln("not valid in flow graph!"))
}

func (g *adjacencyMatrixWithFlow) DeleteEdgeBi(e edge) {
	panic(fmt.Sprintln("not valid in flow graph!"))
}

func newFlowGraph() flowGraph {
	return new(adjacencyMatrixWithFlow).init()
}
