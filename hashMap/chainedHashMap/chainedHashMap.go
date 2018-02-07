package chainedHashMap

import (
	"container/list"
	"algorithms/hashMap"
	"crypto/sha1"
	"bytes"
	"encoding/gob"
	"encoding/binary"
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
		oldBackets := h.backets
		h.Init(h.Cap + hashMap.DEFALUTCAP)
		for _, list := range oldBackets {
			if list != nil {
				for e := list.Front();e != nil; e = e.Next() {
					h.HashInsert(e.Value.(ChainedHashElement).key, e.Value.(ChainedHashElement).value)
				}
			}
		}
	}
}

func (h *ChainedHashMap) hash(key interface{})(uint32) {
	hash := sha1.New()
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	enc.Encode(key)
	hashBytes := hash.Sum(buf.Bytes())
	hashValue := binary.LittleEndian.Uint32(hashBytes)
	return hashValue%h.Cap
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
	hashKey := h.hash(key)
	if h.backets[hashKey] == nil{
		h.backets[hashKey] = list.New()
	}
	e := ChainedHashElement{key:key, value: value}
	le, exist := h.existInList(key, h.backets[hashKey])
	if exist {
		le.Value = e
	} else {
		h.backets[hashKey].PushFront(e)
		h.Count++
		h.resize()
	}
}

func (h *ChainedHashMap) HashGet(key interface{})(interface{}, bool) {
	hashKey := h.hash(key)
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
	hashKey := h.hash(key)
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