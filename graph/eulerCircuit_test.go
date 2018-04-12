package graph

import (
	"reflect"
	"testing"
)

func diEulerCircuitSetup(g graph) {
	g.AddEdge(edge{2, 3})
	g.AddEdge(edge{2, 5})
	g.AddEdge(edge{3, 4})
	g.AddEdge(edge{1, 2})
	g.AddEdge(edge{4, 2})
	g.AddEdge(edge{5, 1})
}

func diEulerCircuitGolden() []edge {
	return []edge{{2, 3}, {3, 4}, {4, 2}, {2, 5}, {5, 1}, {1, 2}}
}

func udEulerCircuitSetup(g graph) {

}

func TestDiEulerCircuitOK(t *testing.T) {
	g := newAdjacencyList()
	diEulerCircuitSetup(g)
	path := eulerCircuit(g, true)
	pathExp := diEulerCircuitGolden()
	if !reflect.DeepEqual(path, pathExp) {
		t.Log("expect:", pathExp, ", actual :", path)
		t.Fail()
	}
}
