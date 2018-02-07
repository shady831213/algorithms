package sort

import (
	"algorithms/heap/arrayHeap"
	"testing"
)

func countingSortWrapper(arr []int) {
	_arr := make([]int, len(arr), cap(arr))
	copy(_arr, arr)
	h := arrayHeap.New(arr)
	max := h.Pop().(int)
	sortedArry := countingSort(arr, max)
	copy(arr, sortedArry)
}

func Test_countingSort(t *testing.T){
	testSort(t, countingSortWrapper)
}