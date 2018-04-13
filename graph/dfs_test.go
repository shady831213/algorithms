package graph

import (
	"sort"
	"testing"
)

func dfsSetupGraph(g graph) {
	g.AddVertex("u")
	g.AddVertex("v")
	g.AddVertex("w")
	g.AddVertex("x")
	g.AddVertex("y")
	g.AddVertex("z")
	g.AddEdge(edge{"u", "v"})
	g.AddEdge(edge{"u", "x"})
	g.AddEdge(edge{"v", "y"})
	g.AddEdge(edge{"x", "v"})
	g.AddEdge(edge{"y", "x"})
	g.AddEdge(edge{"w", "y"})
	g.AddEdge(edge{"w", "z"})
	g.AddEdge(edge{"z", "z"})
}

func dfsGolden(g graph) *dfsForest {
	dfsForest := newDFSForest(g)
	vertexes := make(map[interface{}]*dfsElement)

	vertexes["u"] = newDFSElement("u")
	vertexes["u"].D = 1
	vertexes["u"].F = 8
	vertexes["u"].P = nil
	vertexes["u"].Root = vertexes["u"]

	vertexes["v"] = newDFSElement("v")
	vertexes["v"].D = 2
	vertexes["v"].F = 7
	vertexes["v"].P = vertexes["u"]
	vertexes["v"].Root = vertexes["u"]

	vertexes["y"] = newDFSElement("y")
	vertexes["y"].D = 3
	vertexes["y"].F = 6
	vertexes["y"].P = vertexes["v"]
	vertexes["y"].Root = vertexes["u"]

	vertexes["x"] = newDFSElement("x")
	vertexes["x"].D = 4
	vertexes["x"].F = 5
	vertexes["x"].P = vertexes["y"]
	vertexes["x"].Root = vertexes["u"]

	vertexes["w"] = newDFSElement("w")
	vertexes["w"].D = 9
	vertexes["w"].F = 12
	vertexes["w"].P = nil
	vertexes["w"].Root = vertexes["w"]

	vertexes["z"] = newDFSElement("z")
	vertexes["z"].D = 10
	vertexes["z"].F = 11
	vertexes["z"].P = vertexes["w"]
	vertexes["z"].Root = vertexes["w"]

	for v := range vertexes {
		vertexes[v].Color = black
		dfsForest.AddVertex(vertexes[v])
	}

	for v := range vertexes {
		if vertexes[v].P != nil {
			dfsForest.AddTreeEdge(edge{vertexes[v].P, vertexes[v]})
		}
	}

	dfsForest.AddBackEdge(edge{vertexes["x"], vertexes["v"]})
	dfsForest.AddBackEdge(edge{vertexes["z"], vertexes["z"]})

	dfsForest.AddForwardEdge(edge{vertexes["u"], vertexes["x"]})

	dfsForest.AddCrossEdge(edge{vertexes["w"], vertexes["y"]})
	return dfsForest
}

func TestDFS(t *testing.T) {
	g := newAdjacencyList()
	dfsSetupGraph(g)
	dfsGraph := dfs(g, func(vertices []interface{}) {
		sort.Slice(vertices, func(i, j int) bool {
			return vertices[i].(string) < vertices[j].(string)
		})
	})
	expDfsGraph := dfsGolden(g)
	for _, v := range dfsGraph.Trees.AllVertices() {
		t.Log(v)
	}
	checkDFSGraphOutOfOrder(t, dfsGraph.Trees, expDfsGraph.Trees)
	checkDFSGraphOutOfOrder(t, dfsGraph.BackEdges, expDfsGraph.BackEdges)
	checkDFSGraphOutOfOrder(t, dfsGraph.ForwardEdges, expDfsGraph.ForwardEdges)
	checkDFSGraphOutOfOrder(t, dfsGraph.CrossEdges, expDfsGraph.CrossEdges)
}
