package graph

type iterator interface {
	Len() int
	Next() interface{}
	Value() interface{}
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
	return i.key
}

func (i *linkedMapIterator) Value() interface{} {
	if i.key == nil {
		return nil
	}
	return i.key
}

func newLinkedMapIterator(m *linkedMap) *linkedMapIterator {
	return new(linkedMapIterator).init(m)
}
