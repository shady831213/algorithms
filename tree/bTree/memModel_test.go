package bTree

import (
	"testing"
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
)

func TestDiscModel(t *testing.T) {
	d := newDisc()
	for i := 0; i < 8; i ++ {
		buf := new(bytes.Buffer)
		binary.Write(buf, binary.LittleEndian, uint32(i))
		d.write(i, buf.Bytes())
	}
	for i := 0; i < 8; i ++ {
		data, _ := d.read(i)
		buf := bytes.NewBuffer(data.([]byte))
		var data_i uint32
		binary.Read(buf, binary.LittleEndian, &data_i)
		if uint32(i) != data_i {
			t.Log(i, data_i)
			t.Fail()
		}
	}
}

type testCacheModel struct {
	cacheModel
}

func (c *testCacheModel) init(size int, downStreamModel memory) (*testCacheModel) {
	c.cacheModel.init(size, downStreamModel, c)
	return c
}

func (c *testCacheModel) encIdx(id interface{}) (interface{}) {
	return int(id.(uint32))
}

func (c *testCacheModel) decIdx(id interface{}) (interface{}) {
	return uint32(id.(int))
}

func (c *testCacheModel) encData(data interface{}) (interface{}) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	enc.Encode(data)
	return buf.Bytes()
}

func (c *testCacheModel) decData(data interface{}) (interface{}) {
	var result string
	buf := bytes.NewBuffer(data.([]byte))
	dec := gob.NewDecoder(buf)
	dec.Decode(&result)
	return result
}

func newTestCacheModel(size int, downStreamModel memory) *testCacheModel {
	return new(testCacheModel).init(size, downStreamModel)
}

func TestCacheModel(t *testing.T) {
	d := newDisc()
	c := newTestCacheModel(4, d)
	for i := 0; i < 8; i ++ {
		c.write(uint32(i), string(i))
	}
	//check data
	for i := 0; i < 8; i ++ {
		data, _ := c.read(uint32(i))
		if string(i) != data {
			t.Log(i, data)
			t.Fail()
		}
	}
	//check pq status
	expOrder := []int{7,6,5,4}
	for item, i:= c.pq.Front(), 0; item != nil; item = item.Next() {
		if item.Value.(*cacheItem).key != expOrder[i] {
			t.Log(i, fmt.Sprintf("%+v", item.Value))
			t.Fail()
		}
		i++
	}
}
