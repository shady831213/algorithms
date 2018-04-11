package heap

import (
	"testing"
)

func Test_leftistHeap(t *testing.T) {
	h := newLtHeapArray()
	testHeap(t, h)
}

func Test_leftistHeapUnion(t *testing.T) {
	h, h2 := newLtHeapArray(), newLtHeapArray()
	testHeapUnion(t, h, h2)
}

func Benchmark_leftistHeap(b *testing.B) {
	h := newLtHeapArray()
	benchmarkHeap(b, h)
}
