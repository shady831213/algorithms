package hashMap


const DEFALUTCAP  = 256

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
