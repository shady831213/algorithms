package graph

import (
	"container/list"
)

type dfsElement struct {
	Color   int
	D, F    int
	P, Root *dfsElement
	V       interface{}
	Iter    iterator
}

func (e *dfsElement) Init(v interface{}) *dfsElement {
	e.V = v
	e.Color = white
	e.D = 0
	e.F = 0
	e.P = nil
	e.Root = e
	e.Iter = nil
	return e
}

//path compression, inspiration from disjoint-set
func (e *dfsElement) FindRoot() *dfsElement {
	_e := e
	for _e.Root != _e {
		if _e.Root.Root != _e.Root {
			_e.Root = _e.Root.Root
		}
		_e = _e.Root
	}
	return _e

}

func newDFSElement(v interface{}) *dfsElement {
	return new(dfsElement).Init(v)
}

type dfsForest struct {
	Trees, BackEdges, ForwardEdges, CrossEdges graph
	Comps                                      map[*dfsElement]*dfsForest
}

func (f *dfsForest) Init(g graph) *dfsForest {
	f.Trees = createGraphByType(g)
	f.BackEdges = createGraphByType(g)
	f.ForwardEdges = createGraphByType(g)
	f.CrossEdges = createGraphByType(g)
	f.Comps = make(map[*dfsElement]*dfsForest)
	return f
}

func (f *dfsForest) addVertex(v *dfsElement) {
	f.Trees.AddVertex(v)
	f.BackEdges.AddVertex(v)
	f.ForwardEdges.AddVertex(v)
	f.CrossEdges.AddVertex(v)
}

func (f *dfsForest) addTreeEdge(e edge) {
	f.Trees.AddEdge(e)
}

func (f *dfsForest) addForwardEdge(e edge) {
	f.ForwardEdges.AddEdge(e)
}

func (f *dfsForest) addBackEdge(e edge) {
	f.BackEdges.AddEdge(e)
}

func (f *dfsForest) addCrossEdge(e edge) {
	f.CrossEdges.AddEdge(e)
}

func (f *dfsForest) getRoot(e *dfsElement) *dfsElement {
	root := e.FindRoot()
	if _, ok := f.Comps[root]; !ok {
		f.Comps[root] = newDFSForest(f.Trees)
	}
	return root
}

func (f *dfsForest) AddVertex(v *dfsElement) {
	f.addVertex(v)
	f.Comps[f.getRoot(v)].addVertex(v)
}

func (f *dfsForest) AddTreeEdge(e edge) {
	f.addTreeEdge(e)
	root := f.getRoot(e.Start.(*dfsElement))
	if root == f.getRoot(e.End.(*dfsElement)) {
		f.Comps[root].addTreeEdge(e)
	}
}

func (f *dfsForest) AddForwardEdge(e edge) {
	f.addForwardEdge(e)
	root := f.getRoot(e.Start.(*dfsElement))
	if root == f.getRoot(e.End.(*dfsElement)) {
		f.Comps[root].addForwardEdge(e)
	}
}

func (f *dfsForest) AddBackEdge(e edge) {
	f.addBackEdge(e)
	root := f.getRoot(e.Start.(*dfsElement))
	if root == f.getRoot(e.End.(*dfsElement)) {
		f.Comps[root].addBackEdge(e)
	}
}

func (f *dfsForest) AddCrossEdge(e edge) {
	f.addCrossEdge(e)
	root := f.getRoot(e.Start.(*dfsElement))
	if root == f.getRoot(e.End.(*dfsElement)) {
		f.Comps[root].addCrossEdge(e)
	}
}

func (f *dfsForest) AllTreeEdges() []edge {
	return f.Trees.AllEdges()
}

func (f *dfsForest) AllBackEdges() []edge {
	return f.BackEdges.AllEdges()
}

func (f *dfsForest) AllForwardEdges() []edge {
	return f.ForwardEdges.AllEdges()
}

func (f *dfsForest) AllCrossEdges() []edge {
	return f.CrossEdges.AllEdges()
}

func (f *dfsForest) AllVertices() []interface{} {
	return f.Trees.AllVertices()
}

func (f *dfsForest) AllEdges() []edge {
	edges := f.AllTreeEdges()
	edges = append(edges, f.AllForwardEdges()...)
	edges = append(edges, f.AllBackEdges()...)
	edges = append(edges, f.AllCrossEdges()...)
	return edges
}

func newDFSForest(g graph) *dfsForest {
	return new(dfsForest).Init(g)
}

