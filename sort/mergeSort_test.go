package sort

import "testing"

func Test_mergeSort(t *testing.T)  {
	testSort(t,mergeSort)
}

func Benchmark_mergeSort(b *testing.B) {
	benchmarkSort(b, mergeSort)
}

func Test_mergeSortParallel(t *testing.T)  {
	testSort(t,mergeSortParallel)
}

func Benchmark_mergeSortParallel(b *testing.B) {
	benchmarkSort(b, mergeSortParallel)
}