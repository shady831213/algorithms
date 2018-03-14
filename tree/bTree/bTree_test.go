package bTree

import (
	"testing"
	"algorithms/tree"
	"sort"
	"reflect"
	"fmt"
)

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

func (bt *testBTree) newNode(t int) (*bTreeNode) {
	return &newTestBTreeNode(t).bTreeNode
}

func newTestBTree(t int) (*testBTree) {
	return new(testBTree).init(t)
}

func checkBtree(t *testing.T,exp []int,bt *bTree) {
	result := make([]int, 0, 0)
	bt.inOrderWalk(bt.root,
		func(btree *bTree, node *bTreeNode, idx int) (bool) {
			result = append(result, node.keyValue[idx].key.(int))
			if node != bt.root && node.Len() < bt.t - 1{
				t.Log(fmt.Sprintf("len of non-root node is less than t - 1: \n%+v\n ", node))
				t.Fail()
			}
			if node.Len() > 2*bt.t -1 {
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
	if !reflect.DeepEqual(exp,result) {
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
	arr := tree.RandomSlice(0, 20, 10)
	exp := make([]int, len(arr), cap(arr))
	copy(exp,arr)
	sort.Ints(exp)
	for i := range arr {
		bt.insert(arr[i], arr[i])
	}
	checkBtree(t,exp,&bt.bTree)
	for i := range arr {
		bt.insert(arr[i], arr[i])
	}
	checkBtree(t,exp,&bt.bTree)
}
