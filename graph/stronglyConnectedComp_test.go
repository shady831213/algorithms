package graph

import (
	"testing"
)

func sccSetupGraph(g graph) {
	g.AddVertex("a")
	g.AddVertex("b")
	g.AddVertex("c")
	g.AddVertex("d")
	g.AddVertex("e")
	g.AddVertex("f")
	g.AddVertex("g")
	g.AddVertex("h")
	g.AddEdge(edge{"a", "b"})
	g.AddEdge(edge{"b", "e"})
	g.AddEdge(edge{"e", "a"})
	g.AddEdge(edge{"e", "f"})
	g.AddEdge(edge{"b", "f"})
	g.AddEdge(edge{"b", "c"})
	g.AddEdge(edge{"c", "d"})
	g.AddEdge(edge{"d", "c"})
	g.AddEdge(edge{"c", "g"})
	g.AddEdge(edge{"f", "g"})
	g.AddEdge(edge{"g", "f"})
	g.AddEdge(edge{"g", "h"})
	g.AddEdge(edge{"d", "h"})
	g.AddEdge(edge{"h", "h"})
}

func sccGolden(g graph) (scc graph) {
	scc = createGraphByType(g)
	bea := createGraphByType(g)
	bea.AddVertex("a")
	bea.AddVertex("e")
	bea.AddVertex("b")
	bea.AddEdge(edge{"a", "b"})
	bea.AddEdge(edge{"b", "e"})
	bea.AddEdge(edge{"e", "a"})
	scc.AddVertex(bea)
	cd := createGraphByType(g)
	cd.AddVertex("c")
	cd.AddVertex("d")
	cd.AddEdge(edge{"c", "d"})
	cd.AddEdge(edge{"d", "c"})
	scc.AddVertex(cd)
	gf := createGraphByType(g)
	gf.AddVertex("f")
	gf.AddVertex("g")
	gf.AddEdge(edge{"f", "g"})
	gf.AddEdge(edge{"g", "f"})
	scc.AddVertex(gf)
	h := createGraphByType(g)
	h.AddVertex("h")
	h.AddEdge(edge{"h", "h"})
	scc.AddVertex(h)

	scc.AddEdge(edge{bea, cd})
	scc.AddEdge(edge{bea, gf})
	scc.AddEdge(edge{cd, gf})
	scc.AddEdge(edge{cd, h})
	scc.AddEdge(edge{gf, h})

	return
}

func TestSCC(t *testing.T) {
	g := newAdjacencyMatrix()
	sccSetupGraph(g)
	scc := scc(g)
	expScc := sccGolden(g)
	checkGraphInOrder(t, scc, expScc)
}
