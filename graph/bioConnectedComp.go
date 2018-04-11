package graph

import (
	"container/list"
	"sort"
)

func getCutsAndBridgesFromComponent(dfsForest *dfsForest) (cuts, bridges graph) {
	cuts, bridges = createGraphByType(dfsForest.Trees), createGraphByType(dfsForest.Trees)
	vertices := dfsForest.AllVertices()
	lows := make(map[*dfsElement]int)
	//sort in order of decreasing D. it means from deepest to root
	sort.Slice(vertices, func(i, j int) bool {
		return vertices[i].(*dfsElement).D > vertices[j].(*dfsElement).D
	})
	for i := range vertices {
		connections := dfsForest.Trees.AllConnectedVertices(vertices[i])
		for _, v := range connections {
			//for leaf vertex
			if _, ok := lows[v.(*dfsElement)]; !ok {
				lows[v.(*dfsElement)] = v.(*dfsElement).D
			}
			//update back edges
			for _, bv := range dfsForest.BackEdges.AllConnectedVertices(v) {
				if bv != v.(*dfsElement).P && bv.(*dfsElement).D < lows[v.(*dfsElement)] {
					lows[v.(*dfsElement)] = bv.(*dfsElement).D
				}
			}

			if _, ok := lows[vertices[i].(*dfsElement)]; !ok {
				lows[vertices[i].(*dfsElement)] = vertices[i].(*dfsElement).D
			}
			if lows[v.(*dfsElement)] < lows[vertices[i].(*dfsElement)] {
				lows[vertices[i].(*dfsElement)] = lows[v.(*dfsElement)]
			}

			if lows[v.(*dfsElement)] >= vertices[i].(*dfsElement).D {
				//Cuts, excluding root that has less than 2 children
				if !(vertices[i].(*dfsElement).P == nil && len(connections) < 2) {
					cuts.AddVertex(vertices[i])
				}
			}
			if lows[v.(*dfsElement)] > vertices[i].(*dfsElement).D {
				//bridges
				bridges.AddEdgeBi(edge{vertices[i], v})
			}
		}
	}
	return
}

func vertexBCC(g graph) (cuts graph, comps []graph) {
	cuts = createGraphByType(g)
	comps = make([]graph, 0, 0)
	lows := make(map[interface{}]int)
	disc := make(map[interface{}]int)
	children := make(map[interface{}]int)
	timer := 0
	vertexStack := list.New()
	edgeStack := list.New()
	pushVertexStack := func(v interface{}) {
		timer++
		disc[v] = timer
		lows[v] = disc[v]
		vertexStack.PushBack(v)
	}
	for _, v := range g.AllVertices() {
		if _, ok := disc[v]; !ok {
			pushVertexStack(v)
			for vertexStack.Len() != 0 {
				top := vertexStack.Back().Value
				for e := range g.IterConnectedVertices(top) {
					if _, ok := disc[e]; !ok {
						children[top]++
						pushVertexStack(e)
						edgeStack.PushBack(edge{top, e})
						break
					} else if disc[e] < lows[top] {
						//back edge
						edgeStack.PushBack(edge{top, e})
						lows[top] = disc[e]
					}
				}
				if top == vertexStack.Back().Value {
					vertexStack.Remove(vertexStack.Back())
					if vertexStack.Back() == nil {
						continue
					}
					next := vertexStack.Back().Value
					if lows[top] < lows[next] {
						lows[next] = lows[top]
					}
					if lows[top] >= disc[next] {
						if !(disc[next] == 1 && children[top] < 2) {
							cuts.AddVertex(next)
						}
						comps = append(comps, createGraphByType(g))
						curEdge := edge{next, top}
						comps[len(comps)-1].AddEdgeBi(curEdge)
						for e := edgeStack.Back().Value; e != curEdge; e = edgeStack.Back().Value {
							edgeStack.Remove(edgeStack.Back())
							comps[len(comps)-1].AddEdgeBi(e.(edge))
						}
						edgeStack.Remove(edgeStack.Back())
					}
				}
			}
		}
	}
	return
}

func edgeBCC(g graph) (bridges graph, comps []graph) {
	bridges = createGraphByType(g)
	comps = make([]graph, 0, 0)
	lows := make(map[interface{}]int)
	disc := make(map[interface{}]int)
	parent := make(map[interface{}]interface{})
	timer := 0
	vertexStack := list.New()
	edgeStack := list.New()
	pushVertexStack := func(v interface{}) {
		timer++
		disc[v] = timer
		lows[v] = disc[v]
		vertexStack.PushBack(v)
	}
	for _, v := range g.AllVertices() {
		if _, ok := disc[v]; !ok {
			pushVertexStack(v)
			for vertexStack.Len() != 0 {
				top := vertexStack.Back().Value
				for e := range g.IterConnectedVertices(top) {
					if _, ok := disc[e]; !ok {
						parent[e] = top
						pushVertexStack(e)
						edgeStack.PushBack(edge{top, e})
						break
					} else if disc[e] < lows[top] && parent[top] != e {
						//back edge, excluding bio-dir edge
						edgeStack.PushBack(edge{top, e})
						lows[top] = disc[e]
					}
				}
				if top == vertexStack.Back().Value {
					vertexStack.Remove(vertexStack.Back())
					if vertexStack.Back() == nil {
						continue
					}
					next := vertexStack.Back().Value
					if lows[top] < lows[next] {
						lows[next] = lows[top]
					}

					if lows[top] >= disc[next] {
						comp := createGraphByType(g)
						curEdge := edge{next, top}
						if lows[top] > disc[next] {
							bridges.AddEdgeBi(curEdge)
						} else {
							//excluding bridge
							comp.AddEdgeBi(curEdge)
						}
						for e := edgeStack.Back().Value; e != curEdge; e = edgeStack.Back().Value {
							edgeStack.Remove(edgeStack.Back())
							comp.AddEdgeBi(e.(edge))
						}
						edgeStack.Remove(edgeStack.Back())
						//if component is not empty, push it to the slice
						if len(comp.AllVertices()) != 0 {
							comps = append(comps, comp)
						}
					}
				}
			}
		}
	}
	return
}
