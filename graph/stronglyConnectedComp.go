package graph

import "sort"

//need add edges to scc
func scc(g graph) (scc graph) {
	scc = newGraph()
	//dfs and get forest
	dfsGraph, gT := dfs(g, nil), g.Transpose()
	//dfs of transpose in order of decreasing finish time
	dfsVertices := dfsGraph.AllVertices()
	sort.Slice(dfsVertices, func(i, j int) bool {
		return dfsVertices[i].(*dfsElement).F > dfsVertices[j].(*dfsElement).F
	})
	dfsGraphOfT := dfs(gT, func(vertices []interface{}) {
		for i, v := range dfsVertices {
			vertices[i] = v.(*dfsElement).V
		}
	})
	//shrink all vertices, according to the root(disjoint-set)
	components := make(map[*dfsElement]graph)

	for i := range dfsGraphOfT.Comps {
		components[i] = newGraph()
		//add all sub vertices
		for _, v := range dfsGraphOfT.Comps[i].AllVertices() {
			components[i].AddVertex(v.(*dfsElement).V)
		}
		//add all sub edges
		for _, e := range dfsGraphOfT.Comps[i].AllEdges() {
			components[i].AddEdge(edge{e.End.(*dfsElement).V, e.Start.(*dfsElement).V})
		}
	}

	//keep all cross edges which cross components
	for _, e := range dfsGraphOfT.AllCrossEdges() {
		if e.End.(*dfsElement).FindRoot() != e.Start.(*dfsElement).FindRoot() {
			scc.AddEdge(edge{components[e.End.(*dfsElement).FindRoot()], components[e.Start.(*dfsElement).FindRoot()]})
		}
	}

	return
}
