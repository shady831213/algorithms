package vEBTree

import (
	"testing"
	"fmt"
	"algorithms/tree"
	"math/rand"
)

func TestRsVEBTreeElement_InsertMemberBasic(t *testing.T) {
	datas := map[uint32][]uint32{1: {1},
		2: {2, 3},
		100: {60, 70},
		9: {8, 9},
		111: {79, 86},
		47: {31, 2}}
	vEBT := newRsVEBTreeUint32(8)
	for i := range datas {
		for j := range datas[i] {
			vEBT.Insert(i, datas[i][j])
		}
	}
	//check
	for i := range datas {
		member := vEBT.Member(i)
		if member.key != i {
			t.Log(fmt.Sprintf("key error! exp = %0d, result = %0d", i, member.key))
			t.Fail()
		}
		e := member.value.Front()
		for j := range datas[i] {
			if e.Value != datas[i][j] {
				t.Log(fmt.Sprintf("value error @ key %0d ! exp = %0d, result = %0d", i, datas[i][j], e.Value))
				t.Fail()
			}
			e = e.Next()
		}
	}
}

func TestRsVEBTreeElement_InsertMember(t *testing.T) {
	lgu := rand.Intn(31) + 7
	datasKey := tree.RandomSlice(1, (1<<uint32(lgu))-1, rand.Intn((1<<7)-1)+1)
	datas := make(map[uint32][]int)
	for _, v := range datasKey {
		arr := tree.RandomSlice(0, 64, rand.Intn(10)+1)
		datas[uint32(v)] = arr
	}
	vEBT := newRsVEBTreeUint32(lgu)
	for i := range datas {
		for j := range datas[i] {
			vEBT.Insert(i, datas[i][j])
		}
	}
	//check
	for i := range datas {
		member := vEBT.Member(i)
		if member.key != i {
			t.Log(fmt.Sprintf("key error! exp = %0d, result = %0d", i, member.key))
			t.Fail()
		}
		e := member.value.Front()
		for j := range datas[i] {
			if e.Value != datas[i][j] {
				t.Log(fmt.Sprintf("value error @ key %0d ! exp = %0d, result = %0d", i, datas[i][j], e.Value))
				t.Fail()
			}
			e = e.Next()
		}
	}
}
