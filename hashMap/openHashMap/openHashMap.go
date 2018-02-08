package openHashMap

import (
	"algorithms/hashMap"
	"crypto/sha1"
	"crypto/sha256"
	"math/big"
)

type OpenHashElement struct {
	hashMap.HashElement
	exist bool
}

type OpenHashMap struct {
	hashMap.HashMapBase
	backets []*OpenHashElement
}

func (h *OpenHashMap) Init(cap uint32) {
	h.HashMapBase.Init(cap)
	h.backets = make([]*OpenHashElement, h.Cap, h.Cap)
}

func (h *OpenHashMap) resize() {
	if h.GetAlpha() >= 0.75 {
		oldBackets := h.backets
		h.Init(h.Cap+hashMap.DEFALUTCAP)
		for _, v := range oldBackets {
			if v!=nil {
				h.HashInsert(v.Key, v.Value)
			}
		}
	}
}

func (h *OpenHashMap) hash(key interface{}, i uint32) (uint32) {
	hashValue1, hashValue2 := h.HashFunc(key, sha1.New()), h.HashFunc(key, sha256.New())
	ib:=big.NewInt(int64(i))
	mb:=big.NewInt(int64(h.Cap))
	hashValue2.Mul(hashValue2,ib).Add(hashValue2, hashValue1).Mod(hashValue2,mb)
	return uint32(hashValue2.Uint64())
}

func (h *OpenHashMap) existKey(key uint32) (bool) {
	if h.backets[key] == nil {
		return false
	}
	return h.backets[key].exist
}

func (h *OpenHashMap) HashInsert(key interface{}, value interface{}) {
	for i := 0; i < int(h.Cap); i ++ {
		hashValue := h.hash(key, uint32(i))
		if h.backets[hashValue] == nil {
			h.backets[hashValue] = &OpenHashElement{exist: false}
		}
		exist := h.existKey(hashValue)
		if exist && h.backets[hashValue].Key == key {
			h.backets[hashValue].Value = value
			return
		} else if !exist {
			h.backets[hashValue].Key = key
			h.backets[hashValue].Value = value
			h.backets[hashValue].exist = true
			h.Count++
			h.resize()
			return
		}
	}
}

func (h *OpenHashMap) HashGet(key interface{}) (interface{}, bool) {
	for i := 0; i < int(h.Cap); i ++ {
		hashValue := h.hash(key, uint32(i))
		if h.backets[hashValue] !=nil && h.backets[hashValue].Key == key {
			return h.backets[hashValue].Value, h.backets[hashValue].exist
		}
	}
	return nil, false
}

func (h *OpenHashMap) HashDelete(key interface{}) {
	for i := 0; i < int(h.Cap); i ++ {
		hashValue := h.hash(key, uint32(i))
		if h.existKey(hashValue) && h.backets[hashValue].Key == key  {
			h.backets[hashValue] = &OpenHashElement{exist: false}
			h.Count--
			return
		}
	}
}

func New(cap uint32) (h hashMap.HashMap) {
	h = new(OpenHashMap)
	h.Init(cap)
	return
}
