package graph

import (
	"testing"
	"reflect"
)

func sccSetupGraph(g Graph) {
	g.AddVertex("a")
	g.AddVertex("b")
	g.AddVertex("c")
	g.AddVertex("d")
	g.AddVertex("e")
	g.AddVertex("f")
	g.AddVertex("g")
	g.AddVertex("h")
	g.AddEdge(Edge{"a", "b"})
	g.AddEdge(Edge{"b", "e"})
	g.AddEdge(Edge{"e", "a"})
	g.AddEdge(Edge{"e", "f"})
	g.AddEdge(Edge{"b", "f"})
	g.AddEdge(Edge{"b", "c"})
	g.AddEdge(Edge{"c", "d"})
	g.AddEdge(Edge{"d", "c"})
	g.AddEdge(Edge{"c", "g"})
	g.AddEdge(Edge{"f", "g"})
	g.AddEdge(Edge{"g", "f"})
	g.AddEdge(Edge{"g", "h"})
	g.AddEdge(Edge{"d", "h"})
	g.AddEdge(Edge{"h", "h"})
}

func sccGolden(g Graph) (scc Graph) {
	scc = CreateGraphByType(g)
	scc.AddVertex(&[]interface{}{"b", "e", "a"})
	scc.AddVertex(&[]interface{}{"c", "d"})
	scc.AddVertex(&[]interface{}{"g", "f"})
	scc.AddVertex(&[]interface{}{"h"})
	return
}

func checkSCCGraph(t *testing.T, g Graph, gGloden Graph) {
	edges := g.AllEdges()
	vertexes := g.AllVertices()

	expEdges := gGloden.AllEdges()
	expVertices := gGloden.AllVertices()

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

func TestSCC(t *testing.T) {
	g := NewAdjacencyList()
	sccSetupGraph(g)
	scc := SCC(g)
	expScc := sccGolden(g)
	checkSCCGraph(t, scc, expScc)
}
