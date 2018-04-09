package graph

import "sort"

//need add edges to scc
func SCC(g Graph) (scc Graph) {
	scc = CreateGraphByType(g)
	//DFS and get forest
	dfsGraph, gT := DFS(g, nil), g.Transpose()
	//DFS of transpose in order of decreasing finish time
	dfsVertices := dfsGraph["dfsForest"].AllVertices()
	sort.Slice(dfsVertices, func(i, j int) bool {
		return dfsVertices[i].(*DFSElement).F > dfsVertices[j].(*DFSElement).F
	})
	dfsGraphOfT := DFS(gT, func(vertices []interface{}) {
		for i, v := range dfsVertices {
			vertices[i] = v.(*DFSElement).V
		}
	})
	//shrink all vertices, according to the root(disjoint-set)
	dfsForest := make(map[string]map[*DFSElement]Graph)
	for i := range dfsGraphOfT {
		dfsForest[i] = GetDFSComponent(dfsGraphOfT[i])
	}
	components := make(map[*DFSElement]Graph)
	//add all sub vertices
	for i := range dfsForest["dfsForest"] {
		components[i] = CreateGraphByType(g)
		for _,v := range dfsForest["dfsForest"][i].AllVertices() {
			components[i].AddVertex(v.(*DFSElement).V)
		}
	}
	//add all sub edges
	for i := range dfsForest {
		for j := range dfsForest[i] {
			for _,e := range dfsForest[i][j].AllEdges() {
				components[j].AddEdge(Edge{e.End.(*DFSElement).V,e.Start.(*DFSElement).V})
			}
		}
	}
	//keep all cross edges which cross components
	for _,e := range dfsGraphOfT["dfsCrossEdges"].AllEdges() {
		if e.End.(*DFSElement).FindRoot() != e.Start.(*DFSElement).FindRoot() {
			scc.AddEdge(Edge{components[e.End.(*DFSElement).FindRoot()], components[e.Start.(*DFSElement).FindRoot()]})
		}
	}

	return
}
