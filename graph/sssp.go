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

func addSsspGEdge(g, ssspG weightedGraph, ssspE *ssspElement) {
	if ssspE.P != nil {
		ssspG.AddEdgeWithWeight(edge{ssspE.P, ssspE}, g.Weight(edge{ssspE.P, ssspE}))
	}
}

func checkOrGetSsspGEdge(g weightedGraph, ssspE map[interface{}]*ssspElement, r relax) weightedGraph {
	ssspG := newWeightedGraph()
	for _, e := range g.AllEdges() {
		if ssspE == nil || r.Compare(ssspE[e.Start], ssspE[e.End], g.Weight(e)) {
			return nil
		}
		addSsspGEdge(g, ssspG, ssspE[e.End])
	}
	return ssspG
}

func getSsspGEdge(g weightedGraph, ssspE map[interface{}]*ssspElement) weightedGraph {
	ssspG := newWeightedGraph()
	for _, e := range g.AllEdges() {
		addSsspGEdge(g, ssspG, ssspE[e.End])
	}
	return ssspG
}

func ssspWrapper(core func(weightedGraph, interface{}, relax) map[interface{}]*ssspElement) func(weightedGraph, interface{}, relax) weightedGraph {
	return func(g weightedGraph, s interface{}, r relax) weightedGraph {
		return checkOrGetSsspGEdge(g, core(g, s, r), r)
	}
}

func ssspPosWeightWrapper(core func(weightedGraph, interface{}, relax) map[interface{}]*ssspElement) func(weightedGraph, interface{}, relax) weightedGraph {
	return func(g weightedGraph, s interface{}, r relax) weightedGraph {
		return getSsspGEdge(g, core(g, s, r))
	}
}

func bellmanFordCore(g weightedGraph, s interface{}, r relax) map[interface{}]*ssspElement {
	ssspE := initSingleSource(g, r.InitValue())
	ssspE[s].D = 0
	//dp
	for i := 0; i < len(ssspE)-1; i++ {
		for _, e := range g.AllEdges() {
			r.Relax(ssspE[e.Start], ssspE[e.End], g.Weight(e))
		}
	}

	return ssspE
}
func bellmanFord(g weightedGraph, s interface{}, r relax) weightedGraph {
	return ssspWrapper(bellmanFordCore)(g, s, r)
}

func spfaCore(g weightedGraph, s interface{}, r relax) map[interface{}]*ssspElement {
	ssspE := initSingleSource(g, r.InitValue())
	ssspE[s].D = 0
	visit := make(map[interface{}]int)
	//use queue
	queue := list.New()
	queue.PushBack(ssspE[s])
	for queue.Len() != 0 {
		v := queue.Front().Value.(*ssspElement)
		iter := g.IterConnectedVertices(v.V)
		for e := iter.Value(); e != nil; e = iter.Next() {
			if r.Relax(v, ssspE[e], g.Weight(edge{v.V, e})) {
				if cnt, ok := visit[e]; ok && cnt > len(visit) {
					// if visit cnt larger than vertices number, it means neg loop detected
					return nil
				}
				queue.PushBack(ssspE[e])
				visit[e]++
			}
		}
		queue.Remove(queue.Front())
	}
	return ssspE
}

func spfa(g weightedGraph, s interface{}, r relax) weightedGraph {
	return ssspWrapper(spfaCore)(g, s, r)
}

func dijkstraCore(g weightedGraph, s interface{}, r relax) map[interface{}]*ssspElement {
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
		for e := iter.Value(); e != nil; e = iter.Next() {
			if _, ok := pqElement[e]; ok {
				currentEdge := edge{v, e}
				r.Relax(ssspE[v], ssspE[e], g.Weight(currentEdge))
				pq.ModifyNode(pqElement[e], ssspE[e].D, e)
			}
		}
	}

	return ssspE
}

func dijkstra(g weightedGraph, s interface{}, r relax) weightedGraph {
	return ssspPosWeightWrapper(dijkstraCore)(g, s, r)
}

func gabow(g weightedGraph, s interface{}, r relax, k uint32) weightedGraph {
	degree := k
	if degree == 0 {
		degree = 32
	}

	ssspE := initSingleSource(g, r.InitValue())
	updateSsspE := func(currentSspE map[interface{}]*ssspElement) {
		for v := range currentSspE {
			if ssspE[v].D != r.InitValue() {
				currentSspE[v].D = currentSspE[v].D + (ssspE[v].D << 1)
			}
			ssspE[v] = currentSspE[v]
		}
	}

	gi := newWeightedGraph()
	updateGi := func(j uint32) {
		for _, e := range g.AllEdges() {
			gi.AddEdgeWithWeight(e, (g.Weight(e)>>j)+((ssspE[e.Start].D-ssspE[e.End].D)<<1))
		}
	}

	for i := uint32(0); i < degree; i++ {
		updateGi(degree - i - 1)
		currentSspE := spfaCore(gi, s, r)
		updateSsspE(currentSspE)
	}

	return getSsspGEdge(g, ssspE)
}

/*
problems
*/
//nested Boxes
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
	g := newWeightedGraph()
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
	dijkstraCore(g, root, nestedBoxesR)
	//output sequence
	seq := make([][]int, 0, 0)
	for e := nestedBoxesR.lastE; e.V != root; e = e.P {
		seq = append(seq, *e.V.(*[]int))
	}

	return seq
}

//karp
type karpElement struct {
	k     int
	u     float32
	ssspE []*ssspElement //array keep for each k, dp
}

func (e *karpElement) init(n int, v interface{}, init int) *karpElement {
	e.ssspE = make([]*ssspElement, n+1, n+1)
	for i := range e.ssspE {
		e.ssspE[i] = newSsspElement(v, init)
	}
	e.k = 0
	e.u = math.MinInt32
	return e
}

func (e *karpElement) getSsspE() *ssspElement {
	return e.ssspE[e.k]
}

func (e *karpElement) summary() {
	//max((D_n - D_k)/(n -k)) k >= 0 && k <= n - 1
	for i := range e.ssspE[:len(e.ssspE)-1] {
		if max := float32(e.ssspE[len(e.ssspE)-1].D-e.ssspE[i].D) / float32((len(e.ssspE) - 1 - i)); max > e.u {
			e.u = max
		}

	}
}

func karp(g weightedGraph, s interface{}) float32 {
	karpE := make(map[interface{}]*karpElement)
	r := new(defaultRelax)
	vertices := g.AllVertices()
	for _, v := range vertices {
		karpE[v] = new(karpElement).init(len(vertices), v, r.InitValue())
	}

	karpE[s].ssspE[0].D = 0
	//use spfa and dp
	queue := list.New()
	queue.PushBack(karpE[s])
	for queue.Len() != 0 {
		ke := queue.Front().Value.(*karpElement)
		v := ke.getSsspE()
		iter := g.IterConnectedVertices(v.V)
		for e := iter.Value(); e != nil; e = iter.Next() {
			if ke.k < len(vertices) {
				//count distance
				karpE[e].k = ke.k + 1
				if r.Relax(v, karpE[e].getSsspE(), g.Weight(edge{v.V, e})) {
					queue.PushBack(karpE[e])
				}
			}
		}
		queue.Remove(queue.Front())
	}

	//get result
	//min(karpE.u)
	u := float32(math.MaxInt32)
	for v := range karpE {
		if karpE[v].summary(); karpE[v].u < u {
			u = karpE[v].u
		}

	}

	return u
}
