package genericBinaryTree

import (
	"fmt"
	"testing"
	"algorithms/tree/binaryTree"
	"algorithms/tree/binaryTree/binarySearchTree"
	"reflect"
)

func CheckGBT(t *testing.T, nodeCnt *int, debug bool) (func(binaryTree.BinaryTreeIf, interface{}) (bool)) {
	return func(tree binaryTree.BinaryTreeIf, node interface{}) bool {
		n := node.(*GBTElement)
		if !tree.IsNil(n.Left) && n.Left.Key >= n.Key {
			t.Log(fmt.Sprintf("Left child %+v of node: %+v is more than or equal to n!", n.Left, n))
			t.Fail()
			return true
		}
		if !tree.IsNil(n.Right) && n.Right.Key <= n.Key {
			t.Log(fmt.Sprintf("Right child %+v of node: %+v is less than or equal to n!", n.Right, n))
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

func checkGBTPreOrder(t *testing.T, tree binaryTree.BinaryTreeIf, data []int) {
	resultArr := make([]int, 0, 0)
	expArr := make([]int,0,0)
	tree.PreOrderWalk(tree.Root(), func(tree binaryTree.BinaryTreeIf, node interface{}) bool {
		n := node.(*GBTElement)
		resultArr = append(resultArr,int(n.Key))
		return false
	})

	expBst := binarySearchTree.NewBstRecrusive()
	for _,v := range data {
		expBst.Insert(uint32(v))
	}
	expBst.PreOrderWalk(expBst.Root(), func(tree binaryTree.BinaryTreeIf, node interface{}) bool {
		n := node.(*binarySearchTree.BstElement)
		expArr = append(expArr,int(n.Key))
		return false
	})
	if !reflect.DeepEqual(resultArr, expArr) {
		t.Log(fmt.Sprintf("expect:%v", expArr) + fmt.Sprintf("but get:%v", resultArr))
		t.Fail()
	}
}

func checkGBTPostOrder(t *testing.T, tree binaryTree.BinaryTreeIf,data []int) {
	resultArr := make([]int, 0, 0)
	expArr := make([]int,0,0)
	tree.PostOrderWalk(tree.Root(), func(tree binaryTree.BinaryTreeIf, node interface{}) bool {
		n := node.(*GBTElement)
		resultArr = append(resultArr,int(n.Key))
		return false
	})
	expBst := binarySearchTree.NewBstRecrusive()
	for _,v := range data {
		expBst.Insert(uint32(v))
	}
	expBst.PostOrderWalk(expBst.Root(), func(tree binaryTree.BinaryTreeIf, node interface{}) bool {
		n := node.(*binarySearchTree.BstElement)
		expArr = append(expArr,int(n.Key))
		return false
	})
	if !reflect.DeepEqual(resultArr, expArr) {
		t.Log(fmt.Sprintf("expect:%v", expArr) + fmt.Sprintf("but get:%v", resultArr))
		t.Fail()
	}
}