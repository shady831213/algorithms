package chainedHashMap

import (
	"testing"
	"reflect"
	"fmt"
)

func Test_ChainedHashMap(t *testing.T) {
	cmap := New(4)
	actMap := make(map[string]int)
	expMap := map[string]int{
		"a":   1,
		"b":   2,
		"c":   999,
		"z":   16,
		"bcd": 90,
	}
	for i,v := range expMap {
		cmap.HashInsert(i,v)
		value,_ := cmap.HashGet(i)
		actMap[i] = value.(int)
	}
	if !reflect.DeepEqual(actMap, expMap) {
		t.Log(fmt.Sprintf("expect:%m", expMap) + fmt.Sprintf("but get:%m", actMap))
		t.Fail()
	}
}

func Test_ChainedHashMapResize(t *testing.T) {
	cmap := new(ChainedHashMap)
	cmap.Init(4)
	expMap := map[string]int{
		"a":   1,
		"b":   2,
		"c":   999,
		"z":   16,
		"bcd": 90,
		"ed": 90,
	}
	for i,v := range expMap {
		cmap.HashInsert(i,v)
	}
	if !reflect.DeepEqual(cmap.Cap, uint32(8)) {
		t.Log(fmt.Sprintf("expect:", 8) + fmt.Sprintf("but get:", cmap.Cap))
		t.Fail()
	}
}

func Test_ChainedHashMapDelete(t *testing.T) {
	cmap := new(ChainedHashMap)
	cmap.Init(4)
	expMap := map[string]int{
		"a":   1,
		"b":   2,
		"c":   999,
		"z":   16,
		"bcd": 90,
		"ed": 90,
	}
	for i,v := range expMap {
		cmap.HashInsert(i,v)
		cmap.HashDelete(i)
		value, exist := cmap.HashGet(i)
		if value != nil{
			t.Log(fmt.Sprintf("expect value of key %s should be nil", i) + fmt.Sprintf("but get: %s", value))
			t.Fail()
		}
		if exist {
			t.Log(fmt.Sprintf("expect key %s should not exist", i))
			t.Fail()
		}
	}
	if !reflect.DeepEqual(cmap.Count, uint32(0)) {
		t.Log(fmt.Sprintf("expect:", 0) + fmt.Sprintf("but get:", cmap.Count))
		t.Fail()
	}
}

