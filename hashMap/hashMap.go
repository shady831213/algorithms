package hashMap

import (
	"hash"
	"bytes"
	"encoding/gob"
	"math/big"
)

const DEFALUTCAP  = 4

type HashElement struct {
	Key interface{}
	Value interface{}
}

type HashMap interface {
	HashInsert(interface{},interface{})
	HashDelete(interface{})
	HashGet(interface{})(interface{},bool)
	Init(uint32)
}

type ScaleableMap interface {
	UpScale()
	DownScale()
	Move(uint32)
}

type HashMapBase struct {
	Cap uint32
	Count uint32
	HashMap
	ScaleableMap
}

func (h *HashMapBase) Init (cap uint32) {
	h.Cap = cap
	h.Count = 0
}

func (h *HashMapBase) GetAlpha ()(float64) {
	if h.Cap == 0 {
		return 1.0
	}
	return float64(h.Count)/float64(h.Cap)
}

func (h *HashMapBase) HashFunc(key interface{}, hash hash.Hash) (*big.Int) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	enc.Encode(key)
	hashBytes := hash.Sum(buf.Bytes())
	return new(big.Int).SetBytes(hashBytes)
}

func (h *HashMapBase) UpScale() {
	if h.GetAlpha() >= 0.75 {
		if h.Cap == 0 {
			h.HashMap.Init(DEFALUTCAP)
		} else {
			h.Move(h.Cap << 1)
		}
	}
}

func (h *HashMapBase) DownScale() {
	if h.GetAlpha() <= 0.25 {
		if h.Count == 0 {
			h.HashMap.Init(0)
			return
		}
		if h.Cap > DEFALUTCAP {
			h.Move(h.Cap >> 1)
		}
	}
}