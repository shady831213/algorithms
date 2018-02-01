package sort

import "testing"

func Test_heapSort(t *testing.T)  {
	_TestSort(t,heapSort)
}

func Benchmark_heapSort(b *testing.B)  {
	_BenchmarkSort(b,heapSort)
}
