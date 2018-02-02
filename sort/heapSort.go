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

import ("algorithms/sort/heap")

type heapIntArray []int

func (h *heapIntArray)Swap(i int, j int)() {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *heapIntArray)Key(i int)(int) {
	return (*h)[i]
}

func (h *heapIntArray)Len()(int) {
	return len(*h)
}

func (h *heapIntArray) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *heapIntArray)Pop()(i interface{}) {
	(*h), i = (*h)[:len(*h)-1], (*h)[len(*h)-1]
	return
}

func (h *heapIntArray)Append(i interface{}) {
	(*h) = append((*h), i.(int))
}

func heapSort(arr []int) {
	a := heapIntArray(arr)
	h := heap.Heap{&a}
	h.BuildHeap()
	for i := a.Len() - 1;i>0;i-- {
		h.Pop()
	}
}