package graph

import (
	"container/list"
	"math"
)

type residualGraph interface {
	flowGraph
	RCap(edge) int
}

type adjacencyMatrixResidual struct {
	adjacencyMatrixWithFlow
}

func (g *adjacencyMatrixResidual) init() *adjacencyMatrixResidual {
	g.adjacencyMatrixWithFlow.init()
	return g
}

func (g *adjacencyMatrixResidual) AddEdgeWithFlow(e edge, f int) {
	g.adjacencyMatrixWithFlow.AddEdgeWithFlow(e, f)
	if g.RCap(e) == 0 {
		//delete edge but keep cap and flow, because if residual flow (RCap) change, this edge should be added back
		g.adjacencyMatrix.DeleteEdge(e)
	}
}

func (g *adjacencyMatrixResidual) RCap(e edge) int {
	return g.Cap(e) - g.Flow(e)
}

func newResidualGraph(g flowGraph) residualGraph {
	residualG := new(adjacencyMatrixResidual).init()
	for _, e := range g.AllEdges() {
		residualG.AddEdgeWithCap(e, g.Cap(e))
	}
	return residualG
}

func augmentingPath(g residualGraph, s interface{}, t interface{}) (int, []edge) {
	augmentingEdges := make([]edge, 0, 0)
	minRC := math.MaxInt32
	handler := newBFSVisitHandler()
	handler.EdgeHandler = func(start, end *bfsElement) {
		if end.V == t {
			for v := end; v.P != nil; v = v.P {
				currentEdge := edge{v.P.V, v.V}
				augmentingEdges = append(augmentingEdges, currentEdge)
				if rc := g.RCap(currentEdge); rc < minRC {
					minRC = rc
				}
			}
		}
	}

	bfsVisit(g, s, handler)

	return minRC, augmentingEdges
}

func updateFlow(rg residualGraph, g flowGraph, rc int, edges []edge) {
	for _, e := range edges {
		flow := g.Flow(e) + rc
		g.AddEdgeWithFlow(e, flow)
		rg.AddEdgeWithFlow(e, flow)
		re := edge{e.End, e.Start}
		rg.AddEdgeWithFlow(re, 0-flow)
	}
}

func edmondesKarp(g flowGraph, s interface{}, t interface{}) {
	residualG := newResidualGraph(g)

	for rc, edges := augmentingPath(residualG, s, t); len(edges) > 0; rc, edges = augmentingPath(residualG, s, t) {
		updateFlow(residualG, g, rc, edges)
	}

}

type preFlowGraph interface {
	residualGraph
	SetHeight(interface{}, int)
	SetExcess(interface{}, int)
	Height(interface{}) int
	Excess(interface{}) int
	Push(edge) bool
	Relabel(interface{}) bool
	Overflow(interface{}) bool
}

type adjacencyMatrixPreFlow struct {
	adjacencyMatrixResidual
	height, excess map[interface{}]int
	s, t           interface{}
}

func (g *adjacencyMatrixPreFlow) init(fg flowGraph, s, t interface{}) *adjacencyMatrixPreFlow {
	g.adjacencyMatrixResidual.init()
	g.height = make(map[interface{}]int)
	g.excess = make(map[interface{}]int)
	g.s = s
	g.t = t

	vertices := fg.AllVertices()
	for _, e := range fg.AllEdges() {
		g.AddEdgeWithCap(e, fg.Cap(e))
	}
	g.SetHeight(s, len(vertices))
	iter := fg.IterConnectedVertices(s)
	for v := iter.Value(); v != nil; v = iter.Next() {
		c := fg.Cap(edge{s, v})
		g.AddEdgeWithFlow(edge{s, v}, c)
		g.AddEdgeWithFlow(edge{v, s}, 0-c)
		g.SetExcess(v, c)
		g.SetExcess(s, g.Excess(s)-c)
	}
	return g
}

func (g *adjacencyMatrixPreFlow) SetHeight(v interface{}, h int) {
	g.height[v] = h
}

