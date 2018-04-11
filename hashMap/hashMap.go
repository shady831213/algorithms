package hashMap

import (
	"bytes"
	"encoding/gob"
	"hash"
	"math/big"
)

const defalutcap = 8

type hashElement struct {
	Key   interface{}
	Value interface{}
}

type hashMap interface {
	HashInsert(interface{}, interface{})
	HashDelete(interface{})
	HashGet(interface{}) (interface{}, bool)
	Init(uint32)
}

type scaleableMap interface {
	UpScale()
	DownScale()
	Move(uint32)
}

type hashMapBase struct {
	Cap   uint32
	Count uint32
	hashMap
	scaleableMap
}

func (h *hashMapBase) Init(cap uint32) {
	h.Cap = cap
	h.Count = 0
}

func (h *hashMapBase) GetAlpha() float64 {
	if h.Cap == 0 {
		return 1.0
	}
	return float64(h.Count) / float64(h.Cap)
}

func (h *hashMapBase) HashFunc(key interface{}, hash hash.Hash) *big.Int {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	enc.Encode(key)
	hashBytes := hash.Sum(buf.Bytes())
	return new(big.Int).SetBytes(hashBytes)
}

func (h *hashMapBase) UpScale() {
	if h.GetAlpha() >= 0.75 {
		if h.Cap == 0 {
			h.hashMap.Init(defalutcap)
		} else {
			h.Move(h.Cap << 1)
		}
	}
}

func (h *hashMapBase) DownScale() {
	if h.GetAlpha() <= 0.125 {
		if h.Count == 0 {
			h.hashMap.Init(0)
			return
		}
		if h.Cap > defalutcap {
			h.Move(h.Cap >> 1)
		}
	}
}
