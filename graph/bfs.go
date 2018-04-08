package graph

import (
	"container/list"
	"math"
)

type BFSElement struct {
	Color int
	Dist  int
	P     *BFSElement
	V     interface{}
}

func (e *BFSElement) Init(v interface{}) *BFSElement {
	e.V = v
	e.Color = WHITE
	e.Dist = math.MaxInt32
	e.P = nil
	return e
}

func NewBFSElement(v interface{}) *BFSElement {
	return new(BFSElement).Init(v)
}

func BFS(g Graph, s interface{}) (bfsGraph Graph) {
	bfsGraph = CreateGraphByType(g)

	elements := make(map[interface{}]*BFSElement)
	queue := list.New()
	for _, v := range g.AllVertices() {
		elements[v] = NewBFSElement(v)
		bfsGraph.AddVertex(elements[v])
	}

	elements[s].Color = GRAY
	elements[s].Dist = 0
	queue.PushBack(s)

	for queue.Len() != 0 {
		qe := queue.Front()
		for _, v := range g.AllConnectedVertices(qe.Value) {
			if elements[v].Color == WHITE {
				elements[v].Color = GRAY
				elements[v].Dist = elements[qe.Value].Dist + 1
				elements[v].P = elements[qe.Value]
				bfsGraph.AddEdge(Edge{elements[qe.Value],elements[v]})
				queue.PushBack(v)
			}
		}
		elements[qe.Value].Color = BLACK
		queue.Remove(qe)
	}
	return
}
