package sort

import (
	"testing"
)

func Test_heapSort(t *testing.T)  {
	_TestSort(t,heapSort)
}

func Benchmark_heapSort(b *testing.B)  {
	_BenchmarkSort(b,heapSort)
}

func Test_heapPopAndAppend(t *testing.T)  {
	arr := []int{3,2,5,1,2}
	heap := &heapIntArray{arr}
	heap.buildHeap()
	max := heap.pop().(int)
	if max != 5 {
		t.Log("max value should be 5"+" but get "+string(max))
		t.Fail()
	}
	heap.append(8)
	max = heap.pop().(int)
	if max != 8 {
		t.Log("max value should be 8"+" but get ",max)
		t.Fail()
	}
	heap.append(1)
	max = heap.pop().(int)
	if max != 3 {
		t.Log("max value should be 3"+" but get ",max)
		t.Fail()
	}
	max = heap.pop().(int)
	if max != 2 {
		t.Log("max value should be 2"+" but get ",max)
		t.Fail()
	}
}