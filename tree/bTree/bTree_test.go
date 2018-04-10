package bTree

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func GetRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomSlice(start int, end int, count int) []int {
	if end < start || (end-start) < count {
		return nil
	}
	nums := make([]int, 0)
	for len(nums) < count {
		num := GetRand().Intn((end - start)) + start
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
		}
	}
	return nums
}

type testBTreeNode struct {
	bTreeNode
}

func (n *testBTreeNode) init(t int) *testBTreeNode {
	n.bTreeNode.init(t, n)
	return n
}

func (n *testBTreeNode) LessByKey(key1, key2 interface{}) bool {
	return key1.(int) < key2.(int)
}

func newTestBTreeNode(t int) *testBTreeNode {
	return new(testBTreeNode).init(t)
}

type testBTree struct {
	bTree
}

func (bt *testBTree) init(t int) *testBTree {
	bt.bTree.init(t, bt)
	return bt
}

func (bt *testBTree) newNode(t int) *bTreeNode {
	return &newTestBTreeNode(t).bTreeNode
}

func newTestBTree(t int) *testBTree {
	return new(testBTree).init(t)
}

func checkBtree(t *testing.T, exp []int, bt *bTree) {
	result := make([]int, 0, 0)
	bt.inOrderWalk(bt.root,
		func(btree *bTree, node *bTreeNode, idx int) bool {
			result = append(result, node.keyValue[idx].key.(int))
			if node != bt.root && node.Len() < bt.t-1 {
				t.Log(fmt.Sprintf("len of non-root node is less than t - 1: \n%+v\n ", node))
				t.Fail()
			}
			if node.Len() > 2*bt.t-1 {
				t.Log(fmt.Sprintf("len of node is larger than 2*t - 1: \n%+v\n ", node))
				t.Fail()
			}
			if len(node.c) > 2*bt.t {
				t.Log(fmt.Sprintf("num of children of node is larger than 2*t: \n%+v\n ", node))
				t.Fail()
			}
			if len(node.c) != node.Len()+1 {
				t.Log(fmt.Sprintf("num of children of node is not match the len of node: \n%+v\n ", node))
				t.Fail()
			}
			return false
		})
	if !reflect.DeepEqual(exp, result) {
		t.Log("inOrder walk error!")
		t.Log("exp:")
		t.Log(exp)
		t.Log("result:")
		t.Log(result)
		t.Fail()
	}
}

func TestBTreeInsert(t *testing.T) {
	bt := newTestBTree(2)
	arr := RandomSlice(0, 20, 10)
	exp := make([]int, len(arr), cap(arr))
	copy(exp, arr)
	sort.Ints(exp)
	for i := range arr {
		bt.insert(arr[i], arr[i])
	}
	checkBtree(t, exp, &bt.bTree)
	for i := range arr {
		bt.insert(arr[i], arr[i])
	}
	checkBtree(t, exp, &bt.bTree)
}

func TestBTreeRemove(t *testing.T) {
	deleteExp := func(arr []int, e int) []int {
		//binary search
		i, j := 0, len(arr)-1
		for i != j {
			mid := (j-i)/2 + i
			if arr[mid] == e {
				i = mid
				j = mid
			} else if e < arr[mid] {
				j = mid
			} else {
				i = mid + 1
			}
		}
		if len(arr)-1 == 0 {
			return []int{}
		}
		if i == len(arr)-1 {
			return arr[:len(arr)-1]
		}
		return append(arr[:i], arr[i+1:]...)
	}

	bt := newTestBTree(2)
	arr := RandomSlice(0, 20, 10)
	removeOrder := RandomSlice(0, 10, 10)
	exp := make([]int, len(arr), cap(arr))
	copy(exp, arr)
	sort.Ints(exp)
	for i := range arr {
		bt.insert(arr[i], arr[i])
	}
	checkBtree(t, exp, &bt.bTree)
	for _, v := range removeOrder {
		bt.remove(arr[v])
		exp = deleteExp(exp, arr[v])
		checkBtree(t, exp, &bt.bTree)
	}
}
