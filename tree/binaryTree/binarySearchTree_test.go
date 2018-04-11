package binaryTree

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestBst_Insert(t *testing.T) {
	arr := randomSlice(0, 20, 10)
	nodeCnt := 0
	bst := newBstRecrusive()
	for _, v := range arr {
		bst.Insert(uint32(v))
	}
	bst.InOrderWalk(bst.root, checkBst(t, &nodeCnt, *debug))
	if nodeCnt != len(arr) {
		t.Log(fmt.Sprintf("node cnt expect to %0d but get: %0d", len(arr), nodeCnt))
		t.Fail()
	}
}

func TestBst_Delete(t *testing.T) {
	arr := randomSlice(0, 20, 10)
	nodeCnt := 0
	bst := newBstRecrusive()
	for _, v := range arr {
		bst.Insert(uint32(v))
	}
	delNum := getRand().Intn(10) + 1
	if *debug {
		fmt.Println(delNum)
	}
	for i := 0; i < delNum; i++ {
		bst.Delete(uint32(arr[i]))
	}
	bst.InOrderWalk(bst.root, checkBst(t, &nodeCnt, *debug))
	if nodeCnt != len(arr)-delNum {
		t.Log(fmt.Sprintf("node cnt expect to %0d but get: %0d", len(arr)-delNum, nodeCnt))
		t.Fail()
	}

}

func TestBst_Min(t *testing.T) {
	arr := randomSlice(0, 20, 10)
	bst := newBstRecrusive()
	for _, v := range arr {
		bst.Insert(uint32(v))
	}
	min := int(bst.Min(bst.root).(*bstElement).Key)
	sort.Ints(arr)
	if min != arr[0] {
		t.Log(fmt.Sprintf("min expect to %0d but get:%0d", arr[0], min))
		t.Fail()
	}
}

func TestBst_Max(t *testing.T) {
	arr := randomSlice(0, 20, 10)
	bst := newBstRecrusive()
	for _, v := range arr {
		bst.Insert(uint32(v))
	}
	max := int(bst.Max(bst.root).(*bstElement).Key)
	sort.Ints(arr)
	if max != arr[len(arr)-1] {
		t.Log(fmt.Sprintf("max expect to %0d but get: %0d", arr[len(arr)-1], max))
		t.Fail()
	}
}

func TestBst_Search(t *testing.T) {
	arr := randomSlice(0, 20, 10)
	bst := newBstRecrusive()
	for _, v := range arr {
		bst.Insert(uint32(v))
	}
	key := getRand().Intn(len(arr))
	result := int(bst.Search(uint32(arr[key])).(*bstElement).Key)
	if result != arr[key] {
		t.Log(fmt.Sprintf("search result expect to %0d but get: %0d", arr[key], result))
		t.Fail()
	}
}

func TestBast_Predecessor(t *testing.T) {
	arr := randomSlice(0, 20, 10)
	bst := newBstRecrusive()
	for _, v := range arr {
		bst.Insert(uint32(v))
	}
	sort.Ints(arr)
	key := getRand().Intn(len(arr)-2) + 1
	result := int(bst.Predecessor(bst.Search(uint32(arr[key])), bst.Root()).(*bstElement).Key)
	if result != arr[key-1] {
		t.Log(fmt.Sprintf("Predecessor of %0d  expect to %0d but get:%0d", arr[key], arr[key-1], result))
		t.Fail()
	}
}

func TestBst_Successor(t *testing.T) {
	arr := randomSlice(0, 20, 10)
	bst := newBstRecrusive()
	for _, v := range arr {
		bst.Insert(uint32(v))
	}
	sort.Ints(arr)
	key := getRand().Intn(len(arr)-2) + 1
	result := int(bst.Successor(bst.Search(uint32(arr[key])), bst.Root()).(*bstElement).Key)
	if result != arr[key+1] {
		t.Log(fmt.Sprintf("Successor of %0d expect to %0d but get:%0d", arr[key], arr[key+1], result))
		t.Fail()
	}
}

