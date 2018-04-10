package heap

import (
	"testing"
)

func Test_linkedHeap(t *testing.T) {
	h := NewLinkedHeap()
	TestHeap(t, h)
}

func Test_linkedHeapUnion(t *testing.T) {
	h, h2 := NewLinkedHeap(), NewLinkedHeap()
	TestHeapUnion(t, h, h2)
}

func Benchmark_linkedHeap(b *testing.B) {
	h := NewLinkedHeap()
	BenchmarkHeap(b, h)
}
