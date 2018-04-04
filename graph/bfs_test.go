package graph

import (
	"testing"
	"reflect"
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

func bfsGolden(g Graph)(bfsGraph Graph) {
	if _, isList := g.GetGraph().(*AdjacencyList); isList {
		bfsGraph = NewAdjacencyList()
	} else {
		bfsGraph = NewAdjacencyMatrix()
	}
	vertexes := make(map[interface{}]*BFSElement)
	vertexes["s"] = NewBFSElement("s")
	vertexes["s"].Dist = 0

	vertexes["r"] = NewBFSElement("r")
	vertexes["r"].Dist = 1
	vertexes["r"].P = vertexes["s"]

	vertexes["v"] = NewBFSElement("v")
	vertexes["v"].Dist = 2
	vertexes["v"].P = vertexes["r"]

	vertexes["w"] = NewBFSElement("w")
	vertexes["w"].Dist = 1
	vertexes["w"].P = vertexes["s"]

	vertexes["t"] = NewBFSElement("t")
	vertexes["t"].Dist = 2
	vertexes["t"].P = vertexes["w"]

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
			bfsGraph.AddEdge(Edge{vertexes[v], vertexes[v].P})
		}
	}

	return
}

func checkBFSGraph(t *testing.T, g Graph, gGloden Graph) {
	edges := g.AllEdges()
	vertexes := g.AllVertices()
	sort.Slice(edges, func(i, j int) bool {
		if edges[i].Start == edges[j].Start {
			return edges[i].End.(*BFSElement).V.(string) < edges[j].End.(*BFSElement).V.(string)
		}
		return edges[i].Start.(*BFSElement).V.(string) < edges[j].Start.(*BFSElement).V.(string)
	})
	sort.Slice(vertexes, func(i, j int) bool {
		return vertexes[i].(*BFSElement).V.(string) < vertexes[j].(*BFSElement).V.(string)
	})

	expEdges := gGloden.AllEdges()
	expVertices := gGloden.AllVertices()

	sort.Slice(expEdges, func(i, j int) bool {
		if expEdges[i].Start == expEdges[j].Start {
			return expEdges[i].End.(*BFSElement).V.(string) < expEdges[j].End.(*BFSElement).V.(string)
		}
		return expEdges[i].Start.(*BFSElement).V.(string) < expEdges[j].Start.(*BFSElement).V.(string)
	})
	sort.Slice(expVertices, func(i, j int) bool {
		return expVertices[i].(*BFSElement).V.(string) < expVertices[j].(*BFSElement).V.(string)
	})

	if !reflect.DeepEqual(edges, expEdges) {
		t.Log("get edges error!")
		for i := range expEdges {
			if !reflect.DeepEqual(expEdges[i], edges[i]) {
				t.Log("expect:")
				t.Log(expEdges[i])
				t.Log(expEdges[i].Start)
				t.Log(expEdges[i].End)
				t.Log("but get:")
				t.Log(edges[i])
				t.Log(edges[i].Start)
				t.Log(edges[i].End)
			}
		}

		t.Fail()
	}
	if !reflect.DeepEqual(vertexes, expVertices) {
		t.Log("get vertexes error!")
		for i := range expVertices {
			if !reflect.DeepEqual(expVertices[i], vertexes[i]) {
				t.Log("expect:")
				t.Log(expVertices[i])
				t.Log("but get:")
				t.Log(vertexes[i])
			}

		}
		t.Fail()
	}
}

func TestBFS(t *testing.T) {
	g := NewAdjacencyList()
	bfsSetupGraph(g)
	bfsGraph := BFS(g, "s")
	expBfsGraph := bfsGolden(g)
	checkBFSGraph(t, bfsGraph, expBfsGraph)
}
