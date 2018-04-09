package graph

import "sort"

func getCutsAndBridgesFromComponent(dfsForest *DFSForest)(cutsAndBridges Graph) {
	cutsAndBridges = CreateGraphByType(dfsForest.Trees)
	vertices := dfsForest.AllVertices()
	lows := make([]int, len(vertices), cap(vertices))
	sort.Slice(vertices, func(i, j int) bool {
		return vertices[i].(*DFSElement).D > vertices[j].(*DFSElement).D
	})
	for i := range vertices {
		lows[i] = vertices[i].(*DFSElement).D
		for _, e := range dfsForest.BackEdges.AllConnectedVertices(vertices[i]) {
			if e.(*DFSElement).D < lows[i] {
				lows[i] = e.(*DFSElement).D
			}
		}
	}
	return cutsAndBridges
}