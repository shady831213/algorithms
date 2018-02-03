package sort

import "testing"

func Test_quickSort(t *testing.T)  {
	testSort(t,quickSort)
}

func Benchmark_quickSort(b *testing.B) {
	benchmarkSort(b, quickSort)
}

func Test_randomQuickSort(t *testing.T)  {
	testSort(t,randomQuickSort)
}

func Benchmark_randomQuickSort(b *testing.B) {
	benchmarkSort(b, randomQuickSort)
}