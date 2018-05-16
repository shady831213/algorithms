package graph

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func setupGraph() graph {
	g := newGraph()
	g.AddVertex(2)
	g.AddVertex(4)
	g.AddVertex(6)
	g.AddEdgeBi(edge{6, 4})
	g.AddVertex(6)
	g.AddVertex(8)
	g.AddEdge(edge{1, 2})
	g.AddEdge(edge{2, 1})
	g.AddEdge(edge{2, 4})
	g.AddEdge(edge{8, 7})
	return g
}

func checkGraph(t *testing.T, g graph) {
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
	expEdges := []edge{{1, 2}, {2, 1}, {2, 4}, {4, 6}, {6, 4}, {8, 7}}
	expVertices := []interface{}{1, 2, 4, 6, 7, 8}
	expConnetedVertices := [][]interface{}{{2}, {1, 4}, {6}, {4}, {}, {7}}
	if !reflect.DeepEqual(edges, expEdges) {
		t.Log(fmt.Sprintf("get edges error!expect:%+v;but get:%+v", expEdges, edges))
		t.Fail()
	}
	if !reflect.DeepEqual(vertexes, expVertices) {
		t.Log(fmt.Sprintf("get vertexes error!expect:%+v;but get:%+v", expVertices, vertexes))
		t.Fail()
	}
	if !reflect.DeepEqual(connectedVertices, expConnetedVertices) {
		t.Log(fmt.Sprintf("get connectedVertices error!expect:%+v;but get:%+v", expConnetedVertices, connectedVertices))
		t.Fail()
	}

}

func TestNewAdjacencyMatrix(t *testing.T) {
	checkGraph(t, setupGraph())
}
