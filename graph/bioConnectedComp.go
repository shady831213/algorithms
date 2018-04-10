package graph

import (
	"sort"
	"container/list"
)

func getCutsAndBridgesFromComponent(dfsForest *DFSForest) (cuts, bridges Graph) {
	cuts, bridges = CreateGraphByType(dfsForest.Trees), CreateGraphByType(dfsForest.Trees)
	vertices := dfsForest.AllVertices()
	lows := make(map[*DFSElement]int)
	//sort in order of decreasing D. it means from deepest to root
	sort.Slice(vertices, func(i, j int) bool {
		return vertices[i].(*DFSElement).D > vertices[j].(*DFSElement).D
	})
	for i := range vertices {
		connections := dfsForest.Trees.AllConnectedVertices(vertices[i])
		for _, v := range connections {
			//for leaf vertex
			if _, ok := lows[v.(*DFSElement)]; !ok {
				lows[v.(*DFSElement)] = v.(*DFSElement).D
			}
			//update back edges
			for _, bv := range dfsForest.BackEdges.AllConnectedVertices(v) {
				if bv != v.(*DFSElement).P && bv.(*DFSElement).D < lows[v.(*DFSElement)] {
					lows[v.(*DFSElement)] = bv.(*DFSElement).D
				}
			}

			if _, ok := lows[vertices[i].(*DFSElement)]; !ok {
				lows[vertices[i].(*DFSElement)] = vertices[i].(*DFSElement).D
			}
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
				bridges.AddEdgeBi(Edge{vertices[i], v})
			}
		}
	}
	return
}

func VertexBCC(g Graph) (cuts Graph, comps []Graph) {
	cuts = CreateGraphByType(g)
	comps = make([]Graph, 0, 0)
	lows := make(map[interface{}]int)
	disc := make(map[interface{}]int)
	children := make(map[interface{}]int)
	timer := 0
	vertexStack := list.New()
	edgeStack := list.New()
	for _, v := range g.AllVertices() {
		if _, ok := disc[v]; !ok {
			timer ++
			disc[v] = timer
			lows[v] = disc[v]
			vertexStack.PushBack(v)
			for {
				top := vertexStack.Back().Value
				for _, e := range g.AllConnectedVertices(top) {
					if _, ok := disc[e]; !ok {
						if _, ok := children[top]; !ok {
							children[top] = 0
						}
						children[top]++
						timer ++
						disc[e] = timer
						lows[e] = disc[e]
						vertexStack.PushBack(e)
						edgeStack.PushBack(Edge{top, e})
						break
					} else if disc[e] < lows[top] {
						//back edge
						edgeStack.PushBack(Edge{top, e})
						lows[top] = disc[e]
					}
				}
				if top == vertexStack.Back().Value {
					vertexStack.Remove(vertexStack.Back())
					if vertexStack.Len() != 0 {
						next := vertexStack.Back().Value
						if lows[top] < lows[next] {
							lows[next] = lows[top]
						}
						if lows[top] >= disc[next] {
							if !(disc[next] == 1 && children[next] < 2) {
								cuts.AddVertex(next)
							}
							comps = append(comps, CreateGraphByType(g))
							for edgeStack.Len() != 0 {
								e := edgeStack.Back().Value
								edgeStack.Remove(edgeStack.Back())
								comps[len(comps)-1].AddEdgeBi(e.(Edge))
								if e.(Edge).Start == next && e.(Edge).End == top {
									break
								}
							}
						}
					} else {
						break
					}
				}
			}
		}
	}
	return
}

func EdgeBCC(g Graph) (bridges Graph, comps []Graph) {
	bridges = CreateGraphByType(g)
	comps = make([]Graph, 0, 0)
	lows := make(map[interface{}]int)
	disc := make(map[interface{}]int)
	parent := make(map[interface{}]interface{})
	timer := 0
	vertexStack := list.New()
	edgeStack := list.New()
	for _, v := range g.AllVertices() {
		if _, ok := disc[v]; !ok {
			timer ++
			disc[v] = timer
			lows[v] = disc[v]
			vertexStack.PushBack(v)
			for {
				top := vertexStack.Back().Value
				for _, e := range g.AllConnectedVertices(top) {
					if _, ok := disc[e]; !ok {
						parent[e] = top
						timer ++
						disc[e] = timer
						lows[e] = disc[e]
						vertexStack.PushBack(e)
						edgeStack.PushBack(Edge{top, e})
						break
					} else if disc[e] < lows[top] && parent[top] != e {
						//back edge, excluding bio-dir edge
						edgeStack.PushBack(Edge{top, e})
						lows[top] = disc[e]
					}
				}
				if top == vertexStack.Back().Value {
					vertexStack.Remove(vertexStack.Back())
					if vertexStack.Len() != 0 {
						next := vertexStack.Back().Value
						if lows[top] < lows[next] {
							lows[next] = lows[top]
						}
						if lows[top] > disc[next] {
							bridges.AddEdgeBi(Edge{next, top})
						}
						if lows[top] >= disc[next] {
							comp := CreateGraphByType(g)
							for edgeStack.Len() != 0 {
								e := edgeStack.Back().Value
								edgeStack.Remove(edgeStack.Back())
								//excluding bridge
								if lows[top] == disc[next] || e.(Edge).Start != next || e.(Edge).End != top {
									comp.AddEdgeBi(e.(Edge))
								}
								if e.(Edge).Start == next && e.(Edge).End == top {
									break
								}
							}
							//if component is not empty, push it to the slice
							if len(comp.AllVertices()) != 0 {
								comps = append(comps,comp)
							}
						}
					} else {
						break
					}
				}
			}
		}
	}
	return
}
