package graph

import "container/list"

type linkedMap struct {
	keyL *list.List
	m    map[interface{}]interface{}
}

func (m *linkedMap) init() *linkedMap {
	m.keyL = list.New()
	m.m = make(map[interface{}]interface{})
	return m
}

func (m *linkedMap) exist(key interface{}) bool {
	_, ok := m.m[key]
	return ok
}

func (m *linkedMap) add(key, value interface{}) {
	if !m.exist(key) {
		e := m.keyL.PushBack(key)
		m.m[key] = []interface{}{e, value}
	} else {
		m.m[key].([]interface{})[1] = value
	}
}

func (m *linkedMap) get(key interface{}) interface{} {
	if m.exist(key) {
		return m.m[key].([]interface{})[1]
	}
	return nil
}

func (m *linkedMap) delete(key interface{}) {
	if m.exist(key) {
		i := m.m[key].([]interface{})
		m.keyL.Remove(i[0].(*list.Element))
		delete(m.m, key)
	}
}

func (m *linkedMap) frontKey() interface{} {
	if m.keyL.Len() == 0 {
		return nil
	}
	return m.keyL.Front().Value
}

func (m *linkedMap) backKey() interface{} {
	if m.keyL.Len() == 0 {
		return nil
	}
	return m.keyL.Back().Value
}

func (m *linkedMap) nextKey(key interface{}) interface{} {
	if e := m.m[key].([]interface{})[0].(*list.Element).Next(); e != nil {
		return e.Value
	}
	return nil
}

func (m *linkedMap) prevKey(key interface{}) interface{} {
	if e := m.m[key].([]interface{})[0].(*list.Element).Prev(); e != nil {
		return e.Value
	}
	return nil
}
