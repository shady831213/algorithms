package graph

import (
	"github.com/shady831213/algorithms/tree/disjointSetTree"
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
