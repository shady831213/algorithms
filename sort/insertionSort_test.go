package sort

import (
	"testing"
)

func Test_insertionSort_basic(t *testing.T) {
	_basicTestSort(t,insertionSort)
}

func Benchmark_insertionSort(b *testing.B) {
	_BenchmarkSort(b, insertionSort)
}

func Test_bubbleSort_basic(t *testing.T) {
	_TestSort(t,bubbleSort)
}

func Benchmark_bubbleSort(b *testing.B) {
	_BenchmarkSort(b, bubbleSort)
}