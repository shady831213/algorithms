package graph

import (
	"reflect"
	"sort"
	"testing"
)

func compareGraph(t *testing.T, v, vExp []interface{}, e, eExp []edge) {
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

func checkGraphInOrder(t *testing.T, g graph, gGolden graph) {
	edges := g.AllEdges()
	//finish time increase order
	vertexes := g.AllVertices()

	expEdges := gGolden.AllEdges()
	expVertices := gGolden.AllVertices()

	compareGraph(t, vertexes, expVertices, edges, expEdges)
}

func checkGraphOutOfOrderInString(t *testing.T, g graph, gGolden graph, comparator func(t *testing.T, v, vExp []interface{}, e, eExp []edge)) {
	edges := g.AllEdges()
	//finish time increase order
	vertexes := g.AllVertices()
	sort.Slice(edges, func(i, j int) bool {
		if edges[i].End.(string) == edges[j].End.(string) {
			return edges[i].Start.(string) < edges[j].Start.(string)
		}
		return edges[i].End.(string) < edges[j].End.(string)
	})

	sort.Slice(vertexes, func(i, j int) bool {
		return vertexes[i].(string) < vertexes[j].(string)
	})

	expEdges := gGolden.AllEdges()
	expVertices := gGolden.AllVertices()

	sort.Slice(expEdges, func(i, j int) bool {
		if expEdges[i].End.(string) == expEdges[j].End.(string) {
			return expEdges[i].Start.(string) < expEdges[j].Start.(string)
		}
		return expEdges[i].End.(string) < expEdges[j].End.(string)
	})

	sort.Slice(expVertices, func(i, j int) bool {
		return expVertices[i].(string) < expVertices[j].(string)
	})

	compareGraph(t, vertexes, expVertices, edges, expEdges)
	if comparator != nil {
		comparator(t, vertexes, expVertices, edges, expEdges)
	}
}

func checkGraphOutOfOrderInInt(t *testing.T, g graph, gGolden graph, comparator func(t *testing.T, v, vExp []interface{}, e, eExp []edge)) {
	edges := g.AllEdges()
	//finish time increase order
	vertexes := g.AllVertices()
	sort.Slice(edges, func(i, j int) bool {
		if edges[i].End.(int) == edges[j].End.(int) {
			return edges[i].Start.(int) < edges[j].Start.(int)
		}
		return edges[i].End.(int) < edges[j].End.(int)
	})

	sort.Slice(vertexes, func(i, j int) bool {
		return vertexes[i].(int) < vertexes[j].(int)
	})

	expEdges := gGolden.AllEdges()
	expVertices := gGolden.AllVertices()

	sort.Slice(expEdges, func(i, j int) bool {
		if expEdges[i].End.(int) == expEdges[j].End.(int) {
			return expEdges[i].Start.(int) < expEdges[j].Start.(int)
		}
		return expEdges[i].End.(int) < expEdges[j].End.(int)
	})

	sort.Slice(expVertices, func(i, j int) bool {
		return expVertices[i].(int) < expVertices[j].(int)
	})

	compareGraph(t, vertexes, expVertices, edges, expEdges)
	if comparator != nil {
		comparator(t, vertexes, expVertices, edges, expEdges)
	}
}
