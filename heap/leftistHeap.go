package heap

type LtHeapElement struct {
	left, right *LtHeapElement
	dist int
	Value interface{}
}

type LtHeapArray struct {
	root *LtHeapElement
	len int
}

func (h*LtHeapArray) Left(i interface{})(interface{}) {
	iE := i.(*LtHeapElement)
	return iE.left
}
func (h*LtHeapArray) Right(i interface{})(interface{}) {
	iE := i.(*LtHeapElement)
	return iE.right
}

func (h*LtHeapArray) Head()(interface{}) {
	return h.root
}

func (h *LtHeapArray) Swap(i interface{}, j interface{}) () {
	iE :=i.(**LtHeapElement)
	jE :=j.(**LtHeapElement)
	(*iE),(*jE) = (*jE), (*iE)
}

func (h *LtHeapArray) Key(i interface{}) (int) {
	iE :=i.(*LtHeapElement)
	return iE.Value.(int)
}

func (h *LtHeapArray) Value(i interface{}) (interface{}) {
	iE :=i.(*LtHeapElement)
	return iE.Value
}

func (h *LtHeapArray) Len() (int) {
	return h.len
}

func (h *LtHeapArray) Pop() (i interface{}) {
	i = h.root.Value
	h.root = h.merge(h.root.left,h.root.right).(*LtHeapElement)
	h.len--
	return
}

func (h *LtHeapArray) Append(i interface{}) {
	newE := LtHeapElement{Value:i}
	h.root = h.merge(h.root,&newE).(*LtHeapElement)
	h.len++
}
//merge:O(logn)
func (h *LtHeapArray) merge(i interface{},j interface{})(interface{}) {
	iE :=i.(*LtHeapElement)
	jE :=j.(*LtHeapElement)
	if iE == nil {
		return jE
	}
	if jE == nil {
		return iE
	}
	if h.Key(iE) < h.Key(jE) {
		h.Swap(&iE, &jE)
	}
	iE.right = h.merge(iE.right, jE).(*LtHeapElement)
	if iE.left == nil || iE.right.dist > iE.left.dist {
		h.Swap(&iE.left,&iE.right)
	}
	if iE.right == nil {
		iE.dist = 0
	} else {
		iE.dist = iE.right.dist + 1
	}
	return iE
}

func (h *LtHeapArray) Union(i interface{})(interface{}) {
	_i :=i.(*LtHeapArray)
	h.root = h.merge(h.root,_i.root).(*LtHeapElement)
	h.len+=_i.len
	return h
}

func NewLtHeapArray() *LtHeapArray {
	return new(LtHeapArray)
}