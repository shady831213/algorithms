package linkedHeap

import (
	"testing"
	"algorithms/heap"
)

func Test_linkedHeap(t *testing.T) {
	h := New()
	heap.TestHeap(t, h)
}

func Test_linkedHeapUnion(t *testing.T) {
	h, h2 := New(),New()
	heap.TestHeapUnion(t, h,h2)
}

func Benchmark_linkedHeap(b *testing.B) {
	h := New()
	heap.BenchmarkHeap(b,h)
}
