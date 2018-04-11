package binaryTree

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestGBT_Insert(t *testing.T) {
	arr := randomSlice(0, 20, 10)
	nodeCnt := 0
	gbt := newGBT()
	for _, v := range arr {
		gbt.Insert(uint32(v))
	}
	gbt.InOrderWalk(gbt.Root(), checkGBT(t, &nodeCnt, *debug))
	if nodeCnt != len(arr) {
		t.Log("node cnt expect to ", len(arr), "but get:", nodeCnt)
		t.Fail()
	}
}

func TestGBT_Delete(t *testing.T) {
	arr := randomSlice(0, 20, 10)
	nodeCnt := 0
	gbt := newGBT()
	for _, v := range arr {
		gbt.Insert(uint32(v))
	}
	delNum := getRand().Intn(10) + 1
	if *debug {
		fmt.Println(delNum)
	}
	for i := 0; i < delNum; i++ {
		gbt.Delete(uint32(arr[i]))
	}
	gbt.InOrderWalk(gbt.Root(), checkGBT(t, &nodeCnt, *debug))
	if nodeCnt != len(arr)-delNum {
		t.Log(fmt.Sprintf("node cnt expect to %0d but get:%0d", len(arr)-delNum, nodeCnt))
		t.Fail()
	}

}

func TestGBT_Min(t *testing.T) {
	arr := randomSlice(0, 20, 10)
	gbt := newGBT()
	for _, v := range arr {
		gbt.Insert(uint32(v))
	}
	min := int(gbt.Min(gbt.Root()).(*gbtElement).Key)
	sort.Ints(arr)
	if min != arr[0] {
		t.Log(fmt.Sprintf("min expect to %0d but get:%0d", arr[0], min))
		t.Fail()
	}
}

func TestGBT_Max(t *testing.T) {
	arr := randomSlice(0, 20, 10)
	gbt := newGBT()
	for _, v := range arr {
		gbt.Insert(uint32(v))
	}
	max := int(gbt.Max(gbt.Root()).(*gbtElement).Key)
	sort.Ints(arr)
	if max != arr[len(arr)-1] {
		t.Log(fmt.Sprintf("max expect to %0d but get:%0d", arr[len(arr)-1], max))
		t.Fail()
	}
}

func TestGBT_Search(t *testing.T) {
	arr := randomSlice(0, 20, 10)
	gbt := newGBT()
	for _, v := range arr {
		gbt.Insert(uint32(v))
	}
	key := getRand().Intn(len(arr))
	result := int(gbt.Search(uint32(arr[key])).(*gbtElement).Key)
	if result != arr[key] {
		t.Log(fmt.Sprintf("search result expect to %0d but get:%0d", arr[key], result))
		t.Fail()
	}
}

func TestGET_Predecessor(t *testing.T) {
	arr := randomSlice(0, 20, 10)
	gbt := newGBT()
	for _, v := range arr {
		gbt.Insert(uint32(v))
	}
	sort.Ints(arr)
	key := getRand().Intn(len(arr)-2) + 1
	result := int(gbt.Predecessor(gbt.Search(uint32(arr[key])), gbt.Root()).(*gbtElement).Key)
	if result != arr[key-1] {
		t.Log(fmt.Sprintf("Predecessor of %0d expect to %0d but get:%0d", arr[key], arr[key-1], result))
		t.Fail()
	}
}

func TestGBT_Successor(t *testing.T) {
	arr := randomSlice(0, 20, 10)
	gbt := newGBT()
	for _, v := range arr {
		gbt.Insert(uint32(v))
	}
	sort.Ints(arr)
	key := getRand().Intn(len(arr)-2) + 1
	result := int(gbt.Successor(gbt.Search(uint32(arr[key])), gbt.Root()).(*gbtElement).Key)
	if result != arr[key+1] {
		t.Log(fmt.Sprintf("Successor of %0d expect to %0d but get:%0d", arr[key], arr[key+1], result))
		t.Fail()
	}
}

