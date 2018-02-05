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
	Pop() (interface{})
	Append(interface{})
	Parent(interface{}) (interface{})
	Left(interface{}) (interface{})
	Right(interface{}) (interface{})
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
