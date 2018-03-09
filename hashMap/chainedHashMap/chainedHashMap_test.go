package chainedHashMap

import (
	"testing"
	"algorithms/hashMap"
	"reflect"
	"fmt"
)

func Test_ChainedHashMap(t *testing.T) {
	cmap := New()
	hashMap.BasicTestHashMap(t,cmap)
}

func Test_ChainedHashMapUpScale(t *testing.T) {
	cmap := New()
	hashMap.TestHashMapResize(t, cmap)
	if !reflect.DeepEqual(cmap.Cap, uint32(16)) {
		t.Log(fmt.Sprintf("expect:", uint32(16)) + fmt.Sprintf("but get:", cmap.Cap))
		t.Fail()
	}
}

func Test_ChainedHashMapDelete(t *testing.T) {
	cmap := New()
	for i:=0;i<4; {
		hashMap.TestHashMapDelete(t, cmap)
		if !reflect.DeepEqual(cmap.Count, uint32(0)) {
			t.Log(fmt.Sprintf("expect:", 0) + fmt.Sprintf("but get:", cmap.Count))
			t.Fail()
		}
		if !reflect.DeepEqual(cmap.Cap, uint32(0)) {
			t.Log(fmt.Sprintf("expect:", uint32(0)) + fmt.Sprintf("but get:", cmap.Cap))
			t.Fail()
		}
		i++
	}
}

func BenchmarkChainedHashMap_HashInsert(b *testing.B) {
	hashMap.BenchmarkHashMapInsert(b,New())
}

func BenchmarkChainedHashMap_HashInsertDelete(b *testing.B) {
	hashMap.BenchmarkHashMapInsertDelete(b,New())
}

func BenchmarkChainedHashMap_HashGet(b *testing.B) {
	hashMap.BenchmarkHashMapGet(b,New())
}