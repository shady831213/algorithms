package graph

import (
	"testing"
	"reflect"
	"sort"
)

func compareGraph(t *testing.T, v, vExp []interface{}, e, eExp []Edge) {
	if !reflect.DeepEqual(e, eExp) {
		t.Log("get edges error!")
		for i := range eExp {
			if !reflect.DeepEqual(eExp[i], e[i]) {
				t.Log("expect:")
				t.Log(eExp[i])
				t.Log(eExp[i].Start)
				t.Log(eExp[i].End)
				t.Log("but get:")
				t.Log(e[i])
				t.Log(e[i].Start)
				t.Log(e[i].End)
			}
		}

		t.Fail()
	}
	if !reflect.DeepEqual(v, vExp) {
		t.Log("get vertexes error!")
		for i := range vExp {
			if !reflect.DeepEqual(vExp[i], v[i]) {
				t.Log("expect:")
				t.Log(vExp[i])
				t.Log("but get:")
				t.Log(v[i])
			}

		}
		t.Fail()
	}
}


func checkDFSGraphOutOfOrder(t *testing.T, g Graph, gGloden Graph) {
	edges := g.AllEdges()
	//finish time increase order
	vertexes := g.AllVertices()
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Start.(*DFSElement).D < edges[j].Start.(*DFSElement).D
	})

	sort.Slice(vertexes, func(i, j int) bool {
		return vertexes[i].(*DFSElement).D < vertexes[j].(*DFSElement).D
	})

	expEdges := gGloden.AllEdges()
	expVertices := gGloden.AllVertices()

	sort.Slice(expEdges, func(i, j int) bool {
		return expEdges[i].Start.(*DFSElement).D < expEdges[j].Start.(*DFSElement).D
	})

	sort.Slice(expVertices, func(i, j int) bool {
		return expVertices[i].(*DFSElement).D < expVertices[j].(*DFSElement).D
	})

	compareGraph(t, vertexes, expVertices, edges, expEdges)
}

func checkGraphInOrder(t *testing.T, g Graph, gGloden Graph) {
	edges := g.AllEdges()
	//finish time increase order
	vertexes := g.AllVertices()

	expEdges := gGloden.AllEdges()
	expVertices := gGloden.AllVertices()

	compareGraph(t, vertexes, expVertices, edges, expEdges)
}