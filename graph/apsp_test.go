package graph

import (
	"testing"
)

func apspSetup() weightedGraph {
	g := newWeightedGraph()
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
	return g
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
	comparator := func(t *testing.T, v, vExp []interface{}, e, eExp []edge) {
		for _, e := range eExp {
			if g.Weight(e) != gGolden.Weight(e) {
				t.Log(e, "weight Error! exp :", gGolden.Weight(e), "actual: ", g.Weight(e))
				t.FailNow()
			}
		}
	}
	checkGraphOutOfOrderInInt(t, g, gGolden, comparator)
}

func TestDistFloydWarShall(t *testing.T) {
	g := apspSetup()
	newG := distFloydWarShall(g)
	goldenG := distFloydWarShallGolden()
	checkApspOutOfOrder(t, newG, goldenG)
}

func TestPathFloydWarShall(t *testing.T) {
	g := apspSetup()
	newForest := pathFloydWarShall(g)
	goldenForest := pathFloydWarShallGolden(g)
	for v := range newForest {
		checkApspOutOfOrder(t, newForest[v], goldenForest[v])
	}
}

func TestJohnson(t *testing.T) {
	g := apspSetup()
	newG := johnson(g)
	goldenG := distFloydWarShallGolden()
	checkApspOutOfOrder(t, newG, goldenG)
}
