package heap

type Element struct {
	parent, left, right, next, prev *Element
	Value interface{}
}

type LinkedHeapList struct {
	root Element
	len int
}

type LinkedHeap struct {
	LinkedHeapList
	Heap
}

func (h*LinkedHeapList) Init() *LinkedHeapList {
	h.root.prev = &h.root
	h.root.next = &h.root
	h.root.parent = &h.root
	h.root.left = &h.root
	h.root.right = &h.root
	h.len = 0
	return h
}

func (h*LinkedHeapList) Parent(i interface{})(interface{}) {
	iE := i.(*Element)
	return iE.parent
}
func (h*LinkedHeapList) Left(i interface{})(interface{}) {
	iE := i.(*Element)
	return iE.left
}
func (h*LinkedHeapList) Right(i interface{})(interface{}) {
	iE := i.(*Element)
	return iE.right
}

func (h*LinkedHeapList) Prev(i interface{})(interface{}) {
	iE := i.(*Element)
	return iE.prev
}
func (h*LinkedHeapList) Next(i interface{})(interface{}) {
	iE := i.(*Element)
	return iE.next
}

func (h*LinkedHeapList) Last()(interface{}) {
	return h.root.prev
}
func (h*LinkedHeapList) Head()(interface{}) {
	return h.root.next
}

func (h*LinkedHeapList) Valid(i interface{})(bool){
	iE := i.(*Element)
	return iE != &h.root && iE != nil
}

func (h *LinkedHeapList) Swap(i interface{}, j interface{}) () {
	iE := i.(*Element)
	jE := j.(*Element)
	iE.Value, jE.Value = jE.Value, iE.Value
}

func (h *LinkedHeapList) Key(i interface{}) (int) {
	iE := i.(*Element)
	return iE.Value.(int)
}

func (h *LinkedHeapList) Value(i interface{}) (interface{}) {
	iE := i.(*Element)
	return iE.Value
}

func (h *LinkedHeapList) Len() (int) {
	return h.len
}

func (h *LinkedHeapList) Pop() (i interface{}) {
	last := h.Last().(*Element)
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

func (h *LinkedHeapList) Append(i interface{}) {
	newE := new(Element)
	newE.Value = i
	newE.next = &h.root
	last := h.Last().(*Element)
	lastParent := last.parent
	last.next = newE
	newE.prev = last
	h.root.prev = newE
	if lastParent.right == nil{
		lastParent.right = newE
		newE.parent = lastParent
	} else {
		lastParent.next.left = newE
		newE.parent = lastParent.next
	}
	h.len++
}
//O(n)
func (h *LinkedHeapList) Union(i interface{})(interface{}) {
	var midNode *Element
	_i := i.(*LinkedHeapList)
	if h.Len() > _i.Len() {
		midNode = _i.Head().(*Element)
		h.Last().(*Element).next = midNode
		midNode.prev = h.Last().(*Element)
		h.root.prev = _i.Last().(*Element)
		_i.Last().(*Element).next = &h.root
	} else {
		midNode := h.Head().(*Element)
		_i.Last().(*Element).next = midNode
		midNode.prev = _i.Last().(*Element)
		h.root.next = _i.Head().(*Element)
		_i.Head().(*Element).prev = &h.root
	}
	for iNode := midNode; h.Valid(iNode); iNode = h.Next(iNode).(*Element) {
		iNode.parent = nil
		iNode.left = nil
		iNode.right = nil
		prev := iNode.prev
		prevParent := prev.parent
		if prevParent.right == nil{
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


func (h*LinkedHeap) Init() *LinkedHeap {
	h.LinkedHeapList = LinkedHeapList{}
	h.LinkedHeapList.Init()
	h.Heap.BinHeapArrayIf = &h.LinkedHeapList
	return h
}
//merge:O(n)
//rebuild:O(n)
//T(n)=O(n)
func (h *LinkedHeap) Union(i interface{})(interface{}) {
	h.LinkedHeapList = h.LinkedHeapList.Union(&(i.(*LinkedHeap).LinkedHeapList)).(LinkedHeapList)
	h.Heap.BuildHeap()
	return h
}

func (h *LinkedHeap) Pop() (i interface{}) {
	return h.Heap.Pop()
}

func (h *LinkedHeap) Append(i interface{}) {
	h.Heap.Append(i)
}