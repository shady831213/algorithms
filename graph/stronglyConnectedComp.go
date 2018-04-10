package graph

import "sort"

//need add edges to scc
func SCC(g Graph) (scc Graph) {
	scc = CreateGraphByType(g)
	//DFS and get forest
	dfsGraph, gT := DFS(g, nil), g.Transpose()
	//DFS of transpose in order of decreasing finish time
	dfsVertices := dfsGraph.AllVertices()
	sort.Slice(dfsVertices, func(i, j int) bool {
		return dfsVertices[i].(*DFSElement).F > dfsVertices[j].(*DFSElement).F
	})
	dfsGraphOfT := DFS(gT, func(vertices []interface{}) {
		for i, v := range dfsVertices {
			vertices[i] = v.(*DFSElement).V
		}
	})
	//shrink all vertices, according to the root(disjoint-set)
	components := make(map[*DFSElement]Graph)

	for i := range dfsGraphOfT.Comps {
		components[i] = CreateGraphByType(g)
		//add all sub vertices
		for _, v := range dfsGraphOfT.Comps[i].AllVertices() {
			components[i].AddVertex(v.(*DFSElement).V)
		}
		//add all sub edges
		for _, e := range dfsGraphOfT.Comps[i].AllEdges() {
			components[i].AddEdge(Edge{e.End.(*DFSElement).V, e.Start.(*DFSElement).V})
		}
	}

	//keep all cross edges which cross components
	for _, e := range dfsGraphOfT.AllCrossEdges() {
		if e.End.(*DFSElement).FindRoot() != e.Start.(*DFSElement).FindRoot() {
			scc.AddEdge(Edge{components[e.End.(*DFSElement).FindRoot()], components[e.Start.(*DFSElement).FindRoot()]})
		}
	}

	return
}
