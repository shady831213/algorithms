package sort

import "testing"

func Test_heapSort2(t *testing.T)  {
	_TestSort(t,heapSort2)
}

func Benchmark_heapSort2(b *testing.B)  {
	_BenchmarkSort(b,heapSort2)
}

func Test_heapSort(t *testing.T)  {
	_TestSort(t,heapSort)
}

func Benchmark_heapSort(b *testing.B)  {
	_BenchmarkSort(b,heapSort)
}

