package heap

import (
	"testing"
)

func Test_linkedHeap(t *testing.T) {
	h := newLinkedHeap()
	testHeap(t, h)
}

func Test_linkedHeapUnion(t *testing.T) {
	h, h2 := newLinkedHeap(), newLinkedHeap()
	testHeapUnion(t, h, h2)
}

func Benchmark_linkedHeap(b *testing.B) {
	h := newLinkedHeap()
	benchmarkHeap(b, h)
}
