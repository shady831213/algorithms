package sort

import "testing"

func Test_mergeSort(t *testing.T)  {
	_TestSort(t,mergeSort)
}

func Benchmark_mergeSort(b *testing.B) {
	_BenchmarkSort(b, mergeSort)
}

func Test_mergeSortParallel(t *testing.T)  {
	_TestSort(t,mergeSortParallel)
}

func Benchmark_mergeSortParallel(b *testing.B) {
	_BenchmarkSort(b, mergeSortParallel)
}