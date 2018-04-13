package graph

import (
	"sort"
	"testing"
)

func bccSetupGraph(g graph) {
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
}

func vertexBCCGolden(g graph) (cuts graph, comps []graph) {
	cuts = createGraphByType(g)
	comps = make([]graph, 12, 12)
	for i := range comps {
		comps[i] = createGraphByType(g)
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

func edgeBCCGolden(g graph) (bridges graph, comps []graph) {
	bridges = createGraphByType(g)
	comps = make([]graph, 6, 6)
	for i := range comps {
		comps[i] = createGraphByType(g)
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

func checkBCCGraphOutOfOrder(t *testing.T, g graph, gGloden graph) {
	edges := g.AllEdges()
	//finish time increase order
	vertexes := g.AllVertices()
	sort.Slice(edges, func(i, j int) bool {
		if edges[i].Start.(int) == edges[j].Start.(int) {
			return edges[i].End.(int) < edges[j].End.(int)
		}
		return edges[i].Start.(int) < edges[j].Start.(int)
	})

	sort.Slice(vertexes, func(i, j int) bool {
		return vertexes[i].(int) < vertexes[j].(int)
	})

	expEdges := gGloden.AllEdges()
	expVertices := gGloden.AllVertices()

	sort.Slice(expEdges, func(i, j int) bool {
		if expEdges[i].Start.(int) == expEdges[j].Start.(int) {
			return expEdges[i].End.(int) < expEdges[j].End.(int)
		}
		return expEdges[i].Start.(int) < expEdges[j].Start.(int)
	})

	sort.Slice(expVertices, func(i, j int) bool {
		return expVertices[i].(int) < expVertices[j].(int)
	})

	compareGraph(t, vertexes, expVertices, edges, expEdges)
}

func TestVertexBCC(t *testing.T) {
	g := newAdjacencyList()
	bccSetupGraph(g)
	cuts, comps := vertexBCC(g)
	cutsExp, compsExp := vertexBCCGolden(g)
	checkBCCGraphOutOfOrder(t, cuts, cutsExp)
	for i := range comps {
		checkBCCGraphOutOfOrder(t, comps[i], compsExp[i])
	}
}

func TestEdgeBCC(t *testing.T) {
	g := newAdjacencyList()
	bccSetupGraph(g)
	bridges, comps := edgeBCC(g)
	bridgesExp, compsExp := edgeBCCGolden(g)
	checkBCCGraphOutOfOrder(t, bridges, bridgesExp)
	for i := range comps {
		checkBCCGraphOutOfOrder(t, comps[i], compsExp[i])
	}
	//for _, v := range bridges.AllEdges() {
	//	t.Log(v)
	//	t.Log(v.Start)
	//	t.Log(v.End)
	//}
	//for i := range comps {
	//	t.Log("comps ", i, "vertices:")
	//	for _, v := range comps[i].AllVertices() {
	//		t.Log(v)
	//	}
	//	t.Log("comps ", i, "edges:")
	//	for _, v := range comps[i].AllEdges() {
	//		t.Log(v)
	//		t.Log(v.Start)
	//		t.Log(v.End)
	//	}
	//}
}
