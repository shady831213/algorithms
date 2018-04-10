package heap

type ArrayIf interface {
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

type BinHeapArrayIf interface {
	ArrayIf
	Last() interface{}
	Head() interface{}
	Prev(interface{}) interface{}
	Next(interface{}) interface{}
	Valid(interface{}) bool
	Parent(interface{}) interface{}
}

type HeapIf interface {
	ArrayIf
}

type BinHeapIf interface {
	ArrayIf
	MaxHeaplify(interface{})
	BuildHeap()
}

type Heap struct {
	BinHeapArrayIf
}

func (h *Heap) MaxHeaplify(i interface{}) {
	largest, largest_idx := h.Key(i), i
	if h.Valid(h.Left(i)) && h.Key(h.Left(i)) > largest {
		largest, largest_idx = h.Key(h.Left(i)), h.Left(i)
	}
	if h.Valid(h.Right(i)) && h.Key(h.Right(i)) > largest {
		_, largest_idx = h.Key(h.Right(i)), h.Right(i)
	}
	if i != largest_idx {
		h.Swap(largest_idx, i)
		h.MaxHeaplify(largest_idx)
	}
}

func (h *Heap) BuildHeap() {
	for i := h.Last(); h.Valid(i); i = h.Prev(i) {
		h.MaxHeaplify(i)
	}
}

func (h *Heap) Pop() (i interface{}) {
	h.Swap(h.Head(), h.Last())
	i = h.BinHeapArrayIf.Pop()
	if h.Len() > 0 {
		h.MaxHeaplify(h.Head())
	}
	return
}

func (h *Heap) Append(i interface{}) {
	h.BinHeapArrayIf.Append(i)
	for idx := h.Last(); h.Valid(h.Parent(idx)) && h.Key(idx) > h.Key(h.Parent(idx)); {
		h.Swap(idx, h.Parent(idx))
		idx = h.Parent(idx)
	}
}
