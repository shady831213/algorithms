package graph

import (
	"container/list"
	"math"
)

type bfsElement struct {
	Color int
	Dist  int
	P     *bfsElement
	V     interface{}
}

func (e *bfsElement) Init(v interface{}) *bfsElement {
	e.V = v
	e.Color = WHITE
	e.Dist = math.MaxInt32
	e.P = nil
	return e
}

func newBFSElement(v interface{}) *bfsElement {
	return new(bfsElement).Init(v)
}

func bfs(g graph, s interface{}) (bfsGraph graph) {
	bfsGraph = createGraphByType(g)

	elements := make(map[interface{}]*bfsElement)
	queue := list.New()
	for _, v := range g.AllVertices() {
		elements[v] = newBFSElement(v)
		bfsGraph.AddVertex(elements[v])
	}

	elements[s].Color = GRAY
	elements[s].Dist = 0
	queue.PushBack(s)

	for queue.Len() != 0 {
		qe := queue.Front()
		for v := range g.IterConnectedVertices(qe.Value) {
			if elements[v].Color == WHITE {
				elements[v].Color = GRAY
				elements[v].Dist = elements[qe.Value].Dist + 1
				elements[v].P = elements[qe.Value]
				bfsGraph.AddEdge(edge{elements[qe.Value], elements[v]})
				queue.PushBack(v)
			}
		}
		elements[qe.Value].Color = BLACK
		queue.Remove(qe)
	}
	return
}
