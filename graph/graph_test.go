package graph

import (
	"testing"
	"sort"
	"reflect"
	"fmt"
)

func setupGraph(g Graph) {
	g.AddVertex(2)
	g.AddVertex(4)
	g.AddVertex(6)
	g.AddEdgeBi(Edge{6,4})
	g.AddVertex(6)
	g.AddVertex(8)
	g.AddEdge(Edge{1,2})
	g.AddEdge(Edge{2,1})
	g.AddEdge(Edge{2,4})
	g.AddEdge(Edge{8,7})
}

func checkGrap(t *testing.T, g Graph) {
	edges := g.AllEdges()
	vertexes := g.AllVertices()
	sort.Slice(edges, func(i, j int) bool {
		if edges[i].Start == edges[j].Start {
			return edges[i].End.(int) < edges[j].End.(int)
		}
		return edges[i].Start.(int) < edges[j].Start.(int)
	})
	sort.Slice(vertexes, func(i, j int) bool {
		return vertexes[i].(int) < vertexes[j].(int)
	})
	connectedVertices := make([][]interface{}, len(vertexes), cap(vertexes))
	for i := range vertexes {
		connectedVertices[i] = g.AllConnectedVertices(vertexes[i])
		sort.Slice(connectedVertices[i], func(k, j int) bool {
			return connectedVertices[i][k].(int) < connectedVertices[i][j].(int)
		})
	}
	expEdges := []Edge{Edge{1,2},Edge{2,1},Edge{2,4},Edge{4,6},Edge{6,4},Edge{8,7}}
	expVertices := []interface{}{1,2,4,6,7,8}
	expConnetedVertices := [][]interface{}{[]interface{}{2}, []interface{}{1, 4}, []interface{}{6}, []interface{}{4}, []interface{}{}, []interface{}{7}}
	if !reflect.DeepEqual(edges, expEdges) {
		t.Log(fmt.Sprintf("get edges error!expect:%+v;but get:%+v", expEdges,edges))
		t.Fail()
	}
	if !reflect.DeepEqual(vertexes, expVertices) {
		t.Log(fmt.Sprintf("get vertexes error!expect:%+v;but get:%+v", expVertices,vertexes))
		t.Fail()
	}
	if !reflect.DeepEqual(connectedVertices, expConnetedVertices) {
		t.Log(fmt.Sprintf("get connectedVertices error!expect:%+v;but get:%+v", expConnetedVertices,connectedVertices))
		t.Fail()
	}

}

func testGraph(t *testing.T, g Graph) {
	setupGraph(g)
	checkGrap(t,g)
}

func TestNewAdjacencyList(t *testing.T) {
	testGraph(t, NewAdjacencyList())
}

func TestNewAdjacencyMatrix(t *testing.T) {
	testGraph(t, NewAdjacencyMatrix())
}

func TestAdjacencyList2AdjacencyMatrix(t *testing.T) {
	l := NewAdjacencyList()
	setupGraph(l)
	checkGrap(t, AdjacencyList2AdjacencyMatrix(l))
}

func TestAdjacencyMatrix2AdjacencyList(t *testing.T) {
	m := NewAdjacencyMatrix()
	setupGraph(m)
	checkGrap(t, AdjacencyMatrix2AdjacencyList(m))
}