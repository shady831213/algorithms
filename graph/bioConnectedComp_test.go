package graph

import (
	"sort"
	"testing"
)

func bccSetupGraph(g Graph) {
	for i := 0; i < 23; i++ {
		g.AddVertex(i)
	}

	g.AddEdgeBi(Edge{0, 1})
	g.AddEdgeBi(Edge{1, 2})
	g.AddEdgeBi(Edge{2, 3})
	g.AddEdgeBi(Edge{3, 0})
	g.AddEdgeBi(Edge{0, 4})
	g.AddEdgeBi(Edge{4, 5})
	g.AddEdgeBi(Edge{4, 6})
	g.AddEdgeBi(Edge{5, 6})
	g.AddEdgeBi(Edge{4, 7})
	g.AddEdgeBi(Edge{4, 8})
	g.AddEdgeBi(Edge{7, 8})
	g.AddEdgeBi(Edge{7, 9})
	g.AddEdgeBi(Edge{4, 10})
	g.AddEdgeBi(Edge{10, 11})
	g.AddEdgeBi(Edge{11, 12})
	g.AddEdgeBi(Edge{12, 13})
	g.AddEdgeBi(Edge{13, 10})
	g.AddEdgeBi(Edge{12, 14})
	g.AddEdgeBi(Edge{14, 15})
	g.AddEdgeBi(Edge{15, 16})
	g.AddEdgeBi(Edge{16, 17})
	g.AddEdgeBi(Edge{17, 14})
	g.AddEdgeBi(Edge{14, 18})
	g.AddEdgeBi(Edge{18, 16})
	g.AddEdgeBi(Edge{17, 19})
	g.AddEdgeBi(Edge{19, 20})
	g.AddEdgeBi(Edge{20, 21})
	g.AddEdgeBi(Edge{21, 19})
	g.AddEdgeBi(Edge{17, 22})
}

func cutsAndBridgeGolden(dfsGraph Graph) (cuts, bridges Graph) {
	cuts, bridges = CreateGraphByType(dfsGraph), CreateGraphByType(dfsGraph)
	vertices := make(map[interface{}]*DFSElement)
	for _, v := range dfsGraph.AllVertices() {
		vertices[v.(*DFSElement).V] = v.(*DFSElement)
	}
	cuts.AddVertex(vertices[19])
	cuts.AddVertex(vertices[17])
	cuts.AddVertex(vertices[14])
	cuts.AddVertex(vertices[12])
	cuts.AddVertex(vertices[10])
	cuts.AddVertex(vertices[7])
	cuts.AddVertex(vertices[4])
	cuts.AddVertex(vertices[0])

	bridges.AddEdgeBi(Edge{vertices[19], vertices[17]})
	bridges.AddEdgeBi(Edge{vertices[17], vertices[22]})
	bridges.AddEdgeBi(Edge{vertices[12], vertices[14]})
	bridges.AddEdgeBi(Edge{vertices[10], vertices[4]})
	bridges.AddEdgeBi(Edge{vertices[7], vertices[9]})
	bridges.AddEdgeBi(Edge{vertices[0], vertices[4]})
	return
}

func TestGetCutsAndBridgesFromComponent(t *testing.T) {
	g := NewAdjacencyList()
	bccSetupGraph(g)
	dfsForest := DFS(g, nil)
	cuts, bridges := getCutsAndBridgesFromComponent(dfsForest)
	cutsExp, bridgesExp := cutsAndBridgeGolden(dfsForest.Trees)
	checkDFSGraphOutOfOrder(t, cuts, cutsExp)
	checkDFSGraphOutOfOrder(t, bridges, bridgesExp)
}

