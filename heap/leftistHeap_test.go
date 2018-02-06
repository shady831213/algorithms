package heap

import (
	"testing"
)

func Test_leftistHeap(t *testing.T) {
	h := new(LtHeapArray)
	testHeap(t, h)
}

func Test_leftistHeapUnion(t *testing.T) {
	h,h2 := new(LtHeapArray),new(LtHeapArray)
	testHeapUnion(t, h,h2)
}

func Benchmark_leftistHeap(b *testing.B) {
	h := new(LtHeapArray)
	benchmarkHeap(b,h)
}
