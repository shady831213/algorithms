package hashMap

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_ChainedHashMap(t *testing.T) {
	cmap := NewChainedHashMap()
	BasicTestHashMap(t, cmap)
}

func Test_ChainedHashMapUpScale(t *testing.T) {
	cmap := NewChainedHashMap()
	TestHashMapResize(t, cmap)
	if !reflect.DeepEqual(cmap.Cap, uint32(16)) {
		t.Log(fmt.Sprintf("expect:%0d ", uint32(16)) + fmt.Sprintf("but get:%0d ", cmap.Cap))
		t.Fail()
	}
}

func Test_ChainedHashMapDelete(t *testing.T) {
	cmap := NewChainedHashMap()
	for i := 0; i < 4; {
		TestHashMapDelete(t, cmap)
		if !reflect.DeepEqual(cmap.Count, uint32(0)) {
			t.Log(fmt.Sprintf("expect:%0d ", 0) + fmt.Sprintf("but get:%0d ", cmap.Count))
			t.Fail()
		}
		if !reflect.DeepEqual(cmap.Cap, uint32(0)) {
			t.Log(fmt.Sprintf("expect:%0d ", uint32(0)) + fmt.Sprintf("but get:%0d ", cmap.Cap))
			t.Fail()
		}
		i++
	}
}

func BenchmarkChainedHashMap_HashInsert(b *testing.B) {
	BenchmarkHashMapInsert(b, NewChainedHashMap())
}

func BenchmarkChainedHashMap_HashInsertDelete(b *testing.B) {
	BenchmarkHashMapInsertDelete(b, NewChainedHashMap())
}

func BenchmarkChainedHashMap_HashGet(b *testing.B) {
	BenchmarkHashMapGet(b, NewChainedHashMap())
}
