package chainedHashMap

import (
	"container/list"
	"algorithms/hashMap"
)

type ChainedHashElement struct {
	key interface{}
	value interface{}
}

type ChainedHashMap struct {
	hashMap.HashMapBase
	backets []*list.List
}

func (h *ChainedHashMap) Init (cap uint32) {
	h.HashMapBase.Init(cap)
	h.backets = make([]*list.List, cap, cap)
}

func (h *ChainedHashMap) resize () {
	if h.GetAlpha() >= 0.75 {
		h.backets = append(h.backets, make([]*list.List, h.Cap, h.Cap)...)
		h.Cap = (h.Cap << 1)
	}
}

func (h *ChainedHashMap) existInList(key interface{}, list *list.List)(*list.Element, bool) {
	for e := list.Front();e != nil; e = e.Next() {
		if e.Value.(ChainedHashElement).key == key {
			return e, true
		}
	}
	return nil, false
}

func (h *ChainedHashMap) HashInsert(key interface{},value interface{}) {
	hashKey := h.Hash(key)
	if h.backets[hashKey] == nil{
		h.backets[hashKey] = list.New()
		h.Count++
		h.resize()
	}
	e := ChainedHashElement{key:key, value: value}
	le, exist := h.existInList(key, h.backets[hashKey])
	if exist {
		le.Value = e
	} else {
		h.backets[hashKey].PushFront(e)
	}
}

func (h *ChainedHashMap) HashGet(key interface{})(interface{}, bool) {
	hashKey := h.Hash(key)
	if h.backets[hashKey] == nil {
		return nil,false
	}
	le, exist := h.existInList(key, h.backets[hashKey])
	if exist {
		return le.Value.(ChainedHashElement).value, true
	}
	return nil,false
}

func (h *ChainedHashMap) HashDelete(key interface{}) {
	hashKey := h.Hash(key)
	if h.backets[hashKey] == nil {
		return
	}
	le, exist := h.existInList(key, h.backets[hashKey])
	if exist {
		h.backets[hashKey].Remove(le)
	}
	if h.backets[hashKey].Len() == 0 {
		h.backets[hashKey] = nil
		h.Count--
	}
}

func New(cap uint32)(h hashMap.HashMap) {
	h = new(ChainedHashMap)
	h.Init(cap)
	return
}