package graph

import (
	"container/list"
)

type eulerVertex struct {
	vertex, p        interface{}
	iDegree, oDegree int
	iter             iterator
}

func (e *eulerVertex) init(vertex interface{}) *eulerVertex {
	e.vertex = vertex
	e.iDegree = 0
	e.oDegree = 0
	e.p = nil
	e.iter = nil
	return e
}

func newEulerVertex(vertex interface{}) *eulerVertex {
	return new(eulerVertex).init(vertex)
}

func checkDegree(v *eulerVertex, oriented bool) bool {
	if v.iDegree == 0 || v.oDegree == 0 {
		return false
	}
	if oriented {
		return v.iDegree == v.oDegree
	}
	return v.iDegree%2 == 0
}

func checkVertexAndEdgeCnt(vCnt, eCnt int, oriented bool) bool {
	if oriented {
		return eCnt == vCnt+1
	}
	return eCnt == (vCnt+1)<<1
}

func eulerCircuit(g graph, oriented bool) []edge {
	vertices := make(map[interface{}]*eulerVertex)
	vertexStack := list.New()
	path := make([]edge, 0, 0)
	vCnt, eCnt := 0, 0

	pushVertexStack := func(vertex interface{}) {
		vertices[vertex] = newEulerVertex(vertex)
		vertices[vertex].iter = g.IterConnectedVertices(vertex)
		vertexStack.PushBack(vertex)
	}
	//dfs O(E)
	for _, v := range g.AllVertices() {
		vCnt++
		if _, ok := vertices[v]; !ok {
			pushVertexStack(v)
		}
		for vertexStack.Len() != 0 {
			top := vertexStack.Back().Value
			for e := vertices[top].iter.Value(); e != nil; {
				eCnt++
				//it means top has a new output
				vertices[top].oDegree++
				if _, ok := vertices[e]; !ok {
					//e is first time discovered, and has a input.Then link new vertices to the ring
					pushVertexStack(e)
					vertices[e].iDegree = 1
					vertices[e].p = top
					path = append(path, edge{top, e})
					vertices[top].iter.Next()
					break
				} else {
					vertices[e].iDegree++
					if oriented || vertices[top].p != e {
						//ignore redundant edge of undirectedGraph
						path = append(path, edge{top, e})
					}
					e = vertices[top].iter.Next()
				}

			}
			if top == vertexStack.Back().Value {
				vertices[top].iter = nil
				vertexStack.Remove(vertexStack.Back())
			}
		}
	}
	//for single vertex condition
	if !checkVertexAndEdgeCnt(vCnt, eCnt, oriented) {
		return nil
	}

	//check and output path, O(E)
	for _, e := range path {
		if !checkDegree(vertices[e.Start], oriented) || !checkDegree(vertices[e.End], oriented) {
			//check degree rules
			return nil
		}
	}

	return path
}
