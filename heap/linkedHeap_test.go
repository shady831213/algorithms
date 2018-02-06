package heap

import (
	"testing"
)

func Test_likedHeap(t *testing.T) {
	lh := new(LinkedHeap)
	lh.Init()
	h := Heap{lh}
	testHeap(t, &h)
}

func Test_likedHeapUnion(t *testing.T) {
	lh, lh2 := new(LinkedHeap),new(LinkedHeap)
	lh.Init()
	lh2.Init()
	h, h2 := Heap{lh}, Heap{lh2}
	testHeapUnion(t, &h,&h2)
}

func Benchmark_likedHeap(b *testing.B) {
	lh := new(LinkedHeap)
	lh.Init()
	h := Heap{lh}
	benchmarkHeap(b,&h)
}
