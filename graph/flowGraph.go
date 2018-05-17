package graph

import (
	"math"
)

type residualGraph interface {
	flowGraph
	RCap(edge) int
}

type adjacencyMatrixResidual struct {
	adjacencyMatrixWithFlow
}

func (g *adjacencyMatrixResidual) init() *adjacencyMatrixResidual {
	g.adjacencyMatrixWithFlow.init()
	return g
}

func (g *adjacencyMatrixResidual) AddEdgeWithFlow(e edge, f int) {
	g.adjacencyMatrixWithFlow.AddEdgeWithFlow(e, f)
	if g.RCap(e) == 0 {
		//delete edge but keep cap and flow, because if residual flow (RCap) change, this edge should be added back
		g.adjacencyMatrix.DeleteEdge(e)
	}
}

func (g *adjacencyMatrixResidual) RCap(e edge) int {
	return g.Cap(e) - g.Flow(e)
}

func newResidualGraph(g flowGraph) residualGraph {
	residualG := new(adjacencyMatrixResidual).init()
	for _, e := range g.AllEdges() {
		residualG.AddEdgeWithCap(e, g.Cap(e))
	}
	return residualG
}

func augmentingPath(g residualGraph, s interface{}, t interface{}) (int, []edge) {
	augmentingEdges := make([]edge, 0, 0)
	minRC := math.MaxInt32
	handler := newBFSVisitHandler()
	handler.EdgeHandler = func(start, end *bfsElement) {
		if end.V == t {
			for v := end; v.P != nil; v = v.P {
				currentEdge := edge{v.P.V, v.V}
				augmentingEdges = append(augmentingEdges, currentEdge)
				if rc := g.RCap(currentEdge); rc < minRC {
					minRC = rc
				}
			}
		}
	}

	bfsVisit(g, s, handler)

	return minRC, augmentingEdges
}

func updateFlow(rg residualGraph, g flowGraph, rc int, edges []edge) {
	for _, e := range edges {
		flow := g.Flow(e) + rc
		g.AddEdgeWithFlow(e, flow)
		rg.AddEdgeWithFlow(e, flow)
		re := edge{e.End, e.Start}
		rg.AddEdgeWithFlow(re, 0-flow)
	}
}

func edmondesKarp(g flowGraph, s interface{}, t interface{}) {
	residualG := newResidualGraph(g)

	for rc, edges := augmentingPath(residualG, s, t); len(edges) > 0; rc, edges = augmentingPath(residualG, s, t) {
		updateFlow(residualG, g, rc, edges)
	}

}

type preFlowGraph interface {
	residualGraph
	SetHeight(interface{}, int)
	SetExcess(interface{}, int)
	Height(interface{}) int
	Excess(interface{}) int
	Push(edge) bool
	Relabel(interface{}) bool
	Overflow(interface{}) bool
}

type adjacencyMatrixPreFlow struct {
	adjacencyMatrixResidual
	height, excess map[interface{}]int
	s, t           interface{}
}

func (g *adjacencyMatrixPreFlow) init(s, t interface{}) *adjacencyMatrixPreFlow {
	g.adjacencyMatrixResidual.init()
	g.height = make(map[interface{}]int)
	g.excess = make(map[interface{}]int)
	g.s = s
	g.t = t
	return g
}

func (g *adjacencyMatrixPreFlow) SetHeight(v interface{}, h int) {
	g.height[v] = h
}

func (g *adjacencyMatrixPreFlow) Height(v interface{}) int {
	if _, ok := g.height[v]; !ok {
		return 0
	}
	return g.height[v]
}

func (g *adjacencyMatrixPreFlow) SetExcess(v interface{}, e int) {
	g.excess[v] = e
}

func (g *adjacencyMatrixPreFlow) Excess(v interface{}) int {
	if _, ok := g.excess[v]; !ok {
		return 0
	}
	return g.excess[v]
}

func (g *adjacencyMatrixPreFlow) Push(e edge) bool {
	if g.Overflow(e.Start) && g.RCap(e) > 0 && g.Height(e.Start) == g.Height(e.End)+1 {
		d := g.RCap(e)
		if g.Excess(e.Start) < d {
			d = g.Excess(e.Start)
		}
		flow := g.Flow(e) + d
		g.AddEdgeWithFlow(e, flow)
		re := edge{e.End, e.Start}
		g.AddEdgeWithFlow(re, 0-flow)
		g.SetExcess(e.Start, g.Excess(e.Start)-d)
		g.SetExcess(e.End, g.Excess(e.End)+d)
		return true
	}
	return false
}

func (g *adjacencyMatrixPreFlow) Relabel(v interface{}) bool {
	if !g.Overflow(v) {
		return false
	}
	iter := g.IterConnectedVertices(v)
	minH := math.MaxInt32
	for end := iter.Value(); end != nil; end = iter.Next() {
		if g.Height(end) < minH {
			minH = g.Height(end)
		}
		if g.Height(v) > g.Height(end) {
			return false
		}
	}
	g.SetHeight(v, minH+1)
	return true
}

func (g *adjacencyMatrixPreFlow) Overflow(v interface{}) bool {
	//According to definition, start and target vertex never overflow
	return v != g.s && v != g.t && g.Excess(v) > 0
}

func newPreFlowGraph(g flowGraph, s interface{}, t interface{}) preFlowGraph {
	preFlowG := new(adjacencyMatrixPreFlow).init(s, t)
	vertices := g.AllVertices()
	for _, e := range g.AllEdges() {
		preFlowG.AddEdgeWithCap(e, g.Cap(e))
	}
	preFlowG.SetHeight(s, len(vertices))
	iter := g.IterConnectedVertices(s)
	for v := iter.Value(); v != nil; v = iter.Next() {
		c := g.Cap(edge{s, v})
		preFlowG.AddEdgeWithFlow(edge{s, v}, c)
		preFlowG.AddEdgeWithFlow(edge{v, s}, 0-c)
		preFlowG.SetExcess(v, c)
		preFlowG.SetExcess(s, preFlowG.Excess(s)-c)
	}
	return preFlowG
}

func pushRelabel(g flowGraph, s interface{}, t interface{}) {
	preFlowG := newPreFlowGraph(g, s, t)
	for stop := false; !stop; {
		stop = true
		for _, e := range preFlowG.AllEdges() {
			stop = stop && !preFlowG.Push(e) && !preFlowG.Relabel(e.Start) && !preFlowG.Relabel(e.End)
		}
	}
	for _, e := range g.AllEdges() {
		g.AddEdgeWithFlow(e, preFlowG.Flow(e))
	}
}

func bioGraphMaxMatch(g graph, l []interface{}, flowAlg func(g flowGraph, s interface{}, t interface{})) graph {
	//build flow graph
	fG := newFlowGraph()
	s := struct{ start string }{"s"}
	t := struct{ end string }{"t"}
	for _, vl := range l {
		fG.AddEdgeWithCap(edge{s, vl}, 1)
		iter := g.IterConnectedVertices(vl)
		for rv := iter.Value(); rv != nil; rv = iter.Next() {
			fG.AddEdgeWithCap(edge{vl, rv}, 1)
			fG.AddEdgeWithCap(edge{rv, t}, 1)
		}
	}

	flowAlg(fG, s, t)
	matchG := newGraph()
	for _, e := range fG.AllEdges() {
		if fG.Flow(e) > 0 && e.Start != s && e.End != t {
			matchG.AddEdgeBi(e)
		}
	}
	return matchG
}
