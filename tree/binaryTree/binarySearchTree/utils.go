package binarySearchTree

import (
	"fmt"
	"testing"
	"algorithms/tree/binaryTree"
	"reflect"
)

func checkBst(t *testing.T, nodeCnt *int, debug bool) (func(interface{}) (bool)) {
	return func(node interface{}) bool {
		n := node.(*BstElement)
		if n.left != nil && n.left.Key >= n.Key {
			t.Log(fmt.Sprintf("left child ", n.left, "of node:", n, "is more than or equal to n!"))
			t.Fail()
			return true
		}
		if n.right != nil && n.right.Key <= n.Key {
			t.Log(fmt.Sprintf("right child ", n.right, "of node:", n, "is less than or equal to n!"))
			t.Fail()
			return true
		}
		if debug {
			fmt.Println(n)
		}
		*nodeCnt++
		return false
	}
}

func checkBstPreOrder(t *testing.T, bst binaryTree.BinaryTreeIf) {
	resultArr := make([]int, 0, 0)

	bst.PreOrderWalk(bst.Root(), func(node interface{}) bool {
		n := node.(*BstElement)
		resultArr = append(resultArr,int(n.Key))
		return false
	})

	if bst.Root().(*BstElement).Key != uint32(resultArr[0]) {
		t.Log("first element is expected to be root ", bst.Root().(*BstElement).Key, " but get ", uint32(resultArr[0]))
		t.Fail()
	}

	min := int(bst.Min(bst.Root()).(*BstElement).Key)
	for i := range resultArr[:len(resultArr)-1] {
		if resultArr[i] > resultArr[i+1] {
			//expect decreasing numbers down to min
			if resultArr[i] < min {
				t.Log("element ", resultArr[i], " @", i, "is expected to be more than ", min)
				t.Fail()
				return
			}
		} else {
			//tail of decreasing numbers should be min
			if resultArr[i] != min {
				t.Log("element ", resultArr[i], " @", i, "is expected to equal to ", min)
				t.Fail()
				return
			}
			//next min value should be the min of the nearest right sub tree
			nextMin := int(bst.Min(bst.Search(uint32(resultArr[i+1]))).(*BstElement).Key)
			if nextMin <= min {
				t.Log("next min", nextMin, "of element ", resultArr[i+1], " @", i+1, "is expected to be more than ", min)
				t.Fail()
				return
			} else {
				min = nextMin
			}
		}
	}
}

func checkBstPostOrder(t *testing.T, bst binaryTree.BinaryTreeIf) {
	resultArr := make([]int, 0, 0)
	expArr:= make([]int, 0, 0)
	bst.PostOrderWalk(bst.Root(), func(node interface{}) bool {
		n := node.(*BstElement)
		resultArr = append(resultArr,int(n.Key))
		return false
	})

	curNode:=bst.Min(bst.Root()).(*BstElement)
	for curNode.right != nil {
		curNode = bst.Successor(curNode).(*BstElement)
	}
	expArr = append(expArr,int(curNode.Key))
	for curNode.parent!=nil {
		//if it is left child ,get the successor of which node's right child is not nil from parent, if it is rignt child, get the parent
		nextNode := curNode.parent
		if curNode == nextNode.left {
			for nextNode.right != nil {
				nextNode = bst.Successor(nextNode).(*BstElement)
			}
		}
		curNode = nextNode
		expArr = append(expArr,int(curNode.Key))
	}

	if !reflect.DeepEqual(resultArr, expArr) {
		t.Log(fmt.Sprintf("expect:%v", expArr) + fmt.Sprintf("but get:%v", resultArr))
		t.Fail()
	}
}