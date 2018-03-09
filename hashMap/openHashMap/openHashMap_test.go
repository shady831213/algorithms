package openHashMap
import (
	"testing"
	"algorithms/hashMap"
	"reflect"
	"fmt"
)

func Test_OpenHashMap(t *testing.T) {
	cmap := New()
	hashMap.BasicTestHashMap(t,cmap)
}

func Test_OpenHashMapResize(t *testing.T) {
	cmap := New()
	hashMap.TestHashMapResize(t, cmap)
	if !reflect.DeepEqual(cmap.Cap, uint32(8)) {
		t.Log(fmt.Sprintf("expect:", uint32(8)) + fmt.Sprintf("but get:", cmap.Cap))
		t.Fail()
	}
}

func Test_OpenHashMapDelete(t *testing.T) {
	cmap := New()
	for i:=0;i<4; {
		hashMap.TestHashMapDelete(t, cmap)
		if !reflect.DeepEqual(cmap.Count, uint32(0)) {
			t.Log(fmt.Sprintf("expect:", 0) + fmt.Sprintf("but get:", cmap.Count))
			t.Fail()
		}
		if !reflect.DeepEqual(cmap.Cap, uint32(0)) {
			t.Log(fmt.Sprintf("expect:", 0) + fmt.Sprintf("but get:", cmap.Cap))
			t.Fail()
		}
		i++
	}
}

func BenchmarkOpenHashMap_HashInsert(b *testing.B) {
	hashMap.BenchmarkHashMapInsert(b,New())
}

func BenchmarkOpenHashMap_HashInsertDelete(b *testing.B) {
	hashMap.BenchmarkHashMapInsertDelete(b,New())
}

func BenchmarkOpenHashMap_HashGet(b *testing.B) {
	hashMap.BenchmarkHashMapGet(b,New())
}