func (g *adjacencyMatrixPreFlow) Height(v interface{}) int {
	if _, ok := g.height[v]; !ok {
		return 0
	}
	return g.height[v]
}

func (g *adjacencyMatrixPreFlow) SetExcess(v interface{}, e int) {
	g.excess[v] = e
}

func (g *adjacencyMatrixPreFlow) Excess(v interface{}) int {
	if _, ok := g.excess[v]; !ok {
		return 0
	}
	return g.excess[v]
}

func (g *adjacencyMatrixPreFlow) Push(e edge) bool {
	if g.Overflow(e.Start) && g.RCap(e) > 0 && g.Height(e.Start) == g.Height(e.End)+1 {
		d := g.RCap(e)
		if g.Excess(e.Start) < d {
			d = g.Excess(e.Start)
		}
		flow := g.Flow(e) + d
		g.AddEdgeWithFlow(e, flow)
		re := edge{e.End, e.Start}
		g.AddEdgeWithFlow(re, 0-flow)
		g.SetExcess(e.Start, g.Excess(e.Start)-d)
		g.SetExcess(e.End, g.Excess(e.End)+d)
		return true
	}
	return false
}

func (g *adjacencyMatrixPreFlow) Relabel(v interface{}) bool {
	if !g.Overflow(v) {
		return false
	}
	iter := g.IterConnectedVertices(v)
	minH := math.MaxInt32
	for end := iter.Value(); end != nil; end = iter.Next() {
		if g.Height(end) < minH {
			minH = g.Height(end)
		}
		if g.Height(v) > g.Height(end) {
			return false
		}
	}
	g.SetHeight(v, minH+1)
	return true
}

func (g *adjacencyMatrixPreFlow) Overflow(v interface{}) bool {
	//According to definition, start and target vertex never overflow
	return v != g.s && v != g.t && g.Excess(v) > 0
}

func newPreFlowGraph(g flowGraph, s interface{}, t interface{}) preFlowGraph {
	return new(adjacencyMatrixPreFlow).init(g, s, t)
}

func pushRelabel(g flowGraph, s interface{}, t interface{}) {
	preFlowG := newPreFlowGraph(g, s, t)
	for stop := false; !stop; {
		stop = true
		for _, e := range preFlowG.AllEdges() {
			stop = stop && !preFlowG.Push(e) && !preFlowG.Relabel(e.Start) && !preFlowG.Relabel(e.End)
		}
	}
	for _, e := range g.AllEdges() {
		g.AddEdgeWithFlow(e, preFlowG.Flow(e))
	}
}

type allowedGraph interface {
	preFlowGraph
	Discharge(interface{})
}

type adjacencyMatrixAllowed struct {
	adjacencyMatrixPreFlow
	edges graph
}

func (g *adjacencyMatrixAllowed) init(fg flowGraph, s, t interface{}) *adjacencyMatrixAllowed {
	g.adjacencyMatrixPreFlow.init(fg, s, t)
	g.edges = newGraph()
	for _, e := range fg.AllEdges() {
		g.edges.AddEdgeBi(e)
	}
	return g
}

func (g *adjacencyMatrixAllowed) Discharge(v interface{}) {
	iter := g.edges.IterConnectedVertices(v)
	for g.Overflow(v) {
		if iter.Value() == nil {
			g.Relabel(v)
			iter = g.edges.IterConnectedVertices(v)
		} else if !g.Push(edge{v, iter.Value()}) {
			iter.Next()
		}
	}
}

func newAllowedGraph(g flowGraph, s, t interface{}) allowedGraph {
	return new(adjacencyMatrixAllowed).init(g, s, t)
}

func relabelToFront(g flowGraph, s interface{}, t interface{}) {
	allowedG := newAllowedGraph(g, s, t)
	l := list.New()
	for _, v := range g.AllVertices() {
		if v != s && v != t {
			l.PushBack(v)
		}
	}
	for e := l.Front(); e != nil; {
		oldH := allowedG.Height(e.Value)
		allowedG.Discharge(e.Value)
		if allowedG.Height(e.Value) > oldH {
			l.MoveToFront(e)
		}
		e = e.Next()
	}
	for _, e := range g.AllEdges() {
		g.AddEdgeWithFlow(e, allowedG.Flow(e))
	}
}

