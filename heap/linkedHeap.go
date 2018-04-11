package heap

type linkedHeapListElement struct {
	parent, left, right, next, prev *linkedHeapListElement
	Value                           interface{}
}

type linkedHeapList struct {
	root linkedHeapListElement
	len  int
}

type linkedHeap struct {
	linkedHeapList
	heap
}

func (h *linkedHeapList) Init() {
	h.root.prev = &h.root
	h.root.next = &h.root
	h.root.parent = &h.root
	h.root.left = &h.root
	h.root.right = &h.root
	h.len = 0
}

func (h *linkedHeapList) Parent(i interface{}) interface{} {
	iE := i.(*linkedHeapListElement)
	return iE.parent
}
func (h *linkedHeapList) Left(i interface{}) interface{} {
	iE := i.(*linkedHeapListElement)
	return iE.left
}
func (h *linkedHeapList) Right(i interface{}) interface{} {
	iE := i.(*linkedHeapListElement)
	return iE.right
}

func (h *linkedHeapList) Prev(i interface{}) interface{} {
	iE := i.(*linkedHeapListElement)
	return iE.prev
}
func (h *linkedHeapList) Next(i interface{}) interface{} {
	iE := i.(*linkedHeapListElement)
	return iE.next
}

func (h *linkedHeapList) Last() interface{} {
	return h.root.prev
}
func (h *linkedHeapList) Head() interface{} {
	return h.root.next
}

func (h *linkedHeapList) Valid(i interface{}) bool {
	iE := i.(*linkedHeapListElement)
	return iE != &h.root && iE != nil
}

func (h *linkedHeapList) Swap(i interface{}, j interface{}) {
	iE := i.(*linkedHeapListElement)
	jE := j.(*linkedHeapListElement)
	iE.Value, jE.Value = jE.Value, iE.Value
}

func (h *linkedHeapList) Key(i interface{}) int {
	iE := i.(*linkedHeapListElement)
	return iE.Value.(int)
}

func (h *linkedHeapList) Value(i interface{}) interface{} {
	iE := i.(*linkedHeapListElement)
	return iE.Value
}

func (h *linkedHeapList) Len() int {
	return h.len
}

func (h *linkedHeapList) Pop() (i interface{}) {
	last := h.Last().(*linkedHeapListElement)
	lastParent := last.parent
	lastPrev := last.prev
	lastPrev.next = &h.root
	h.root.prev = lastPrev
	if lastParent.left == last {
		lastParent.left = nil
	} else {
		lastParent.right = nil
	}
	h.len--
	return last.Value
}

func (h *linkedHeapList) Append(i interface{}) {
	newE := new(linkedHeapListElement)
	newE.Value = i
	newE.next = &h.root
	last := h.Last().(*linkedHeapListElement)
	lastParent := last.parent
	last.next = newE
	newE.prev = last
	h.root.prev = newE
	if lastParent.right == nil {
		lastParent.right = newE
		newE.parent = lastParent
	} else {
		lastParent.next.left = newE
		newE.parent = lastParent.next
	}
	h.len++
}

//O(n)
func (h *linkedHeapList) Union(i interface{}) interface{} {
	var midNode *linkedHeapListElement
	_i := i.(*linkedHeapList)
	if h.Len() > _i.Len() {
		midNode = _i.Head().(*linkedHeapListElement)
		h.Last().(*linkedHeapListElement).next = midNode
		midNode.prev = h.Last().(*linkedHeapListElement)
		h.root.prev = _i.Last().(*linkedHeapListElement)
		_i.Last().(*linkedHeapListElement).next = &h.root
	} else {
		midNode := h.Head().(*linkedHeapListElement)
		_i.Last().(*linkedHeapListElement).next = midNode
		midNode.prev = _i.Last().(*linkedHeapListElement)
		h.root.next = _i.Head().(*linkedHeapListElement)
		_i.Head().(*linkedHeapListElement).prev = &h.root
	}
	for iNode := midNode; h.Valid(iNode); iNode = h.Next(iNode).(*linkedHeapListElement) {
		iNode.parent = nil
		iNode.left = nil
		iNode.right = nil
		prev := iNode.prev
		prevParent := prev.parent
		if prevParent.right == nil {
			prevParent.right = iNode
			iNode.parent = prevParent
		} else {
			prevParent.next.left = iNode
			iNode.parent = prevParent.next
		}
	}
	h.len += _i.Len()
	return *h
}

func newLinkedHeap() *linkedHeap {
	h := new(linkedHeap)
	h.linkedHeapList = linkedHeapList{}
	h.linkedHeapList.Init()
	h.heap.binHeapArrayIf = &h.linkedHeapList
	return h
}

//merge:O(n)
//rebuild:O(n)
//T(n)=O(n)
func (h *linkedHeap) Union(i interface{}) interface{} {
	h.linkedHeapList = h.linkedHeapList.Union(&(i.(*linkedHeap).linkedHeapList)).(linkedHeapList)
	h.heap.BuildHeap()
	return h
}

func (h *linkedHeap) Pop() (i interface{}) {
	return h.heap.Pop()
}

func (h *linkedHeap) Append(i interface{}) {
	h.heap.Append(i)
}
