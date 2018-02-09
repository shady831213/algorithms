package binarySearchTree

import (
	"fmt"
	"testing"
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
