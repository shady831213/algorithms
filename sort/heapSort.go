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

import (
	"github.com/shady831213/algorithms/heap"
)

func heapSort(arr []int) {
	h := heap.NewHeapIntArray(arr)
	for i := h.Len() - 1; i > 0; i-- {
		h.Pop()
	}
}

/*not generic heap*/
type intArrayForHeapSort []int

func (h *intArrayForHeapSort) parent(i int) int {
	return i >> 1
}
func (h *intArrayForHeapSort) left(i int) int {
	return (i << 1) + 1
}
func (h *intArrayForHeapSort) right(i int) int {
	return (i << 1) + 2
}

func (h *intArrayForHeapSort) maxHeaplify(i int) {
	largest, largestIdx := (*h)[i], i
	if (*h).left(i) < len((*h)) && (*h)[(*h).left(i)] > largest {
		largest, largestIdx = (*h)[(*h).left(i)], (*h).left(i)
	}
	if h.right(i) < len((*h)) && (*h)[h.right(i)] > largest {
		_, largestIdx = (*h)[h.right(i)], h.right(i)
	}
	if i != largestIdx {
		(*h)[largestIdx], (*h)[i] = (*h)[i], (*h)[largestIdx]
		h.maxHeaplify(largestIdx)
	}
}

func (h *intArrayForHeapSort) buildHeap() {
	for i := (len((*h)) >> 1) - 1; i >= 0; i-- {
		h.maxHeaplify(i)
	}
}

func heapSort2(arr []int) {
	h := intArrayForHeapSort(arr)
	h.buildHeap()
	for i := len(h) - 1; i > 0; i-- {
		h[0], h[i] = h[i], h[0]
		h = h[:i]
		h.maxHeaplify(0)
	}
}
