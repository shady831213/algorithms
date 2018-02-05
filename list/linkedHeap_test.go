package list

import (
	"testing"
	"reflect"
	"fmt"
	"algorithms/heap"
	"sort"
	"math/rand"
)

func Test_likedHeap(t *testing.T) {
	arrSize := rand.Intn(100) + 50
	arr := make([]int, arrSize, arrSize)
	for i := range arr {
		arr[i] = rand.Intn(100)
	}
	sortedArr := make([]int, 0, 0)
	lh := new(LinkedHeap)
	lh.Init()
	h := heap.Heap{lh}
	for _, v := range arr {
		h.Append(v)
	}
	for h.Len() > 0 {
		sortedArr = append(sortedArr, h.Pop().(int))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	if !reflect.DeepEqual(sortedArr, arr) {
		t.Log(fmt.Sprintf("expect:%v", arr) + fmt.Sprintf("but get:%v", sortedArr))
		t.Fail()
	}
}

func Test_likedHeapUnion(t *testing.T) {
	arrSize := rand.Intn(100) + 50
	arrSize1 := rand.Intn(arrSize)
	arr := make([]int, arrSize, arrSize)
	for i := range arr {
		arr[i] = rand.Intn(100)
	}
	sortedArr := make([]int, 0, 0)
	lh, lh2 := new(LinkedHeap),new(LinkedHeap)
	lh.Init()
	lh2.Init()
	h, h2 := heap.Heap{lh}, heap.Heap{lh2}
	for i, v := range arr {
		if i < arrSize1 {
			h.Append(v)
		} else {
			h2.Append(v)
		}
	}
	h.Union(&h2)
	for h.Len() > 0 {
		sortedArr = append(sortedArr, h.Pop().(int))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	if !reflect.DeepEqual(sortedArr, arr) {
		t.Log(fmt.Sprintf("expect:%v", arr) + fmt.Sprintf("but get:%v", sortedArr))
		t.Fail()
	}
}

func Benchmark_likedHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrSize := 10000
		arr := make([]int, arrSize, arrSize)
		for i := range arr {
			arr[i] = rand.Intn(arrSize)
		}
		b.StartTimer()

		sortedArr := make([]int, 0, 0)
		lh := new(LinkedHeap)
		lh.Init()
		h := heap.Heap{lh}
		for _, v := range arr {
			h.Append(v)
		}
		for h.Len() > 0 {
			sortedArr = append(sortedArr, h.Pop().(int))
		}
	}

}
