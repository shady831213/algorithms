package graph

import "fmt"

type vertexDegree struct {
	iDegree, oDegree int
}

func (e *vertexDegree) init(vertex interface{}) *vertexDegree {
	e.iDegree = 0
	e.oDegree = 0
	return e
}

func newEulerVertex(vertex interface{}) *vertexDegree {
	return new(vertexDegree).init(vertex)
}

func checkDegree(degree *vertexDegree, e *dfsElement) bool {
	if degree.iDegree == 0 || degree.oDegree == 0 {
		return false
	}
	if e.Iter.Len() == 1 && e.Iter.Value() == e.V {
		//check single vertex loop
		return false
	}
	return degree.iDegree == degree.oDegree
}

func checkVertexAndEdgeCnt(vCnt, eCnt int) bool {
	return eCnt == vCnt+1
}

func eulerCircuit(g graph) []edge {
	degrees := make(map[interface{}]*vertexDegree)
	path := make([]edge, 0, 0)
	vCnt, eCnt := 0, 0
	nonTreeEdgeHandler := func(start, end *dfsElement) {
		if start.P != nil && start.P != end {
			degrees[start.V].oDegree++
			degrees[end.V].iDegree++
			eCnt++
			path = append(path, edge{start.V, end.V})
		}
	}
	handler := newDFSVisitHandler()
	handler.BeforeBfsHandler = func(v *dfsElement) {
		degrees[v.V] = newEulerVertex(v.V)
	}
	handler.TreeEdgeHandler = func(start, end *dfsElement) {
		eCnt++
		degrees[start.V].oDegree++
		degrees[end.V].iDegree = 1
		path = append(path, edge{start.V, end.V})
	}
	handler.BackEdgeHandler = nonTreeEdgeHandler

	//dfs O(E)
	for _, v := range g.AllVertices() {
		vCnt++
		if !handler.Elements.exist(v) {
			dfsVisit(g, v, handler)
		}
	}
	fmt.Println(vCnt, eCnt)
	//for single vertex condition
	if !checkVertexAndEdgeCnt(vCnt, eCnt) {
		return nil
	}

	//check and output path, O(E)
	for _, e := range path {
		if !checkDegree(degrees[e.Start], handler.Elements.get(e.Start).(*dfsElement)) ||
			!checkDegree(degrees[e.End], handler.Elements.get(e.End).(*dfsElement)) {
			//check degree rules
			return nil
		}
	}

	return path
}
