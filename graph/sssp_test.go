package graph

import (
	"reflect"
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

func ssspPosSetup(g weightedGraph) {
	g.AddEdgeWithWeight(edge{"s", "t"}, 10)
	g.AddEdgeWithWeight(edge{"s", "y"}, 5)
	g.AddEdgeWithWeight(edge{"t", "y"}, 2)
	g.AddEdgeWithWeight(edge{"y", "t"}, 3)
	g.AddEdgeWithWeight(edge{"t", "x"}, 1)
	g.AddEdgeWithWeight(edge{"y", "x"}, 9)
	g.AddEdgeWithWeight(edge{"y", "z"}, 2)
	g.AddEdgeWithWeight(edge{"x", "z"}, 4)
	g.AddEdgeWithWeight(edge{"z", "x"}, 6)
	g.AddEdgeWithWeight(edge{"z", "s"}, 7)
}

func ssspGolden() weightedGraph {
	ssspG := newWeightedGraph()
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

func ssspPosGolden() weightedGraph {
	ssspG := newWeightedGraph()
	ssspE := make(map[interface{}]*ssspElement)

	ssspE["s"] = newSsspElement("s", 0)

	ssspE["y"] = newSsspElement("y", 5)
	ssspE["y"].P = ssspE["s"]

	ssspE["t"] = newSsspElement("t", 8)
	ssspE["t"].P = ssspE["y"]

	ssspE["x"] = newSsspElement("x", 9)
	ssspE["x"].P = ssspE["t"]

	ssspE["z"] = newSsspElement("z", 7)
	ssspE["z"].P = ssspE["y"]

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
	g := newWeightedGraph()
	ssspSetup(g)
	ssspG := bellmanFord(g, "s", new(defaultRelax))
	ssspGExp := ssspGolden()
	checkSsspOutOfOrder(t, ssspG, ssspGExp)
}

func TestSpfa(t *testing.T) {
	g := newWeightedGraph()
	ssspSetup(g)
	ssspG := spfa(g, "s", new(defaultRelax))
	ssspGExp := ssspGolden()
	checkSsspOutOfOrder(t, ssspG, ssspGExp)
}

func TestDijkstra(t *testing.T) {
	g := newWeightedGraph()
	ssspPosSetup(g)
	ssspG := dijkstra(g, "s", new(defaultRelax))
	ssspGExp := ssspPosGolden()
	checkSsspOutOfOrder(t, ssspG, ssspGExp)
}

func TestGabow(t *testing.T) {
	g := newWeightedGraph()
	ssspPosSetup(g)
	ssspG := gabow(g, "s", new(defaultRelax), 4)
	ssspGExp := ssspPosGolden()
	checkSsspOutOfOrder(t, ssspG, ssspGExp)
}

/*
problems
*/

func TestNestedBoxes(t *testing.T) {
	boxes := make([][]int, 9, 9)
	boxes[0] = []int{0, 10, 20, 30, 40}
	boxes[1] = []int{9, 10, 11, 31, 41}
	boxes[2] = []int{1, 11, 22, 33, 41}
	boxes[3] = []int{2, 13, 24, 34, 42}
	boxes[4] = []int{1, 13, 22, 34, 42}
	boxes[5] = []int{10, 11, 22, 34, 42}
	boxes[6] = []int{9, 14, 15, 32, 44}
	boxes[7] = []int{10, 11, 12, 32, 45}
	boxes[8] = []int{0, 10, 12, 13, 8}

	expBoxes := make([][]int, 3, 3)
	expBoxes[2] = []int{0, 10, 20, 30, 40}
	expBoxes[1] = []int{1, 11, 22, 33, 41}
	expBoxes[0] = []int{2, 13, 24, 34, 42}

	seqs := nestedBoxes(boxes)

	if !reflect.DeepEqual(seqs, expBoxes) {
		t.Log("exp :", expBoxes, "actaul : ", seqs)
		t.Fail()
	}

}

func TestKarp(t *testing.T) {
	g := newWeightedGraph()
	g.AddEdgeWithWeight(edge{1, 2}, 1)
	g.AddEdgeWithWeight(edge{2, 3}, 3)
	g.AddEdgeWithWeight(edge{1, 3}, 10)
	g.AddEdgeWithWeight(edge{3, 4}, 2)
	g.AddEdgeWithWeight(edge{4, 1}, 8)
	g.AddEdgeWithWeight(edge{4, 2}, 0)
	u := karp(g, 1)
	if u != 5.0/3.0 {
		t.Log("exp :", 5.0/3.0, "actual:", u)
		t.Fail()
	}
}
