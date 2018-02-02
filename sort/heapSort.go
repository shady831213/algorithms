/*
build heap :
max node num in hight h is n/2^(h+1), max h is lgn,
T(n) = n*sum(0/2+1/2^1+...+lgn/2^(lgn+1))= (n/2)*sum(2/2^1+...+lgn/2^lgn)
=(n/2)*(2-((lgn+2)/2^(lgn))) <= (n/2)*2=n
d heap:
n*(d-1)/d^(h+1), max h is logd((d-1)n), T(n) = (n*(d-1)/d) *sum(1/d^1+2/d^2+...logd((d-1)n)/d^logd(d-1)n)
= (n*(d-1)/d) <= O(n)

about golang slice:
a := []int{...}
b := a

b and a point to the same memory, but are different object.
so they both can modify shared data.But they have different lenth,index and so on
*/

package sort

type heapIntArray struct {
	heap []int
}

type heapIf interface {
	parent(int)(int)
	left(int)(int)
	right(int)(int)
	maxHeaplify(int)
	buildHeap()
	sort()
	pop()(interface{})
	append(interface{})
}

func (h *heapIntArray) parent(i int)(int) {
	return i>>1
}
func (h *heapIntArray) left(i int)(int) {
	return (i<<1)+1
}
func (h *heapIntArray) right(i int)(int) {
	return (i<<1)+2
}

func (h *heapIntArray) maxHeaplify(i int) {
	largest, largest_idx := h.heap[i], i
	if h.left(i) < len(h.heap) && h.heap[h.left(i)] > largest {
		largest,largest_idx = h.heap[h.left(i)],h.left(i)
	}
	if h.right(i) < len(h.heap) && h.heap[h.right(i)] > largest {
		largest,largest_idx = h.heap[h.right(i)],h.right(i)
	}
	if i != largest_idx {
		h.heap[largest_idx], h.heap[i] = h.heap[i], h.heap[largest_idx]
		h.maxHeaplify(largest_idx)
	}
}

func (h *heapIntArray) buildHeap() {

	for i := (len(h.heap)>>1) - 1; i >= 0; i-- {
		h.maxHeaplify(i)
	}
}

func (h *heapIntArray) sort() {
	temp_slice := h.heap
	for i := len(h.heap) - 1;i>0;i-- {
		h.heap[0], h.heap[i] = h.heap[i], h.heap[0]
		h.heap = h.heap[:i]
		h.maxHeaplify(0)
	}
	h.heap = temp_slice
}

func (h *heapIntArray) pop()(i interface{}) {
	i = h.heap[0]
	h.heap[0], h.heap[len(h.heap) - 1] = h.heap[len(h.heap) - 1], h.heap[0]
	h.heap = h.heap[:len(h.heap) - 1]
	h.maxHeaplify(0)
	return
}

func (h *heapIntArray) append(i interface{}) {
	h.heap = append(h.heap, i.(int))
	for idx := len(h.heap)-1; h.heap[idx] > h.heap[h.parent(idx)] && idx > 0; {
		h.heap[idx], h.heap[h.parent(idx)] = h.heap[h.parent(idx)], h.heap[idx]
		idx = h.parent(idx)
	}
}

func heapSort(arr []int) {
	var heap heapIf
	heap = &heapIntArray{heap:arr}
	heap.buildHeap()
	heap.sort()
}