package graph

import (
	"container/list"
)

type DFSElement struct {
	Color int
	D, F  int
	P     *DFSElement
	V     interface{}
}

func (e *DFSElement) Init(v interface{}) *DFSElement {
	e.V = v
	e.Color = WHITE
	e.D = 0
	e.F = 0
	e.P = nil
	return e
}

//path compression, inspiration from disjoint-set
func (e *DFSElement) FindRoot() *DFSElement {
	_e := e
	for _e.P != nil {
		if _e.P.P != nil {
			_e.P = _e.P.P
		}
		_e = _e.P
	}
	return _e

}

func NewDFSElement(v interface{}) *DFSElement {
	return new(DFSElement).Init(v)
}


func DFS(g Graph, sorter func([]interface{})) (dfsGraph map[string]Graph) {
	dfsGraph = make(map[string]Graph)
	dfsGraph["dfsForest"] = CreateGraphByType(g)
	dfsGraph["dfsBackEdges"] = CreateGraphByType(g)
	dfsGraph["dfsForwardEdges"] = CreateGraphByType(g)
	dfsGraph["dfsCrossEdges"] = CreateGraphByType(g)

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
			for i := range dfsGraph {
				dfsGraph[i].AddVertex(elements.get(v).(*DFSElement))
			}
			stack.PushBack(elements.get(v).(*DFSElement))

			for stack.Len() != 0 {
				e := stack.Back().Value.(*DFSElement);
				for _, c := range g.AllConnectedVertices(e.V) {
					if elements.get(c).(*DFSElement).Color == WHITE {
						// parent in deeper path always override that in shallower
						elements.get(c).(*DFSElement).Color = GRAY
						timer ++
						elements.get(c).(*DFSElement).D = timer
						elements.get(c).(*DFSElement).P = e
						for i := range dfsGraph {
							dfsGraph[i].AddVertex(elements.get(c).(*DFSElement))
						}
						//tree edge definition. First time visit
						dfsGraph["dfsForest"].AddEdge(Edge{e, elements.get(c).(*DFSElement)})
						stack.PushBack(elements.get(c).(*DFSElement))
						break
					} else if elements.get(c).(*DFSElement).Color == GRAY {
						// if color is already gray, it's a back edge
						dfsGraph["dfsBackEdges"].AddEdge(Edge{e, elements.get(c).(*DFSElement)})
					} else if e.D > elements.get(c).(*DFSElement).D {
						// if color is already black, it's a cross edge,d(e) > d(elements.get(c).(*DFSElement))
						dfsGraph["dfsCrossEdges"].AddEdge(Edge{e, elements.get(c).(*DFSElement)})
					} else if elements.get(c).(*DFSElement).D - e.D > 1{
						// if color is already black, it's a cross edge,d(e) < d(elements.get(c).(*DFSElement)) - 1
						dfsGraph["dfsForwardEdges"].AddEdge(Edge{e, elements.get(c).(*DFSElement)})
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

func GetDFSComponent(dfsGraph Graph)(map[*DFSElement]Graph) {
	forest := make(map[*DFSElement]Graph)
	for _, v := range dfsGraph.AllVertices() {
		root := v.(*DFSElement).FindRoot()
		if _,ok := forest[root];!ok {
			forest[root] = CreateGraphByType(dfsGraph)
		}
		forest[root].AddVertex(v)
		for _, e := range dfsGraph.AllConnectedVertices(v) {
			if e.(*DFSElement).FindRoot() == root {
				forest[root].AddEdge(Edge{v, e})
			}
		}
	}
	return forest
}