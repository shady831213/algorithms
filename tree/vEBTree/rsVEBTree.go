package vEBTree

import (
	"container/list"
)

type rsVEBTreeItem struct {
	key   interface{}
	value *list.List
}

func (m *rsVEBTreeItem) init(key, value interface{}) *rsVEBTreeItem {
	m.key = key
	m.value = list.New()
	m.addValue(value)
	return m
}

func (m *rsVEBTreeItem) getListElement(value interface{}) *list.Element {
	for e := m.value.Front(); e != nil; e = e.Next() {
		if e.Value == value {
			return e
		}
	}
	return nil
}

func (m *rsVEBTreeItem) addValue(value interface{}) {
	if _value, isList := value.(*list.List); isList {
		m.value = _value
	} else {
		m.value.PushBack(value)
	}
}

func (m *rsVEBTreeItem) removeByValue(value interface{}) int {
	if value == nil {
		m.value.Init()
	} else {
		if listE := m.getListElement(value); listE != nil {
			m.value.Remove(listE)
		}
	}
	return m.value.Len()
}

func newRsVEBTreeItem(key, value interface{}) *rsVEBTreeItem {
	return new(rsVEBTreeItem).init(key, value)
}

func copyRsVEBTreeItem(i *rsVEBTreeItem) *rsVEBTreeItem {
	if i == nil {
		return nil
	}
	newI := new(rsVEBTreeItem)
	newI.key = i.key
	newI.value = list.New()
	newI.value.PushBackList(i.value)
	return newI
}

type rsVEBTreeMixin interface {
	Less(int, interface{}, interface{}) bool
	High(int, interface{}) interface{}
	Low(int, interface{}) interface{}
	Index(int, interface{}, interface{}) interface{}
}

type rsVEBTreeElement struct {
	lgu, summaryLgu, clusterLgu int
	min                         *rsVEBTreeItem
	max                         *rsVEBTreeItem
	summary                     *rsVEBTreeElement
	cluster                     map[interface{}]*rsVEBTreeElement
	mixin                       rsVEBTreeMixin
}

func (e *rsVEBTreeElement) init(lgu int, mixin rsVEBTreeMixin) *rsVEBTreeElement {
	e.lgu = lgu
	e.mixin = mixin
	e.cluster = make(map[interface{}]*rsVEBTreeElement)
	if e.lgu > 1 {
		e.summaryLgu = (e.lgu + 1) / 2
		e.clusterLgu = e.lgu - e.summaryLgu
		e.summary = new(rsVEBTreeElement).init(e.summaryLgu, e.mixin)
	} else {
		e.summaryLgu = 0
		e.clusterLgu = 0
	}
	return e
}

func (e *rsVEBTreeElement) addCluster(key interface{}) *rsVEBTreeElement {
	if e.lgu > 1 {
		e.cluster[key] = new(rsVEBTreeElement).init(e.clusterLgu, e.mixin)
		return e.cluster[key]
	}
	return nil
}


func (e *rsVEBTreeElement) Min() (interface{}, *list.List) {
	if e.min == nil {
		return nil, nil
	}
	return e.min.key, e.min.value
}

func (e *rsVEBTreeElement) Max() (interface{}, *list.List) {
	if e.max == nil {
		return nil, nil
	}
	return e.max.key, e.max.value
}

func (e *rsVEBTreeElement) Member(key interface{}) *list.List {
	if e.min == nil {
		return nil
	} else if key == e.min.key {
		return e.min.value
	} else if key == e.max.key {
		return e.max.value
	} else if e.lgu == 1 {
		return nil
	} else {
		if cluster, ok := e.cluster[e.mixin.High(e.lgu, key)]; ok {
			return cluster.Member(e.mixin.Low(e.lgu, key))
		}
		return nil
	}

}

func (e *rsVEBTreeElement) Successor(key interface{}) (interface{}, *list.List) {
	if e.lgu == 1 {
		if key == e.min.key && key != e.max.key {
			return e.Max()
		} else {
			return nil, nil
		}
	} else if e.min != nil && e.mixin.Less(e.lgu, key, e.min.key) {
		return e.Min()
	} else {
		if maxLow, _ := e.cluster[e.mixin.High(e.lgu, key)].Max(); maxLow != nil && e.mixin.Less(e.lgu, e.mixin.Low(e.lgu, key), maxLow) {
			successorK, successorV := e.cluster[e.mixin.High(e.lgu, key)].Successor(e.mixin.Low(e.lgu, key))
			successorK = e.mixin.Index(e.lgu, e.mixin.High(e.lgu, key), successorK)
			return successorK, successorV
		} else {
			if summaryK, _ := e.summary.Successor(e.mixin.High(e.lgu, key)); summaryK != nil {
				successorK, successorV := e.cluster[summaryK].Min()
				successorK = e.mixin.Index(e.lgu, summaryK, successorK)
				return successorK, successorV
			}
			return nil, nil
		}
	}
}

