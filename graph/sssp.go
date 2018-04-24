package graph

import (
	"container/list"
	"github.com/shady831213/algorithms/heap"
	"math"
)

type relax interface {
	InitValue() int
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

func (r *defaultRelax) InitValue() int {
	return math.MaxInt32
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

func bellmanFord(g weightedGraph, s interface{}, r relax) weightedGraph {
	ssspG := createGraphByType(g).(weightedGraph)
	ssspE := initSingleSource(g, r.InitValue())
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

func spfa(g weightedGraph, s interface{}, r relax) weightedGraph {
	ssspG := createGraphByType(g).(weightedGraph)
	ssspE := initSingleSource(g, r.InitValue())
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

func dijkstra(g weightedGraph, s interface{}, r relax) weightedGraph {
	ssspG := createGraphByType(g).(weightedGraph)
	ssspE := initSingleSource(g, r.InitValue())
	ssspE[s].D = 0

	//use fibonacci heap
	pq := newFibHeapKeyInt()
	pqElement := make(map[interface{}]*heap.FibHeapElement)

	for v := range ssspE {
		pqElement[v] = pq.Insert(ssspE[v].D, v)
	}

	for pq.Len() != 0 {
		minElement := pq.ExtractMin()
		v := minElement.Value
		delete(pqElement, v)
		iter := g.IterConnectedVertices(v)
		if ssspE[v].P != nil {
			ssspG.AddEdgeWithWeight(edge{ssspE[v].P, ssspE[v]}, g.Weight(edge{ssspE[v].P.V, v}))
		}
		for e := iter.Value(); e != nil; e = iter.Next() {
			if _, ok := pqElement[e]; ok {
				currentEdge := edge{v, e}
				r.Relax(ssspE[v], ssspE[e], g.Weight(currentEdge))
				pq.ModifyNode(pqElement[e], ssspE[e].D, e)
			}
		}
	}

	return ssspG
}

/*
problems
*/

type nestedBoxesRelax struct {
	maxLen int
	lastE  *ssspElement
	defaultRelax
}

func (r *nestedBoxesRelax) init() *nestedBoxesRelax {
	r.maxLen = 0
	r.lastE = nil
	return r
}

func (r *nestedBoxesRelax) Relax(start, end *ssspElement, weight int) bool {
	update := r.defaultRelax.Relax(start, end, weight)
	if update {
		if end.D < r.maxLen {
			r.maxLen, r.lastE = end.D, end
		}
	}
	return update
}

func nestedBoxes(boxes [][]int) [][]int {
	g := newAdjacencyListWithWeight()
	nested := func(box1, box2 []int) bool {
		if len(box1) != len(box2) {
			return false
		}
		for i := range box1 {
			if box1[i] >= box2[i] {
				return false
			}
		}
		return true
	}
	//build graph
	root := struct{}{}
	for i := range boxes {
		g.AddEdgeWithWeight(edge{root, &boxes[i]}, 0)
		for j := i + 1; j < len(boxes); j++ {
			if nested(boxes[i], boxes[j]) {
				g.AddEdgeWithWeight(edge{&boxes[i], &boxes[j]}, -1)
			} else if nested(boxes[j], boxes[i]) {
				g.AddEdgeWithWeight(edge{&boxes[j], &boxes[i]}, -1)
			}
		}
	}

	//dijkstra
	nestedBoxesR := new(nestedBoxesRelax).init()
	dijkstra(g, root, nestedBoxesR)
	//output sequence
	seq := make([][]int, 0, 0)
	for e := nestedBoxesR.lastE; e.V != root; e = e.P {
		seq = append(seq, *e.V.(*[]int))
	}

	return seq
}