func vertexBCCGolden(g Graph) (cuts Graph, comps []Graph) {
	cuts = CreateGraphByType(g)
	comps = make([]Graph, 12, 12)
	for i := range comps {
		comps[i] = CreateGraphByType(g)
	}

	cuts.AddVertex(19)
	cuts.AddVertex(17)
	cuts.AddVertex(14)
	cuts.AddVertex(12)
	cuts.AddVertex(10)
	cuts.AddVertex(7)
	cuts.AddVertex(4)
	cuts.AddVertex(0)

	comps[0].AddEdgeBi(Edge{0, 1})
	comps[0].AddEdgeBi(Edge{1, 2})
	comps[0].AddEdgeBi(Edge{2, 3})
	comps[0].AddEdgeBi(Edge{3, 0})
	comps[11].AddEdgeBi(Edge{0, 4})
	comps[1].AddEdgeBi(Edge{4, 5})
	comps[1].AddEdgeBi(Edge{4, 6})
	comps[1].AddEdgeBi(Edge{5, 6})
	comps[3].AddEdgeBi(Edge{4, 7})
	comps[3].AddEdgeBi(Edge{4, 8})
	comps[3].AddEdgeBi(Edge{7, 8})
	comps[2].AddEdgeBi(Edge{7, 9})
	comps[10].AddEdgeBi(Edge{4, 10})
	comps[9].AddEdgeBi(Edge{10, 11})
	comps[9].AddEdgeBi(Edge{11, 12})
	comps[9].AddEdgeBi(Edge{12, 13})
	comps[9].AddEdgeBi(Edge{13, 10})
	comps[8].AddEdgeBi(Edge{12, 14})
	comps[7].AddEdgeBi(Edge{14, 15})
	comps[7].AddEdgeBi(Edge{15, 16})
	comps[7].AddEdgeBi(Edge{16, 17})
	comps[7].AddEdgeBi(Edge{17, 14})
	comps[7].AddEdgeBi(Edge{14, 18})
	comps[7].AddEdgeBi(Edge{18, 16})
	comps[5].AddEdgeBi(Edge{17, 19})
	comps[4].AddEdgeBi(Edge{19, 20})
	comps[4].AddEdgeBi(Edge{20, 21})
	comps[4].AddEdgeBi(Edge{21, 19})
	comps[6].AddEdgeBi(Edge{17, 22})
	return
}

func edgeBCCGolden(g Graph) (bridges Graph, comps []Graph) {
	bridges = CreateGraphByType(g)
	comps = make([]Graph, 6, 6)
	for i := range comps {
		comps[i] = CreateGraphByType(g)
	}

	bridges.AddEdgeBi(Edge{19, 17})
	bridges.AddEdgeBi(Edge{17, 22})
	bridges.AddEdgeBi(Edge{12, 14})
	bridges.AddEdgeBi(Edge{10, 4})
	bridges.AddEdgeBi(Edge{7, 9})
	bridges.AddEdgeBi(Edge{0, 4})

	comps[0].AddEdgeBi(Edge{0, 1})
	comps[0].AddEdgeBi(Edge{1, 2})
	comps[0].AddEdgeBi(Edge{2, 3})
	comps[0].AddEdgeBi(Edge{3, 0})
	comps[1].AddEdgeBi(Edge{4, 5})
	comps[1].AddEdgeBi(Edge{4, 6})
	comps[1].AddEdgeBi(Edge{5, 6})
	comps[2].AddEdgeBi(Edge{4, 7})
	comps[2].AddEdgeBi(Edge{4, 8})
	comps[2].AddEdgeBi(Edge{7, 8})
	comps[5].AddEdgeBi(Edge{10, 11})
	comps[5].AddEdgeBi(Edge{11, 12})
	comps[5].AddEdgeBi(Edge{12, 13})
	comps[5].AddEdgeBi(Edge{13, 10})
	comps[4].AddEdgeBi(Edge{14, 15})
	comps[4].AddEdgeBi(Edge{15, 16})
	comps[4].AddEdgeBi(Edge{16, 17})
	comps[4].AddEdgeBi(Edge{17, 14})
	comps[4].AddEdgeBi(Edge{14, 18})
	comps[4].AddEdgeBi(Edge{18, 16})
	comps[3].AddEdgeBi(Edge{19, 20})
	comps[3].AddEdgeBi(Edge{20, 21})
	comps[3].AddEdgeBi(Edge{21, 19})
	return
}

func checkBCCGraphOutOfOrder(t *testing.T, g Graph, gGloden Graph) {
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
	g := NewAdjacencyList()
	bccSetupGraph(g)
	cuts, comps := VertexBCC(g)
	cutsExp, compsExp := vertexBCCGolden(g)
	checkBCCGraphOutOfOrder(t, cuts, cutsExp)
	for i := range comps {
		checkBCCGraphOutOfOrder(t, comps[i], compsExp[i])
	}
}

func TestEdgeBCC(t *testing.T) {
	g := NewAdjacencyList()
	bccSetupGraph(g)
	bridges, comps := EdgeBCC(g)
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
