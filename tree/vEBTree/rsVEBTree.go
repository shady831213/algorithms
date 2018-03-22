package vEBTree

import (
	"container/list"
	"math"
)

type rsVEBTreeItem struct {
	key   interface{}
	value *list.List
}

func (m *rsVEBTreeItem) init(key, value interface{}) *rsVEBTreeItem {
	m.key = key
	m.value = list.New()
	m.value.PushBack(value)
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
	newI.value = i.value
	return newI
}

type rsVEBTreeMixin interface {
	Less(int, interface{}, interface{}) bool
	High(int, interface{}) interface{}
	Low(int, interface{}) interface{}
	Index(int, interface{}, interface{}) interface{}
}

type rsVEBTreeElement struct {
	u, summaryU, clusterU int
	min                   *rsVEBTreeItem
	max                   *rsVEBTreeItem
	summary               *rsVEBTreeElement
	cluster               map[interface{}]*rsVEBTreeElement
	mixin                 rsVEBTreeMixin
}

func (e *rsVEBTreeElement) init(u int, mixin rsVEBTreeMixin) *rsVEBTreeElement {
	e.u = u
	e.mixin = mixin
	e.cluster = make(map[interface{}]*rsVEBTreeElement)
	if e.u > 2 {
		e.summaryU = int(math.Ceil(math.Sqrt(float64(u))))
		e.clusterU = int(math.Floor(math.Sqrt(float64(e.u))))
		e.summary = new(rsVEBTreeElement).init(e.summaryU, e.mixin)
	} else {
		e.summaryU = 0
		e.clusterU = 0
	}
	return e
}

func (e *rsVEBTreeElement) addCluster(key interface{}) *rsVEBTreeElement {
	if e.u > 2 {
		e.cluster[key] = new(rsVEBTreeElement).init(e.clusterU, e.mixin)
		return e.cluster[key]
	}
	return nil
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
	} else if e.u == 2 {
		return nil
	} else {
		if m := e.cluster[e.mixin.High(e.u, key)].Member(e.mixin.Low(e.u, key)); m != nil {
			m.key = e.mixin.Index(e.u, e.mixin.High(e.u, key), m.key)
		}
		return nil
	}

}

func (e *rsVEBTreeElement) Successor(key interface{}) *rsVEBTreeItem {
	if e.u == 2 {
		if key == e.min.key && e.max != nil {
			return copyRsVEBTreeItem(e.max)
		} else {
			return nil
		}
	} else if e.min != nil && e.mixin.Less(e.u, key, e.min.key) {
		return copyRsVEBTreeItem(e.min)
	} else {
		if maxLow := e.cluster[e.mixin.High(e.u, key)].Max(); maxLow != nil && e.mixin.Less(e.u, e.mixin.Low(e.u, key), maxLow.key) {
			successor := e.cluster[e.mixin.High(e.u, key)].Successor(e.mixin.Low(e.u, key))
			successor.key = e.mixin.Index(e.u, e.mixin.High(e.u, key), successor.key)
			return successor
		} else {
			if clusterItem := e.summary.Successor(e.mixin.High(e.u, key)); clusterItem != nil {
				successor := e.cluster[clusterItem.key].Min()
				successor.key = e.mixin.Index(e.u, clusterItem.key, successor.key)
				return successor
			}
			return nil
		}
	}
}

func (e *rsVEBTreeElement) Predecessor(key interface{}) *rsVEBTreeItem {
	if e.u == 2 {
		if key == e.max.key && key != e.min.key {
			return copyRsVEBTreeItem(e.min)
		} else {
			return nil
		}
	} else if e.max != nil && e.mixin.Less(e.u, e.max.key, key) {
		return copyRsVEBTreeItem(e.max)
	} else {
		if minLow := e.cluster[e.mixin.High(e.u, key)].Min(); minLow != nil && e.mixin.Less(e.u, minLow.key, e.mixin.Low(e.u, key)) {
			predecessor := e.cluster[e.mixin.High(e.u, key)].Predecessor(e.mixin.Low(e.u, key))
			predecessor.key = e.mixin.Index(e.u, e.mixin.High(e.u, key), predecessor.key)
			return predecessor
		} else {
			if clusterItem := e.summary.Predecessor(e.mixin.High(e.u, key)); clusterItem != nil {
				predecessor := e.cluster[clusterItem.key].Max()
				predecessor.key = e.mixin.Index(e.u, clusterItem.key, predecessor.key)
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
	} else if e.u == 2 {
		if _key == e.max.key && e.max.removeByValue(_value) == 0 {
			e.max = e.min
		} else if _key == e.min.key && e.min.removeByValue(_value) == 0 {
			e.min = e.max
		}
	} else {
		if _key == e.min.key && e.min.removeByValue(_value) == 0 {
			cluster := e.summary.Min()
			e.min = e.cluster[cluster.key].Min()
			e.min.key = e.mixin.Index(e.u, cluster.key, e.min.key)
			_key, _value = e.min.key, nil
		}

		e.cluster[e.mixin.High(e.u, _key)].Delete(e.mixin.Low(e.u, _key), _value)
		e.summary.Delete(e.mixin.High(e.u, _key), _value)
		if e.cluster[e.mixin.High(e.u, _key)].Min() == nil {
			delete(e.cluster, e.mixin.High(e.u, _key))
			if _key == e.max.key && e.max.removeByValue(_value) == 0 {
				if summaryMax := e.summary.Max(); summaryMax == nil {
					e.max = e.min
				} else {
					e.max = e.cluster[summaryMax.key].Max()
					e.max.key = e.mixin.Index(e.u, summaryMax.key, e.max.key)
				}
			}
		} else if _key == e.max.key && e.max.removeByValue(_value) == 0 {
			e.max = e.cluster[e.mixin.High(e.u, _key)].Max()
			e.max.key = e.mixin.Index(e.u, e.mixin.High(e.u, _key), e.max.key)
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
		e.min.value.PushBack(value)
	} else {
		_key, _value := key, value
		if e.mixin.Less(e.u, _key, e.min.key) {
			_key, _value, e.min = e.min.key, e.min.value, newRsVEBTreeItem(_key, _value)
		}

		if e.u > 2 {
			if cluster, ok := e.cluster[e.mixin.High(e.u, key)]; !ok {
				cluster = e.addCluster(e.mixin.High(e.u, key))
				e.summary.Insert(e.mixin.High(e.u, key), value)
				cluster.Insert(e.mixin.Low(e.u, key), value)
			} else {
				cluster.Insert(e.mixin.Low(e.u, key), value)
			}
		}

		if _key == e.max.key {
			e.max.value.PushBack(_value)
		} else if e.mixin.Less(e.u, e.max.key, _key) {
			e.max = newRsVEBTreeItem(_key, _value)
		}
	}
}

func newRsVEBTreeUint32(u int) *rsVEBTreeElement {
	return new(rsVEBTreeElement).init(u, new(rsVEBTreeUInt32Mixin))
}
