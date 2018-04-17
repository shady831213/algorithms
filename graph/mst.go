package graph

import (
	"github.com/shady831213/algorithms/heap"
	"github.com/shady831213/algorithms/tree/disjointSetTree"
	"math"
	"sort"
)

type graphWeightily interface {
	graph
	Weight(edge) int
	AddEdgeWithWeight(edge, int)
	AddEdgeWithWeightBi(edge, int)
}

type adjacencyMatrixWithWeight struct {
	adjacencyMatrix
	weights map[edge]int
}

func (g *adjacencyMatrixWithWeight) Init() *adjacencyMatrixWithWeight {
	g.adjacencyMatrix.init()
	g.weights = make(map[edge]int)
	return g
}

func (g *adjacencyMatrixWithWeight) AddEdgeWithWeight(e edge, w int) {
	g.adjacencyMatrix.AddEdge(e)
	g.weights[e] = w
}

func (g *adjacencyMatrixWithWeight) AddEdgeWithWeightBi(e edge, w int) {
	g.adjacencyMatrix.AddEdgeBi(e)
	g.weights[e] = w
	g.weights[edge{e.End, e.Start}] = w
}

func (g *adjacencyMatrixWithWeight) Weight(e edge) int {
	if value, ok := g.weights[e]; ok {
		return value
	}
	return -1
}

func newAdjacencyMatrixWithWeight() *adjacencyMatrixWithWeight {
	return new(adjacencyMatrixWithWeight).Init()
}

type adjacencyListWithWeight struct {
	weights map[edge]int
	adjacencyList
}

func (g *adjacencyListWithWeight) Init() *adjacencyListWithWeight {
	g.adjacencyList.init()
	g.weights = make(map[edge]int)
	return g
}

func (g *adjacencyListWithWeight) AddEdgeWithWeight(e edge, w int) {
	g.adjacencyList.AddEdge(e)
	g.weights[e] = w
}

func (g *adjacencyListWithWeight) AddEdgeWithWeightBi(e edge, w int) {
	g.adjacencyList.AddEdgeBi(e)
	g.weights[e] = w
	g.weights[edge{e.End, e.Start}] = w
}

func (g *adjacencyListWithWeight) Weight(e edge) int {
	if value, ok := g.weights[e]; ok {
		return value
	}
	return -1
}

func newAdjacencyListWithWeight() *adjacencyListWithWeight {
	return new(adjacencyListWithWeight).Init()
}

func mstKruskal(g graphWeightily) graph {
	t := createGraphByType(g.GetGraph())
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
			t.AddEdgeBi(edge{e.Start.(*dfsElement).V, e.End.(*dfsElement).V})
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

func mstPrim(g graphWeightily) graph {
	t := createGraphByType(g.GetGraph())
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
			if element, ok := elements[e]; ok && g.Weight(edge{v, e}) < element.Key.(int) {
				p[e] = v
				pq.ModifyNode(element, g.Weight(edge{v, e}), element.Value)
			}
		}
	}

	for i := range p {
		t.AddEdgeBi(edge{p[i], i})
	}

	return t
}
