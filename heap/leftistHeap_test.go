package heap

import (
	"testing"
)

func Test_leftistHeap(t *testing.T) {
	lh := new(LtHeapArray)
	h := LtHeap{lh}
	testHeap(t, &h)
}

func Test_leftistHeapUnion(t *testing.T) {
	lh,lh2 := new(LtHeapArray),new(LtHeapArray)
	h,h2 := LtHeap{lh},LtHeap{lh2}
	testHeapUnion(t, &h,&h2)
}

func Benchmark_leftistHeap(b *testing.B) {
	lh := new(LtHeapArray)
	h := LtHeap{lh}
	benchmarkHeap(b,&h)
}
