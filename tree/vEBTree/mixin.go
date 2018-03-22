package vEBTree

type rsVEBTreeUInt32Mixin struct {
	rsVEBTreeMixin
}

func (m *rsVEBTreeUInt32Mixin) Less(k1, k2 interface{}) bool {
	return k1.(uint32) < k2.(uint32)
}

func (m *rsVEBTreeUInt32Mixin) High(u int, key interface{}) interface{} {
	return key.(uint32) >> uint32(u>>1)
}

func (m *rsVEBTreeUInt32Mixin) Low(u int, key interface{}) interface{} {
	return key.(uint32) & ((1<<uint32(u>>1)) - 1)
}

func (m *rsVEBTreeUInt32Mixin) Index(u int, high,low interface{}) interface{} {
	return (high.(uint32) << uint32(u>>1)) | m.Low(u, low).(uint32)
}