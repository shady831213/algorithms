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
	Iter  iterator
}

func (e *bfsElement) Init(v interface{}) *bfsElement {
	e.V = v
	e.Color = white
	e.Dist = math.MaxInt32
	e.P = nil
	e.Iter = nil
	return e
}

func newBFSElement(v interface{}) *bfsElement {
	return new(bfsElement).Init(v)
}

type bfsVisitHandler struct {
	BeforeBfsHandler, AfterBfsHandler func(*bfsElement)
	EdgeHandler                       func(*bfsElement, *bfsElement)
	Elements                          map[interface{}]*bfsElement
}

func (h *bfsVisitHandler) init() *bfsVisitHandler {
	h.EdgeHandler = func(start *bfsElement, end *bfsElement) {}
	h.BeforeBfsHandler = func(*bfsElement) {}
	h.AfterBfsHandler = func(*bfsElement) {}
	h.Elements = make(map[interface{}]*bfsElement)
	return h
}

func newBFSVisitHandler() *bfsVisitHandler {
	return new(bfsVisitHandler).init()
}

func bfsVisit(g graph, s interface{}, handler *bfsVisitHandler) {
	if handler == nil {
		panic("handler is nil!")
	}
	queue := list.New()

	pushQueue := func(v interface{}) *bfsElement {
		handler.Elements[v] = newBFSElement(v)
		handler.Elements[v].Color = gray
		handler.Elements[v].Iter = g.IterConnectedVertices(v)
		queue.PushBack(handler.Elements[v])
		handler.BeforeBfsHandler(handler.Elements[v])
		return handler.Elements[v]
	}

	pushQueue(s).Dist = 0

	for queue.Len() != 0 {
		v := queue.Front().Value.(*bfsElement)
		for c := v.Iter.Value(); c != nil; c = v.Iter.Next() {
			if _, ok := handler.Elements[c]; !ok {
				newE := pushQueue(c)
				newE.Dist = v.Dist + 1
				newE.P = v
				handler.EdgeHandler(v, newE)
			}
		}
		v.Color = black
		v.Iter = g.IterConnectedVertices(v.V)
		queue.Remove(queue.Front())
		handler.AfterBfsHandler(v)

	}
}

func bfs(g graph, s interface{}) (bfsGraph graph) {
	bfsGraph = newGraph()
	handler := newBFSVisitHandler()
	handler.EdgeHandler = func(start, end *bfsElement) {
		bfsGraph.AddEdge(edge{start, end})
	}
	bfsVisit(g, s, handler)
	return
}
