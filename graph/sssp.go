package graph

import "container/list"

type relax interface {
	Compare(*ssspElement, *ssspElement, int) bool
	Relax(*ssspElement, *ssspElement, int) bool
}

type ssspElement struct {
	D int
	P *ssspElement
	V interface{}
}

func (e *ssspElement) init(v interface{}, d int) *ssspElement {
	e.V = v
	e.D = d
	e.P = nil
	return e
}

func newSsspElement(v interface{}, d int) *ssspElement {
	return new(ssspElement).init(v, d)
}

func initSingleSource(g graph, d int) map[interface{}]*ssspElement {
	ssspE := make(map[interface{}]*ssspElement)
	for _, v := range g.AllVertices() {
		ssspE[v] = newSsspElement(v, d)
	}
	return ssspE
}

type defaultRelax struct {
	relax
}

func (r *defaultRelax) Compare(start, end *ssspElement, weight int) bool {
	return end.D > start.D+weight
}

func (r *defaultRelax) Relax(start, end *ssspElement, weight int) bool {
	if r.Compare(start, end, weight) {
		end.D = start.D + weight
		end.P = start
		return true
	}
	return false
}

func bellmanFord(g weightedGraph, s interface{}, init int, r relax) weightedGraph {
	ssspG := createGraphByType(g).(weightedGraph)
	ssspE := initSingleSource(g, init)
	ssspE[s].D = 0
	//dp
	for i := 0; i < len(ssspE)-1; i++ {
		for _, e := range g.AllEdges() {
			r.Relax(ssspE[e.Start], ssspE[e.End], g.Weight(e))
		}
	}
	for _, e := range g.AllEdges() {
		if !r.Compare(ssspE[e.Start], ssspE[e.End], g.Weight(e)) {
			if ssspE[e.End].P != nil {
				ssspG.AddEdgeWithWeight(edge{ssspE[e.End].P, ssspE[e.End]}, g.Weight(e))
			}
		} else {
			return nil
		}
	}
	return ssspG
}

func spfa(g weightedGraph, s interface{}, init int, r relax) weightedGraph {
	ssspG := createGraphByType(g).(weightedGraph)
	ssspE := initSingleSource(g, init)
	ssspE[s].D = 0
	//use queue
	queue := list.New()
	queue.PushBack(ssspE[s])
	for queue.Len() != 0 {
		v := queue.Front().Value.(*ssspElement)
		iter := g.IterConnectedVertices(v.V)
		for e := iter.Value(); e != nil; e = iter.Next() {
			if r.Relax(v, ssspE[e], g.Weight(edge{v.V, e})) {
				queue.PushBack(ssspE[e])
			}
		}
		queue.Remove(queue.Front())
	}
	for _, e := range g.AllEdges() {
		if !r.Compare(ssspE[e.Start], ssspE[e.End], g.Weight(e)) {
			if ssspE[e.End].P != nil {
				ssspG.AddEdgeWithWeight(edge{ssspE[e.End].P, ssspE[e.End]}, g.Weight(e))
			}
		} else {
			return nil
		}
	}
	return ssspG
}
