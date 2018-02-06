package heap

type ArrayIf interface {
	Swap(interface{}, interface{})
	Len() (int)
	Last() (interface{})
	Head() (interface{})
	Prev(interface{}) (interface{})
	Next(interface{}) (interface{})
	Valid(interface{}) (bool)
	Key(interface{}) (int)
	Value(interface{}) (interface{})
	Pop() (interface{})
	Append(interface{})
	Parent(interface{}) (interface{})
	Left(interface{}) (interface{})
	Right(interface{}) (interface{})
	Merge(interface{},interface{})(interface{})
}

type HeapIf interface {
	Len() (int)
	MaxHeaplify(interface{})
	BuildHeap()
	Pop() (interface{})
	Append(interface{})
	Union(interface{})
}

type Heap struct {
	Arr ArrayIf
}

func (h *Heap) Len() (int) {
	return h.Arr.Len()
}

func (h *Heap) MaxHeaplify(i interface{}) {
	largest, largest_idx := h.Arr.Key(i), i
	if h.Arr.Valid(h.Arr.Left(i)) && h.Arr.Key(h.Arr.Left(i)) > largest {
		largest, largest_idx = h.Arr.Key(h.Arr.Left(i)), h.Arr.Left(i)
	}
	if h.Arr.Valid(h.Arr.Right(i)) && h.Arr.Key(h.Arr.Right(i)) > largest {
		largest, largest_idx = h.Arr.Key(h.Arr.Right(i)), h.Arr.Right(i)
	}
	if i != largest_idx {
		h.Arr.Swap(largest_idx, i)
		h.MaxHeaplify(largest_idx)
	}
}

func (h *Heap) BuildHeap() {
	for i := h.Arr.Last();  h.Arr.Valid(i); i = h.Arr.Prev(i) {
		h.MaxHeaplify(i)
	}
}

func (h *Heap) Pop() (i interface{}) {
	h.Arr.Swap(h.Arr.Head(), h.Arr.Last())
	i = h.Arr.Pop()
	if h.Arr.Len() > 0 {
		h.MaxHeaplify(h.Arr.Head())
	}
	return
}

func (h *Heap) Append(i interface{}) {
	h.Arr.Append(i)
	for idx := h.Arr.Last(); h.Arr.Valid(h.Arr.Parent(idx)) && h.Arr.Key(idx) > h.Arr.Key(h.Arr.Parent(idx)); {
		h.Arr.Swap(idx, h.Arr.Parent(idx))
		idx = h.Arr.Parent(idx)
	}
}

//merge:O(n)
//rebuild:O(n)
//T(n)=O(n)
func (h *Heap) Union(h1 interface{}) {
	h.Arr = h.Arr.Merge(h.Arr, h1.(*Heap).Arr).(ArrayIf)
	h.BuildHeap()
}

type LtHeap struct {
	Arr ArrayIf
}

func (h *LtHeap) Len() (int) {
	return h.Arr.Len()
}

func (h *LtHeap) MaxHeaplify(i interface{}) {
}

func (h *LtHeap) BuildHeap() {
}

func (h *LtHeap) Pop() (i interface{}) {
	return h.Arr.Pop()
}

func (h *LtHeap) Append(i interface{}) {
	h.Arr.Append(i)
}

//merge:O(n)
//rebuild:O(n)
//T(n)=O(n)
func (h *LtHeap) Union(h1 interface{}) {
	h.Arr = h.Arr.Merge(h.Arr, h1.(*LtHeap).Arr).(ArrayIf)
}
