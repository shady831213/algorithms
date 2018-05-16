package graph

import (
	"sort"
	"testing"
)

func bfsSetupGraph(g graph) {
	g.AddVertex("r")
	g.AddVertex("s")
	g.AddVertex("t")
	g.AddVertex("u")
	g.AddVertex("v")
	g.AddVertex("w")
	g.AddVertex("x")
	g.AddVertex("y")
	g.AddEdgeBi(edge{"r", "s"})
	g.AddEdgeBi(edge{"s", "w"})
	g.AddEdgeBi(edge{"r", "v"})
	g.AddEdgeBi(edge{"w", "t"})
	g.AddEdgeBi(edge{"w", "x"})
	g.AddEdgeBi(edge{"t", "x"})
	g.AddEdgeBi(edge{"t", "u"})
	g.AddEdgeBi(edge{"x", "u"})
	g.AddEdgeBi(edge{"x", "y"})
	g.AddEdgeBi(edge{"u", "y"})
}

func bfsGolden(g graph) (bfsGraph graph) {
	bfsGraph = newGraph()
	vertexes := make(map[interface{}]*bfsElement)
	vertexes["s"] = newBFSElement("s")
	vertexes["s"].Dist = 0
	vertexes["s"].Iter = g.IterConnectedVertices("s")

	vertexes["r"] = newBFSElement("r")
	vertexes["r"].Dist = 1
	vertexes["r"].P = vertexes["s"]
	vertexes["r"].Iter = g.IterConnectedVertices("r")

	vertexes["w"] = newBFSElement("w")
	vertexes["w"].Dist = 1
	vertexes["w"].P = vertexes["s"]
	vertexes["w"].Iter = g.IterConnectedVertices("w")

	vertexes["t"] = newBFSElement("t")
	vertexes["t"].Dist = 2
	vertexes["t"].P = vertexes["w"]
	vertexes["t"].Iter = g.IterConnectedVertices("t")

	vertexes["v"] = newBFSElement("v")
	vertexes["v"].Dist = 2
	vertexes["v"].P = vertexes["r"]
	vertexes["v"].Iter = g.IterConnectedVertices("v")

	vertexes["x"] = newBFSElement("x")
	vertexes["x"].Dist = 2
	vertexes["x"].P = vertexes["w"]
	vertexes["x"].Iter = g.IterConnectedVertices("x")

	vertexes["u"] = newBFSElement("u")
	vertexes["u"].Dist = 3
	vertexes["u"].P = vertexes["t"]
	vertexes["u"].Iter = g.IterConnectedVertices("u")

	vertexes["y"] = newBFSElement("y")
	vertexes["y"].Dist = 3
	vertexes["y"].P = vertexes["x"]
	vertexes["y"].Iter = g.IterConnectedVertices("y")

	for v := range vertexes {
		vertexes[v].Color = black
		bfsGraph.AddVertex(vertexes[v])
		if vertexes[v].P != nil {
			bfsGraph.AddEdge(edge{vertexes[v].P, vertexes[v]})
		}
	}

	return
}

func checkBFSGraphOutOfOrder(t *testing.T, g graph, gGolden graph) {
	edges := g.AllEdges()
	//finish time increase order
	vertexes := g.AllVertices()
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].End.(*bfsElement).V.(string) < edges[j].End.(*bfsElement).V.(string)
	})

	sort.Slice(vertexes, func(i, j int) bool {
		return vertexes[i].(*bfsElement).V.(string) < vertexes[j].(*bfsElement).V.(string)
	})

	expEdges := gGolden.AllEdges()
	expVertices := gGolden.AllVertices()

	sort.Slice(expEdges, func(i, j int) bool {
		return expEdges[i].End.(*bfsElement).V.(string) < expEdges[j].End.(*bfsElement).V.(string)
	})

	sort.Slice(expVertices, func(i, j int) bool {
		return expVertices[i].(*bfsElement).V.(string) < expVertices[j].(*bfsElement).V.(string)
	})

	compareGraph(t, vertexes, expVertices, edges, expEdges)
}

func TestBFS(t *testing.T) {
	g := newGraph()
	bfsSetupGraph(g)
	bfsGraph := bfs(g, "s")
	expBfsGraph := bfsGolden(g)
	checkBFSGraphOutOfOrder(t, bfsGraph, expBfsGraph)
}
