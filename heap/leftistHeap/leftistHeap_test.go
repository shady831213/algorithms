package leftistHeap

import (
	"testing"
	"algorithms/heap"
)

func Test_leftistHeap(t *testing.T) {
	h := New()
	heap.TestHeap(t, h)
}

func Test_leftistHeapUnion(t *testing.T) {
	h,h2 := New(),New()
	heap.TestHeapUnion(t, h,h2)
}

func Benchmark_leftistHeap(b *testing.B) {
	h := New()
	heap.BenchmarkHeap(b,h)
}
