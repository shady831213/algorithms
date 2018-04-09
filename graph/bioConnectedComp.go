package graph

import (
	"sort"
)

func getCutsAndBridgesFromComponent(dfsForest *DFSForest)(cuts,bridges Graph) {
	cuts,bridges = CreateGraphByType(dfsForest.Trees),CreateGraphByType(dfsForest.Trees)
	vertices := dfsForest.AllVertices()
	lows := make(map[*DFSElement]int)
	//sort in order of decreasing D. it means from deepest to root
	sort.Slice(vertices, func(i, j int) bool {
		return vertices[i].(*DFSElement).D > vertices[j].(*DFSElement).D
	})
	for i := range vertices {
		lows[vertices[i].(*DFSElement)] = vertices[i].(*DFSElement).D
		connections := dfsForest.Trees.AllConnectedVertices(vertices[i])
		for _, v := range connections {
			if lows[v.(*DFSElement)] < lows[vertices[i].(*DFSElement)] {
				lows[vertices[i].(*DFSElement)] = lows[v.(*DFSElement)]
			}
			if lows[v.(*DFSElement)] >= vertices[i].(*DFSElement).D {
				//Cuts, excluding root that has less than 2 children
				if !(vertices[i].(*DFSElement).P == nil && len(connections) < 2) {
					cuts.AddVertex(vertices[i])
				}
			}
			if lows[v.(*DFSElement)] > vertices[i].(*DFSElement).D {
				//bridges
				bridges.AddEdgeBi(Edge{vertices[i],v})
			}
		}
		for _, v := range dfsForest.BackEdges.AllConnectedVertices(vertices[i]) {
			if v != vertices[i].(*DFSElement).P && v.(*DFSElement).D < lows[vertices[i].(*DFSElement)] {
				lows[vertices[i].(*DFSElement)] = v.(*DFSElement).D
			}
		}
	}
	return
}