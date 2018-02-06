package heap

type HeapIntArray []int

func (h*HeapIntArray) Parent(i interface{})(interface{}) {
	return i.(int)>>1
}
func (h*HeapIntArray) Left(i interface{})(interface{}) {
	return (i.(int)<<1)+1
}
func (h*HeapIntArray) Right(i interface{})(interface{}) {
	return (i.(int)<<1)+2
}

func (h*HeapIntArray) Prev(i interface{})(interface{}) {
	return i.(int)-1
}
func (h*HeapIntArray) Next(i interface{})(interface{}) {
	return i.(int)+1
}

func (h*HeapIntArray) Last()(interface{}) {
	return len(*h)-1
}
func (h*HeapIntArray) Head()(interface{}) {
	return 0
}
func (h*HeapIntArray) Valid(i interface{})(bool){
	return i.(int) >= 0 && i.(int) < len(*h)
}

func (h *HeapIntArray) Swap(i interface{}, j interface{}) () {
	(*h)[i.(int)], (*h)[j.(int)] = (*h)[j.(int)], (*h)[i.(int)]
}

func (h *HeapIntArray) Key(i interface{}) (int) {
	return (*h)[i.(int)]
}

func (h *HeapIntArray) Value(i interface{}) (interface{}) {
	return (*h)[i.(int)]
}

func (h *HeapIntArray) Len() (int) {
	return len(*h)
}

func (h *HeapIntArray) Pop() (i interface{}) {
	(*h), i = (*h)[:len(*h)-1], (*h)[len(*h)-1]
	return
}

func (h *HeapIntArray) Append(i interface{}) {
	(*h) = append((*h), i.(int))
}

func (h *HeapIntArray) Merge(i interface{},j interface{})(interface{}) {
	_i := i.(*HeapIntArray)
	_j := j.(*HeapIntArray)
	(*_i) = append((*_i), (*_j)...)
	return _i
}

