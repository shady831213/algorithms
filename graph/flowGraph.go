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
		g.DeleteEdge(e)
	}
}

func (g *adjacencyMatrixResidual) RCap(e edge) int {
	return g.Cap(e) - g.Flow(e)
}

func newResidualGraph() residualGraph {
	return new(adjacencyMatrixResidual).init()
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

func edmondsKarp(g flowGraph, s interface{}, t interface{}) {
	residualG := newResidualGraph()
	for _, e := range g.AllEdges() {
		g.AddEdgeWithFlow(e, 0)
		residualG.AddEdgeWithCap(e, g.Cap(e))
	}

	for rc, edges := augmentingPath(residualG, s, t); len(edges) > 0; rc, edges = augmentingPath(residualG, s, t) {
		updateFlow(residualG, g, rc, edges)
	}

}

func bioGraphMaxMatch(g graph, l []interface{}) graph {
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

	edmondsKarp(fG, s, t)
	matchG := newGraph()
	for _, e := range fG.AllEdges() {
		if fG.Flow(e) > 0 && e.Start != s && e.End != t {
			matchG.AddEdgeBi(e)
		}
	}
	return matchG
}
