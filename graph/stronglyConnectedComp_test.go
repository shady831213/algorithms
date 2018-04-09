package graph

import (
	"testing"
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
	bea := CreateGraphByType(g)
	bea.AddVertex("a")
	bea.AddVertex("e")
	bea.AddVertex("b")
	bea.AddEdge(Edge{"a", "b"})
	bea.AddEdge(Edge{"b", "e"})
	bea.AddEdge(Edge{"e", "a"})
	scc.AddVertex(bea)
	cd := CreateGraphByType(g)
	cd.AddVertex("c")
	cd.AddVertex("d")
	cd.AddEdge(Edge{"c", "d"})
	cd.AddEdge(Edge{"d", "c"})
	scc.AddVertex(cd)
	gf := CreateGraphByType(g)
	gf.AddVertex("f")
	gf.AddVertex("g")
	gf.AddEdge(Edge{"f", "g"})
	gf.AddEdge(Edge{"g", "f"})
	scc.AddVertex(gf)
	h := CreateGraphByType(g)
	h.AddVertex("h")
	h.AddEdge(Edge{"h", "h"})
	scc.AddVertex(h)

	scc.AddEdge(Edge{bea, cd})
	scc.AddEdge(Edge{bea, gf})
	scc.AddEdge(Edge{cd, gf})
	scc.AddEdge(Edge{cd, h})
	scc.AddEdge(Edge{gf, h})

	return
}

func TestSCC(t *testing.T) {
	g := NewAdjacencyList()
	sccSetupGraph(g)
	scc := SCC(g)
	expScc := sccGolden(g)
	checkGraphInOrder(t, scc, expScc)
}
