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

func dfsGolden(g Graph) (dfsGraph map[string]Graph) {
	dfsGraph = make(map[string]Graph)
	dfsGraph["dfsForest"] = CreateGraphByType(g)
	dfsGraph["dfsBackEdges"] = CreateGraphByType(g)
	dfsGraph["dfsForwardEdges"] = CreateGraphByType(g)
	dfsGraph["dfsCrossEdges"] = CreateGraphByType(g)
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
		for i := range dfsGraph {
			dfsGraph[i].AddVertex(vertexes[v])
		}
	}

	for v := range vertexes {
		if vertexes[v].P != nil {
			dfsGraph["dfsForest"].AddEdge(Edge{vertexes[v].P, vertexes[v]})
		}
	}

	dfsGraph["dfsBackEdges"].AddEdge(Edge{vertexes["y"], vertexes["x"]})
	dfsGraph["dfsBackEdges"].AddEdge(Edge{vertexes["z"], vertexes["z"]})

	dfsGraph["dfsForwardEdges"].AddEdge(Edge{vertexes["u"], vertexes["v"]})

	dfsGraph["dfsCrossEdges"].AddEdge(Edge{vertexes["w"], vertexes["y"]})
	return
}

func compareDFSGraph(t *testing.T, v, vExp []interface{}, e, eExp []Edge) {
	if !reflect.DeepEqual(e, eExp) {
		t.Log("get edges error!")
		for i := range eExp {
			if !reflect.DeepEqual(eExp[i], e[i]) {
				t.Log("expect:")
				t.Log(eExp[i])
				t.Log(eExp[i].Start)
				t.Log(eExp[i].End)
				t.Log("but get:")
				t.Log(eExp[i])
				t.Log(eExp[i].Start)
				t.Log(eExp[i].End)
			}
		}

		t.Fail()
	}
	if !reflect.DeepEqual(v, vExp) {
		t.Log("get vertexes error!")
		for i := range vExp {
			if !reflect.DeepEqual(vExp[i], v[i]) {
				t.Log("expect:")
				t.Log(vExp[i])
				t.Log("but get:")
				t.Log(v[i])
			}

		}
		t.Fail()
	}
}

//dfs Forest keep vertices order
func checkDFSForestGraph(t *testing.T, g Graph, gGloden Graph) {
	edges := g.AllEdges()
	//finish time increase order
	vertexes := g.AllVertices()
	sort.Slice(edges, func(i, j int) bool {
		if edges[i].Start == edges[j].Start {
			return edges[i].End.(*DFSElement).V.(string) < edges[j].End.(*DFSElement).V.(string)
		}
		return edges[i].Start.(*DFSElement).V.(string) < edges[j].Start.(*DFSElement).V.(string)
	})

	expEdges := gGloden.AllEdges()
	expVertices := gGloden.AllVertices()

	sort.Slice(expEdges, func(i, j int) bool {
		if expEdges[i].Start == expEdges[j].Start {
			return expEdges[i].End.(*DFSElement).V.(string) < expEdges[j].End.(*DFSElement).V.(string)
		}
		return expEdges[i].Start.(*DFSElement).V.(string) < expEdges[j].Start.(*DFSElement).V.(string)
	})
	//finish time increase order
	sort.Slice(expVertices, func(i, j int) bool {
		return expVertices[i].(*DFSElement).F < expVertices[j].(*DFSElement).F
	})
	compareDFSGraph(t, vertexes,expVertices,edges,expEdges)
}

//Edges graphs do not keep vertices order
func checkDFSEdgesGraph(t *testing.T, g Graph, gGloden Graph) {
	edges := g.AllEdges()
	//finish time increase order
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

	compareDFSGraph(t, vertexes,expVertices,edges,expEdges)
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
	for i := range dfsGraph {
		if i == "dfsForest" {
			checkDFSForestGraph(t, dfsGraph[i], expDfsGraph[i])
		} else {
			checkDFSEdgesGraph(t, dfsGraph[i], expDfsGraph[i])
		}
	}
}
