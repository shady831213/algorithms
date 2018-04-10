package hashMap

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func BasicTestHashMap(t *testing.T, hmap HashMap) {
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

func TestHashMapResize(t *testing.T, hmap interface{}) {
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
		hmap.(HashMap).HashInsert(i, v)
	}
}

func TestHashMapDelete(t *testing.T, hmap interface{}) {
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
		hmap.(HashMap).HashInsert(i, v)
		hmap.(HashMap).HashDelete(i)
		value, exist := hmap.(HashMap).HashGet(i)
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

func BenchmarkHashMapInsert(b *testing.B, hmap HashMap) {
	for i := 0; i < b.N; i++ {
		hmap.HashInsert(rand.Intn(128), i)
	}
}

func BenchmarkHashMapInsertDelete(b *testing.B, hmap HashMap) {
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

func BenchmarkHashMapGet(b *testing.B, hmap HashMap) {
	b.StopTimer()
	for i := 0; i < 128; i++ {
		hmap.HashInsert(rand.Intn(128), i)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		hmap.HashGet(rand.Intn(128))
	}
}
