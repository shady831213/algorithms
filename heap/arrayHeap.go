package heap

type heapIntArrays []int

//IntArray cross package
type IntArray struct {
	heap
	heapIntArrays
}

func (h *heapIntArrays) Parent(i interface{}) interface{} {
	return i.(int) >> 1
}
func (h *heapIntArrays) Left(i interface{}) interface{} {
	return (i.(int) << 1) + 1
}
func (h *heapIntArrays) Right(i interface{}) interface{} {
	return (i.(int) << 1) + 2
}

func (h *heapIntArrays) Prev(i interface{}) interface{} {
	return i.(int) - 1
}
func (h *heapIntArrays) Next(i interface{}) interface{} {
	return i.(int) + 1
}

func (h *heapIntArrays) Last() interface{} {
	return len(*h) - 1
}
func (h *heapIntArrays) Head() interface{} {
	return 0
}
func (h *heapIntArrays) Valid(i interface{}) bool {
	return i.(int) >= 0 && i.(int) < len(*h)
}

func (h *heapIntArrays) Swap(i interface{}, j interface{}) {
	(*h)[i.(int)], (*h)[j.(int)] = (*h)[j.(int)], (*h)[i.(int)]
}

func (h *heapIntArrays) Key(i interface{}) int {
	return (*h)[i.(int)]
}

func (h *heapIntArrays) Value(i interface{}) interface{} {
	return (*h)[i.(int)]
}

func (h *heapIntArrays) Len() int {
	return len(*h)
}

func (h *heapIntArrays) Pop() (i interface{}) {
	(*h), i = (*h)[:len(*h)-1], (*h)[len(*h)-1]
	return
}

func (h *heapIntArrays) Append(i interface{}) {
	(*h) = append((*h), i.(int))
}

func (h *heapIntArrays) Union(i interface{}) interface{} {
	_i := i.(*heapIntArrays)
	(*h) = append((*h), (*_i)...)
	return h
}

//merge:O(n)
//rebuild:O(n)
//T(n)=O(n)

//Union cross package
func (h *IntArray) Union(i interface{}) interface{} {
	h.heapIntArrays = h.heapIntArrays.Union(&(i.(*IntArray).heapIntArrays)).(heapIntArrays)
	h.heap.BuildHeap()
	return h
}

//Pop cross package
func (h *IntArray) Pop() (i interface{}) {
	return h.heap.Pop()
}

//Append cross package
func (h *IntArray) Append(i interface{}) {
	h.heap.Append(i)
}

//NewHeapIntArray cross package
func NewHeapIntArray(arr []int) *IntArray {
	h := new(IntArray)
	h.heapIntArrays = arr
	h.heap.binHeapArrayIf = &h.heapIntArrays
	h.BuildHeap()
	return h
}
