package graph
//need add edges to scc
func SCC(g Graph) (scc Graph) {
	scc = CreateGraphByType(g)
	//DFS and get forest
	dfsGraph, gT := DFS(g, nil), g.Transpose()
	//DFS of transpose in order of decreasing finish time
	dfsGraphOfT := DFS(gT, func(vertices []interface{}) {
		for i, v := range dfsGraph["dfsForest"].AllVertices() {
			vertices[len(vertices)-1-i] = v.(*DFSElement).V
		}
	})
	//shrink all vertices, according to the root(disjoint-set)
	forest := make(map[*DFSElement]*[]interface{})
	vertices := dfsGraphOfT["dfsForest"].AllVertices()
	for i := range vertices {
		root := vertices[i].(*DFSElement).FindRoot()
		if _, ok := forest[root]; !ok {
			v := []interface{}{vertices[i].(*DFSElement).V}
			forest[root] = &v
			scc.AddVertex(&v)
		} else {
			*forest[root] = append(*forest[root], vertices[i].(*DFSElement).V)
		}
	}
	//keep all cross edges
	for _,e := range dfsGraphOfT["dfsCrossEdges"].AllEdges() {
		scc.AddEdge(Edge{forest[e.End.(*DFSElement).FindRoot()], forest[e.Start.(*DFSElement).FindRoot()]})
	}

	return
}
