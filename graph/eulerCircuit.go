package graph

import (
	"container/list"
	"fmt"
)

type eulerVertex struct {
	vertex, p        interface{}
	iDegree, oDegree int
}

func (e *eulerVertex) init(vertex interface{}) *eulerVertex {
	e.vertex = vertex
	e.iDegree = 0
	e.oDegree = 0
	e.p = nil
	return e
}

func newEulerVertex(vertex interface{}) *eulerVertex {
	return new(eulerVertex).init(vertex)
}

func eulerCircuit(g graph, oriented bool) []edge {
	vertices := make(map[interface{}]*eulerVertex)
	vertexStack := list.New()
	edgeQueue := list.New()
	vCnt, eCnt := 0, 0
	checkDegree := func(v interface{}) bool {
		if vertices[v].iDegree == 0 || vertices[v].oDegree == 0 {
			return false
		}
		if oriented {
			return vertices[v].iDegree == vertices[v].oDegree
		}
		return vertices[v].iDegree%2 == 0
	}

	checkVertexAndEdgeCnt := func() bool {
		if oriented {
			return eCnt == vCnt+1
		}
		return eCnt == (vCnt+1)<<1
	}

	pushVertexStack := func(vertex interface{}) {
		vertices[vertex] = newEulerVertex(vertex)
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
			for e := range g.IterConnectedVertices(top) {
				eCnt++
				//it means top has a new output
				vertices[top].oDegree++
				if _, ok := vertices[e]; !ok {
					//e is first time discovered, and has a input.Then link new vertices to the ring
					pushVertexStack(e)
					vertices[e].iDegree = 1
					vertices[e].p = top
					edgeQueue.PushBack(edge{top, e})
					break
				} else {
					vertices[e].iDegree++
					if oriented || vertices[top].p != e {
						//ignore redundant edge of undirectedGraph
						edgeQueue.PushBack(edge{top, e})
					}
				}
			}
			if top == vertexStack.Back().Value {
				vertexStack.Remove(vertexStack.Back())
			}
		}
	}

	//for single vertex condition
	if !checkVertexAndEdgeCnt() {
		return nil
	}
	path := make([]edge, 0, 0)
	//check and output path, O(E)
	for top := edgeQueue.Front(); top != nil; top = edgeQueue.Front() {
		e := top.Value.(edge)
		edgeQueue.Remove(top)
		fmt.Println(e)
		if !checkDegree(e.Start) || !checkDegree(e.End) {
			//check degree rules
			return nil
		}
		path = append(path, e)
	}

	return path
}