func (e *rsVEBTreeElement) Predecessor(key interface{}) (interface{}, *list.List) {
	if e.lgu == 1 {
		if key == e.max.key && key != e.min.key {
			return e.Min()
		} else {
			return nil, nil
		}
	} else if e.max != nil && e.mixin.Less(e.lgu, e.max.key, key) {
		return e.Max()
	} else {
		if minLow, _ := e.cluster[e.mixin.High(e.lgu, key)].Min(); minLow != nil && e.mixin.Less(e.lgu, minLow, e.mixin.Low(e.lgu, key)) {
			predecessorK, predecessorV := e.cluster[e.mixin.High(e.lgu, key)].Predecessor(e.mixin.Low(e.lgu, key))
			predecessorK = e.mixin.Index(e.lgu, e.mixin.High(e.lgu, key), predecessorK)
			return predecessorK, predecessorV
		} else {
			if summaryK, _ := e.summary.Predecessor(e.mixin.High(e.lgu, key)); summaryK != nil {
				predecessorK, predecessorV := e.cluster[summaryK].Max()
				predecessorK = e.mixin.Index(e.lgu, summaryK, predecessorK)
				return predecessorK, predecessorV
			}
			return nil, nil
		}
	}
}

func (e *rsVEBTreeElement) Delete(key interface{}) {
	_key := key
	if e.min == e.max {
		e.min = nil
		e.max = nil
	} else if e.lgu == 1 {
		if _key == e.max.key {
			e.max = e.min
		} else if _key == e.min.key {
			e.min = e.max
		}
	} else {
		if _key == e.min.key {
			clusterK, _ := e.summary.Min()
			e.min.key, e.min.value = e.cluster[clusterK].Min()
			e.min.key = e.mixin.Index(e.lgu, clusterK, e.min.key)
			_key = e.min.key
		}
		e.cluster[e.mixin.High(e.lgu, _key)].Delete(e.mixin.Low(e.lgu, _key))
		if lowMin, _ := e.cluster[e.mixin.High(e.lgu, _key)].Min(); lowMin == nil {
			e.summary.Delete(e.mixin.High(e.lgu, _key))
			delete(e.cluster, e.mixin.High(e.lgu, _key))
			if _key == e.max.key {
				if summaryMax, _ := e.summary.Max(); summaryMax == nil {
					e.max = e.min
				} else {
					e.max.key, e.max.value = e.cluster[summaryMax].Max()
					e.max.key = e.mixin.Index(e.lgu, summaryMax, e.max.key)
				}
			}
		} else if _key == e.max.key {
			e.max.key, e.max.value = e.cluster[e.mixin.High(e.lgu, _key)].Max()
			e.max.key = e.mixin.Index(e.lgu, e.mixin.High(e.lgu, _key), e.max.key)
		}
	}
}

func (e *rsVEBTreeElement) insertEmptyTree(key, value interface{}) {
	e.min = newRsVEBTreeItem(key, value)
	e.max = e.min
}

func (e *rsVEBTreeElement) Insert(key, value interface{}) {
	if e.min == nil {
		e.insertEmptyTree(key, value)
	} else if key == e.min.key {
		e.min.addValue(value)
	} else {
		_key, _value := key, value
		if e.mixin.Less(e.lgu, _key, e.min.key) {
			_key, _value, e.min = e.min.key, e.min.value, newRsVEBTreeItem(_key, _value)
		}
		if e.lgu > 1 {
			if _, ok := e.cluster[e.mixin.High(e.lgu, _key)]; !ok {
				e.addCluster(e.mixin.High(e.lgu, _key))
			}
			e.summary.Insert(e.mixin.High(e.lgu, _key), _value)
			e.cluster[e.mixin.High(e.lgu, _key)].Insert(e.mixin.Low(e.lgu, _key), _value)
		}
		if _key == e.max.key {
			e.max.addValue(_value)
		} else if e.mixin.Less(e.lgu, e.max.key, _key) {
			e.max = newRsVEBTreeItem(_key, _value)
		}
	}

}

func newRsVEBTreeUint32(lgu int) *rsVEBTreeElement {
	return new(rsVEBTreeElement).init(lgu, new(rsVEBTreeUInt32Mixin))
}
