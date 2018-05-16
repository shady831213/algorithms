package graph

import (
	"testing"
)

func bccSetupGraph() graph {
	g := newGraph()
	for i := 0; i < 23; i++ {
		g.AddVertex(i)
	}

	g.AddEdgeBi(edge{0, 1})
	g.AddEdgeBi(edge{1, 2})
	g.AddEdgeBi(edge{2, 3})
	g.AddEdgeBi(edge{3, 0})
	g.AddEdgeBi(edge{0, 4})
	g.AddEdgeBi(edge{4, 5})
	g.AddEdgeBi(edge{4, 6})
	g.AddEdgeBi(edge{5, 6})
	g.AddEdgeBi(edge{4, 7})
	g.AddEdgeBi(edge{4, 8})
	g.AddEdgeBi(edge{7, 8})
	g.AddEdgeBi(edge{7, 9})
	g.AddEdgeBi(edge{4, 10})
	g.AddEdgeBi(edge{10, 11})
	g.AddEdgeBi(edge{11, 12})
	g.AddEdgeBi(edge{12, 13})
	g.AddEdgeBi(edge{13, 10})
	g.AddEdgeBi(edge{12, 14})
	g.AddEdgeBi(edge{14, 15})
	g.AddEdgeBi(edge{15, 16})
	g.AddEdgeBi(edge{16, 17})
	g.AddEdgeBi(edge{17, 14})
	g.AddEdgeBi(edge{14, 18})
	g.AddEdgeBi(edge{18, 16})
	g.AddEdgeBi(edge{17, 19})
	g.AddEdgeBi(edge{19, 20})
	g.AddEdgeBi(edge{20, 21})
	g.AddEdgeBi(edge{21, 19})
	g.AddEdgeBi(edge{17, 22})
	return g
}

func vertexBCCGolden() (cuts graph, comps []graph) {
	cuts = newGraph()
	comps = make([]graph, 12, 12)
	for i := range comps {
		comps[i] = newGraph()
	}

	cuts.AddVertex(19)
	cuts.AddVertex(17)
	cuts.AddVertex(14)
	cuts.AddVertex(12)
	cuts.AddVertex(10)
	cuts.AddVertex(7)
	cuts.AddVertex(4)
	cuts.AddVertex(0)

	comps[0].AddEdgeBi(edge{0, 1})
	comps[0].AddEdgeBi(edge{1, 2})
	comps[0].AddEdgeBi(edge{2, 3})
	comps[0].AddEdgeBi(edge{3, 0})
	comps[11].AddEdgeBi(edge{0, 4})
	comps[1].AddEdgeBi(edge{4, 5})
	comps[1].AddEdgeBi(edge{4, 6})
	comps[1].AddEdgeBi(edge{5, 6})
	comps[3].AddEdgeBi(edge{4, 7})
	comps[3].AddEdgeBi(edge{4, 8})
	comps[3].AddEdgeBi(edge{7, 8})
	comps[2].AddEdgeBi(edge{7, 9})
	comps[10].AddEdgeBi(edge{4, 10})
	comps[9].AddEdgeBi(edge{10, 11})
	comps[9].AddEdgeBi(edge{11, 12})
	comps[9].AddEdgeBi(edge{12, 13})
	comps[9].AddEdgeBi(edge{13, 10})
	comps[8].AddEdgeBi(edge{12, 14})
	comps[7].AddEdgeBi(edge{14, 15})
	comps[7].AddEdgeBi(edge{15, 16})
	comps[7].AddEdgeBi(edge{16, 17})
	comps[7].AddEdgeBi(edge{17, 14})
	comps[7].AddEdgeBi(edge{14, 18})
	comps[7].AddEdgeBi(edge{18, 16})
	comps[5].AddEdgeBi(edge{17, 19})
	comps[4].AddEdgeBi(edge{19, 20})
	comps[4].AddEdgeBi(edge{20, 21})
	comps[4].AddEdgeBi(edge{21, 19})
	comps[6].AddEdgeBi(edge{17, 22})
	return
}

func edgeBCCGolden() (bridges graph, comps []graph) {
	bridges = newGraph()
	comps = make([]graph, 6, 6)
	for i := range comps {
		comps[i] = newGraph()
	}

	bridges.AddEdgeBi(edge{19, 17})
	bridges.AddEdgeBi(edge{17, 22})
	bridges.AddEdgeBi(edge{12, 14})
	bridges.AddEdgeBi(edge{10, 4})
	bridges.AddEdgeBi(edge{7, 9})
	bridges.AddEdgeBi(edge{0, 4})

	comps[0].AddEdgeBi(edge{0, 1})
	comps[0].AddEdgeBi(edge{1, 2})
	comps[0].AddEdgeBi(edge{2, 3})
	comps[0].AddEdgeBi(edge{3, 0})
	comps[1].AddEdgeBi(edge{4, 5})
	comps[1].AddEdgeBi(edge{4, 6})
	comps[1].AddEdgeBi(edge{5, 6})
	comps[2].AddEdgeBi(edge{4, 7})
	comps[2].AddEdgeBi(edge{4, 8})
	comps[2].AddEdgeBi(edge{7, 8})
	comps[5].AddEdgeBi(edge{10, 11})
	comps[5].AddEdgeBi(edge{11, 12})
	comps[5].AddEdgeBi(edge{12, 13})
	comps[5].AddEdgeBi(edge{13, 10})
	comps[4].AddEdgeBi(edge{14, 15})
	comps[4].AddEdgeBi(edge{15, 16})
	comps[4].AddEdgeBi(edge{16, 17})
	comps[4].AddEdgeBi(edge{17, 14})
	comps[4].AddEdgeBi(edge{14, 18})
	comps[4].AddEdgeBi(edge{18, 16})
	comps[3].AddEdgeBi(edge{19, 20})
	comps[3].AddEdgeBi(edge{20, 21})
	comps[3].AddEdgeBi(edge{21, 19})
	return
}

func TestVertexBCC(t *testing.T) {
	g := bccSetupGraph()
	cuts, comps := vertexBCC(g)
	cutsExp, compsExp := vertexBCCGolden()
	checkGraphOutOfOrderInInt(t, cuts, cutsExp, nil)
	for i := range comps {
		checkGraphOutOfOrderInInt(t, comps[i], compsExp[i], nil)
	}
}

func TestEdgeBCC(t *testing.T) {
	g := bccSetupGraph()
	bridges, comps := edgeBCC(g)
	bridgesExp, compsExp := edgeBCCGolden()
	checkGraphOutOfOrderInInt(t, bridges, bridgesExp, nil)
	for i := range comps {
		checkGraphOutOfOrderInInt(t, comps[i], compsExp[i], nil)
	}
}
