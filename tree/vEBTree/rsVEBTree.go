package vEBTree

import (
	"container/list"
	"math"
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
	u,summaryU,clusterU       int
	min     *rsVEBTreeMinMax
	max     *rsVEBTreeMinMax
	summary *rsVEBTreeElement
	cluster map[interface{}]*rsVEBTreeElement
	Less    func(interface{}, interface{}) bool
	High    func(int, interface{}) interface{}
	Low     func(int, interface{}) interface{}
}

func (e *rsVEBTreeElement) init(u int,
	Less func(interface{}, interface{}) bool,
	High func(int, interface{}) interface{},
	Low func(int, interface{}) interface{}) *rsVEBTreeElement {
	e.u = u
	e.Less = Less
	e.High = High
	e.Low = Low
	e.cluster = make(map[interface{}]*rsVEBTreeElement)
	if e.u > 2 {
		e.summaryU = int(math.Ceil(math.Sqrt(float64(u))))
		e.clusterU = int(math.Floor(math.Sqrt(float64(e.u))))
		e.summary = newRsVEBTreeElement(e.summaryU, Less, High, Low)
	} else {
		e.summaryU = 0
		e.clusterU = 0
	}
	return e
}

func (e *rsVEBTreeElement) addCluster(key interface{}) {
	if e.u > 2 {
		e.cluster[key] = newRsVEBTreeElement(e.clusterU, e.Less, e.High, e.Low)
	}
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
			_key, _value, e.min = e.min.key, e.min.value, newRsVEBTreeMinMax(_key, _value)
		}

		if e.u > 2 {
		}

		if _key == e.max.key {
			e.max.value.PushBack(_value)
		} else if e.Less(e.max.key, _key) {
			e.max = newRsVEBTreeMinMax(_key, _value)
		}

	}
}

func newRsVEBTreeElement(u int,
	Less func(interface{}, interface{}) bool,
	High func(int, interface{}) interface{},
	Low func(int, interface{}) interface{}) *rsVEBTreeElement {
	return new(rsVEBTreeElement).init(u, Less, High, Low)
}
