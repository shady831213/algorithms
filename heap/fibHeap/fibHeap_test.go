package fibHeap

import (
	"testing"
	"reflect"
	"fmt"
	"sort"
	"math/rand"
)

func TestFibHeapBasic(t *testing.T) {
	arr := []int{1,3,2,2,4,5}
	sortedArr := make([]int, 0, 0)
	h := NewFibHeap()
	for _, v := range arr {
		h.Insert(v, v)
	}
	for h.n > 0 {
		min := h.ExtractMin()
		sortedArr = append(sortedArr, min.key.(int))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	if !reflect.DeepEqual(sortedArr, arr) {
		t.Log(fmt.Sprintf("expect:%v", arr) + fmt.Sprintf("but get:%v", sortedArr))
		t.Fail()
	}
}

func TestFibHeap(t *testing.T) {
	arrSize := rand.Intn(100) + 50
	arr := make([]int, arrSize, arrSize)
	for i := range arr {
		arr[i] = rand.Intn(100)
	}
	sortedArr := make([]int, 0, 0)
	h := NewFibHeap()
	for _, v := range arr {
		h.Insert(v, v)
	}
	for h.n > 0 {
		min := h.ExtractMin()
		sortedArr = append(sortedArr, min.key.(int))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	if !reflect.DeepEqual(sortedArr, arr) {
		t.Log(fmt.Sprintf("expect:%v", arr) + fmt.Sprintf("but get:%v", sortedArr))
		t.Fail()
	}
}


func TestFibHeapUnion(t *testing.T) {
	arrSize := rand.Intn(100) + 50
	arrSize1 := rand.Intn(arrSize)
	arr := make([]int, arrSize, arrSize)
	for i := range arr {
		arr[i] = rand.Intn(100)
	}
	sortedArr := make([]int, 0, 0)
	h1, h2 := NewFibHeap(), NewFibHeap()
	for i, v := range arr {
		if i < arrSize1 {
			h1.Insert(v,v)
		} else {
			h2.Insert(v,v)
		}
	}
	h1 = h1.Union(h2)
	h2 = nil
	for h1.n > 0 {
		min := h1.ExtractMin()
		sortedArr = append(sortedArr, min.key.(int))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	if !reflect.DeepEqual(sortedArr, arr) {
		t.Log(fmt.Sprintf("expect:%v", arr) + fmt.Sprintf("but get:%v", sortedArr))
		t.Fail()
	}
}