type dfsVisitHandler struct {
	TreeEdgeHandler, BackEdgeHandler, ForwardEdgeHandler, CrossEdgeHandler func(*dfsElement, *dfsElement)
	BeforeDfsHandler, AfterDfsHandler                                      func(*dfsElement)
	Elements                                                               *linkedMap
	timer                                                                  int
}

func (h *dfsVisitHandler) init() *dfsVisitHandler {
	h.TreeEdgeHandler = func(start *dfsElement, end *dfsElement) {}
	h.BackEdgeHandler = func(start *dfsElement, end *dfsElement) {}
	h.CrossEdgeHandler = func(start *dfsElement, end *dfsElement) {}
	h.ForwardEdgeHandler = func(start *dfsElement, end *dfsElement) {}
	h.BeforeDfsHandler = func(*dfsElement) {}
	h.AfterDfsHandler = func(*dfsElement) {}
	h.Elements = new(linkedMap).init()
	h.timer = 0
	return h
}

func (h *dfsVisitHandler) Counting() int {
	h.timer++
	return h.timer
}

func newDFSVisitHandler() *dfsVisitHandler {
	return new(dfsVisitHandler).init()
}

func dfsVisit(g graph, v interface{}, handler *dfsVisitHandler) {
	if handler == nil {
		panic("handler is nil!")
	}
	stack := list.New()
	//handle handlers
	pushStack := func(v interface{}) *dfsElement {
		//push root vertex to stack

		newE := newDFSElement(v)
		newE.Color = gray
		newE.D = handler.Counting()
		newE.Iter = g.IterConnectedVertices(v)
		handler.Elements.add(v, newE)
		stack.PushBack(newE)
		handler.BeforeDfsHandler(newE)
		return newE
	}

	popStack := func(e *dfsElement) {
		e.Color = black
		e.F = handler.Counting()
		e.Iter = g.IterConnectedVertices(e.V)
		stack.Remove(stack.Back())
		handler.AfterDfsHandler(e)
	}

	pushStack(v)
	for stack.Len() != 0 {
		e := stack.Back().Value.(*dfsElement)
		for c := handler.Elements.get(e.V).(*dfsElement).Iter.Value(); c != nil; {
			if !handler.Elements.exist(c) {
				// parent in deeper path always override that in shallower
				newE := pushStack(c)
				newE.P = e
				newE.Root = e
				//tree edge definition. First time visit
				if handler != nil {
					handler.TreeEdgeHandler(e, newE)
				}
				break
			} else if handler.Elements.get(c).(*dfsElement).Color == gray {
				// if color is already gray, it's a back edge
				handler.BackEdgeHandler(e, handler.Elements.get(c).(*dfsElement))
			} else if e.D > handler.Elements.get(c).(*dfsElement).D {
				// if color is already black, it's a cross edge,d(e) > d(Elements.get(c).(*dfsElement))
				handler.CrossEdgeHandler(e, handler.Elements.get(c).(*dfsElement))
			} else if handler.Elements.get(c).(*dfsElement).D-e.D > 1 {
				// if color is already black, it's a forward edge,d(e) < d(Elements.get(c).(*dfsElement)) - 1
				handler.ForwardEdgeHandler(e, handler.Elements.get(c).(*dfsElement))
			}
			c = handler.Elements.get(e.V).(*dfsElement).Iter.Next()
		}
		if e == stack.Back().Value.(*dfsElement) {
			// if the stack did not grow, it is end-point vertex, finish visit process and pop stack
			popStack(e)
		}
	}
}

func dfs(g graph, sorter func([]interface{})) (dfsForest *dfsForest) {
	dfsForest = newDFSForest(g)
	handler := newDFSVisitHandler()
	handler.BeforeDfsHandler = func(v *dfsElement) {
		dfsForest.AddVertex(v)
	}
	handler.TreeEdgeHandler = func(start, end *dfsElement) {
		dfsForest.AddTreeEdge(edge{start, end})
	}
	handler.BackEdgeHandler = func(start, end *dfsElement) {
		dfsForest.AddBackEdge(edge{start, end})
	}
	handler.ForwardEdgeHandler = func(start, end *dfsElement) {
		dfsForest.AddForwardEdge(edge{start, end})
	}
	handler.CrossEdgeHandler = func(start, end *dfsElement) {
		dfsForest.AddCrossEdge(edge{start, end})
	}
	//init
	vertices := g.AllVertices()
	if sorter != nil {
		sorter(vertices)
	}
	for _, v := range vertices {
		if !handler.Elements.exist(v) {
			dfsVisit(g, v, handler)
		}
	}
	return
}

func checkConnectivity(g graph) bool {
	dfsForest := dfs(g, nil)
	return len(dfsForest.Comps) == 1
}

