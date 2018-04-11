package hashMap

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_OpenHashMap(t *testing.T) {
	cmap := newOpenHashMap()
	basicTestHashMap(t, cmap)
}

func Test_OpenHashMapResize(t *testing.T) {
	cmap := newOpenHashMap()
	testHashMapResize(t, cmap)
	if !reflect.DeepEqual(cmap.Cap, uint32(16)) {
		t.Log(fmt.Sprintf("expect: %0d ", uint32(16)) + fmt.Sprintf("but get: %0d ", cmap.Cap))
		t.Fail()
	}
}

func Test_OpenHashMapDelete(t *testing.T) {
	cmap := newOpenHashMap()
	for i := 0; i < 4; {
		testHashMapDelete(t, cmap)
		if !reflect.DeepEqual(cmap.Count, uint32(0)) {
			t.Log(fmt.Sprintf("expect: %0d ", 0) + fmt.Sprintf("but get:%0d ", cmap.Count))
			t.Fail()
		}
		if !reflect.DeepEqual(cmap.Cap, uint32(0)) {
			t.Log(fmt.Sprintf("expect: %0d ", 0) + fmt.Sprintf("but get:%0d ", cmap.Cap))
			t.Fail()
		}
		i++
	}
}

func BenchmarkOpenHashMap_HashInsert(b *testing.B) {
	benchmarkHashMapInsert(b, newOpenHashMap())
}

func BenchmarkOpenHashMap_HashInsertDelete(b *testing.B) {
	benchmarkHashMapInsertDelete(b, newOpenHashMap())
}

func BenchmarkOpenHashMap_HashGet(b *testing.B) {
	benchmarkHashMapGet(b, newOpenHashMap())
}
