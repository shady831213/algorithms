package graph

import (
	"container/list"
)

func vertexBCC(g graph) (cuts graph, comps []graph) {
	cuts = createGraphByType(g)
	comps = make([]graph, 0, 0)
	lows := make(map[interface{}]int)
	children := make(map[interface{}]int)
	edgeStack := list.New()
	handler := newDFSVisitHandler()
	handler.BeforeBfsHandler = func(v *dfsElement) {
		lows[v.V] = v.D
	}
	handler.TreeEdgeHandler = func(start, end *dfsElement) {
		children[start.V]++
		edgeStack.PushBack(edge{start.V, end.V})
	}
	handler.BackEdgeHandler = func(start, end *dfsElement) {
		if end.D < lows[start.V] && start.P != end {
			edgeStack.PushBack(edge{start.V, end.V})
			lows[start.V] = end.D
		}
	}
	handler.AfterBfsHandler = func(v *dfsElement) {
		p := v.P
		if p == nil {
			return
		}
		if lows[v.V] < lows[p.V] {
			lows[p.V] = lows[v.V]
		}
		if lows[v.V] >= p.D {
			if !(p.D == 1 && children[p.V] < 2) {
				cuts.AddVertex(p.V)
			}
			comps = append(comps, createGraphByType(g))
			curEdge := edge{p.V, v.V}
			comps[len(comps)-1].AddEdgeBi(curEdge)
			for e := edgeStack.Back().Value; e != curEdge; e = edgeStack.Back().Value {
				edgeStack.Remove(edgeStack.Back())
				comps[len(comps)-1].AddEdgeBi(e.(edge))
			}
			edgeStack.Remove(edgeStack.Back())
		}
	}

	for _, v := range g.AllVertices() {
		if !handler.Elements.exist(v) {
			dfsVisit(g, v, handler)
		}
	}
	return
}

func edgeBCC(g graph) (bridges graph, comps []graph) {
	bridges = createGraphByType(g)
	comps = make([]graph, 0, 0)
	lows := make(map[interface{}]int)
	edgeStack := list.New()
	handler := newDFSVisitHandler()
	handler.BeforeBfsHandler = func(v *dfsElement) {
		lows[v.V] = v.D
	}
	handler.TreeEdgeHandler = func(start, end *dfsElement) {
		edgeStack.PushBack(edge{start.V, end.V})
	}
	handler.BackEdgeHandler = func(start, end *dfsElement) {
		if end.D < lows[start.V] && start.P != end {
			edgeStack.PushBack(edge{start.V, end.V})
			lows[start.V] = end.D
		}
	}
	handler.AfterBfsHandler = func(v *dfsElement) {
		p := v.P
		if p == nil {
			return
		}
		if lows[v.V] < lows[p.V] {
			lows[p.V] = lows[v.V]
		}
		if lows[v.V] >= p.D {
			comp := createGraphByType(g)
			curEdge := edge{p.V, v.V}
			if lows[v.V] > p.D {
				bridges.AddEdgeBi(curEdge)
			} else {
				//excluding bridge
				comp.AddEdgeBi(curEdge)
			}
			for e := edgeStack.Back().Value; e != curEdge; e = edgeStack.Back().Value {
				edgeStack.Remove(edgeStack.Back())
				comp.AddEdgeBi(e.(edge))
			}
			edgeStack.Remove(edgeStack.Back())
			//if component is not empty, push it to the slice
			if len(comp.AllVertices()) != 0 {
				comps = append(comps, comp)
			}
		}

	}

	for _, v := range g.AllVertices() {
		if !handler.Elements.exist(v) {
			dfsVisit(g, v, handler)
		}
	}
	return
}
