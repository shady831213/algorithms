package graph

type relax interface {
	Compare(*ssspElement, *ssspElement, int) bool
	Relax(*ssspElement, *ssspElement, int)
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

func (r *defaultRelax) Relax(start, end *ssspElement, weight int) {
	if r.Compare(start, end, weight) {
		end.D = start.D + weight
		end.P = start
	}
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
