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
	arr := []int{3,2,10,1,7}
	heap := &heapIntArray{arr}
	heap.buildHeap()
	max := heap.pop().(int)
	if max != 10 {
		t.Log("max value should be 10"+" but get "+string(max))
		t.Fail()
	}
	heap.append(8)
	heap.append(4)
	max = heap.pop().(int)
	if max != 8 {
		t.Log("max value should be 8"+" but get ",max)
		t.Fail()
	}
	heap.append(1)
	max = heap.pop().(int)
	if max != 7 {
		t.Log("max value should be 7"+" but get ",max)
		t.Fail()
	}
	max = heap.pop().(int)
	if max != 4 {
		t.Log("max value should be 4"+" but get ",max)
		t.Fail()
	}
}