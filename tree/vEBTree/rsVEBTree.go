package vEBTree

import "container/list"

type rsVEBTreeElementIf interface {
	getCluster(interface{}) (interface{}, bool)
}

type rsVEBTreeElement struct {
	u        int
	summary  *rsVEBTreeElement
	min, max interface{}
	Less     func(interface{}, interface{}) bool
	rsVEBTreeElementIf
}

func (e *rsVEBTreeElement) init(self rsVEBTreeElementIf, u int, Less func(interface{}, interface{}) bool) *rsVEBTreeElement {
	e.rsVEBTreeElementIf = self
	e.u = u
	e.Less = Less
	return e
}

func (e *rsVEBTreeElement) getCluster (key interface{}) (interface{}, bool) {
	return e.rsVEBTreeElementIf.getCluster(key)
}

type rsVEBTreeLeafElement struct {
	cluster map[interface{}]list.List
	rsVEBTreeElement
}

func (e *rsVEBTreeLeafElement) init(Less func(interface{}, interface{}) bool) *rsVEBTreeLeafElement {
	e.rsVEBTreeElement.init(e, 2, Less)
	e.cluster = make(map[interface{}]list.List)
	return e
}

func (e *rsVEBTreeLeafElement) getCluster (key interface{}) (interface{}, bool) {
	value, ok := e.cluster[key]
	return value, ok
}

type rsVEBTreeNonLeafElement struct {
	cluster map[interface{}]*rsVEBTreeElement
	rsVEBTreeElement
}

func (e *rsVEBTreeNonLeafElement) init(u int, Less func(interface{}, interface{}) bool) *rsVEBTreeNonLeafElement {
	e.rsVEBTreeElement.init(e, u, Less)
	e.cluster = make(map[interface{}]*rsVEBTreeElement)
	return e
}

func (e *rsVEBTreeNonLeafElement) getCluster (key interface{}) (interface{}, bool) {
	value, ok := e.cluster[key]
	return value, ok
}

type rsVEBTreeIf interface {
	createLeafElement() *rsVEBTreeLeafElement
	createNonLeafElement(int) *rsVEBTreeNonLeafElement
}

type rsVEBTree struct {
	root rsVEBTreeElement
	rsVEBTreeIf
}

func (t *rsVEBTree) init (self rsVEBTreeIf, u int) *rsVEBTree {
	t.rsVEBTreeIf = self
	if u < 2 {
		panic("u must be more than 2!")
	} else if u == 2 {
		t.root = t.rsVEBTreeIf.createLeafElement().rsVEBTreeElement
	} else {
		t.root = t.rsVEBTreeIf.createNonLeafElement(u).rsVEBTreeElement
	}
	return t
}