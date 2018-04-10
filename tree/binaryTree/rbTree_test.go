package rbTree

import (
	"flag"
	"testing"
	"math/rand"
	"algorithms/tree/binaryTree/genericBinaryTree"
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

func TestRBT_Insert(t *testing.T) {
	arr := RandomSlice(0, 20, 10)
	nodeCnt := 0
	rbt := New()
	for i, v := range arr {
		rbt.Insert(uint32(v))
		stop := rbt.PreOrderWalk(rbt.Root(), genericBinaryTree.CheckGBT(t, &nodeCnt, *debug))
		if stop {
			return
		}
		if nodeCnt != i+1 {
			t.Log("node cnt expect to ", i+1, "but get:", nodeCnt)
			t.Fail()
		}
		nodeCnt = 0
		stop = checkRBT(t, rbt)
		if stop {
			return
		}
	}
}

func TestRBT_Delete(t *testing.T) {
	arr := RandomSlice(0, 20, 10)
	deleteSequence := RandomSlice(0, 10, 10)
	nodeCnt := 0
	rbt := New()
	for _, v := range arr {
		rbt.Insert(uint32(v))
	}
	for i, v := range deleteSequence {
		rbt.Delete(uint32(arr[v]))
		stop := rbt.PreOrderWalk(rbt.Root(), genericBinaryTree.CheckGBT(t, &nodeCnt, *debug))
		if stop {
			return
		}
		if nodeCnt != len(deleteSequence)-1-i {
			t.Log("node cnt expect to ", len(deleteSequence)-1-i, "but get:", nodeCnt)
			t.Fail()
		}
		nodeCnt = 0
		if i != len(deleteSequence)-1 {
			stop = checkRBT(t, rbt)
			if stop {
				return
			}
		}
	}
}
