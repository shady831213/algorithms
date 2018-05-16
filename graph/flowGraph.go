package graph

import (
	"math"
)

func augmentingPath(g flowGraph, s interface{}, t interface{}) (int, []edge) {
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

func updateFlow(rg, g flowGraph, rc int, edges []edge) {
	updateResidualGraph := func(g flowGraph, flow int, e edge) {
		g.AddEdgeWithFlow(e, flow)
		if g.RCap(e) == 0 {
			g.DeleteEdge(e)
		}
		re := edge{e.End, e.Start}
		g.AddEdgeWithFlow(re, 0-flow)
		if g.RCap(re) == 0 {
			g.DeleteEdge(re)
		}
	}

	for _, e := range edges {
		flow := g.Flow(e) + rc
		g.AddEdgeWithFlow(e, flow)
		updateResidualGraph(rg, flow, e)
	}
}

func edmondsKarp(g flowGraph, s interface{}, t interface{}) {
	residualG := newFlowGraph()
	for _, e := range g.AllEdges() {
		g.AddEdgeWithFlow(e, 0)
		residualG.AddEdgeWithCap(e, g.Cap(e))
	}

	for rc, edges := augmentingPath(residualG, s, t); len(edges) > 0; rc, edges = augmentingPath(residualG, s, t) {
		updateFlow(residualG, g, rc, edges)
	}

}
