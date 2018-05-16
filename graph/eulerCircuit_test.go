package graph

import (
	"reflect"
	"testing"
)

func diEulerCircuitSetup() graph {
	g := newGraph()
	g.AddEdge(edge{2, 3})
	g.AddEdge(edge{2, 5})
	g.AddEdge(edge{3, 4})
	g.AddEdge(edge{1, 2})
	g.AddEdge(edge{4, 2})
	g.AddEdge(edge{5, 1})
	return g
}

func eulerCircuitGolden() []edge {
	return []edge{{2, 3}, {3, 4}, {4, 2}, {2, 5}, {5, 1}, {1, 2}}
}

func udEulerCircuitSetup() graph {
	g := newGraph()
	g.AddEdgeBi(edge{2, 3})
	g.AddEdgeBi(edge{2, 5})
	g.AddEdgeBi(edge{3, 4})
	g.AddEdgeBi(edge{1, 2})
	g.AddEdgeBi(edge{4, 2})
	g.AddEdgeBi(edge{5, 1})
	return g
}

func TestDiEulerCircuitOK(t *testing.T) {
	g := diEulerCircuitSetup()
	for i := 0; i < 3; i++ {
		path := eulerCircuit(g)
		pathExp := eulerCircuitGolden()
		if !reflect.DeepEqual(path, pathExp) {
			t.Log("expect:", pathExp, ", actual :", path)
			t.Fail()
		}
	}
}

func TestDiEulerCircuitWithSingleVertex(t *testing.T) {
	g := diEulerCircuitSetup()
	g.AddVertex(6)
	path := eulerCircuit(g)
	if path != nil {
		t.Log("expect: nil", ", actual :", path)
		t.Fail()
	}
}

func TestDiEulerCircuitWithSingleVertexLoop(t *testing.T) {
	g := diEulerCircuitSetup()
	g.AddEdgeBi(edge{6, 6})
	path := eulerCircuit(g)
	if path != nil {
		t.Log("expect: nil", ", actual :", path)
		t.Fail()
	}
}

func TestDiEulerCircuitWithNonCircuit(t *testing.T) {
	g := diEulerCircuitSetup()
	g.AddEdge(edge{6, 1})
	path := eulerCircuit(g)
	if path != nil {
		t.Log("expect: nil", ", actual :", path)
		t.Fail()
	}

	g = diEulerCircuitSetup()
	g.AddEdge(edge{1, 6})
	path = eulerCircuit(g)
	if path != nil {
		t.Log("expect: nil", ", actual :", path)
		t.Fail()
	}
}

func TestUdEulerCircuitOK(t *testing.T) {
	g := udEulerCircuitSetup()
	path := eulerCircuit(g)
	pathExp := eulerCircuitGolden()
	if !reflect.DeepEqual(path, pathExp) {
		t.Log("expect:", pathExp, ", actual :", path)
		t.Fail()
	}
}

func TestUdEulerCircuitWithSingleVertex(t *testing.T) {
	g := udEulerCircuitSetup()
	g.AddVertex(6)
	path := eulerCircuit(g)
	if path != nil {
		t.Log("expect: nil", ", actual :", path)
		t.Fail()
	}
}

func TestUdEulerCircuitWithNonCircuit(t *testing.T) {
	g := udEulerCircuitSetup()
	g.AddEdgeBi(edge{1, 6})
	path := eulerCircuit(g)
	if path != nil {
		t.Log("expect: nil", ", actual :", path)
		t.Fail()
	}
}
