package rbTree

import (
	"testing"
	"algorithms/tree/binaryTree"
	"algorithms/tree/binaryTree/genericBinaryTree"
	"fmt"
)

func checkRBT(t *testing.T, rbt *RBT)(bool) {
	blackCntQ := make([]int, 0, 0)
	stop := rbt.PreOrderWalk(rbt.Root(), func(tree binaryTree.BinaryTreeIf, node interface{}) bool {
		_tree := tree.(*RBT)
		n := node.(*genericBinaryTree.GBTElement)
		//root must be black
		if n == _tree.Root() {
			if !_tree.color(n) {
				t.Log(fmt.Sprintf("root node  ", n, "is not black!"))
				t.Fail()
				return true
			}
		}
		//children of red node must be both black
		if !_tree.color(n) {
			if !_tree.IsNil(n.Left) && !_tree.color(n.Left) {
				t.Log(fmt.Sprintf("left node  ", n.Left, "of a red node", n, "is not black!"))
				t.Fail()
				return true
			}
			if !_tree.IsNil(n.Right) && !_tree.color(n.Right) {
				t.Log(fmt.Sprintf("right node  ", n.Right, "of a red node", n, "is not black!"))
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
					t.Log(fmt.Sprintf("black cnt", blackCnt, " in the path from node  ", n, " to root does not equal to the previous black cnt ", blackCntQ[0]))
					t.Fail()
					return true
				}
			} else {
				if blackCnt == 0 {
					t.Log(fmt.Sprintf("black cnt in the path from node  ", n, " to root is 0!"))
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
