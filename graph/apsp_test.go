package graph

import (
	"sort"
	"testing"
)

func apspSetup(g weightedGraph) {
	for i := 0; i < 5; i++ {
		g.AddVertex(i)
	}
	g.AddEdgeWithWeight(edge{0, 1}, 3)
	g.AddEdgeWithWeight(edge{0, 2}, 8)
	g.AddEdgeWithWeight(edge{0, 4}, -4)
	g.AddEdgeWithWeight(edge{1, 3}, 1)
	g.AddEdgeWithWeight(edge{1, 4}, 7)
	g.AddEdgeWithWeight(edge{2, 1}, 4)
	g.AddEdgeWithWeight(edge{3, 0}, 2)
	g.AddEdgeWithWeight(edge{3, 2}, -5)
	g.AddEdgeWithWeight(edge{4, 3}, 6)
}

func distFloydWarShallGolden() weightedGraph {
	golden := newWeightedGraph()
	for i := 0; i < 5; i++ {
		golden.AddVertex(i)
		golden.AddEdgeWithWeight(edge{i, i}, 0)
	}
	golden.AddEdgeWithWeight(edge{0, 1}, 1)
	golden.AddEdgeWithWeight(edge{0, 2}, -3)
	golden.AddEdgeWithWeight(edge{0, 3}, 2)
	golden.AddEdgeWithWeight(edge{0, 4}, -4)

	golden.AddEdgeWithWeight(edge{1, 0}, 3)
	golden.AddEdgeWithWeight(edge{1, 2}, -4)
	golden.AddEdgeWithWeight(edge{1, 3}, 1)
	golden.AddEdgeWithWeight(edge{1, 4}, -1)

	golden.AddEdgeWithWeight(edge{2, 0}, 7)
	golden.AddEdgeWithWeight(edge{2, 1}, 4)
	golden.AddEdgeWithWeight(edge{2, 3}, 5)
	golden.AddEdgeWithWeight(edge{2, 4}, 3)

	golden.AddEdgeWithWeight(edge{3, 0}, 2)
	golden.AddEdgeWithWeight(edge{3, 1}, -1)
	golden.AddEdgeWithWeight(edge{3, 2}, -5)
	golden.AddEdgeWithWeight(edge{3, 4}, -2)

	golden.AddEdgeWithWeight(edge{4, 0}, 8)
	golden.AddEdgeWithWeight(edge{4, 1}, 5)
	golden.AddEdgeWithWeight(edge{4, 2}, 1)
	golden.AddEdgeWithWeight(edge{4, 3}, 6)

	return golden
}

func pathFloydWarShallGolden(g weightedGraph) map[interface{}]weightedGraph {
	golden := make(map[interface{}]weightedGraph)
	for i := 0; i < 5; i++ {
		golden[i] = newWeightedGraph()
	}
	golden[0].AddEdgeWithWeight(edge{2, 1}, g.Weight(edge{2, 1}))
	golden[0].AddEdgeWithWeight(edge{3, 2}, g.Weight(edge{3, 2}))
	golden[0].AddEdgeWithWeight(edge{4, 3}, g.Weight(edge{4, 3}))
	golden[0].AddEdgeWithWeight(edge{0, 4}, g.Weight(edge{0, 4}))

	golden[1].AddEdgeWithWeight(edge{3, 0}, g.Weight(edge{3, 0}))
	golden[1].AddEdgeWithWeight(edge{3, 2}, g.Weight(edge{3, 2}))
	golden[1].AddEdgeWithWeight(edge{0, 4}, g.Weight(edge{0, 4}))
	golden[1].AddEdgeWithWeight(edge{1, 3}, g.Weight(edge{1, 3}))

	golden[2].AddEdgeWithWeight(edge{3, 0}, g.Weight(edge{3, 0}))
	golden[2].AddEdgeWithWeight(edge{0, 4}, g.Weight(edge{0, 4}))
	golden[2].AddEdgeWithWeight(edge{2, 1}, g.Weight(edge{2, 1}))
	golden[2].AddEdgeWithWeight(edge{1, 3}, g.Weight(edge{1, 3}))

	golden[3].AddEdgeWithWeight(edge{3, 0}, g.Weight(edge{3, 0}))
	golden[3].AddEdgeWithWeight(edge{3, 2}, g.Weight(edge{3, 2}))
	golden[3].AddEdgeWithWeight(edge{0, 4}, g.Weight(edge{0, 4}))
	golden[3].AddEdgeWithWeight(edge{2, 1}, g.Weight(edge{2, 1}))

	golden[4].AddEdgeWithWeight(edge{3, 0}, g.Weight(edge{3, 0}))
	golden[4].AddEdgeWithWeight(edge{3, 2}, g.Weight(edge{3, 2}))
	golden[4].AddEdgeWithWeight(edge{2, 1}, g.Weight(edge{2, 1}))
	golden[4].AddEdgeWithWeight(edge{4, 3}, g.Weight(edge{4, 3}))

	return golden
}

func checkApspOutOfOrder(t *testing.T, g, gGolden weightedGraph) {
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
	for _, e := range expEdges {
		if g.Weight(e) != gGolden.Weight(e) {
			t.Log(e, "weight Error! exp :", gGolden.Weight(e), "actual: ", g.Weight(e))
			t.FailNow()
		}
	}
}

func TestDistFloydWarShall(t *testing.T) {
	g := newWeightedGraph()
	apspSetup(g)
	newG := distFloydWarShall(g)
	goldenG := distFloydWarShallGolden()
	checkApspOutOfOrder(t, newG, goldenG)
}

func TestPathFloydWarShall(t *testing.T) {
	g := newWeightedGraph()
	apspSetup(g)
	newForest := pathFloydWarShall(g)
	goldenForest := pathFloydWarShallGolden(g)
	for v := range newForest {
		checkApspOutOfOrder(t, newForest[v], goldenForest[v])
	}
}

func TestJohnson(t *testing.T) {
	g := newWeightedGraph()
	apspSetup(g)
	newG := johnson(g)
	goldenG := distFloydWarShallGolden()
	checkApspOutOfOrder(t, newG, goldenG)
}
