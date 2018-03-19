package fibHeap

//element
type fibHeapElement struct {
	p, l, r    *fibHeapElement
	c, list    *fibHeapElementList
	mark       bool
	key, value interface{}
}

func (e *fibHeapElement) Init(key, value interface{}) *fibHeapElement {
	e.p = nil
	e.l = e
	e.r = e
	e.c = NewFabHeapElementList(e)
	e.mark = false
	e.key = key
	e.value = value
	return e
}

func (e *fibHeapElement) Degree() int {
	return e.c.Len()
}

func (e *fibHeapElement) AddChild(child *fibHeapElement) {
	child.p = e
	e.c.PushRight(child)
}

func NewFabHeapElement(key, value interface{}) *fibHeapElement {
	return new(fibHeapElement).Init(key, value)
}

//list container
type fibHeapElementList struct {
	p, leftist *fibHeapElement
	len        int
}

func (l *fibHeapElementList) Init(p *fibHeapElement) *fibHeapElementList {
	l.p = p
	l.len = 0
	return l
}

func (l *fibHeapElementList) Len() int {
	return l.len
}

func (l *fibHeapElementList) insert(e, at *fibHeapElement) *fibHeapElement {
	n := at.r
	at.r = e
	e.l = at
	e.r = n
	n.l = e
	e.p = l.p
	l.len++
	return e
}

func (l *fibHeapElementList) Remove(e *fibHeapElement) {
	if e.p == l.p {
		e.l.r = e.r
		e.r.l = e.l
		l.len--
	}
}

func (l *fibHeapElementList) PushLeft(e *fibHeapElement) {
	if l.leftist == nil {
		l.leftist = e
		l.len++
	} else {
		l.insert(e, l.leftist)
	}
}

func (l *fibHeapElementList) PushRight(e *fibHeapElement) {
	if l.leftist == nil {
		l.leftist = e
		l.len++
	} else {
		l.insert(e, l.leftist.l)
	}
}

func (l *fibHeapElementList) Leftist() *fibHeapElement {
	if l.len == 0 {
		return nil
	}
	return l.leftist
}

func (l *fibHeapElementList) Rightist() *fibHeapElement {
	if l.len == 0 {
		return nil
	}
	return l.leftist.l
}

func (l *fibHeapElementList) MergeRightList(other *fibHeapElementList) *fibHeapElementList {
	for i, e := other.Len(), other.Leftist(); i > 0; i, e = i-1, e.r {
		l.insert(e, l.leftist.l)
	}
	return l
}

func (l *fibHeapElementList) MergeLeftList(other *fibHeapElementList) *fibHeapElementList {
	for i, e := other.Len(), other.Rightist(); i > 0; i, e = i-1, e.l {
		l.insert(e, l.leftist)
	}
	return l
}

func NewFabHeapElementList(p *fibHeapElement) *fibHeapElementList {
	return new(fibHeapElementList).Init(p)
}

//heap
type fibHeapIf interface {
	Less(*fibHeapElement, *fibHeapElement) bool
	Swap(*fibHeapElement, *fibHeapElement)
}

type fibHeap struct {
	root      *fibHeapElementList
	min       *fibHeapElement
	maxDegree int
	n         int
	fibHeapIf
}

func (h *fibHeap) Init(self fibHeapIf) *fibHeap {
	h.root = NewFabHeapElementList(nil)
	h.min = nil
	h.maxDegree = 0
	h.n = 0
	h.fibHeapIf = self
	return h
}

//default Less function
func (h *fibHeap) Less(n1, n2 *fibHeapElement) bool {
	if n1 == nil && n2 == nil {
		panic("both nodes are nil!")
	} else if n1 == nil {
		return false
	} else if n2 == nil {
		return true
	}
	return n1.key.(int) < n2.key.(int)
}

func (h *fibHeap) Insert(key, value interface{}) *fibHeapElement {
	n := NewFabHeapElement(key, value)
	h.root.PushRight(NewFabHeapElement(key, value))
	if h.fibHeapIf.Less(n, h.min) {
		h.min = n
	}
	h.n ++
	return n
}

func (h *fibHeap) Union(h1 *fibHeap) *fibHeap {
	h.root = h.root.MergeRightList(h1.root)
	h1.root = nil
	if h.fibHeapIf.Less(h1.min, h.min) {
		h.min = h1.min
	}
	h1.min = nil
	if h.maxDegree < h1.maxDegree {
		h.maxDegree = h1.maxDegree
	}
	h.n += h1.n
	h1 = nil
	return h
}

func (h *fibHeap) consolidate() {

}

func (h *fibHeap) ExtractMin() *fibHeapElement {
	n := h.min
	if n != nil {
		for i, e := n.Degree(), n.c.Leftist(); i > 0; i, e = i-1, e.r {
			h.root.PushLeft(e)
		}
		h.root.Remove(n)
		if n == n.r {
			h.min = nil
		} else {
			h.min = n.r
			h.consolidate()
		}
		h.n--
	}
	return n
}
