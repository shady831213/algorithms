package graph

import (
	"sort"
	"testing"
)

func dfsSetupGraph(g Graph) {
	g.AddVertex("u")
	g.AddVertex("v")
	g.AddVertex("w")
	g.AddVertex("x")
	g.AddVertex("y")
	g.AddVertex("z")
	g.AddEdge(Edge{"u", "v"})
	g.AddEdge(Edge{"u", "x"})
	g.AddEdge(Edge{"v", "y"})
	g.AddEdge(Edge{"x", "v"})
	g.AddEdge(Edge{"y", "x"})
	g.AddEdge(Edge{"w", "y"})
	g.AddEdge(Edge{"w", "z"})
	g.AddEdge(Edge{"z", "z"})
}

func dfsGolden(g Graph) *DFSForest {
	dfsForest := NewDFSForest(g)
	vertexes := make(map[interface{}]*DFSElement)

	vertexes["u"] = NewDFSElement("u")
	vertexes["u"].D = 1
	vertexes["u"].F = 8
	vertexes["u"].P = nil
	vertexes["u"].Root = vertexes["u"]

	vertexes["v"] = NewDFSElement("v")
	vertexes["v"].D = 2
	vertexes["v"].F = 7
	vertexes["v"].P = vertexes["u"]
	vertexes["v"].Root = vertexes["u"]

	vertexes["y"] = NewDFSElement("y")
	vertexes["y"].D = 3
	vertexes["y"].F = 6
	vertexes["y"].P = vertexes["v"]
	vertexes["y"].Root = vertexes["u"]

	vertexes["x"] = NewDFSElement("x")
	vertexes["x"].D = 4
	vertexes["x"].F = 5
	vertexes["x"].P = vertexes["y"]
	vertexes["x"].Root = vertexes["u"]

	vertexes["w"] = NewDFSElement("w")
	vertexes["w"].D = 9
	vertexes["w"].F = 12
	vertexes["w"].P = nil
	vertexes["w"].Root = vertexes["w"]

	vertexes["z"] = NewDFSElement("z")
	vertexes["z"].D = 10
	vertexes["z"].F = 11
	vertexes["z"].P = vertexes["w"]
	vertexes["z"].Root = vertexes["w"]

	for v := range vertexes {
		vertexes[v].Color = BLACK
		dfsForest.AddVertex(vertexes[v])
	}

	for v := range vertexes {
		if vertexes[v].P != nil {
			dfsForest.AddTreeEdge(Edge{vertexes[v].P, vertexes[v]})
		}
	}

	dfsForest.AddBackEdge(Edge{vertexes["x"], vertexes["v"]})
	dfsForest.AddBackEdge(Edge{vertexes["z"], vertexes["z"]})

	dfsForest.AddForwardEdge(Edge{vertexes["u"], vertexes["x"]})

	dfsForest.AddCrossEdge(Edge{vertexes["w"], vertexes["y"]})
	return dfsForest
}

func TestDFS(t *testing.T) {
	g := NewAdjacencyList()
	dfsSetupGraph(g)
	dfsGraph := DFS(g, func(vertices []interface{}) {
		sort.Slice(vertices, func(i, j int) bool {
			return vertices[i].(string) < vertices[j].(string)
		})
	})
	expDfsGraph := dfsGolden(g)
	checkDFSGraphOutOfOrder(t, dfsGraph.Trees, expDfsGraph.Trees)
	checkDFSGraphOutOfOrder(t, dfsGraph.BackEdges, expDfsGraph.BackEdges)
	checkDFSGraphOutOfOrder(t, dfsGraph.ForwardEdges, expDfsGraph.ForwardEdges)
	checkDFSGraphOutOfOrder(t, dfsGraph.CrossEdges, expDfsGraph.CrossEdges)
}
