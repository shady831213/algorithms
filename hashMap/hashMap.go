package hashMap
type HashMap interface {
	HashInsert(interface{},interface{})
	HashDelete(interface{})
	HashGet(interface{})(interface{})
}