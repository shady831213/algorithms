package sort

import (
	"algorithms/heap"
	"testing"
)

func countingSortWrapper(arr []int) {
	_arr := make([]int, len(arr), cap(arr))
	copy(_arr, arr)
	a := heapIntArray(arr)
	h := heap.Heap{&a}
	h.BuildHeap()
	max := h.Pop().(int)
	sortedArry := countingSort(arr, max)
	copy(arr, sortedArry)
}

func Test_countingSort(t *testing.T){
	testSort(t, countingSortWrapper)
}