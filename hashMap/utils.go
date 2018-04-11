package hashMap

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func basicTestHashMap(t *testing.T, hmap hashMap) {
	actMap := make(map[string]int)
	expMap := map[string]int{
		"a":   1,
		"b":   2,
		"c":   999,
		"z":   16,
		"bcd": 90,
	}
	for i, v := range expMap {
		hmap.HashInsert(i, v)
		value, _ := hmap.HashGet(i)
		actMap[i] = value.(int)
	}
	if !reflect.DeepEqual(actMap, expMap) {
		t.Log(fmt.Sprintf("expect:%+v", expMap) + fmt.Sprintf("but get:%+v", actMap))
		t.Fail()
	}
}

func testHashMapResize(t *testing.T, hmap interface{}) {
	expMap := map[string]int{
		"a":    1,
		"b":    2,
		"c":    999,
		"z":    16,
		"bcd":  90,
		"ed":   90,
		"ab":   1,
		"ba":   2,
		"ca":   999,
		"za":   16,
		"bcda": 90,
		"eda":  90,
	}
	for i, v := range expMap {
		hmap.(hashMap).HashInsert(i, v)
	}
}

func testHashMapDelete(t *testing.T, hmap interface{}) {
	expMap := map[string]int{
		"a":    1,
		"b":    2,
		"c":    999,
		"z":    16,
		"bcd":  90,
		"ed":   90,
		"ab":   1,
		"ba":   2,
		"ca":   999,
		"za":   16,
		"bcda": 90,
		"eda":  90,
	}
	for i, v := range expMap {
		hmap.(hashMap).HashInsert(i, v)
		hmap.(hashMap).HashDelete(i)
		value, exist := hmap.(hashMap).HashGet(i)
		if value != nil {
			t.Log(fmt.Sprintf("expect value of key %s should be nil", i) + fmt.Sprintf("but get: %s", value))
			t.Fail()
		}
		if exist {
			t.Log(fmt.Sprintf("expect key %s should not exist", i))
			t.Fail()
		}
	}
}

func benchmarkHashMapInsert(b *testing.B, hmap hashMap) {
	for i := 0; i < b.N; i++ {
		hmap.HashInsert(rand.Intn(128), i)
	}
}

func benchmarkHashMapInsertDelete(b *testing.B, hmap hashMap) {
	b.StopTimer()
	for i := 0; i < 128; i++ {
		hmap.HashInsert(rand.Intn(128), i)
	}
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		hmap.HashDelete(rand.Intn(128))
		b.StopTimer()
		hmap.HashInsert(rand.Intn(128), i)
	}
}

func benchmarkHashMapGet(b *testing.B, hmap hashMap) {
	b.StopTimer()
	for i := 0; i < 128; i++ {
		hmap.HashInsert(rand.Intn(128), i)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		hmap.HashGet(rand.Intn(128))
	}
}
