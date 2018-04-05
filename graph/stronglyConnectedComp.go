package graph
//need add edges to scc
func SCC(g Graph) (scc Graph) {
	scc = CreateGraphByType(g)

	dfsGraph, gT := DFS(g, nil), g.Transpose()
	dfsGraphOfT := DFS(gT, func(vertices []interface{}) {
		for i, v := range dfsGraph.AllVertices() {
			vertices[len(vertices)-1-i] = v.(*DFSElement).V
		}
	})
	forest := make(map[*DFSElement]*[]interface{})
	vertices := dfsGraphOfT.AllVertices()
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

	return
}
