package graph

import (
	"testing"
)

func flowGraphSetup() flowGraph {
	g := newFlowGraph()
	g.AddEdgeWithCap(edge{"s", "v1"}, 16)
	g.AddEdgeWithCap(edge{"s", "v2"}, 13)
	g.AddEdgeWithCap(edge{"v2", "v1"}, 4)
	g.AddEdgeWithCap(edge{"v1", "v2"}, 10)
	g.AddEdgeWithCap(edge{"v1", "v3"}, 12)
	g.AddEdgeWithCap(edge{"v3", "v2"}, 9)
	g.AddEdgeWithCap(edge{"v2", "v4"}, 14)
	g.AddEdgeWithCap(edge{"v4", "v3"}, 7)
	g.AddEdgeWithCap(edge{"v3", "t"}, 20)
	g.AddEdgeWithCap(edge{"v4", "t"}, 4)
	return g
}

func flowGraphGolden(g flowGraph) flowGraph {
	flowG := newFlowGraph()
	for _, e := range g.AllEdges() {
		flowG.AddEdgeWithCap(e, g.Cap(e))
	}
	flowG.AddEdgeWithFlow(edge{"s", "v1"}, 12)
	flowG.AddEdgeWithFlow(edge{"s", "v2"}, 11)
	flowG.AddEdgeWithFlow(edge{"v2", "v1"}, 0)
	flowG.AddEdgeWithFlow(edge{"v1", "v2"}, 0)
	flowG.AddEdgeWithFlow(edge{"v1", "v3"}, 12)
	flowG.AddEdgeWithFlow(edge{"v3", "v2"}, 0)
	flowG.AddEdgeWithFlow(edge{"v2", "v4"}, 11)
	flowG.AddEdgeWithFlow(edge{"v4", "v3"}, 7)
	flowG.AddEdgeWithFlow(edge{"v3", "t"}, 19)
	flowG.AddEdgeWithFlow(edge{"v4", "t"}, 4)
	return flowG
}

func pushRelabelGolden(g flowGraph) flowGraph {
	flowG := newFlowGraph()
	for _, e := range g.AllEdges() {
		flowG.AddEdgeWithCap(e, g.Cap(e))
	}
	flowG.AddEdgeWithFlow(edge{"s", "v1"}, 13)
	flowG.AddEdgeWithFlow(edge{"s", "v2"}, 10)
	flowG.AddEdgeWithFlow(edge{"v2", "v1"}, -1)
	flowG.AddEdgeWithFlow(edge{"v1", "v2"}, 1)
	flowG.AddEdgeWithFlow(edge{"v1", "v3"}, 12)
	flowG.AddEdgeWithFlow(edge{"v3", "v2"}, 0)
	flowG.AddEdgeWithFlow(edge{"v2", "v4"}, 11)
	flowG.AddEdgeWithFlow(edge{"v4", "v3"}, 7)
	flowG.AddEdgeWithFlow(edge{"v3", "t"}, 19)
	flowG.AddEdgeWithFlow(edge{"v4", "t"}, 4)
	return flowG
}

func bipGraphMaxMatchSetup() (graph, []interface{}) {
	g := newGraph()
	g.AddEdgeBi(edge{"l0", "r0"})
	g.AddEdgeBi(edge{"l1", "r0"})
	g.AddEdgeBi(edge{"l1", "r2"})
	g.AddEdgeBi(edge{"l2", "r1"})
	g.AddEdgeBi(edge{"l2", "r2"})
	g.AddEdgeBi(edge{"l2", "r3"})
	g.AddEdgeBi(edge{"l3", "r2"})
	g.AddEdgeBi(edge{"l4", "r2"})
	return g, []interface{}{"l0", "l1", "l2", "l3", "l4"}
}

func bipGraphMaxMatchGoldenByEdmondesKarp() graph {
	g := newGraph()
	g.AddEdgeBi(edge{"l0", "r0"})
	g.AddEdgeBi(edge{"l1", "r2"})
	g.AddEdgeBi(edge{"l2", "r1"})
	return g
}

func bipGraphMaxMatchGoldenByPushRelabel() graph {
	g := newGraph()
	g.AddEdgeBi(edge{"l1", "r0"})
	g.AddEdgeBi(edge{"l4", "r2"})
	g.AddEdgeBi(edge{"l2", "r1"})
	return g
}

func hopcraftKarpSetup() graph {
	g := newGraph()
	g.AddEdge(edge{"l0", "r0"})
	g.AddEdge(edge{"l1", "r0"})
	g.AddEdge(edge{"l1", "r2"})
	g.AddEdge(edge{"l2", "r1"})
	g.AddEdge(edge{"l2", "r2"})
	g.AddEdge(edge{"l2", "r3"})
	g.AddEdge(edge{"l3", "r2"})
	g.AddEdge(edge{"l4", "r2"})
	return g
}

func bipGraphMaxMatchGoldenByRelabelToFront() graph {
	g := newGraph()
	g.AddEdgeBi(edge{"l0", "r0"})
	g.AddEdgeBi(edge{"l1", "r2"})
	g.AddEdgeBi(edge{"l2", "r1"})
	return g
}

func checkFlowGraphOutOfOrder(t *testing.T, g, gGolden flowGraph) {
	comparator := func(t *testing.T, v, vExp []interface{}, e, eExp []edge) {
		for _, e := range eExp {
			if g.Cap(e) != gGolden.Cap(e) {
				t.Log(e, "Cap Error! exp :", gGolden.Cap(e), "actual: ", g.Cap(e))
				t.FailNow()
			}
			if g.Flow(e) != gGolden.Flow(e) {
				t.Log(e, "Flow Error! exp :", gGolden.Flow(e), "actual: ", g.Flow(e))
				t.FailNow()
			}
		}
	}
	checkGraphOutOfOrderInString(t, g, gGolden, comparator)
}

func TestEdmondesKarp(t *testing.T) {
	g := flowGraphSetup()
	edmondesKarp(g, "s", "t")
	gGolden := flowGraphGolden(g)
	checkFlowGraphOutOfOrder(t, g, gGolden)
}

func TestPushRelabel(t *testing.T) {
	g := flowGraphSetup()
	pushRelabel(g, "s", "t")
	gGolden := pushRelabelGolden(g)
	checkFlowGraphOutOfOrder(t, g, gGolden)
}

func TestRelabelToFront(t *testing.T) {
	g := flowGraphSetup()
	relabelToFront(g, "s", "t")
	gGolden := pushRelabelGolden(g)
	checkFlowGraphOutOfOrder(t, g, gGolden)
}

func TestBipGraphMaxMatch(t *testing.T) {
	bioG, l := bipGraphMaxMatchSetup()
	gGolden := bipGraphMaxMatchGoldenByEdmondesKarp()
	g := bipGraphMaxMatch(bioG, l, edmondesKarp)
	checkGraphOutOfOrderInString(t, g, gGolden, nil)

	g = bipGraphMaxMatch(bioG, l, pushRelabel)
	gGolden = bipGraphMaxMatchGoldenByPushRelabel()
	checkGraphOutOfOrderInString(t, g, gGolden, nil)

	g = bipGraphMaxMatch(bioG, l, relabelToFront)
	gGolden = bipGraphMaxMatchGoldenByRelabelToFront()
	checkGraphOutOfOrderInString(t, g, gGolden, nil)
}

func TestHopcraftKarp(t *testing.T) {
	matches := new(hopcraftKarp).init(hopcraftKarpSetup()).maxMatch()
	if matches != 3 {
		t.Log("exp : 3 but get :", matches)
		t.Fail()
	}
}
