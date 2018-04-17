package graph

import (
	"github.com/shady831213/algorithms/heap"
	"github.com/shady831213/algorithms/tree/disjointSetTree"
	"math"
	"sort"
)

func mstKruskal(g graphWeightily) graphWeightily {
	t := createGraphWithWeightByType(g)
	dfsForest := dfs(g.GetGraph(), nil)
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

func mstPrim(g graphWeightily) graphWeightily {
	t := createGraphWithWeightByType(g)
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
		for keyValue := iter.Value(); keyValue != nil; keyValue = iter.Next() {
			e := keyValue.(struct{ key, value interface{} }).key
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
