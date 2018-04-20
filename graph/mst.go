package graph

import (
	"github.com/shady831213/algorithms/heap"
	"github.com/shady831213/algorithms/tree/disjointSetTree"
	"math"
	"sort"
)

func mstKruskal(g weightedGraph) weightedGraph {
	t := createGraphByType(g).(weightedGraph)
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

func mstPrim(g weightedGraph) weightedGraph {
	t := createGraphByType(g).(weightedGraph)
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

func secondaryMst(g weightedGraph) weightedGraph {
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

func reduceGraphLessWeight(g weightedGraph, origin map[edge]edge, getRoot func(interface{}) interface{}) (weightedGraph, map[edge]edge) {
	//origin points to edge in origin graph
	newG := createGraphByType(g).(weightedGraph)
	newOrigin := make(map[edge]edge)
	for _, e := range g.AllEdges() {
		rootStart := getRoot(e.Start)
		rootEnd := getRoot(e.End)
		if rootStart != rootEnd {
			newE := edge{rootStart, rootEnd}
			//if start and end don't have same root
			if !newG.CheckEdge(newE) {
				// if the edge is not in new graph, add to new graph, use origin weight
				newG.AddEdgeWithWeightBi(newE, g.Weight(e))
				newOrigin[newE] = origin[e]
				newOrigin[edge{newE.End, newE.Start}] = origin[edge{e.End, e.Start}]
			} else if g.Weight(e) < newG.Weight(newE) {
				//update new Weight which is less , and update new origin
				newG.AddEdgeWithWeightBi(newE, g.Weight(e))
				newOrigin[newE] = origin[e]
				newOrigin[edge{newE.End, newE.Start}] = origin[edge{e.End, e.Start}]
			}
		}
	}
	return newG, newOrigin
}

func mstReduceOnce(g, t weightedGraph, origin map[edge]edge) (weightedGraph, map[edge]edge) {
	//build the disjoint-sets as components by minimum weight path
	set := make(map[interface{}]*disjointSetTree.DisjointSet)
	mark := make(map[interface{}]bool)
	for _, v := range g.AllVertices() {
		if _, ok := mark[v]; !ok {
			set[v] = disjointSetTree.MakeSet(v)
			iter := g.IterConnectedVertices(v)
			minWeight := math.MaxInt32
			var minEnd interface{}
			for e := iter.Value(); e != nil; e = iter.Next() {
				if _, ok := set[e]; !ok {
					set[e] = disjointSetTree.MakeSet(e)
				}
				if g.Weight(edge{v, e}) < minWeight {
					minWeight = g.Weight(edge{v, e})
					minEnd = e
				}
			}
			//add minimum weight edge to sub graph and tree
			t.AddEdgeWithWeightBi(origin[edge{v, minEnd}], minWeight)
			//shrink
			disjointSetTree.Union(set[v], set[minEnd])
			mark[v] = true
			mark[minEnd] = true
		}
	}

	return reduceGraphLessWeight(g, origin, func(v interface{}) interface{} {
		return disjointSetTree.FindSet(set[v]).Value
	})
}

func mstReducedPrim(g weightedGraph, k int) weightedGraph {

	t := createGraphByType(g).(weightedGraph)

	origin := make(map[edge]edge)
	for _, e := range g.AllEdges() {
		origin[e] = e
	}

	newG := g

	for i := 0; i < k; i++ {
		newG, origin = mstReduceOnce(newG, t, origin)
	}

	newT := mstPrim(newG)

	for _, e := range newT.AllEdges() {
		t.AddEdgeWithWeight(origin[e], newT.Weight(origin[e]))
	}
	return t
}

/*
bottleNeck Spanning Tree
*/

//partitionGraph : O(E)
func partitionGraph(g weightedGraph) (g1, g2 weightedGraph) {
	edges := g.AllEdges()
	if len(edges) == 2 {
		return g, nil
	}

	g1, g2 = createGraphByType(g).(weightedGraph), createGraphByType(g).(weightedGraph)

	mid, idx := len(edges)/2, 0
	for end, start := len(edges)-1, 0; idx != mid; {
		sanity := g.Weight(edges[end])
		idx = start
		for i := start; i <= end; i++ {
			if g.Weight(edges[i]) < sanity {
				edges[i], edges[idx] = edges[idx], edges[i]
				idx++
			}
		}
		edges[idx], edges[end] = edges[end], edges[idx]
		if idx < mid {
			start = idx + 1
		} else {
			end = idx - 1
		}
	}

	for _, v := range g.AllVertices() {
		g1.AddVertex(v)
		g2.AddVertex(v)
	}

	for i := range edges {
		if i < mid {
			g1.AddEdgeWithWeight(edges[i], g.Weight(edges[i]))
		} else {
			g2.AddEdgeWithWeight(edges[i], g.Weight(edges[i]))
		}
	}
	return
}

func bottleNeckSpanningTreeHandle(g, t weightedGraph, origin map[edge]edge) weightedGraph {
	//partition by medium
	g1, g2 := partitionGraph(g)
	if g2 == nil {
		for _, e := range g1.AllEdges() {
			t.AddEdgeWithWeight(origin[e], g1.Weight(e))
		}
		return g1
	}

	//dfs g1
	comps := make(map[interface{}]weightedGraph)
	roots := make(map[interface{}]interface{})
	treeEdge := createGraphByType(g).(weightedGraph)
	handler := newDFSVisitHandler()
	handler.TreeEdgeHandler = func(start, end *dfsElement) {
		e := edge{start.V, end.V}
		comps[end.FindRoot().V].AddEdgeWithWeightBi(e, g1.Weight(e))
		treeEdge.AddEdgeWithWeightBi(e, g1.Weight(e))
	}
	handler.AfterDfsHandler = func(v *dfsElement) {
		roots[v.V] = v.FindRoot().V
	}

	for _, v := range g1.AllVertices() {
		if !handler.Elements.exist(v) {
			comps[v] = createGraphByType(g1).(weightedGraph)
			dfsVisit(g1, v, handler)
		}
	}
	//reclusive if all are connected
	if len(comps) == 1 {
		return bottleNeckSpanningTreeHandle(g1, t, origin)
	}

	//add comps' edges to the tree, reclusive other edges which are cross comps

	for _, e := range treeEdge.AllEdges() {
		t.AddEdgeWithWeight(origin[e], g.Weight(e))
	}
	reducedG, newOrigin := reduceGraphLessWeight(g, origin, func(v interface{}) interface{} {
		return roots[v]
	})
	return bottleNeckSpanningTreeHandle(reducedG, t, newOrigin)
}

func bottleNeckSpanningTree(g weightedGraph) weightedGraph {
	t := createGraphByType(g).(weightedGraph)
	origin := make(map[edge]edge)
	for _, e := range g.AllEdges() {
		origin[e] = e
	}
	bottleNeckSpanningTreeHandle(g, t, origin)
	return t
}
