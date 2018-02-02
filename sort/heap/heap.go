package heap

import "sort"

type ArrayIf interface {
	sort.Interface
	Key(int)(int)
	Pop()(interface{})
	Append(interface{})
}

type Heap struct {
	Arr ArrayIf
}

func (h*Heap) parent(i int)(int) {
	return i>>1
}
func (h*Heap) left(i int)(int) {
	return (i<<1)+1
}
func (h*Heap) right(i int)(int) {
	return (i<<1)+2
}

func (h *Heap) MaxHeaplify(i int) {
	largest, largest_idx := h.Arr.Key(i), i
	if h.left(i) < h.Arr.Len() && h.Arr.Key(h.left(i)) > largest {
		largest,largest_idx = h.Arr.Key(h.left(i)),h.left(i)
	}
	if h.right(i) < h.Arr.Len() && h.Arr.Key(h.right(i)) > largest {
		largest,largest_idx = h.Arr.Key(h.right(i)),h.right(i)
	}
	if i != largest_idx {
		h.Arr.Swap(largest_idx, i)
		h.MaxHeaplify(largest_idx)
	}
}

func (h *Heap) BuildHeap() {
	for i := (h.Arr.Len()>>1) - 1; i >= 0; i-- {
		h.MaxHeaplify(i)
	}
}

func (h *Heap) Pop()(i interface{}) {
	h.Arr.Swap(0, h.Arr.Len()-1)
	i = h.Arr.Pop()
	h.MaxHeaplify(0)
	return
}

func (h *Heap) Append(i interface{}) {
	h.Arr.Append(i)
	for idx := h.Arr.Len()-1; h.Arr.Key(idx) > h.Arr.Key(h.parent(idx)) && idx > 0; {
		h.Arr.Swap(idx, h.parent(idx))
		idx = h.parent(idx)
	}
}