package hashMap

import (
	"crypto/sha1"
	"bytes"
	"encoding/gob"
	"encoding/binary"
)

const DEFALUTCAP  = 256

type HashMap interface {
	HashInsert(interface{},interface{})
	HashDelete(interface{})
	HashGet(interface{})(interface{},bool)
	Init(uint32)
}

type HashMapBase struct {
	Cap uint32
	Count uint32
}

func (h *HashMapBase) Init (cap uint32) {
	if cap == 0 {
		h.Cap = DEFALUTCAP
	} else {
		h.Cap = cap
	}
	h.Count = 0
}

func (h *HashMapBase) GetAlpha ()(float64) {
	return float64(h.Count)/float64(h.Cap)
}

func (h *HashMapBase) Hash (key interface{})(uint32) {
	hash := sha1.New()
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	enc.Encode(key)
	hashBytes := hash.Sum(buf.Bytes())
	hashValue := binary.LittleEndian.Uint32(hashBytes)
	return hashValue%h.Cap
}