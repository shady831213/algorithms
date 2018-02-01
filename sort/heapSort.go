package sort

import (
	"reflect"
	"math"
	"errors"
)

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
	leftNode, rightNode := node.l, node.r
	if leftNode != nil && leftNode.getIntValue() < node.getIntValue() {
		node.swapValue(leftNode)
		h.minHeaplify(leftNode)
	}
	if rightNode != nil && rightNode.getIntValue() < node.getIntValue() {
		node.swapValue(rightNode)
		h.minHeaplify(rightNode)
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
