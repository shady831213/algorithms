package heap

//FibHeapElement:Cross package
type FibHeapElement struct {
	p, l, r    *FibHeapElement
	c, list    *fibHeapElementList
	mark       bool
	key, value interface{}
}

func (e *FibHeapElement) Init(key, value interface{}) *FibHeapElement {
	e.p = nil
	e.l = nil
	e.r = nil
	e.c = newFabHeapElementList(e)
	e.mark = false
	e.key = key
	e.value = value
	return e
}

func (e *FibHeapElement) Right() *FibHeapElement {
	if p := e.r; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

func (e *FibHeapElement) Left() *FibHeapElement {
	if p := e.l; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

func (e *FibHeapElement) Degree() int {
	return e.c.Len()
}

func (e *FibHeapElement) AddChild(child *FibHeapElement) {
	child.p = e
	e.c.PushRight(child)
}

func newFabHeapElement(key, value interface{}) *FibHeapElement {
	return new(FibHeapElement).Init(key, value)
}

//list container
type fibHeapElementList struct {
	p    *FibHeapElement
	root FibHeapElement
	len  int
}

func (l *fibHeapElementList) Init(p *FibHeapElement) *fibHeapElementList {
	l.root.l = &l.root
	l.root.r = &l.root
	l.p = p
	l.len = 0
	return l
}

func (l *fibHeapElementList) Len() int {
	return l.len
}

func (l *fibHeapElementList) insert(e, at *FibHeapElement) *FibHeapElement {
	n := at.r
	at.r = e
	e.l = at
	e.r = n
	n.l = e
	e.p = l.p
	e.list = l
	l.len++
	return e
}

func (l *fibHeapElementList) Remove(e *FibHeapElement) {
	if e.list == l {
		e.l.r = e.r
		e.r.l = e.l
		e.r = nil // avoid memory leaks
		e.l = nil // avoid memory leaks
		e.list = nil
		e.p = nil
		l.len--
	}
}

func (l *fibHeapElementList) PushLeft(e *FibHeapElement) {
	l.insert(e, &l.root)
}

func (l *fibHeapElementList) PushRight(e *FibHeapElement) {
	l.insert(e, l.root.l)
}

func (l *fibHeapElementList) Leftist() *FibHeapElement {
	if l.len == 0 {
		return nil
	}
	return l.root.r
}

func (l *fibHeapElementList) Rightist() *FibHeapElement {
	if l.len == 0 {
		return nil
	}
	return l.root.l
}

func (l *fibHeapElementList) MergeRightList(other *fibHeapElementList) *fibHeapElementList {
	for i, e := other.Len(), other.Leftist(); i > 0; i = i - 1 {
		nextE := e.Right()
		other.Remove(e)
		l.insert(e, l.root.l)
		e = nextE
	}
	return l
}

func (l *fibHeapElementList) MergeLeftList(other *fibHeapElementList) *fibHeapElementList {
	for i, e := other.Len(), other.Rightist(); i > 0; i = i - 1 {
		nextE := e.Left()
		other.Remove(e)
		l.insert(e, l.root.r)
		e = nextE
	}
	return l
}

func newFabHeapElementList(p *FibHeapElement) *fibHeapElementList {
	return new(fibHeapElementList).Init(p)
}

//heap
type fibHeapIf interface {
	LessKey(interface{}, interface{}) bool
}

//FibHeap:Cross package
type FibHeap struct {
	root *fibHeapElementList
	min  *FibHeapElement
	n    int
	fibHeapIf
}

func (h *FibHeap) Init(self fibHeapIf) *FibHeap {
	h.root = newFabHeapElementList(nil)
	h.min = nil
	h.n = 0
	h.fibHeapIf = self
	return h
}

//default Less function, max heap
func (h *FibHeap) Less(n1, n2 *FibHeapElement) bool {
	if n1 == nil && n2 == nil {
		panic("both nodes are nil!")
	} else if n1 == nil {
		return false
	} else if n2 == nil {
		return true
	}
	return h.fibHeapIf.LessKey(n1.key, n2.key)
}

func (h *FibHeap) LessKey(key1, key2 interface{}) bool {
	return key1.(int) > key2.(int)
}

//floor(lg n)
func (h *FibHeap) Degree() int {
	if h.n == 0 {
		return 0
	}
	i := 1
	for diff := h.n; diff != 1; i++ {
		diff = h.n >> uint32(i)
	}
	return i - 1
}

func (h *FibHeap) Insert(key, value interface{}) *FibHeapElement {
	n := newFabHeapElement(key, value)
	h.root.PushRight(n)
	if h.Less(n, h.min) {
		h.min = n
	}
	h.n++
	return n
}

func (h *FibHeap) Union(h1 *FibHeap) *FibHeap {
	h.root = h.root.MergeRightList(h1.root)
	if h.Less(h1.min, h.min) {
		h.min = h1.min
	}
	h.n += h1.n
	return h
}

func (h *FibHeap) consolidate() {
	degreeArray := make([]*FibHeapElement, h.Degree()+1, h.Degree()+1)
	for i, e := h.root.Len(), h.root.Leftist(); i > 0; i = i - 1 {
		nextE := e.Right()
		for e1 := degreeArray[e.Degree()]; e1 != nil && e.Degree() < h.Degree(); e1 = degreeArray[e.Degree()] {
			degreeArray[e.Degree()] = nil
			if h.Less(e1, e) {
				e, e1 = e1, e
			}
			e1.mark = false
			h.root.Remove(e1)
			e.c.PushRight(e1)
		}
		degreeArray[e.Degree()] = e
		e = nextE
	}
	h.min = nil
	h.root = newFabHeapElementList(nil)
	for i := range degreeArray {
		if degreeArray[i] != nil {
			h.root.PushRight(degreeArray[i])
			if h.Less(degreeArray[i], h.min) {
				h.min = degreeArray[i]
			}
		}
	}
}

func (h *FibHeap) ExtractMin() *FibHeapElement {
	n := h.min
	if n != nil {
		for i, e := n.Degree(), n.c.Leftist(); i > 0; i = i - 1 {
			nextE := e.Right()
			n.c.Remove(e)
			h.root.PushLeft(e)
			e = nextE
		}
		h.root.Remove(n)
		if n == n.Right() {
			h.min = nil
		} else {
			h.min = n.Right()
			h.consolidate()
		}
		h.n--
	}
	return n
}

func (h *FibHeap) cascadingCut(n *FibHeapElement) {
	if p := n.p; p != nil {
		if n.mark {
			p.c.Remove(n)
			h.root.PushLeft(n)
			n.mark = false
			h.cascadingCut(p)
		} else {
			n.mark = true
		}
	}
}

func (h *FibHeap) ModifyNode(n *FibHeapElement, key, value interface{}) {
	if h.fibHeapIf.LessKey(n.key, key) {
		panic("key violated")
	}
	n.key = key
	n.value = value
	if p := n.p; n.p != nil && h.Less(n, n.p) {
		p.c.Remove(n)
		h.root.PushLeft(n)
		n.mark = false
		h.cascadingCut(p)
	}
	if h.Less(n, h.min) {
		h.min = n
	}
}

//NewFibHeap:Cross package
func NewFibHeap() *FibHeap {
	h := new(FibHeap)
	return h.Init(h)
}
