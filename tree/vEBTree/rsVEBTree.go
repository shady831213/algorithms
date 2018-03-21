package vEBTree

import (
	"container/list"
)

type rsVEBTreeMinMax struct {
	key   interface{}
	value *list.List
}

func (m *rsVEBTreeMinMax) init(key, value interface{}) *rsVEBTreeMinMax {
	m.key = key
	m.value = list.New()
	m.value.PushBack(value)
	return m
}

func newRsVEBTreeMinMax(key, value interface{}) *rsVEBTreeMinMax {
	return new(rsVEBTreeMinMax).init(key, value)
}

type rsVEBTreeElement struct {
	u       int
	min     *rsVEBTreeMinMax
	max     *rsVEBTreeMinMax
	summary *rsVEBTreeElement
	cluster map[interface{}]*rsVEBTreeElement
	Less    func(interface{}, interface{}) bool
}

func (e *rsVEBTreeElement) init(u int, Less func(interface{}, interface{}) bool) *rsVEBTreeElement {
	e.u = u
	e.Less = Less
	e.cluster = make(map[interface{}]*rsVEBTreeElement)
	return e
}

func (e *rsVEBTreeElement) insertEmptyTree(key, value interface{}) {
	e.min = newRsVEBTreeMinMax(key, value)
	e.max = e.min
}

func (e *rsVEBTreeElement) insert(key, value interface{}) {
	if e.min == nil {
		e.insertEmptyTree(key, value)
	} else if key == e.min.key {
		e.min.value.PushBack(value)
	} else {
		_key, _value := key, value
		if e.Less(_key, e.min.key) {
			_key, _value, e.min = e.min.key, e.min.value, newRsVEBTreeMinMax(_key,_value)
		}

		if e.u > 2 {
		}

		if _key == e.max.key {
			e.max.value.PushBack(_value)
		} else if e.Less(e.max.key, _key) {
			e.max = newRsVEBTreeMinMax(_key,_value)
		}

	}
}

type rsVEBTreeIf interface {
	createElement() *rsVEBTreeElement
}

type rsVEBTree struct {
	root rsVEBTreeElement
	rsVEBTreeIf
}

func (t *rsVEBTree) Init(self rsVEBTreeIf, u int) *rsVEBTree {
	t.rsVEBTreeIf = self
	if u < 2 {
		panic("u must be more than 2!")
	}
	t.root = *t.rsVEBTreeIf.createElement()
	return t
}
