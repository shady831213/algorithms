package graph

import (
	"testing"
)

func mstSetup() weightedGraph {
	g := newWeightedGraph()
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
	return g
}

func mstGolden() weightedGraph {
	t := newWeightedGraph()
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

func secondaryMstGolden() weightedGraph {
	t := newWeightedGraph()
	t.AddEdgeWithWeightBi(edge{"a", "b"}, 4)
	t.AddEdgeWithWeightBi(edge{"a", "h"}, 8)
	t.AddEdgeWithWeightBi(edge{"c", "l"}, 2)
	t.AddEdgeWithWeightBi(edge{"g", "h"}, 1)
	t.AddEdgeWithWeightBi(edge{"g", "f"}, 2)
	t.AddEdgeWithWeightBi(edge{"c", "f"}, 4)
	t.AddEdgeWithWeightBi(edge{"c", "d"}, 7)
	t.AddEdgeWithWeightBi(edge{"e", "f"}, 10)

	return t
}

func bottleNeckSpanningTreeGolden() weightedGraph {
	t := newWeightedGraph()
	t.AddEdgeWithWeightBi(edge{"a", "b"}, 4)
	t.AddEdgeWithWeightBi(edge{"a", "h"}, 8)
	t.AddEdgeWithWeightBi(edge{"c", "l"}, 2)
	t.AddEdgeWithWeightBi(edge{"g", "h"}, 1)
	t.AddEdgeWithWeightBi(edge{"g", "f"}, 2)
	t.AddEdgeWithWeightBi(edge{"l", "g"}, 6)
	t.AddEdgeWithWeightBi(edge{"c", "d"}, 7)
	t.AddEdgeWithWeightBi(edge{"d", "e"}, 9)

	return t
}

func checkMstOutOfOrder(t *testing.T, g, gGolden weightedGraph) {
	checkGraphOutOfOrderInString(t, g, gGolden, nil)
	if g.TotalWeight() != gGolden.TotalWeight() {
		t.Log("expect totalWeight :", gGolden.TotalWeight(), "actaul :", g.TotalWeight())
		t.Fail()
	}
}

func TestMstKruskal(t *testing.T) {
	g := mstSetup()
	tree := mstKruskal(g)
	treeExp := mstGolden()
	checkMstOutOfOrder(t, tree, treeExp)
}

func TestMstPrim(t *testing.T) {
	g := mstSetup()
	tree := mstPrim(g)
	treeExp := mstGolden()
	checkMstOutOfOrder(t, tree, treeExp)
}

func TestSecondaryMst(t *testing.T) {
	g := mstSetup()
	tree := secondaryMst(g)
	treeExp := secondaryMstGolden()
	checkMstOutOfOrder(t, tree, treeExp)
}

func TestMstReducedPrim(t *testing.T) {
	g := mstSetup()
	tree := mstReducedPrim(g, 2)
	treeExp := mstGolden()
	checkMstOutOfOrder(t, tree, treeExp)
}

func TestBottleNeckSpanningTree(t *testing.T) {
	g := mstSetup()
	tree := bottleNeckSpanningTree(g)
	treeExp := bottleNeckSpanningTreeGolden()
	checkMstOutOfOrder(t, tree, treeExp)
}
