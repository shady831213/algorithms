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

func getRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func randomSlice(start int, end int, count int) []int {
	if end < start || (end-start) < count {
		return nil
	}
	nums := make([]int, 0)
	for len(nums) < count {
		num := getRand().Intn((end - start)) + start
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

func checkBst(t *testing.T, nodeCnt *int, debug bool) func(binaryTreeIf, interface{}) bool {
	return func(tree binaryTreeIf, node interface{}) bool {
		n := node.(*bstElement)
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

func checkBstPreOrder(t *testing.T, bst binaryTreeIf) {
	resultArr := make([]int, 0, 0)

	bst.PreOrderWalk(bst.Root(), func(tree binaryTreeIf, node interface{}) bool {
		n := node.(*bstElement)
		resultArr = append(resultArr, int(n.Key))
		return false
	})

	if bst.Root().(*bstElement).Key != uint32(resultArr[0]) {
		t.Log("first element is expected to be root ", bst.Root().(*bstElement).Key, " but get ", uint32(resultArr[0]))
		t.Fail()
	}

	min := int(bst.Min(bst.Root()).(*bstElement).Key)
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
			nextMin := int(bst.Min(bst.Search(uint32(resultArr[i+1]))).(*bstElement).Key)
			if nextMin <= min {
				t.Log("next min", nextMin, "of element ", resultArr[i+1], " @", i+1, "is expected to be more than ", min)
				t.Fail()
				return
			}
			min = nextMin
		}
	}
}

func checkBstPostOrder(t *testing.T, bst binaryTreeIf) {
	resultArr := make([]int, 0, 0)
	expArr := make([]int, 0, 0)
	bst.PostOrderWalk(bst.Root(), func(tree binaryTreeIf, node interface{}) bool {
		n := node.(*bstElement)
		resultArr = append(resultArr, int(n.Key))
		return false
	})

	curNode := bst.Min(bst.Root()).(*bstElement)
	for curNode.right != nil {
		curNode = bst.Min(curNode.right).(*bstElement)
	}
	expArr = append(expArr, int(curNode.Key))
	for curNode.parent != nil {
		//if it is left child ,get the min of which node's right child is not nil from parent, if it is rignt child, get the parent
		nextNode := curNode.parent
		if curNode == nextNode.left {
			for nextNode.right != nil {
				nextNode = bst.Min(nextNode.right).(*bstElement)
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

func checkGBT(t *testing.T, nodeCnt *int, debug bool) func(binaryTreeIf, interface{}) bool {
	return func(tree binaryTreeIf, node interface{}) bool {
		n := node.(*gbtElement)
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

func checkGBTPreOrder(t *testing.T, tree binaryTreeIf, data []int) {
	resultArr := make([]int, 0, 0)
	expArr := make([]int, 0, 0)
	tree.PreOrderWalk(tree.Root(), func(tree binaryTreeIf, node interface{}) bool {
		n := node.(*gbtElement)
		resultArr = append(resultArr, int(n.Key))
		return false
	})

	expBst := newBstRecrusive()
	for _, v := range data {
		expBst.Insert(uint32(v))
	}
	expBst.PreOrderWalk(expBst.Root(), func(tree binaryTreeIf, node interface{}) bool {
		n := node.(*bstElement)
		expArr = append(expArr, int(n.Key))
		return false
	})
	if !reflect.DeepEqual(resultArr, expArr) {
		t.Log(fmt.Sprintf("expect:%v", expArr) + fmt.Sprintf("but get:%v", resultArr))
		t.Fail()
	}
}

func checkGBTPostOrder(t *testing.T, tree binaryTreeIf, data []int) {
	resultArr := make([]int, 0, 0)
	expArr := make([]int, 0, 0)
	tree.PostOrderWalk(tree.Root(), func(tree binaryTreeIf, node interface{}) bool {
		n := node.(*gbtElement)
		resultArr = append(resultArr, int(n.Key))
		return false
	})
	expBst := newBstRecrusive()
	for _, v := range data {
		expBst.Insert(uint32(v))
	}
	expBst.PostOrderWalk(expBst.Root(), func(tree binaryTreeIf, node interface{}) bool {
		n := node.(*bstElement)
		expArr = append(expArr, int(n.Key))
		return false
	})
	if !reflect.DeepEqual(resultArr, expArr) {
		t.Log(fmt.Sprintf("expect:%v", expArr) + fmt.Sprintf("but get:%v", resultArr))
		t.Fail()
	}
}

func checkRBTRoot(t *testing.T, tree *rbt, n *gbtElement) {
	//root must be black
	if n == tree.Root() {
		if !tree.color(n) {
			t.Log(fmt.Sprintf("root node  %+v is not black!", n))
			t.FailNow()
		}
	}
}

func checkRBTRedNode(t *testing.T, tree *rbt, n *gbtElement) {
	//children of red node must be both black
	if !tree.color(n) {
		if !tree.IsNil(n.Left) && !tree.color(n.Left) {
			t.Log(fmt.Sprintf("left node  %+v of a red node %+v is not black!", n.Left, n))
			t.FailNow()
		}
		if !tree.IsNil(n.Right) && !tree.color(n.Right) {
			t.Log(fmt.Sprintf("right node  %+v of a red node %+v is not black!", n.Right, n))
			t.FailNow()
		}
	}
}

func checkRBTBlackPath(t *testing.T, tree *rbt, n *gbtElement, blackCntQ *[]int) {
	//check blackcnt, leaves to root path must be the same
	blackCnt := 0
	if tree.IsNil(n.Right) && tree.IsNil(n.Left) {
		for curNode := n; !tree.IsNil(curNode); curNode = curNode.Parent {
			if tree.color(curNode) {
				blackCnt++
			}
		}
		if len(*blackCntQ) != 0 {
			if blackCnt != (*blackCntQ)[0] {
				t.Log(fmt.Sprintf("black cnt %0d in the path from node  %+v to root does not equal to the previous black cnt %0d", blackCnt, n, (*blackCntQ)[0]))
				t.FailNow()
			}
		} else {
			if blackCnt == 0 {
				t.Log(fmt.Sprintf("black cnt in the path from node  %0d  to root is 0!", n))
				t.FailNow()
			}
			*blackCntQ = append(*blackCntQ, blackCnt)
		}
	}
}

func checkRBT(t *testing.T, rbTree *rbt) bool {
	blackCntQ := make([]int, 0, 0)

	stop := rbTree.PreOrderWalk(rbTree.Root(), func(tree binaryTreeIf, node interface{}) bool {
		checkRBTRoot(t, tree.(*rbt), node.(*gbtElement))
		checkRBTRedNode(t, tree.(*rbt), node.(*gbtElement))
		checkRBTBlackPath(t, tree.(*rbt), node.(*gbtElement), &blackCntQ)
		return false
	})

	if len(blackCntQ) == 0 {
		t.Error("black cnt collect error!")
		return true
	}
	return stop
}
