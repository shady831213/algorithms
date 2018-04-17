package graph

import (
	"sort"
	"testing"
)

func mstSetup(g graphWeightily) {
	g.AddVertex("a")
	g.AddVertex("b")
	g.AddVertex("c")
	g.AddVertex("d")
	g.AddVertex("e")
	g.AddVertex("f")
	g.AddVertex("g")
	g.AddVertex("h")
	g.AddVertex("l")

	g.AddEdgeWithWeightBi(edge{"a", "b"}, 4)
	g.AddEdgeWithWeightBi(edge{"b", "c"}, 8)
	g.AddEdgeWithWeightBi(edge{"c", "d"}, 7)
	g.AddEdgeWithWeightBi(edge{"d", "e"}, 9)
	g.AddEdgeWithWeightBi(edge{"e", "f"}, 10)
	g.AddEdgeWithWeightBi(edge{"f", "g"}, 2)
	g.AddEdgeWithWeightBi(edge{"g", "h"}, 1)
	g.AddEdgeWithWeightBi(edge{"h", "l"}, 7)
	g.AddEdgeWithWeightBi(edge{"l", "c"}, 2)
	g.AddEdgeWithWeightBi(edge{"b", "h"}, 11)
	g.AddEdgeWithWeightBi(edge{"c", "f"}, 4)
	g.AddEdgeWithWeightBi(edge{"d", "f"}, 14)
	g.AddEdgeWithWeightBi(edge{"g", "l"}, 6)
	g.AddEdgeWithWeightBi(edge{"a", "h"}, 8)
}

func mstGolden(g graphWeightily) graphWeightily {
	t := createGraphWithWeightByType(g)
	t.AddEdgeWithWeightBi(edge{"a", "b"}, 4)
	t.AddEdgeWithWeightBi(edge{"a", "h"}, 8)
	t.AddEdgeWithWeightBi(edge{"c", "l"}, 2)
	t.AddEdgeWithWeightBi(edge{"g", "h"}, 1)
	t.AddEdgeWithWeightBi(edge{"g", "f"}, 2)
	t.AddEdgeWithWeightBi(edge{"c", "f"}, 4)
	t.AddEdgeWithWeightBi(edge{"c", "d"}, 7)
	t.AddEdgeWithWeightBi(edge{"d", "e"}, 9)

	return t
}

func checkMstOutOfOrder(t *testing.T, g, gGolden graphWeightily) {
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

	if g.TotalWeight() != gGolden.TotalWeight() {
		t.Log("expect totalWeight :", gGolden.TotalWeight(), "actaul :", g.TotalWeight())
		t.Fail()
	}
}

func TestMstKruskal(t *testing.T) {
	g := newAdjacencyListWithWeight()
	mstSetup(g)
	tree := mstKruskal(g)
	treeExp := mstGolden(g)
	checkMstOutOfOrder(t, tree, treeExp)
}

func TestMstPrim(t *testing.T) {
	g := newAdjacencyMatrixWithWeight()
	mstSetup(g)
	tree := mstPrim(g)
	treeExp := mstGolden(g)
	checkMstOutOfOrder(t, tree, treeExp)
}
