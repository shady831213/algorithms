package binaryTree

import (
	"flag"
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

var debug = flag.Bool("debug", false, "debug flag")

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

func checkBst(t *testing.T, nodeCnt *int, debug bool) func(BinaryTreeIf, interface{}) bool {
	return func(tree BinaryTreeIf, node interface{}) bool {
		n := node.(*BstElement)
		if n.left != nil && n.left.Key >= n.Key {
			t.Log(fmt.Sprintf("left child %+v of node: %+v is more than or equal to n!", n.left, n))
			t.Fail()
			return true
		}
		if n.right != nil && n.right.Key <= n.Key {
			t.Log(fmt.Sprintf("right child %+v of node: %+v is less than or equal to n!", n.right, n))
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

func checkBstPreOrder(t *testing.T, bst BinaryTreeIf) {
	resultArr := make([]int, 0, 0)

	bst.PreOrderWalk(bst.Root(), func(tree BinaryTreeIf, node interface{}) bool {
		n := node.(*BstElement)
		resultArr = append(resultArr, int(n.Key))
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

func checkBstPostOrder(t *testing.T, bst BinaryTreeIf) {
	resultArr := make([]int, 0, 0)
	expArr := make([]int, 0, 0)
	bst.PostOrderWalk(bst.Root(), func(tree BinaryTreeIf, node interface{}) bool {
		n := node.(*BstElement)
		resultArr = append(resultArr, int(n.Key))
		return false
	})

	curNode := bst.Min(bst.Root()).(*BstElement)
	for curNode.right != nil {
		curNode = bst.Min(curNode.right).(*BstElement)
	}
	expArr = append(expArr, int(curNode.Key))
	for curNode.parent != nil {
		//if it is left child ,get the min of which node's right child is not nil from parent, if it is rignt child, get the parent
		nextNode := curNode.parent
		if curNode == nextNode.left {
			for nextNode.right != nil {
				nextNode = bst.Min(nextNode.right).(*BstElement)
			}
		}
		curNode = nextNode
		expArr = append(expArr, int(curNode.Key))
	}

	if !reflect.DeepEqual(resultArr, expArr) {
		t.Log(fmt.Sprintf("expect:%v", expArr) + fmt.Sprintf("but get:%v", resultArr))
		t.Fail()
	}
}

func CheckGBT(t *testing.T, nodeCnt *int, debug bool) func(BinaryTreeIf, interface{}) bool {
	return func(tree BinaryTreeIf, node interface{}) bool {
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

func checkGBTPreOrder(t *testing.T, tree BinaryTreeIf, data []int) {
	resultArr := make([]int, 0, 0)
	expArr := make([]int, 0, 0)
	tree.PreOrderWalk(tree.Root(), func(tree BinaryTreeIf, node interface{}) bool {
		n := node.(*GBTElement)
		resultArr = append(resultArr, int(n.Key))
		return false
	})

	expBst := NewBstRecrusive()
	for _, v := range data {
		expBst.Insert(uint32(v))
	}
	expBst.PreOrderWalk(expBst.Root(), func(tree BinaryTreeIf, node interface{}) bool {
		n := node.(*BstElement)
		expArr = append(expArr, int(n.Key))
		return false
	})
	if !reflect.DeepEqual(resultArr, expArr) {
		t.Log(fmt.Sprintf("expect:%v", expArr) + fmt.Sprintf("but get:%v", resultArr))
		t.Fail()
	}
}

func checkGBTPostOrder(t *testing.T, tree BinaryTreeIf, data []int) {
	resultArr := make([]int, 0, 0)
	expArr := make([]int, 0, 0)
	tree.PostOrderWalk(tree.Root(), func(tree BinaryTreeIf, node interface{}) bool {
		n := node.(*GBTElement)
		resultArr = append(resultArr, int(n.Key))
		return false
	})
	expBst := NewBstRecrusive()
	for _, v := range data {
		expBst.Insert(uint32(v))
	}
	expBst.PostOrderWalk(expBst.Root(), func(tree BinaryTreeIf, node interface{}) bool {
		n := node.(*BstElement)
		expArr = append(expArr, int(n.Key))
		return false
	})
	if !reflect.DeepEqual(resultArr, expArr) {
		t.Log(fmt.Sprintf("expect:%v", expArr) + fmt.Sprintf("but get:%v", resultArr))
		t.Fail()
	}
}

func checkRBT(t *testing.T, rbt *RBT) bool {
	blackCntQ := make([]int, 0, 0)
	stop := rbt.PreOrderWalk(rbt.Root(), func(tree BinaryTreeIf, node interface{}) bool {
		_tree := tree.(*RBT)
		n := node.(*GBTElement)
		//root must be black
		if n == _tree.Root() {
			if !_tree.color(n) {
				t.Log(fmt.Sprintf("root node  %+v is not black!", n))
				t.Fail()
				return true
			}
		}
		//children of red node must be both black
		if !_tree.color(n) {
			if !_tree.IsNil(n.Left) && !_tree.color(n.Left) {
				t.Log(fmt.Sprintf("left node  %+v of a red node %+v is not black!", n.Left, n))
				t.Fail()
				return true
			}
			if !_tree.IsNil(n.Right) && !_tree.color(n.Right) {
				t.Log(fmt.Sprintf("right node  %+v of a red node %+v is not black!", n.Right, n))
				t.Fail()
				return true
			}
		}
		//check blackcnt, leaves to root path must be the same
		blackCnt := 0
		if _tree.IsNil(n.Right) && _tree.IsNil(n.Left) {
			for curNode := n; !_tree.IsNil(curNode); curNode = curNode.Parent {
				if _tree.color(curNode) {
					blackCnt++
				}
			}
			if len(blackCntQ) != 0 {
				if blackCnt != blackCntQ[0] {
					t.Log(fmt.Sprintf("black cnt %0d in the path from node  %+v to root does not equal to the previous black cnt %0d", blackCnt, n, blackCntQ[0]))
					t.Fail()
					return true
				}
			} else {
				if blackCnt == 0 {
					t.Log(fmt.Sprintf("black cnt in the path from node  %0d  to root is 0!", n))
					t.Fail()
					return true
				}
				blackCntQ = append(blackCntQ, blackCnt)
			}
		}
		return false
	})

	if len(blackCntQ) == 0 {
		t.Error("black cnt collect error!")
		return true
	}
	return stop
}
