package genericBinaryTree

import (
	"fmt"
	"testing"
	"algorithms/tree/binaryTree"
	"reflect"
)

func checkGBT(t *testing.T, nodeCnt *int, debug bool) (func(binaryTree.BinaryTreeIf, interface{}) (bool)) {
	return func(tree binaryTree.BinaryTreeIf, node interface{}) bool {
		n := node.(*GBTElement)
		if !tree.IsNil(n.left) && n.left.Key >= n.Key {
			t.Log(fmt.Sprintf("left child ", n.left, "of node:", n, "is more than or equal to n!"))
			t.Fail()
			return true
		}
		if !tree.IsNil(n.right) && n.right.Key <= n.Key {
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

func checkGBTPreOrder(t *testing.T, tree binaryTree.BinaryTreeIf) {
	resultArr := make([]int, 0, 0)

	tree.PreOrderWalk(tree.Root(), func(tree binaryTree.BinaryTreeIf, node interface{}) bool {
		n := node.(*GBTElement)
		resultArr = append(resultArr,int(n.Key))
		return false
	})

	if tree.Root().(*GBTElement).Key != uint32(resultArr[0]) {
		t.Log("first element is expected to be root ", tree.Root().(*GBTElement).Key, " but get ", uint32(resultArr[0]))
		t.Fail()
	}

	min := int(tree.Min(tree.Root()).(*GBTElement).Key)
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
			nextMin := int(tree.Min(tree.Search(uint32(resultArr[i+1]))).(*GBTElement).Key)
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

func checkGBTPostOrder(t *testing.T, tree binaryTree.BinaryTreeIf) {
	resultArr := make([]int, 0, 0)
	expArr:= make([]int, 0, 0)
	tree.PostOrderWalk(tree.Root(), func(tree binaryTree.BinaryTreeIf, node interface{}) bool {
		n := node.(*GBTElement)
		resultArr = append(resultArr,int(n.Key))
		return false
	})

	curNode:= tree.Min(tree.Root()).(*GBTElement)
	for !tree.IsNil(curNode.right) {
		curNode = tree.Min(curNode.right).(*GBTElement)
	}
	expArr = append(expArr,int(curNode.Key))
	for !tree.IsNil(curNode.parent) {
		//if it is left child ,get the min of which node's right child is not nil from parent, if it is rignt child, get the parent
		nextNode := curNode.parent
		if curNode == nextNode.left {
			for !tree.IsNil(nextNode.right) {
				nextNode = tree.Min(nextNode.right).(*GBTElement)
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