package heap

import (
	"testing"
)

func Test_linkedHeap(t *testing.T) {
	h := new(LinkedHeap)
	h.Init()
	testHeap(t, h)
}

func Test_linkedHeapUnion(t *testing.T) {
	h, h2 := new(LinkedHeap),new(LinkedHeap)
	h.Init()
	h2.Init()
	testHeapUnion(t, h,h2)
}

func Benchmark_linkedHeap(b *testing.B) {
	h := new(LinkedHeap)
	h.Init()
	benchmarkHeap(b,h)
}