func TestGBT_LeftRotate(t *testing.T) {
	arr := randomSlice(0, 20, 10)
	nodeCnt := 0
	gbt := newGBT()
	for _, v := range arr {
		gbt.Insert(uint32(v))
	}
	//Left rotate randomly
	leftRotateNodes := make([]*gbtElement, 0, 0)
	gbt.InOrderWalk(gbt.Root(), func(GBT binaryTreeIf, node interface{}) bool {
		rotate := getRand().Intn(2)
		if rotate == 1 {
			leftRotateNodes = append(leftRotateNodes, node.(*gbtElement))
			if *debug {
				return true
			}
		}
		return false
	})
	for _, v := range leftRotateNodes {
		gbt.LeftRotate(v)
	}
	gbt.InOrderWalk(gbt.Root(), checkGBT(t, &nodeCnt, *debug))
	if nodeCnt != len(arr) {
		t.Log("node cnt expect to ", len(arr), "but get:", nodeCnt)
		t.Fail()
	}
	if *debug {
		resultArr := make([]int, 0, 0)
		sort.Ints(arr)
		gbt.InOrderWalk(gbt.Root(), func(GBT binaryTreeIf, node interface{}) bool {
			n := node.(*gbtElement)
			resultArr = append(resultArr, int(n.Key))
			return false
		})
		if !reflect.DeepEqual(resultArr, arr) {
			t.Log(fmt.Sprintf("expect:%v", arr) + fmt.Sprintf("but get:%v", resultArr))
			t.Fail()
		}
	}
}

func TestGBT_RightRotate(t *testing.T) {
	arr := randomSlice(0, 20, 10)
	nodeCnt := 0
	gbt := newGBT()
	for _, v := range arr {
		gbt.Insert(uint32(v))
	}
	//Right rotate randomly
	rightRotateNodes := make([]*gbtElement, 0, 0)
	gbt.InOrderWalk(gbt.Root(), func(GBT binaryTreeIf, node interface{}) bool {
		rotate := getRand().Intn(2)
		if rotate == 1 {
			rightRotateNodes = append(rightRotateNodes, node.(*gbtElement))
		}
		return false
	})
	for _, v := range rightRotateNodes {
		gbt.RightRotate(v)
	}
	gbt.InOrderWalk(gbt.Root(), checkGBT(t, &nodeCnt, *debug))
	if nodeCnt != len(arr) {
		t.Log("node cnt expect to ", len(arr), "but get:", nodeCnt)
		t.Fail()
	}
}

func TestGBT_Rotate(t *testing.T) {
	arr := randomSlice(0, 20, 10)
	nodeCnt := 0
	gbt := newGBT()
	for _, v := range arr {
		gbt.Insert(uint32(v))
	}
	//Right rotate randomly
	rotateNodes := make([]*gbtElement, 0, 0)
	gbt.InOrderWalk(gbt.Root(), func(GBT binaryTreeIf, node interface{}) bool {
		rotate := getRand().Intn(2)
		if rotate == 1 {
			rotateNodes = append(rotateNodes, node.(*gbtElement))
		}
		return false
	})
	for _, v := range rotateNodes {
		rotate := getRand().Intn(2)
		if rotate == 1 {
			gbt.LeftRotate(v)
		} else {
			gbt.RightRotate(v)
		}
	}
	gbt.InOrderWalk(gbt.Root(), checkGBT(t, &nodeCnt, *debug))
	if nodeCnt != len(arr) {
		t.Log(fmt.Sprintf("node cnt expect to %0d but get:%0d", len(arr), nodeCnt))
		t.Fail()
	}
}

func TestGBTRecrusive_InOrderWalk(t *testing.T) {
	arr := randomSlice(0, 20, 10)
	resultArr := make([]int, 0, 0)
	gbt := newGBT()
	for _, v := range arr {
		gbt.Insert(uint32(v))
	}
	sort.Ints(arr)
	gbt.InOrderWalk(gbt.Root(), func(GBT binaryTreeIf, node interface{}) bool {
		n := node.(*gbtElement)
		resultArr = append(resultArr, int(n.Key))
		return false
	})
	if !reflect.DeepEqual(resultArr, arr) {
		t.Log(fmt.Sprintf("expect:%v", arr) + fmt.Sprintf("but get:%v", resultArr))
		t.Fail()
	}
}

func TestGBTRecrusive_PreOrderWalk(t *testing.T) {
	arr := randomSlice(0, 20, 10)
	gbt := newGBT()
	for _, v := range arr {
		gbt.Insert(uint32(v))
	}
	checkGBTPreOrder(t, gbt, arr)
}

func TestGBTRecrusive_PostOrderWalk(t *testing.T) {
	arr := randomSlice(0, 20, 10)
	gbt := newGBT()
	for _, v := range arr {
		gbt.Insert(uint32(v))
	}
	checkGBTPostOrder(t, gbt, arr)
}
