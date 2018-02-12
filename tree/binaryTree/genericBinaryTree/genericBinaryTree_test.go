package genericBinaryTree
import (
	"testing"
	"algorithms/tree"
	"fmt"
	"flag"
	"sort"
	"reflect"
	"algorithms/tree/binaryTree"
)

var debug  = flag.Bool("debug", false, "debug flag")

func TestGBT_Insert(t *testing.T) {
	arr := tree.RandomSlice(0,20, 10)
	nodeCnt := 0
	gbt := New()
	for _,v := range arr {
		gbt.Insert(uint32(v))
	}
	gbt.InOrderWalk(gbt.Root(), checkGBT(t, &nodeCnt, *debug))
	if nodeCnt != len(arr) {
		t.Log(fmt.Sprintf("node cnt expect to ", len(arr), "but get:",nodeCnt))
		t.Fail()
	}
}

func TestGBT_Delete(t *testing.T) {
	arr := tree.RandomSlice(0,20, 10)
	nodeCnt := 0
	gbt := New()
	for _,v := range arr {
		gbt.Insert(uint32(v))
	}
	delNum := tree.GetRand().Intn(10)+1
	if *debug {
		fmt.Println(delNum)
	}
	for i:=0;i < delNum;i++{
		gbt.Delete(uint32(arr[i]))
	}
	gbt.InOrderWalk(gbt.Root(), checkGBT(t, &nodeCnt, *debug))
	if nodeCnt != len(arr) - delNum {
		t.Log(fmt.Sprintf("node cnt expect to ", len(arr) - delNum, "but get:",nodeCnt))
		t.Fail()
	}

}

func TestGBT_Min(t *testing.T) {
	arr := tree.RandomSlice(0,20, 10)
	gbt := New()
	for _,v := range arr {
		gbt.Insert(uint32(v))
	}
	min := int(gbt.Min(gbt.Root()).(*GBTElement).Key)
	sort.Ints(arr)
	if min != arr[0] {
		t.Log(fmt.Sprintf("min expect to ", arr[0], "but get:",min))
		t.Fail()
	}
}

func TestGBT_Max(t *testing.T) {
	arr := tree.RandomSlice(0,20, 10)
	gbt := New()
	for _,v := range arr {
		gbt.Insert(uint32(v))
	}
	max := int(gbt.Max(gbt.Root()).(*GBTElement).Key)
	sort.Ints(arr)
	if max != arr[len(arr)-1] {
		t.Log(fmt.Sprintf("max expect to ", arr[len(arr)-1], "but get:",max))
		t.Fail()
	}
}

func TestGBT_Search(t *testing.T) {
	arr := tree.RandomSlice(0,20, 10)
	gbt := New()
	for _,v := range arr {
		gbt.Insert(uint32(v))
	}
	key := tree.GetRand().Intn(len(arr))
	result := int(gbt.Search(uint32(arr[key])).(*GBTElement).Key)
	if result != arr[key] {
		t.Log(fmt.Sprintf("search result expect to ", arr[key], "but get:",result))
		t.Fail()
	}
}

func TestGBT_Predecesor(t *testing.T) {
	arr := tree.RandomSlice(0,20, 10)
	gbt := New()
	for _,v := range arr {
		gbt.Insert(uint32(v))
	}
	sort.Ints(arr)
	key := tree.GetRand().Intn(len(arr) - 2) + 1
	result := int(gbt.Predecesor(gbt.Search(uint32(arr[key]))).(*GBTElement).Key)
	if result != arr[key-1] {
		t.Log(fmt.Sprintf("Predecesor of",arr[key], " expect to ", arr[key-1] , "but get:",result))
		t.Fail()
	}
}

func TestGBT_Successor(t *testing.T) {
	arr := tree.RandomSlice(0,20, 10)
	gbt := New()
	for _,v := range arr {
		gbt.Insert(uint32(v))
	}
	sort.Ints(arr)
	key := tree.GetRand().Intn(len(arr) - 2) + 1
	result := int(gbt.Successor(gbt.Search(uint32(arr[key]))).(*GBTElement).Key)
	if result != arr[key+1] {
		t.Log(fmt.Sprintf("Successor of",arr[key], " expect to ", arr[key+1] , "but get:",result))
		t.Fail()
	}
}

func TestGBTRecrusive_InOrderWalk(t *testing.T) {
	arr := tree.RandomSlice(0,20, 10)
	resultArr := make([]int, 0, 0)
	gbt := New()
	for _,v := range arr {
		gbt.Insert(uint32(v))
	}
	sort.Ints(arr)
	gbt.InOrderWalk(gbt.Root(), func(tree binaryTree.BinaryTreeIf,node interface{}) bool {
		n := node.(*GBTElement)
		resultArr = append(resultArr,int(n.Key))
		return false
	})
	if !reflect.DeepEqual(resultArr, arr) {
		t.Log(fmt.Sprintf("expect:%v", arr) + fmt.Sprintf("but get:%v", resultArr))
		t.Fail()
	}
}

func TestGBTRecrusive_PreOrderWalk(t *testing.T) {
	arr := tree.RandomSlice(0,20, 10)
	gbt := New()
	for _,v := range arr {
		gbt.Insert(uint32(v))
	}
	checkGBTPreOrder(t, gbt)
}

func TestGBTRecrusive_PostOrderWalk(t *testing.T) {
	arr := tree.RandomSlice(0,20, 10)
	gbt := New()
	for _,v := range arr {
		gbt.Insert(uint32(v))
	}
	checkGBTPostOrder(t, gbt)
}
