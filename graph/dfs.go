package graph

import (
	"container/list"
)

type dfsElement struct {
	Color   int
	D, F    int
	P, Root *dfsElement
	V       interface{}
}

func (e *dfsElement) Init(v interface{}) *dfsElement {
	e.V = v
	e.Color = white
	e.D = 0
	e.F = 0
	e.P = nil
	e.Root = e
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

func dfs(g graph, sorter func([]interface{})) (dfsForest *dfsForest) {
	dfsForest = newDFSForest(g)
	timer := 0
	stack := list.New()
	//to keep vertices order
	elements := new(linkedMap).init()
	//init
	vertices := g.AllVertices()
	if sorter != nil {
		sorter(vertices)
	}
	for _, v := range vertices {
		elements.add(v, newDFSElement(v))
	}
	pushStack := func(v interface{}) {
		//push root vertex to stack
		elements.get(v).(*dfsElement).Color = gray
		timer++
		elements.get(v).(*dfsElement).D = timer
		dfsForest.AddVertex(elements.get(v).(*dfsElement))
		stack.PushBack(elements.get(v).(*dfsElement))
	}
	for v := elements.frontKey(); v != nil; v = elements.nextKey(v) {
		if elements.get(v).(*dfsElement).Color == white {
			pushStack(v)

			for stack.Len() != 0 {
				e := stack.Back().Value.(*dfsElement)
				for c := range g.IterConnectedVertices(e.V) {
					if elements.get(c).(*dfsElement).Color == white {
						// parent in deeper path always override that in shallower
						elements.get(c).(*dfsElement).P = e
						elements.get(c).(*dfsElement).Root = e
						pushStack(c)
						//tree edge definition. First time visit
						dfsForest.AddTreeEdge(edge{e, elements.get(c).(*dfsElement)})
						break
					} else if elements.get(c).(*dfsElement).Color == gray {
						// if color is already gray, it's a back edge
						dfsForest.AddBackEdge(edge{e, elements.get(c).(*dfsElement)})
					} else if e.D > elements.get(c).(*dfsElement).D {
						// if color is already black, it's a cross edge,d(e) > d(elements.get(c).(*dfsElement))
						dfsForest.AddCrossEdge(edge{e, elements.get(c).(*dfsElement)})
					} else if elements.get(c).(*dfsElement).D-e.D > 1 {
						// if color is already black, it's a cross edge,d(e) < d(elements.get(c).(*dfsElement)) - 1
						dfsForest.AddForwardEdge(edge{e, elements.get(c).(*dfsElement)})
					}
				}
				if e == stack.Back().Value.(*dfsElement) {
					// if the stack did not grow, it is end-point vertex, finish visit process and pop stack
					e.Color = black
					timer++
					e.F = timer
					stack.Remove(stack.Back())
				}
			}
		}
	}
	return
}
