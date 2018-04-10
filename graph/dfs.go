package graph

import (
	"container/list"
)

type DFSElement struct {
	Color   int
	D, F    int
	P, Root *DFSElement
	V       interface{}
}

func (e *DFSElement) Init(v interface{}) *DFSElement {
	e.V = v
	e.Color = WHITE
	e.D = 0
	e.F = 0
	e.P = nil
	e.Root = e
	return e
}

//path compression, inspiration from disjoint-set
func (e *DFSElement) FindRoot() *DFSElement {
	_e := e
	for _e.Root != _e {
		if _e.Root.Root != _e.Root {
			_e.Root = _e.Root.Root
		}
		_e = _e.Root
	}
	return _e

}

func NewDFSElement(v interface{}) *DFSElement {
	return new(DFSElement).Init(v)
}

type DFSForest struct {
	Trees, BackEdges, ForwardEdges, CrossEdges Graph
	Comps                                      map[*DFSElement]*DFSForest
}

func (f *DFSForest) Init(g Graph) *DFSForest {
	f.Trees = CreateGraphByType(g)
	f.BackEdges = CreateGraphByType(g)
	f.ForwardEdges = CreateGraphByType(g)
	f.CrossEdges = CreateGraphByType(g)
	f.Comps = make(map[*DFSElement]*DFSForest)
	return f
}

func (f *DFSForest) addVertex(v *DFSElement) {
	f.Trees.AddVertex(v)
	f.BackEdges.AddVertex(v)
	f.ForwardEdges.AddVertex(v)
	f.CrossEdges.AddVertex(v)
}

func (f *DFSForest) addTreeEdge(e Edge) {
	f.Trees.AddEdge(e)
}

func (f *DFSForest) addForwardEdge(e Edge) {
	f.ForwardEdges.AddEdge(e)
}

func (f *DFSForest) addBackEdge(e Edge) {
	f.BackEdges.AddEdge(e)
}

func (f *DFSForest) addCrossEdge(e Edge) {
	f.CrossEdges.AddEdge(e)
}

func (f *DFSForest) getRoot(e *DFSElement) *DFSElement {
	root := e.FindRoot()
	if _, ok := f.Comps[root]; !ok {
		f.Comps[root] = NewDFSForest(f.Trees)
	}
	return root
}

func (f *DFSForest) AddVertex(v *DFSElement) {
	f.addVertex(v)
	f.Comps[f.getRoot(v)].addVertex(v)
}

func (f *DFSForest) AddTreeEdge(e Edge) {
	f.addTreeEdge(e)
	root := f.getRoot(e.Start.(*DFSElement))
	if root == f.getRoot(e.End.(*DFSElement)) {
		f.Comps[root].addTreeEdge(e)
	}
}

func (f *DFSForest) AddForwardEdge(e Edge) {
	f.addForwardEdge(e)
	root := f.getRoot(e.Start.(*DFSElement))
	if root == f.getRoot(e.End.(*DFSElement)) {
		f.Comps[root].addForwardEdge(e)
	}
}

func (f *DFSForest) AddBackEdge(e Edge) {
	f.addBackEdge(e)
	root := f.getRoot(e.Start.(*DFSElement))
	if root == f.getRoot(e.End.(*DFSElement)) {
		f.Comps[root].addBackEdge(e)
	}
}

func (f *DFSForest) AddCrossEdge(e Edge) {
	f.addCrossEdge(e)
	root := f.getRoot(e.Start.(*DFSElement))
	if root == f.getRoot(e.End.(*DFSElement)) {
		f.Comps[root].addCrossEdge(e)
	}
}

func (f *DFSForest) AllTreeEdges() []Edge {
	return f.Trees.AllEdges()
}

func (f *DFSForest) AllBackEdges() []Edge {
	return f.BackEdges.AllEdges()
}

func (f *DFSForest) AllForwardEdges() []Edge {
	return f.ForwardEdges.AllEdges()
}

func (f *DFSForest) AllCrossEdges() []Edge {
	return f.CrossEdges.AllEdges()
}

func (f *DFSForest) AllVertices() []interface{} {
	return f.Trees.AllVertices()
}

func (f *DFSForest) AllEdges() []Edge {
	edges := f.AllTreeEdges()
	edges = append(edges, f.AllForwardEdges()...)
	edges = append(edges, f.AllBackEdges()...)
	edges = append(edges, f.AllCrossEdges()...)
	return edges
}

func NewDFSForest(g Graph) *DFSForest {
	return new(DFSForest).Init(g)
}

func DFS(g Graph, sorter func([]interface{})) (dfsForest *DFSForest) {
	dfsForest = NewDFSForest(g)
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
		elements.add(v, NewDFSElement(v))
	}
	for v := elements.frontKey(); v != nil; v = elements.nextKey(v) {
		if elements.get(v).(*DFSElement).Color == WHITE {
			//push root vertex to stack
			elements.get(v).(*DFSElement).Color = GRAY
			timer ++
			elements.get(v).(*DFSElement).D = timer
			dfsForest.AddVertex(elements.get(v).(*DFSElement))
			stack.PushBack(elements.get(v).(*DFSElement))

			for stack.Len() != 0 {
				e := stack.Back().Value.(*DFSElement);
				for c := range g.IterConnectedVertices(e.V) {
					if elements.get(c).(*DFSElement).Color == WHITE {
						// parent in deeper path always override that in shallower
						elements.get(c).(*DFSElement).Color = GRAY
						timer ++
						elements.get(c).(*DFSElement).D = timer
						elements.get(c).(*DFSElement).P = e
						elements.get(c).(*DFSElement).Root = e
						dfsForest.AddVertex(elements.get(c).(*DFSElement))
						//tree edge definition. First time visit
						dfsForest.AddTreeEdge(Edge{e, elements.get(c).(*DFSElement)})
						stack.PushBack(elements.get(c).(*DFSElement))
						break
					} else if elements.get(c).(*DFSElement).Color == GRAY {
						// if color is already gray, it's a back edge
						dfsForest.AddBackEdge(Edge{e, elements.get(c).(*DFSElement)})
					} else if e.D > elements.get(c).(*DFSElement).D {
						// if color is already black, it's a cross edge,d(e) > d(elements.get(c).(*DFSElement))
						dfsForest.AddCrossEdge(Edge{e, elements.get(c).(*DFSElement)})
					} else if elements.get(c).(*DFSElement).D-e.D > 1 {
						// if color is already black, it's a cross edge,d(e) < d(elements.get(c).(*DFSElement)) - 1
						dfsForest.AddForwardEdge(Edge{e, elements.get(c).(*DFSElement)})
					}
				}
				if e == stack.Back().Value.(*DFSElement) {
					// if the stack did not grow, it is end-point vertex, finish visit process and pop stack
					e.Color = BLACK
					timer ++
					e.F = timer
					stack.Remove(stack.Back())
				}
			}
		}
	}
	return
}
