package bTree

import (
	"container/list"
	"fmt"
)

type memory interface {
	write(interface{}, interface{})
	read(interface{}) (interface{}, bool)
}

type Memory interface {
	memory
	encIdx(interface{}) interface{}
	decIdx(interface{}) interface{}
	encData(interface{}) interface{}
	decData(interface{}) interface{}
}

type discModel struct {
	mem map[interface{}][]byte // (id, []byte)
	memory
}

func (m *discModel) init() *discModel {
	m.mem = make(map[interface{}][]byte)
	return m
}

func (m *discModel) write(id interface{}, data interface{}) {
	m.mem[id] = data.([]byte)
}

func (m *discModel) read(id interface{}) (interface{}, bool) {
	data, ok := m.mem[id]
	return data, ok
}

func newDisc() *discModel {
	return new(discModel).init()
}

//LRU
type cacheItem struct {
	key, value interface{}
	dirty      bool
}

type cacheModel struct {
	mem             map[interface{}]*list.Element
	pq              *list.List
	downStreamModel memory
	Memory
}

func (m *cacheModel) init(size int, downStreamModel memory, self Memory) *cacheModel {
	m.mem = make(map[interface{}]*list.Element)
	m.pq = list.New()
	for i := 0; i < size; i++ {
		cacheItem := new(cacheItem)
		cacheItem.dirty = false
		m.pq.PushBack(cacheItem)
	}
	m.downStreamModel = downStreamModel
	m.Memory = self
	return m
}

func (m *cacheModel) update(id interface{}) *list.Element {
	var item *list.Element
	var ok bool
	if item, ok = m.mem[id]; !ok {
		if item = m.pq.Back(); item.Value.(*cacheItem).dirty {
			m.downStreamModel.write(item.Value.(*cacheItem).key, item.Value.(*cacheItem).value)
			item.Value.(*cacheItem).dirty = false
		}
		delete(m.mem, item.Value.(*cacheItem).key)
		m.mem[id] = item
		item.Value.(*cacheItem).key = id
	}
	m.pq.MoveToFront(item)
	return item
}

func (m *cacheModel) write(id interface{}, data interface{}) {
	sId, sData := m.Memory.encIdx(id), m.Memory.encData(data)
	item := m.update(sId)
	item.Value.(*cacheItem).value = sData
	item.Value.(*cacheItem).dirty = true
}

func (m *cacheModel) read(id interface{}) (interface{}, bool) {
	sId := m.Memory.encIdx(id)
	item := m.update(sId)
	var ok bool
	if item.Value.(*cacheItem).value, ok = m.downStreamModel.read(sId); !ok {
		panic(fmt.Sprintf("%v is not initialized in downStreamModel", sId))
	}
	item.Value.(*cacheItem).dirty = false
	return m.Memory.decData(item.Value.(*cacheItem).value), ok
}
