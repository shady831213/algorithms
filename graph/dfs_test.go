package graph

import (
	"testing"
	"reflect"
	"sort"
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

func dfsGolden(g Graph) (dfsGraph Graph) {
	if _, isList := g.GetGraph().(*AdjacencyList); isList {
		dfsGraph = NewAdjacencyList()
	} else {
		dfsGraph = NewAdjacencyMatrix()
	}
	vertexes := make(map[interface{}]*DFSElement)

	vertexes["u"] = NewDFSElement("u")
	vertexes["u"].D = 1
	vertexes["u"].F = 8
	vertexes["u"].P = nil

	vertexes["x"] = NewDFSElement("x")
	vertexes["x"].D = 2
	vertexes["x"].F = 7
	vertexes["x"].P = vertexes["u"]

	vertexes["v"] = NewDFSElement("v")
	vertexes["v"].D = 3
	vertexes["v"].F = 6
	vertexes["v"].P = vertexes["x"]

	vertexes["y"] = NewDFSElement("y")
	vertexes["y"].D = 4
	vertexes["y"].F = 5
	vertexes["y"].P = vertexes["v"]

	vertexes["w"] = NewDFSElement("w")
	vertexes["w"].D = 9
	vertexes["w"].F = 12
	vertexes["w"].P = nil

	vertexes["z"] = NewDFSElement("z")
	vertexes["z"].D = 10
	vertexes["z"].F = 11
	vertexes["z"].P = vertexes["w"]

	for v := range vertexes {
		vertexes[v].Color = BLACK
		dfsGraph.AddVertex(vertexes[v])
		if vertexes[v].P != nil {
			dfsGraph.AddEdge(Edge{vertexes[v], vertexes[v].P})
		}
	}

	return
}

func checkDFSGraph(t *testing.T, g Graph, gGloden Graph) {
	edges := g.AllEdges()
	vertexes := g.AllVertices()
	sort.Slice(edges, func(i, j int) bool {
		if edges[i].Start == edges[j].Start {
			return edges[i].End.(*DFSElement).V.(string) < edges[j].End.(*DFSElement).V.(string)
		}
		return edges[i].Start.(*DFSElement).V.(string) < edges[j].Start.(*DFSElement).V.(string)
	})
	sort.Slice(vertexes, func(i, j int) bool {
		return vertexes[i].(*DFSElement).V.(string) < vertexes[j].(*DFSElement).V.(string)
	})

	expEdges := gGloden.AllEdges()
	expVertices := gGloden.AllVertices()

	sort.Slice(expEdges, func(i, j int) bool {
		if expEdges[i].Start == expEdges[j].Start {
			return expEdges[i].End.(*DFSElement).V.(string) < expEdges[j].End.(*DFSElement).V.(string)
		}
		return expEdges[i].Start.(*DFSElement).V.(string) < expEdges[j].Start.(*DFSElement).V.(string)
	})
	sort.Slice(expVertices, func(i, j int) bool {
		return expVertices[i].(*DFSElement).V.(string) < expVertices[j].(*DFSElement).V.(string)
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

func TestDFS(t *testing.T) {
	g := NewAdjacencyList()
	dfsSetupGraph(g)
	dfsGraph := DFS(g)
	expDfsGraph := dfsGolden(g)
	checkDFSGraph(t, dfsGraph, expDfsGraph)
}
