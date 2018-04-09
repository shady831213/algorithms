package graph

import (
	"testing"
	"sort"
)

func bfsSetupGraph(g Graph) {
	g.AddVertex("r")
	g.AddVertex("s")
	g.AddVertex("t")
	g.AddVertex("u")
	g.AddVertex("v")
	g.AddVertex("w")
	g.AddVertex("x")
	g.AddVertex("y")
	g.AddEdgeBi(Edge{"r", "s"})
	g.AddEdgeBi(Edge{"s", "w"})
	g.AddEdgeBi(Edge{"r", "v"})
	g.AddEdgeBi(Edge{"w", "t"})
	g.AddEdgeBi(Edge{"w", "x"})
	g.AddEdgeBi(Edge{"t", "x"})
	g.AddEdgeBi(Edge{"t", "u"})
	g.AddEdgeBi(Edge{"x", "u"})
	g.AddEdgeBi(Edge{"x", "y"})
	g.AddEdgeBi(Edge{"u", "y"})
}

func bfsGolden(g Graph) (bfsGraph Graph) {
	bfsGraph = CreateGraphByType(g)
	vertexes := make(map[interface{}]*BFSElement)
	vertexes["s"] = NewBFSElement("s")
	vertexes["s"].Dist = 0

	vertexes["r"] = NewBFSElement("r")
	vertexes["r"].Dist = 1
	vertexes["r"].P = vertexes["s"]

	vertexes["w"] = NewBFSElement("w")
	vertexes["w"].Dist = 1
	vertexes["w"].P = vertexes["s"]

	vertexes["t"] = NewBFSElement("t")
	vertexes["t"].Dist = 2
	vertexes["t"].P = vertexes["w"]

	vertexes["v"] = NewBFSElement("v")
	vertexes["v"].Dist = 2
	vertexes["v"].P = vertexes["r"]

	vertexes["x"] = NewBFSElement("x")
	vertexes["x"].Dist = 2
	vertexes["x"].P = vertexes["w"]

	vertexes["u"] = NewBFSElement("u")
	vertexes["u"].Dist = 3
	vertexes["u"].P = vertexes["t"]

	vertexes["y"] = NewBFSElement("y")
	vertexes["y"].Dist = 3
	vertexes["y"].P = vertexes["x"]

	for v := range vertexes {
		vertexes[v].Color = BLACK
		bfsGraph.AddVertex(vertexes[v])
		if vertexes[v].P != nil {
			bfsGraph.AddEdge(Edge{vertexes[v].P, vertexes[v]})
		}
	}

	return
}

func checkBFSGraphOutOfOrder(t *testing.T, g Graph, gGloden Graph) {
	edges := g.AllEdges()
	//finish time increase order
	vertexes := g.AllVertices()
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].End.(*BFSElement).V.(string) < edges[j].End.(*BFSElement).V.(string)
	})

	sort.Slice(vertexes, func(i, j int) bool {
		return vertexes[i].(*BFSElement).V.(string) < vertexes[j].(*BFSElement).V.(string)
	})

	expEdges := gGloden.AllEdges()
	expVertices := gGloden.AllVertices()

	sort.Slice(expEdges, func(i, j int) bool {
		return expEdges[i].End.(*BFSElement).V.(string) < expEdges[j].End.(*BFSElement).V.(string)
	})

	sort.Slice(expVertices, func(i, j int) bool {
		return expVertices[i].(*BFSElement).V.(string) < expVertices[j].(*BFSElement).V.(string)
	})

	compareGraph(t, vertexes, expVertices, edges, expEdges)
}

func TestBFS(t *testing.T) {
	g := NewAdjacencyList()
	bfsSetupGraph(g)
	bfsGraph := BFS(g, "s")
	expBfsGraph := bfsGolden(g)
	checkBFSGraphOutOfOrder(t, bfsGraph, expBfsGraph)
}
