package graph

import (
	"github.com/shady831213/algorithms/tree/disjointSetTree"
	"math"
)

func augmentingPath(g flowGraph, s interface{}, t interface{}) (int, []edge) {
	set := make(map[interface{}]*disjointSetTree.DisjointSet)
	handler := newBFSVisitHandler()
	handler.EdgeHandler = func(start, end *bfsElement) {
		if _, ok := set[start.V]; !ok {
			set[start.V] = disjointSetTree.MakeSet(start)
		}
		if _, ok := set[end.V]; !ok {
			set[end.V] = disjointSetTree.MakeSet(end)
		}
		if g.RCap(edge{start.V, end.V}) > 0 {
			disjointSetTree.Union(set[start.V], set[end.V])
		}
	}

	bfsVisit(g, s, handler)

	if _, ok := set[t]; ok && disjointSetTree.FindSet(set[t]) == disjointSetTree.FindSet(set[s]) {
		augmentingEdges := make([]edge, 0, 0)
		minRC := math.MaxInt32
		for v := set[t].Value.(*bfsElement); v.P != nil; v = v.P {
			currentEdge := edge{v.P.V, v.V}
			augmentingEdges = append(augmentingEdges, currentEdge)
			if rc := g.RCap(currentEdge); rc < minRC {
				minRC = rc
			}
		}
		return minRC, augmentingEdges
	}
	return 0, make([]edge, 0, 0)
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
	residualG := createGraphByType(g).(flowGraph)
	for _, e := range g.AllEdges() {
		g.AddEdgeWithFlow(e, 0)
		residualG.AddEdgeWithCap(e, g.Cap(e))
	}

	for rc, edges := augmentingPath(residualG, s, t); len(edges) > 0; rc, edges = augmentingPath(residualG, s, t) {
		updateFlow(residualG, g, rc, edges)
	}

}
