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
		if m.value != _value {
			m.value = _value
		}
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

func (e *rsVEBTreeElement) IsEmpty() bool {
	return e.min == nil
}

func (e *rsVEBTreeElement) Min() *rsVEBTreeItem {
	return copyRsVEBTreeItem(e.min)
}

func (e *rsVEBTreeElement) Max() *rsVEBTreeItem {
	return copyRsVEBTreeItem(e.max)
}

func (e *rsVEBTreeElement) Member(key interface{}) *rsVEBTreeItem {
	if key == e.min.key {
		return copyRsVEBTreeItem(e.min)
	} else if key == e.max.key {
		return copyRsVEBTreeItem(e.max)
	} else if e.lgu == 1 {
		return nil
	} else {
		if m := e.cluster[e.mixin.High(e.lgu, key)].Member(e.mixin.Low(e.lgu, key)); m != nil {
			m.key = e.mixin.Index(e.lgu, e.mixin.High(e.lgu, key), m.key)
			return m
		}
		return nil
	}

}

func (e *rsVEBTreeElement) Successor(key interface{}) *rsVEBTreeItem {
	if e.lgu == 1 {
		if key == e.min.key && e.max != nil {
			return copyRsVEBTreeItem(e.max)
		} else {
			return nil
		}
	} else if e.min != nil && e.mixin.Less(e.lgu, key, e.min.key) {
		return copyRsVEBTreeItem(e.min)
	} else {
		if maxLow := e.cluster[e.mixin.High(e.lgu, key)].Max(); maxLow != nil && e.mixin.Less(e.lgu, e.mixin.Low(e.lgu, key), maxLow.key) {
			successor := e.cluster[e.mixin.High(e.lgu, key)].Successor(e.mixin.Low(e.lgu, key))
			successor.key = e.mixin.Index(e.lgu, e.mixin.High(e.lgu, key), successor.key)
			return successor
		} else {
			if clusterItem := e.summary.Successor(e.mixin.High(e.lgu, key)); clusterItem != nil {
				successor := e.cluster[clusterItem.key].Min()
				successor.key = e.mixin.Index(e.lgu, clusterItem.key, successor.key)
				return successor
			}
			return nil
		}
	}
}

func (e *rsVEBTreeElement) Predecessor(key interface{}) *rsVEBTreeItem {
	if e.lgu == 1 {
		if key == e.max.key && key != e.min.key {
			return copyRsVEBTreeItem(e.min)
		} else {
			return nil
		}
	} else if e.max != nil && e.mixin.Less(e.lgu, e.max.key, key) {
		return copyRsVEBTreeItem(e.max)
	} else {
		if minLow := e.cluster[e.mixin.High(e.lgu, key)].Min(); minLow != nil && e.mixin.Less(e.lgu, minLow.key, e.mixin.Low(e.lgu, key)) {
			predecessor := e.cluster[e.mixin.High(e.lgu, key)].Predecessor(e.mixin.Low(e.lgu, key))
			predecessor.key = e.mixin.Index(e.lgu, e.mixin.High(e.lgu, key), predecessor.key)
			return predecessor
		} else {
			if clusterItem := e.summary.Predecessor(e.mixin.High(e.lgu, key)); clusterItem != nil {
				predecessor := e.cluster[clusterItem.key].Max()
				predecessor.key = e.mixin.Index(e.lgu, clusterItem.key, predecessor.key)
				return predecessor
			}
			return nil
		}
	}
}

func (e *rsVEBTreeElement) Delete(key, value interface{}) {
	_key, _value := key, value
	if e.min == e.max {
		if e.min.removeByValue(_value) == 0 {
			e.min = nil
			e.max = nil
		}
	} else if e.lgu == 1 {
		if _key == e.max.key && e.max.removeByValue(_value) == 0 {
			e.max = e.min
		} else if _key == e.min.key && e.min.removeByValue(_value) == 0 {
			e.min = e.max
		}
	} else {
		if _key == e.min.key && e.min.removeByValue(_value) == 0 {
			cluster := e.summary.Min()
			e.min = e.cluster[cluster.key].Min()
			e.min.key = e.mixin.Index(e.lgu, cluster.key, e.min.key)
			_key, _value = e.min.key, nil
		}

		e.cluster[e.mixin.High(e.lgu, _key)].Delete(e.mixin.Low(e.lgu, _key), _value)
		e.summary.Delete(e.mixin.High(e.lgu, _key), _value)
		if e.cluster[e.mixin.High(e.lgu, _key)].Min() == nil {
			delete(e.cluster, e.mixin.High(e.lgu, _key))
			if _key == e.max.key && e.max.removeByValue(_value) == 0 {
				if summaryMax := e.summary.Max(); summaryMax == nil {
					e.max = e.min
				} else {
					e.max = e.cluster[summaryMax.key].Max()
					e.max.key = e.mixin.Index(e.lgu, summaryMax.key, e.max.key)
				}
			}
		} else if _key == e.max.key && e.max.removeByValue(_value) == 0 {
			e.max = e.cluster[e.mixin.High(e.lgu, _key)].Max()
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
			if cluster, ok := e.cluster[e.mixin.High(e.lgu, _key)]; !ok {
				cluster = e.addCluster(e.mixin.High(e.lgu, _key))
				e.summary.Insert(e.mixin.High(e.lgu, _key), _value)
				cluster.Insert(e.mixin.Low(e.lgu, _key), _value)
			} else {
				cluster.Insert(e.mixin.Low(e.lgu, _key), _value)
			}
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
