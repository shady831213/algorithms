package heap

type ltHeapElement struct {
	left, right *ltHeapElement
	dist        int
	Value       interface{}
}

type ltHeapArray struct {
	root *ltHeapElement
	len  int
}

func (h *ltHeapArray) Left(i interface{}) interface{} {
	iE := i.(*ltHeapElement)
	return iE.left
}
func (h *ltHeapArray) Right(i interface{}) interface{} {
	iE := i.(*ltHeapElement)
	return iE.right
}

func (h *ltHeapArray) Head() interface{} {
	return h.root
}

func (h *ltHeapArray) Swap(i interface{}, j interface{}) {
	iE := i.(**ltHeapElement)
	jE := j.(**ltHeapElement)
	(*iE), (*jE) = (*jE), (*iE)
}

func (h *ltHeapArray) Key(i interface{}) int {
	iE := i.(*ltHeapElement)
	return iE.Value.(int)
}

func (h *ltHeapArray) Value(i interface{}) interface{} {
	iE := i.(*ltHeapElement)
	return iE.Value
}

func (h *ltHeapArray) Len() int {
	return h.len
}

func (h *ltHeapArray) Pop() (i interface{}) {
	i = h.root.Value
	h.root = h.merge(h.root.left, h.root.right).(*ltHeapElement)
	h.len--
	return
}

func (h *ltHeapArray) Append(i interface{}) {
	newE := ltHeapElement{Value: i}
	h.root = h.merge(h.root, &newE).(*ltHeapElement)
	h.len++
}

//merge:O(logn)
func (h *ltHeapArray) merge(i interface{}, j interface{}) interface{} {
	iE := i.(*ltHeapElement)
	jE := j.(*ltHeapElement)
	if iE == nil {
		return jE
	}
	if jE == nil {
		return iE
	}
	if h.Key(iE) < h.Key(jE) {
		h.Swap(&iE, &jE)
	}
	iE.right = h.merge(iE.right, jE).(*ltHeapElement)
	if iE.left == nil || iE.right.dist > iE.left.dist {
		h.Swap(&iE.left, &iE.right)
	}
	if iE.right == nil {
		iE.dist = 0
	} else {
		iE.dist = iE.right.dist + 1
	}
	return iE
}

func (h *ltHeapArray) Union(i interface{}) interface{} {
	_i := i.(*ltHeapArray)
	h.root = h.merge(h.root, _i.root).(*ltHeapElement)
	h.len += _i.len
	return h
}

func newLtHeapArray() *ltHeapArray {
	return new(ltHeapArray)
}
