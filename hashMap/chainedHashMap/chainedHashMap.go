package chainedHashMap

import (
	"container/list"
	"algorithms/hashMap"
	"crypto/sha1"
	"math/big"
)

type ChainedHashMap struct {
	hashMap.HashMapBase
	backets []*list.List
}

func (h *ChainedHashMap) Init (cap uint32) {
	h.HashMapBase.Init(cap)
	h.backets = make([]*list.List, h.Cap, h.Cap)
}

func (h *ChainedHashMap) resize () {
	if h.GetAlpha() >= 0.75 {
		oldBackets := h.backets
		h.Init(h.Cap + hashMap.DEFALUTCAP)
		for _, list := range oldBackets {
			if list != nil {
				for e := list.Front();e != nil; e = e.Next() {
					h.HashInsert(e.Value.(hashMap.HashElement).Key, e.Value.(hashMap.HashElement).Value)
				}
			}
		}
	}
}

func (h *ChainedHashMap) hash(key interface{})(uint32) {
	hashValue :=  h.HashFunc(key,sha1.New())
	mb:=big.NewInt(int64(h.Cap))
	hashValue.Mod(hashValue,mb)
	return uint32(hashValue.Uint64())
}

func (h *ChainedHashMap) existInList(key interface{}, list *list.List)(*list.Element, bool) {
	for e := list.Front();e != nil; e = e.Next() {
		if e.Value.(hashMap.HashElement).Key == key {
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
	e := hashMap.HashElement{Key:key, Value: value}
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
		return le.Value.(hashMap.HashElement).Value, true
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