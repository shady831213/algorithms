package graph

import (
	"github.com/shady831213/algorithms/heap"
	"github.com/shady831213/algorithms/tree/disjointSetTree"
	"math"
	"sort"
)

func mstKruskal(g graphWeightily) graphWeightily {
	t := createGraphByType(g).(graphWeightily)
	dfsForest := dfs(g, nil)
	edges := append(dfsForest.AllTreeEdges(), dfsForest.AllForwardEdges()...)
	verticesSet := make(map[interface{}]*disjointSetTree.DisjointSet)
	sort.Slice(edges, func(i, j int) bool {
		return g.Weight(edge{edges[i].Start.(*dfsElement).V, edges[i].End.(*dfsElement).V}) <
			g.Weight(edge{edges[j].Start.(*dfsElement).V, edges[j].End.(*dfsElement).V})
	})

	for _, e := range edges {
		if _, ok := verticesSet[e.Start]; !ok {
			verticesSet[e.Start] = disjointSetTree.MakeSet(e.Start)
		}
		if _, ok := verticesSet[e.End]; !ok {
			verticesSet[e.End] = disjointSetTree.MakeSet(e.End)
		}
		if disjointSetTree.FindSet(verticesSet[e.Start]) != disjointSetTree.FindSet(verticesSet[e.End]) {
			edge := edge{e.Start.(*dfsElement).V, e.End.(*dfsElement).V}
			t.AddEdgeWithWeightBi(edge, g.Weight(edge))
			disjointSetTree.Union(verticesSet[e.Start], verticesSet[e.End])
		}
	}

	return t
}

type fibHeapLessIntMixin struct {
	heap.FibHeapMixin
}

func (m *fibHeapLessIntMixin) LessKey(i, j interface{}) bool {
	return i.(int) < j.(int)
}

func newFibHeapKeyInt() *heap.FibHeap {
	return new(heap.FibHeap).Init(new(fibHeapLessIntMixin))
}

type mstElement struct {
	p         *mstElement
	pqElement *heap.FibHeapElement
}

func mstPrim(g graphWeightily) graphWeightily {
	t := createGraphByType(g).(graphWeightily)
	pq := newFibHeapKeyInt()
	elements := make(map[interface{}]*heap.FibHeapElement)
	p := make(map[interface{}]interface{})
	for i, v := range g.AllVertices() {
		if i == 0 {
			pq.Insert(-1, v)
		} else {
			elements[v] = pq.Insert(math.MaxInt32, v)
		}
	}

	for pq.Degree() != 0 {
		minElement := pq.ExtractMin()
		v := minElement.Value
		delete(elements, v)
		iter := g.IterConnectedVertices(v)
		for e := iter.Value(); e != nil; e = iter.Next() {
			currentEdge := edge{v, e}
			if element, ok := elements[e]; ok && g.Weight(currentEdge) < element.Key.(int) {
				if _, ok := p[e]; ok {
					t.DeleteEdgeBi(edge{p[e], e})
				}
				p[e] = v
				t.AddEdgeWithWeightBi(currentEdge, g.Weight(currentEdge))
				pq.ModifyNode(element, g.Weight(currentEdge), element.Value)
			}
		}
	}

	return t
}

func secondaryMst(g graphWeightily) graphWeightily {
	t := mstPrim(g)
	//dynamic programming, use BFS to visit all the [i,j] path
	maxEdge := make(map[interface{}]map[interface{}]edge)

	for _, v := range t.AllVertices() {
		if _, ok := maxEdge[v]; !ok {
			maxEdge[v] = make(map[interface{}]edge)
		}
		handler := newBFSVisitHandler()
		handler.EdgeHandler = func(start, end *bfsElement) {
			maxEdge[v][end.V] = edge{v, end.V}
			if t.Weight(maxEdge[v][end.V]) < t.Weight(maxEdge[v][start.V]) {
				maxEdge[v][end.V] = maxEdge[v][start.V]
			}
			if _, ok := maxEdge[end.V]; !ok {
				maxEdge[end.V] = make(map[interface{}]edge)
			}
			maxEdge[end.V][v] = edge{maxEdge[v][end.V].End, maxEdge[v][end.V].Start}
		}
		bfsVisit(t, v, handler)
	}

	//iterate all the edge and find the minimum total weight
	minWeight := math.MaxInt32
	var edgePair struct{ origin, replace edge }

	for _, v := range g.AllVertices() {
		iter := g.IterConnectedVertices(v)
		for e := iter.Value(); e != nil; e = iter.Next() {
			currentEdge := edge{v, e}
			if !t.CheckEdge(currentEdge) {
				origin := maxEdge[v][e]
				t.DeleteEdgeBi(origin)
				t.AddEdgeWithWeightBi(currentEdge, g.Weight(currentEdge))
				if t.TotalWeight() < minWeight {
					minWeight = t.TotalWeight()
					edgePair = struct{ origin, replace edge }{origin, currentEdge}
				}
				t.DeleteEdgeBi(currentEdge)
				t.AddEdgeWithWeightBi(origin, g.Weight(origin))
			}
		}
	}
	t.DeleteEdgeBi(edgePair.origin)
	t.AddEdgeWithWeightBi(edgePair.replace, g.Weight(edgePair.replace))
	return t
}
