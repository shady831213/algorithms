package linkedHeap

import (
	"testing"
)

func Test_linkedHeap(t *testing.T) {
	h := New()
	TestHeap(t, h)
}

func Test_linkedHeapUnion(t *testing.T) {
	h, h2 := New(),New()
	TestHeapUnion(t, h,h2)
}

func Benchmark_linkedHeap(b *testing.B) {
	h := New()
	BenchmarkHeap(b,h)
}
