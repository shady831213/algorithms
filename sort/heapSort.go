/*
build heap :
max node num in hight h is n/2^(h+1), max h is lgn,
T(n) = n*sum(0/2+1/2^1+...+lgn/2^(lgn+1))= (n/2)*sum(2/2^1+...+lgn/2^lgn)
=(n/2)*(2-((lgn+2)/2^(lgn))) <= (n/2)*2=n
d heap:
n*(d-1)/d^(h+1), max h is logd((d-1)n), T(n) = (n*(d-1)/d) *sum(1/d^1+2/d^2+...logd((d-1)n)/d^logd(d-1)n)
= (n*(d-1)/d) <= O(n)
*/

package sort

import (
	"reflect"
	"math"
	"errors"
)

/*
use physical heap
*/
type heapNode struct {
	p     *heapNode
	l     *heapNode
	r     *heapNode
	value reflect.Value
}

func (n *heapNode) getValue() (interface{}) {
	return n.value.Interface()
}

func (n *heapNode) setValue(i interface{}) {
	n.value = reflect.ValueOf(i)
}

func (h *heapNode) swapValue(node *heapNode) {
	tempValue := node.getValue()
	node.setValue(h.getValue())
	h.setValue(tempValue)
}

func (n *heapNode) getIntValue() (int) {
	return n.getValue().(int)
}

type heap struct {
	nodes []*heapNode
}

func (h *heap) minHeaplify(node *heapNode) {
	smallestNode := node
	leftNode, rightNode := node.l, node.r
	if leftNode != nil && leftNode.getIntValue() < smallestNode.getIntValue() {
		smallestNode = leftNode
	}
	if rightNode != nil && rightNode.getIntValue() < smallestNode.getIntValue() {
		smallestNode = rightNode
	}
	if smallestNode != node {
		node.swapValue(smallestNode)
		h.minHeaplify(smallestNode)
	}
}

func (h *heap) getHight() (int) {
	return int(math.Log2(float64(h.getSize())))
}

func (h *heap) getSize() (int) {
	return len(h.nodes)
}

func (h *heap) pop() (v interface{}) {
	if h.getSize() == 0 {
		panic(errors.New("underflow!"))
	}
	v = h.nodes[0].getValue()

	h.nodes[0].swapValue(h.nodes[h.getSize()-1])
	if h.nodes[h.getSize()-1].p != nil && h.nodes[h.getSize()-1].p.l == h.nodes[h.getSize()-1] {
		h.nodes[h.getSize()-1].p.l = nil
	}
	if h.nodes[h.getSize()-1].p != nil && h.nodes[h.getSize()-1].p.r == h.nodes[h.getSize()-1] {
		h.nodes[h.getSize()-1].p.r = nil
	}
	h.nodes = h.nodes[:h.getSize()-1]
	if h.getSize() != 0 {
		h.minHeaplify(h.nodes[0])
	}
	return
}

func (h *heap) buildHeap(arr []int) {
	h.nodes = make([]*heapNode, len(arr), len(arr))
	for i := range arr {
		h.nodes[i] = new(heapNode)
		h.nodes[i].setValue(arr[i])
	}
	for i := h.getSize()/2 - 1; i >= 0; i-- {
		h.nodes[i].l = h.nodes[2*i+1]
		h.nodes[i].l.p = h.nodes[i]
		if h.getSize()%2 == 1 {
			h.nodes[i].r = h.nodes[2*i+2]
			h.nodes[i].r.p = h.nodes[i]
		}
		h.minHeaplify(h.nodes[i])
	}
}

func heapSort(arr []int) {
	heap := new(heap)
	heap.buildHeap(arr)
	for i := range arr {
		arr[i] = heap.pop().(int)
	}
}

/*
use virtual heap
*/

type heapIntArray []int
func (h heapIntArray) parent(i int)(int) {
	return i>>1
}
func (h heapIntArray) left(i int)(int) {
	return (i<<1)+1
}
func (h heapIntArray) right(i int)(int) {
	return (i<<1)+2
}

func (h heapIntArray) maxHeaplify(i int) {
	largest, largest_idx := h[i], i
	if h.left(i) < len(h) && h[h.left(i)] > largest {
		largest,largest_idx = h[h.left(i)],h.left(i)
	}
	if h.right(i) < len(h) && h[h.right(i)] > largest {
		largest,largest_idx = h[h.right(i)],h.right(i)
	}
	if i != largest_idx {
		h[largest_idx], h[i] = h[i], h[largest_idx]
		h.maxHeaplify(largest_idx)
	}
}

func (h heapIntArray) buildHeap() {
	for i := (len(h)>>1) - 1; i >= 0; i-- {
		h.maxHeaplify(i)
	}
}

func heapSort2(arr []int) {
	heap := heapIntArray(arr)
	heap.buildHeap()
	for i := len(heap) - 1;i>0;i-- {
		heap[0], heap[i] = heap[i], heap[0]
		heap = heap[:i]
		heap.maxHeaplify(0)
	}
}