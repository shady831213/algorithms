package graph

import (
	"sort"
	"testing"
)

func flowGraphSetup(g flowGraph) {
	g.AddEdgeWithCap(edge{"s", "v1"}, 16)
	g.AddEdgeWithCap(edge{"s", "v2"}, 13)
	g.AddEdgeWithCap(edge{"v2", "v1"}, 4)
	g.AddEdgeWithCap(edge{"v1", "v2"}, 10)
	g.AddEdgeWithCap(edge{"v1", "v3"}, 12)
	g.AddEdgeWithCap(edge{"v3", "v2"}, 9)
	g.AddEdgeWithCap(edge{"v2", "v4"}, 14)
	g.AddEdgeWithCap(edge{"v4", "v3"}, 7)
	g.AddEdgeWithCap(edge{"v3", "t"}, 20)
	g.AddEdgeWithCap(edge{"v4", "t"}, 4)
}

func flowGraphGolden(g flowGraph) flowGraph {
	flowG := newFlowGraph()
	for _, e := range g.AllEdges() {
		flowG.AddEdgeWithCap(e, g.Cap(e))
	}
	flowG.AddEdgeWithFlow(edge{"s", "v1"}, 12)
	flowG.AddEdgeWithFlow(edge{"s", "v2"}, 11)
	flowG.AddEdgeWithFlow(edge{"v2", "v1"}, 0)
	flowG.AddEdgeWithFlow(edge{"v1", "v2"}, 0)
	flowG.AddEdgeWithFlow(edge{"v1", "v3"}, 12)
	flowG.AddEdgeWithFlow(edge{"v3", "v2"}, 0)
	flowG.AddEdgeWithFlow(edge{"v2", "v4"}, 11)
	flowG.AddEdgeWithFlow(edge{"v4", "v3"}, 7)
	flowG.AddEdgeWithFlow(edge{"v3", "t"}, 19)
	flowG.AddEdgeWithFlow(edge{"v4", "t"}, 4)
	return flowG
}

func checkFlowGraphOutOfOrder(t *testing.T, g, gGolden flowGraph) {
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
	for _, e := range expEdges {
		if g.Cap(e) != gGolden.Cap(e) {
			t.Log(e, "Cap Error! exp :", gGolden.Cap(e), "actual: ", g.Cap(e))
			t.FailNow()
		}
		if g.Flow(e) != gGolden.Flow(e) {
			t.Log(e, "Flow Error! exp :", gGolden.Flow(e), "actual: ", g.Flow(e))
			t.FailNow()
		}
	}
}

func TestEdmondsKarp(t *testing.T) {
	g := newFlowGraph()
	flowGraphSetup(g)
	edmondsKarp(g, "s", "t")
	gGolden := flowGraphGolden(g)
	checkFlowGraphOutOfOrder(t, g, gGolden)
}
