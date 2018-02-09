package binarySearchTree

import (
	"testing"
	"algorithms/tree"
	"fmt"
	"flag"
	"sort"
)

var debug  = flag.Bool("debug", false, "debug flag")

func TestBst_InsertAndInOrderWalk(t *testing.T) {
	arr := tree.RandomSlice(0,20, 10)
	nodeCnt := 0
	bst := NewBstRecrusive()
	for _,v := range arr {
		bst.Insert(uint32(v))
	}
	bst.InOrderWalk(bst.Root, checkBst(t, &nodeCnt, *debug))
	if nodeCnt != len(arr) {
		t.Log(fmt.Sprintf("node cnt expect to ", len(arr), "but get:",nodeCnt))
		t.Fail()
	}
}

func TestBst_Delete(t *testing.T) {
	arr := tree.RandomSlice(0,20, 10)
	nodeCnt := 0
	bst := NewBstRecrusive()
	for _,v := range arr {
		bst.Insert(uint32(v))
	}
	delNum := tree.GetRand().Intn(10)+1
	if *debug {
		fmt.Println(delNum)
	}
	for i:=0;i < delNum;i++{
		bst.Delete(uint32(arr[i]))
	}
	bst.InOrderWalk(bst.Root, checkBst(t, &nodeCnt, *debug))
	if nodeCnt != len(arr) - delNum {
		t.Log(fmt.Sprintf("node cnt expect to ", len(arr) - delNum, "but get:",nodeCnt))
		t.Fail()
	}

}

func TestBst_Min(t *testing.T) {
	arr := tree.RandomSlice(0,20, 10)
	bst := NewBstRecrusive()
	for _,v := range arr {
		bst.Insert(uint32(v))
	}
	min := int(bst.Min(bst.Root).(*BstElement).Key)
	sort.Ints(arr)
	if min != arr[0] {
		t.Log(fmt.Sprintf("min expect to ", arr[0], "but get:",min))
		t.Fail()
	}
}

func TestBst_Max(t *testing.T) {
	arr := tree.RandomSlice(0,20, 10)
	bst := NewBstRecrusive()
	for _,v := range arr {
		bst.Insert(uint32(v))
	}
	max := int(bst.Max(bst.Root).(*BstElement).Key)
	sort.Ints(arr)
	if max != arr[len(arr)-1] {
		t.Log(fmt.Sprintf("max expect to ", arr[len(arr)-1], "but get:",max))
		t.Fail()
	}
}

func TestBst_Search(t *testing.T) {
	arr := tree.RandomSlice(0,20, 10)
	bst := NewBstRecrusive()
	for _,v := range arr {
		bst.Insert(uint32(v))
	}
	key := tree.GetRand().Intn(len(arr))
	result := int(bst.Search(uint32(arr[key])).(*BstElement).Key)
	if result != arr[key] {
		t.Log(fmt.Sprintf("search result expect to ", arr[key], "but get:",result))
		t.Fail()
	}
}

func TestBst_Predecesor(t *testing.T) {
	arr := tree.RandomSlice(0,20, 10)
	bst := NewBstRecrusive()
	for _,v := range arr {
		bst.Insert(uint32(v))
	}
	sort.Ints(arr)
	key := tree.GetRand().Intn(len(arr) - 2) + 1
	result := int(bst.Predecesor(bst.Search(uint32(arr[key]))).(*BstElement).Key)
	if result != arr[key-1] {
		t.Log(fmt.Sprintf("Predecesor of",arr[key], " expect to ", arr[key-1] , "but get:",result))
		t.Fail()
	}
}

func TestBst_Successor(t *testing.T) {
	arr := tree.RandomSlice(0,20, 10)
	bst := NewBstRecrusive()
	for _,v := range arr {
		bst.Insert(uint32(v))
	}
	sort.Ints(arr)
	key := tree.GetRand().Intn(len(arr) - 2) + 1
	result := int(bst.Successor(bst.Search(uint32(arr[key]))).(*BstElement).Key)
	if result != arr[key+1] {
		t.Log(fmt.Sprintf("Successor of",arr[key], " expect to ", arr[key+1] , "but get:",result))
		t.Fail()
	}
}