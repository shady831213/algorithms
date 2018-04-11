package bTree

import (
	"container/list"
	"fmt"
)

type access interface {
	write(interface{}, interface{})
	read(interface{}) (interface{}, bool)
}

type memory interface {
	access
	encIdx(interface{}) interface{}
	decIdx(interface{}) interface{}
	encData(interface{}) interface{}
	decData(interface{}) interface{}
}

type discModel struct {
	mem map[interface{}][]byte // (id, []byte)
	access
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
	downStreamModel access
	memory
}

func (m *cacheModel) init(size int, downStreamModel access, self memory) *cacheModel {
	m.mem = make(map[interface{}]*list.Element)
	m.pq = list.New()
	for i := 0; i < size; i++ {
		cacheItem := new(cacheItem)
		cacheItem.dirty = false
		m.pq.PushBack(cacheItem)
	}
	m.downStreamModel = downStreamModel
	m.memory = self
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
	sID, sData := m.memory.encIdx(id), m.memory.encData(data)
	item := m.update(sID)
	item.Value.(*cacheItem).value = sData
	item.Value.(*cacheItem).dirty = true
}

func (m *cacheModel) read(id interface{}) (interface{}, bool) {
	sID := m.memory.encIdx(id)
	item := m.update(sID)
	var ok bool
	if item.Value.(*cacheItem).value, ok = m.downStreamModel.read(sID); !ok {
		panic(fmt.Sprintf("%v is not initialized in downStreamModel", sID))
	}
	item.Value.(*cacheItem).dirty = false
	return m.memory.decData(item.Value.(*cacheItem).value), ok
}
