package heap

type arrayIf interface {
	Swap(interface{}, interface{})
	Len() int
	Key(interface{}) int
	Value(interface{}) interface{}
	Pop() interface{}
	Append(interface{})
	Left(interface{}) interface{}
	Right(interface{}) interface{}
	Union(interface{}) interface{}
}

type binHeapArrayIf interface {
	arrayIf
	Last() interface{}
	Head() interface{}
	Prev(interface{}) interface{}
	Next(interface{}) interface{}
	Valid(interface{}) bool
	Parent(interface{}) interface{}
}

type heapIf interface {
	arrayIf
}

type binHeapIf interface {
	arrayIf
	MaxHeaplify(interface{})
	BuildHeap()
}

type heap struct {
	binHeapArrayIf
}

func (h *heap) MaxHeaplify(i interface{}) {
	largest, largestIdx := h.Key(i), i
	if h.Valid(h.Left(i)) && h.Key(h.Left(i)) > largest {
		largest, largestIdx = h.Key(h.Left(i)), h.Left(i)
	}
	if h.Valid(h.Right(i)) && h.Key(h.Right(i)) > largest {
		_, largestIdx = h.Key(h.Right(i)), h.Right(i)
	}
	if i != largestIdx {
		h.Swap(largestIdx, i)
		h.MaxHeaplify(largestIdx)
	}
}

func (h *heap) BuildHeap() {
	for i := h.Last(); h.Valid(i); i = h.Prev(i) {
		h.MaxHeaplify(i)
	}
}

func (h *heap) Pop() (i interface{}) {
	h.Swap(h.Head(), h.Last())
	i = h.binHeapArrayIf.Pop()
	if h.Len() > 0 {
		h.MaxHeaplify(h.Head())
	}
	return
}

func (h *heap) Append(i interface{}) {
	h.binHeapArrayIf.Append(i)
	for idx := h.Last(); h.Valid(h.Parent(idx)) && h.Key(idx) > h.Key(h.Parent(idx)); {
		h.Swap(idx, h.Parent(idx))
		idx = h.Parent(idx)
	}
}