func bipGraphMaxMatch(g graph, l []interface{}, flowAlg func(g flowGraph, s interface{}, t interface{})) graph {
	//build flow graph
	fG := newFlowGraph()
	s := struct{ start string }{"s"}
	t := struct{ end string }{"t"}
	for _, vl := range l {
		fG.AddEdgeWithCap(edge{s, vl}, 1)
		iter := g.IterConnectedVertices(vl)
		for rv := iter.Value(); rv != nil; rv = iter.Next() {
			fG.AddEdgeWithCap(edge{vl, rv}, 1)
			fG.AddEdgeWithCap(edge{rv, t}, 1)
		}
	}

	flowAlg(fG, s, t)
	matchG := newGraph()
	for _, e := range fG.AllEdges() {
		if fG.Flow(e) > 0 && e.Start != s && e.End != t {
			matchG.AddEdgeBi(e)
		}
	}
	return matchG
}

type hopcraftKarp struct {
	g              graph // must be directed graph
	dis            int
	xMatch, yMatch map[interface{}]interface{}
	xLevel, yLevel map[interface{}]int
	matches        int
}

func (a *hopcraftKarp) init(g graph) *hopcraftKarp {
	a.g = g
	a.xMatch, a.yMatch = make(map[interface{}]interface{}), make(map[interface{}]interface{})
	a.matches = 0
	return a
}

func (a *hopcraftKarp) bfs() bool {
	a.dis = math.MaxInt32
	a.xLevel, a.yLevel = make(map[interface{}]int), make(map[interface{}]int)
	//use queue
	queue := list.New()
	for _, x := range a.g.AllVertices() {
		if _, ok := a.xMatch[x]; !ok {
			queue.PushBack(x)
			a.xLevel[x] = 0
		}
	}
	for queue.Len() != 0 {
		s := queue.Front().Value
		queue.Remove(queue.Front())
		if v, ok := a.xLevel[s]; ok && v > a.dis {
			break
		}
		iter := a.g.IterConnectedVertices(s)
		for y := iter.Value(); y != nil; y = iter.Next() {
			if _, ok := a.yLevel[y]; !ok {
				a.yLevel[y] = a.xLevel[s] + 1
				if _, ok := a.yMatch[y]; !ok {
					a.dis = a.yLevel[y]
				} else {
					a.xLevel[a.yMatch[y]] = a.yLevel[y] + 1
					queue.PushBack(a.yMatch[y])
				}
			}
		}
	}
	return a.dis != math.MaxInt32
}

func (a *hopcraftKarp) dfs(x interface{}, yVisit map[interface{}]bool) bool {
	iter := a.g.IterConnectedVertices(x)
	for y := iter.Value(); y != nil; y = iter.Next() {
		//fmt.Println(x, y, yVisit, g.yLevel[y], g.xLevel[x], g.xMatch, g.yMatch)
		if _, ok := yVisit[y]; !ok && a.yLevel[y] == a.xLevel[x]+1 {
			yVisit[y] = true
			if _, ok := a.yMatch[y]; ok && a.yLevel[y] == a.dis {
				continue
			}
			if _, ok := a.yMatch[y]; !ok {
				a.xMatch[x] = y
				a.yMatch[y] = x
				return true
			} else if a.dfs(a.yMatch[y], yVisit) {
				a.xMatch[x] = y
				a.yMatch[y] = x
				return true
			}
		}
	}
	return false
}

func (a *hopcraftKarp) maxMatch() int {
	for a.bfs() {
		yVisit := make(map[interface{}]bool)
		for _, x := range a.g.AllVertices() {
			if _, ok := a.xMatch[x]; !ok {
				if a.dfs(x, yVisit) {
					a.matches++
				}
			}
		}
	}
	return a.matches
}
