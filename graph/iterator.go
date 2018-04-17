package graph

import "container/list"

type iterator interface {
	Len() int
	Next() interface{}
	Value() interface{}
}

type listIterator struct {
	list *list.List
	e    *list.Element
	iterator
}

func (i *listIterator) init(l *list.List) *listIterator {
	i.list = l
	i.e = i.list.Front()
	return i
}

func (i *listIterator) Len() int {
	return i.list.Len()
}

func (i *listIterator) Next() interface{} {
	if i.e == nil {
		return nil
	}
	if i.e = i.e.Next(); i.e == nil {
		return nil
	}
	return i.e.Value
}

func (i *listIterator) Value() interface{} {
	if i.e == nil {
		return nil
	}
	return i.e.Value
}

func newListIterator(list *list.List) *listIterator {
	return new(listIterator).init(list)
}

type linkedMapIterator struct {
	m   *linkedMap
	key interface{}
	iterator
}

func (i *linkedMapIterator) init(m *linkedMap) *linkedMapIterator {
	i.m = m
	i.key = i.m.frontKey()
	return i
}

func (i *linkedMapIterator) Len() int {
	return i.m.keyL.Len()
}

func (i *linkedMapIterator) Next() interface{} {
	if i.key == nil {
		return nil
	}
	if i.key = i.m.nextKey(i.key); i.key == nil {
		return nil
	}
	return struct {
		key, value interface{}
	}{i.key, i.m.get(i.key)}
}

func (i *linkedMapIterator) Value() interface{} {
	if i.key == nil {
		return nil
	}
	return struct {
		key, value interface{}
	}{i.key, i.m.get(i.key)}
}

func newLinkedMapIterator(m *linkedMap) *linkedMapIterator {
	return new(linkedMapIterator).init(m)
}
