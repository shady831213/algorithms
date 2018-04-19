# Reduced-Space van Emde Boas Tree
------------

[https://en.wikipedia.org/wiki/Van_Emde_Boas_tree](https://en.wikipedia.org/wiki/Van_Emde_Boas_tree)

------------

Support single key multi value.Lazy hashtable is used to instead of array to reduce space complexity.

use log2(upper) instead upper.

[Code](https://github.com/shady831213/algorithms/blob/master/tree/vEBTree/rsVEBTree.go)

[Test](https://github.com/shady831213/algorithms/blob/master/tree/vEBTree/rsVEBTree_test.go)

# Go Mixin design pattern
```go
//define mixin interface
type rsVEBTreeMixin interface {
	Less(int, interface{}, interface{}) bool
	High(int, interface{}) interface{}
	Low(int, interface{}) interface{}
	Index(int, interface{}, interface{}) interface{}
}
//struct calling mixin method
type rsVEBTreeElement struct {
	//...
	mixin                       rsVEBTreeMixin
}

func (e *rsVEBTreeElement) init(lgu int, mixin rsVEBTreeMixin) *rsVEBTreeElement {
  //...
	e.mixin = mixin
	return e
}
//mixin struct
type rsVEBTreeUInt32Mixin struct {
	rsVEBTreeMixin
}

func (m *rsVEBTreeUInt32Mixin) Less(lgu int, k1, k2 interface{}) bool {
	//...
}

func (m *rsVEBTreeUInt32Mixin) High(lgu int, key interface{}) interface{} {
	//...
}

func (m *rsVEBTreeUInt32Mixin) Low(lgu int, key interface{}) interface{} {
	//...
}

func (m *rsVEBTreeUInt32Mixin) Index(lgu int, high, low interface{}) interface{} {
	//...
}

//composition
func NewRsVEBTreeUint32(lgu int) *rsVEBTreeElement {
  //...
	return new(rsVEBTreeElement).init(lgu, new(rsVEBTreeUInt32Mixin))
}

```
