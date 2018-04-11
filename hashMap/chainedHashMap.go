package hashMap

import (
	"container/list"
	"crypto/sha1"
	"math/big"
)

type chainedHashMap struct {
	hashMapBase
	backets []*list.List
}

func (h *chainedHashMap) Init(cap uint32) {
	h.hashMapBase.Init(cap)
	if cap == 0 {
		h.backets = nil
	} else {
		h.backets = make([]*list.List, h.Cap, h.Cap)
	}
}

func (h *chainedHashMap) Move(cap uint32) {
	oldBackets := h.backets
	h.Init(cap)
	for _, list := range oldBackets {
		if list != nil {
			for e := list.Front(); e != nil; e = e.Next() {
				h.HashInsert(e.Value.(hashElement).Key, e.Value.(hashElement).Value)
			}
		}
	}
}

func (h *chainedHashMap) hash(key interface{}) uint32 {
	hashValue := h.HashFunc(key, sha1.New())
	mb := big.NewInt(int64(h.Cap))
	hashValue.Mod(hashValue, mb)
	return uint32(hashValue.Uint64())
}

func (h *chainedHashMap) existInList(key interface{}, list *list.List) (*list.Element, bool) {
	for e := list.Front(); e != nil; e = e.Next() {
		if e.Value.(hashElement).Key == key {
			return e, true
		}
	}
	return nil, false
}

func (h *chainedHashMap) HashInsert(key interface{}, value interface{}) {
	h.UpScale()
	hashKey := h.hash(key)
	if h.backets[hashKey] == nil {
		h.backets[hashKey] = list.New()
	}
	e := hashElement{Key: key, Value: value}
	le, exist := h.existInList(key, h.backets[hashKey])
	if exist {
		le.Value = e
	} else {
		h.backets[hashKey].PushFront(e)
		h.Count++
	}
}

func (h *chainedHashMap) HashGet(key interface{}) (interface{}, bool) {
	if h.Count != 0 {
		hashKey := h.hash(key)
		if h.backets[hashKey] == nil {
			return nil, false
		}
		le, exist := h.existInList(key, h.backets[hashKey])
		if exist {
			return le.Value.(hashElement).Value, true
		}
	}
	return nil, false
}

func (h *chainedHashMap) HashDelete(key interface{}) {
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
	h.DownScale()
}

func newChainedHashMap() *chainedHashMap {
	h := new(chainedHashMap)
	h.hashMapBase.hashMap = h
	h.hashMapBase.scaleableMap = h
	return h
}