func TestBstRecrusive_InOrderWalk(t *testing.T) {
	arr := randomSlice(0, 20, 10)
	resultArr := make([]int, 0, 0)
	bst := newBstRecrusive()
	for _, v := range arr {
		bst.Insert(uint32(v))
	}
	sort.Ints(arr)
	bst.InOrderWalk(bst.root, func(tree binaryTreeIf, node interface{}) bool {
		n := node.(*bstElement)
		resultArr = append(resultArr, int(n.Key))
		return false
	})
	if !reflect.DeepEqual(resultArr, arr) {
		t.Log(fmt.Sprintf("expect:%v", arr) + fmt.Sprintf("but get:%v", resultArr))
		t.Fail()
	}
}

func TestBstRecrusive_PreOrderWalk(t *testing.T) {
	arr := randomSlice(0, 20, 10)
	bst := newBstRecrusive()
	for _, v := range arr {
		bst.Insert(uint32(v))
	}
	checkBstPreOrder(t, bst)
}

func TestBstRecrusive_PostOrderWalk(t *testing.T) {
	arr := randomSlice(0, 20, 10)
	bst := newBstRecrusive()
	for _, v := range arr {
		bst.Insert(uint32(v))
	}
	checkBstPostOrder(t, bst)
}

func TestBstIterative_InOrderWalk(t *testing.T) {
	arr := randomSlice(0, 20, 10)
	resultArr := make([]int, 0, 0)
	bst := newBstIterative()
	for _, v := range arr {
		bst.Insert(uint32(v))
	}
	sort.Ints(arr)
	bst.InOrderWalk(bst.root, func(tree binaryTreeIf, node interface{}) bool {
		n := node.(*bstElement)
		resultArr = append(resultArr, int(n.Key))
		return false
	})
	if !reflect.DeepEqual(resultArr, arr) {
		t.Log(fmt.Sprintf("expect:%v", arr) + fmt.Sprintf("but get:%v", resultArr))
		t.Fail()
	}
}

func TestBstIterative_PreOrderWalk(t *testing.T) {
	arr := randomSlice(0, 20, 10)
	resultArr := make([]int, 0, 0)
	expArr := make([]int, 0, 0)
	bst := newBstIterative()
	expBst := newBstRecrusive()
	for _, v := range arr {
		bst.Insert(uint32(v))
		expBst.Insert(uint32(v))
	}
	expBst.PreOrderWalk(expBst.Root(), func(tree binaryTreeIf, node interface{}) bool {
		n := node.(*bstElement)
		expArr = append(expArr, int(n.Key))
		return false
	})
	bst.PreOrderWalk(bst.Root(), func(tree binaryTreeIf, node interface{}) bool {
		n := node.(*bstElement)
		if *debug {
			fmt.Println(n)
		}
		resultArr = append(resultArr, int(n.Key))
		return false
	})
	if !reflect.DeepEqual(resultArr, expArr) {
		t.Log(fmt.Sprintf("expect:%v", expArr) + fmt.Sprintf("but get:%v", resultArr))
		t.Fail()
	}
}

func TestBstIterative_PostOrderWalk(t *testing.T) {
	arr := randomSlice(0, 20, 10)
	resultArr := make([]int, 0, 0)
	expArr := make([]int, 0, 0)
	bst := newBstIterative()
	expBst := newBstRecrusive()
	for _, v := range arr {
		bst.Insert(uint32(v))
		expBst.Insert(uint32(v))
	}
	expBst.PostOrderWalk(expBst.Root(), func(tree binaryTreeIf, node interface{}) bool {
		n := node.(*bstElement)
		expArr = append(expArr, int(n.Key))
		return false
	})
	bst.PostOrderWalk(bst.Root(), func(tree binaryTreeIf, node interface{}) bool {
		n := node.(*bstElement)
		if *debug {
			fmt.Println(n)
		}
		resultArr = append(resultArr, int(n.Key))
		return false
	})
	if !reflect.DeepEqual(resultArr, expArr) {
		t.Log(fmt.Sprintf("expect:%v", expArr) + fmt.Sprintf("but get:%v", resultArr))
		t.Fail()
	}
}
