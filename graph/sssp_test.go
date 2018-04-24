package graph

import (
	"math"
	"sort"
	"testing"
)

func ssspSetup(g weightedGraph) {
	g.AddEdgeWithWeight(edge{"s", "t"}, 6)
	g.AddEdgeWithWeight(edge{"s", "y"}, 7)
	g.AddEdgeWithWeight(edge{"t", "y"}, 8)
	g.AddEdgeWithWeight(edge{"t", "x"}, 5)
	g.AddEdgeWithWeight(edge{"x", "t"}, -2)
	g.AddEdgeWithWeight(edge{"t", "z"}, -4)
	g.AddEdgeWithWeight(edge{"y", "x"}, -3)
	g.AddEdgeWithWeight(edge{"y", "z"}, 9)
	g.AddEdgeWithWeight(edge{"z", "x"}, 7)
	g.AddEdgeWithWeight(edge{"z", "s"}, 2)
}

func ssspGolden(g weightedGraph) weightedGraph {
	ssspG := createGraphByType(g).(weightedGraph)
	ssspE := make(map[interface{}]*ssspElement)

	ssspE["s"] = newSsspElement("s", 0)

	ssspE["y"] = newSsspElement("y", 7)
	ssspE["y"].P = ssspE["s"]

	ssspE["x"] = newSsspElement("x", 4)
	ssspE["x"].P = ssspE["y"]

	ssspE["t"] = newSsspElement("t", 2)
	ssspE["t"].P = ssspE["x"]

	ssspE["z"] = newSsspElement("z", -2)
	ssspE["z"].P = ssspE["t"]

	for v := range ssspE {
		if ssspE[v].P != nil {
			ssspG.AddEdgeWithWeight(edge{ssspE[v].P, ssspE[v]}, ssspE[v].D-ssspE[v].P.D)
		}
	}

	return ssspG
}

func checkSsspOutOfOrder(t *testing.T, g, gGolden weightedGraph) {
	edges := g.AllEdges()
	//finish time increase order
	vertexes := g.AllVertices()
	sort.Slice(edges, func(i, j int) bool {
		if edges[i].End.(*ssspElement).V.(string) == edges[j].End.(*ssspElement).V.(string) {
			return edges[i].Start.(*ssspElement).V.(string) < edges[j].Start.(*ssspElement).V.(string)
		}
		return edges[i].End.(*ssspElement).V.(string) < edges[j].End.(*ssspElement).V.(string)
	})

	sort.Slice(vertexes, func(i, j int) bool {
		return vertexes[i].(*ssspElement).V.(string) < vertexes[j].(*ssspElement).V.(string)
	})

	expEdges := gGolden.AllEdges()
	expVertices := gGolden.AllVertices()

	sort.Slice(expEdges, func(i, j int) bool {
		if expEdges[i].End.(*ssspElement).V.(string) == expEdges[j].End.(*ssspElement).V.(string) {
			return expEdges[i].Start.(*ssspElement).V.(string) < expEdges[j].Start.(*ssspElement).V.(string)
		}
		return expEdges[i].End.(*ssspElement).V.(string) < expEdges[j].End.(*ssspElement).V.(string)
	})

	sort.Slice(expVertices, func(i, j int) bool {
		return expVertices[i].(*ssspElement).V.(string) < expVertices[j].(*ssspElement).V.(string)
	})

	compareGraph(t, vertexes, expVertices, edges, expEdges)

}

func TestBellManFord(t *testing.T) {
	g := newAdjacencyListWithWeight()
	ssspSetup(g)
	ssspG := bellmanFord(g, "s", math.MaxInt32, new(defaultRelax))
	ssspGExp := ssspGolden(g)
	checkSsspOutOfOrder(t, ssspG, ssspGExp)
}
