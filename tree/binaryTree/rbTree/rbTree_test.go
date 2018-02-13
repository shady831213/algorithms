package rbTree

import (
	"flag"
	"testing"
	"algorithms/tree"
	"algorithms/tree/binaryTree/genericBinaryTree"
)

var debug  = flag.Bool("debug", false, "debug flag")

func TestRBT_Insert(t *testing.T) {
	arr := tree.RandomSlice(0,20, 10)
	nodeCnt := 0
	rbt := New()
	for i,v := range arr {
		rbt.Insert(uint32(v))
		stop := rbt.InOrderWalk(rbt.Root(), genericBinaryTree.CheckGBT(t, &nodeCnt, *debug))
		if stop {
			return
		}
		if nodeCnt != i+1 {
			t.Log("node cnt expect to ", i+1, "but get:",nodeCnt)
			t.Fail()
		}
		nodeCnt = 0
		checkRBT(t,rbt)
	}
}
