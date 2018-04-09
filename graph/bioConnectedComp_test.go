package graph

import "testing"

func bccSetupGraph(g Graph) {
	for i := 0; i < 23; i ++ {
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
