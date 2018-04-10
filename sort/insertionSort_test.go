package sort

import (
	"testing"
)

func Test_insertionSort_basic(t *testing.T) {
	basicTestSort(t, insertionSort)
}

func Benchmark_insertionSort(b *testing.B) {
	benchmarkSort(b, insertionSort)
}

func Test_bubbleSort_basic(t *testing.T) {
	testSort(t, bubbleSort)
}

func Benchmark_bubbleSort(b *testing.B) {
	benchmarkSort(b, bubbleSort)
}
