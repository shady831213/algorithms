package sort

import (
	"testing"
	"algorithms/sort/heap"
)

func Test_heapSort(t *testing.T)  {
	testSort(t,heapSort)
}

func Benchmark_heapSort(b *testing.B)  {
	benchmarkSort(b,heapSort)
}

func Test_heapPopAndAppend(t *testing.T)  {
	arr:= heapIntArray([]int{3,2,10,1,7})
	h := heap.Heap{&arr}
	h.BuildHeap()
	max := h.Pop().(int)
	if max != 10 {
		t.Log("max value should be 10"+" but get "+string(max))
		t.Fail()
	}
	h.Append(8)
	h.Append(4)
	max = h.Pop().(int)
	if max != 8 {
		t.Log("max value should be 8"+" but get ",max)
		t.Fail()
	}
	h.Append( 1)
	max = h.Pop().(int)
	if max != 7 {
		t.Log("max value should be 7"+" but get ",max)
		t.Fail()
	}
	max = h.Pop().(int)
	if max != 4 {
		t.Log("max value should be 4"+" but get ",max)
		t.Fail()
	}
}
