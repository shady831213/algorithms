package graph

import "sort"

func getCutsAndBridges(dfsGraphs map[string]Graph)(cutsAndBridges Graph) {
	cutsAndBridges = CreateGraphByType(dfsGraphs["dfsForest"])
	vertices := dfsGraphs["dfsForest"].AllVertices()
	lows := make([]int, len(vertices), cap(vertices))
	sort.Slice(vertices, func(i, j int) bool {
		return vertices[i].(*DFSElement).D > vertices[j].(*DFSElement).D
	})
	for i := range vertices {
		lows[i] = vertices[i].(*DFSElement).D
		for _, e := range dfsGraphs["dfsBackEdges"].AllConnectedVertices(vertices[i]) {
			if e.(*DFSElement).D < lows[i] {
				lows[i] = e.(*DFSElement).D
			}
		}
	}
	return cutsAndBridges
}