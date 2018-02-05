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
	"algorithms/heap"
)

type heapIntArray []int

func (h*heapIntArray) Parent(i interface{})(interface{}) {
	return i.(int)>>1
}
func (h*heapIntArray) Left(i interface{})(interface{}) {
	return (i.(int)<<1)+1
}
func (h*heapIntArray) Right(i interface{})(interface{}) {
	return (i.(int)<<1)+2
}

func (h*heapIntArray) Prev(i interface{})(interface{}) {
	return i.(int)-1
}
func (h*heapIntArray) Next(i interface{})(interface{}) {
	return i.(int)+1
}

func (h*heapIntArray) Last()(interface{}) {
	return len(*h)-1
}
func (h*heapIntArray) Head()(interface{}) {
	return 0
}
func (h*heapIntArray) Valid(i interface{})(bool){
	return i.(int) >= 0 && i.(int) < len(*h)
}

func (h *heapIntArray) Swap(i interface{}, j interface{}) () {
	(*h)[i.(int)], (*h)[j.(int)] = (*h)[j.(int)], (*h)[i.(int)]
}

func (h *heapIntArray) Key(i interface{}) (int) {
	return (*h)[i.(int)]
}

func (h *heapIntArray) Len() (int) {
	return len(*h)
}

func (h *heapIntArray) Pop() (i interface{}) {
	(*h), i = (*h)[:len(*h)-1], (*h)[len(*h)-1]
	return
}

func (h *heapIntArray) Append(i interface{}) {
	(*h) = append((*h), i.(int))
}

func heapSort(arr []int) {
	a := heapIntArray(arr)
	h := heap.Heap{&a}
	h.BuildHeap()
	for i := a.Len() - 1; i > 0; i-- {
		h.Pop()
	}
}

/*not generic heap*/
type intArrayForHeapSort []int

func (h *intArrayForHeapSort) parent(i int) (int) {
	return i >> 1
}
func (h *intArrayForHeapSort) left(i int) (int) {
	return (i << 1) + 1
}
func (h *intArrayForHeapSort) right(i int) (int) {
	return (i << 1) + 2
}

func (h *intArrayForHeapSort) maxHeaplify(i int) {
	largest, largest_idx := (*h)[i], i
	if (*h).left(i) < len((*h)) && (*h)[(*h).left(i)] > largest {
		largest, largest_idx = (*h)[(*h).left(i)], (*h).left(i)
	}
	if h.right(i) < len((*h)) && (*h)[h.right(i)] > largest {
		largest, largest_idx = (*h)[h.right(i)], h.right(i)
	}
	if i != largest_idx {
		(*h)[largest_idx], (*h)[i] = (*h)[i], (*h)[largest_idx]
		h.maxHeaplify(largest_idx)
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
