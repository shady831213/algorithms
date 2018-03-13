package bTree

import (
	"sort"
	"reflect"
)

//Node
type keyValue struct {
	key, value interface{}
}

type bTreeNodeSort interface {
	sort.Interface
	LessByKey(interface{}, interface{}) (bool)
}

type bTreeNode struct {
	p        *bTreeNode
	c        []*bTreeNode
	keyValue []*keyValue
	isLeaf   bool
	t        int
	bTreeNodeSort
}

func (n *bTreeNode) init(t int, self bTreeNodeSort) (*bTreeNode) {
	n.t = t
	n.c = make([]*bTreeNode, 1, 1)
	n.keyValue = make([]*keyValue, 0, 0)
	n.isLeaf = true
	n.bTreeNodeSort = self
	return n
}

func (n *bTreeNode) Len() int {
	return len(n.keyValue)
}

func (n *bTreeNode) Swap(i, j int) {
	n.keyValue[i], n.keyValue[j] = n.keyValue[j], n.keyValue[i]
}

func (n *bTreeNode) Less(i, j int) bool {
	return n.bTreeNodeSort.LessByKey(n.keyValue[i].key,n.keyValue[j].key)
}

func (n *bTreeNode) hitOrGetChild(key interface{}) (*bTreeNode, *keyValue) {
	if len(n.keyValue) == 0 {
		return n, nil
	}
	//binary search
	i, j := 0, len(n.keyValue)-1
	for i != j {
		mid := (j-i)/2 + i
		if reflect.DeepEqual(key, n.keyValue[mid].key) {
			return n, n.keyValue[mid]
		} else if n.LessByKey(key, n.keyValue[mid].key) {
			j = mid
		} else {
			i = mid + 1
		}
	}
	if n.LessByKey(key, n.keyValue[i].key) {
		return n.c[i], nil
	} else {
		return n.c[i+1], nil
	}
}

func (n *bTreeNode) addKeyValue(key, value interface{}) (int) {
	//insert sort
	n.keyValue = append(n.keyValue, &keyValue{key, value})
	i := len(n.keyValue) - 1
	for ; i > 0 && n.Less(i, i-1); i -- {
		n.Swap(i-1, i)
	}
	temp := n.c[i:]
	n.c = append(n.c[:i], nil)
	n.c = append(n.c, temp...)
	return i
}

func (n *bTreeNode) split(n2 *bTreeNode) (int) {
	//because it is up-down, parent must be not full
	if !n.isFull() {
		panic("split when not full!")
	}
	n2.p = n.p
	n2.isLeaf = n.isLeaf
	n2.keyValue = append(make([]*keyValue, 0, 0), n.keyValue[n.t:]...)
	n2.c = append(make([]*bTreeNode, 0, 0), n.c[n.t:]...)
	for _, v := range n2.c {
		if v != nil {
			v.p = n2
		}
	}
	i := n.p.addKeyValue(n.keyValue[n.t-1].key, n.keyValue[n.t-1].value)
	n.keyValue = n.keyValue[:n.t-1]
	n.c = n.c[:n.t]
	n.p.c[i] = n
	n.p.c[i+1] = n2
	return i
}

func (n *bTreeNode) splitAndGetChild(n2 *bTreeNode, key interface{}) *bTreeNode {
	i := n.split(n2)
	if n.LessByKey(key, n.p.keyValue[i].key) {
		return n
	} else {
		return n2
	}
}

func (n *bTreeNode) isFull() (bool) {
	return len(n.keyValue) == 2*n.t-1
}

//Tree
type bTreeIf interface {
	newNode(int) (*bTreeNode)
}

type bTree struct {
	root      *bTreeNode
	t, height int
	bTreeIf
}

func (bt *bTree) init(t int, self bTreeIf) (*bTree) {
	bt.t = t
	bt.height = 0
	bt.bTreeIf = self
	return bt
}

func (bt *bTree) insert(key, value interface{}) {
	//empty tree
	if bt.root == nil {
		bt.root = bt.bTreeIf.newNode(bt.t)
		bt.height++
	}

	//hit root
	if _, keyValue := bt.root.hitOrGetChild(key); keyValue != nil {
		keyValue.value = value
		return
	}

	n := bt.root

	//root is full
	if n.isFull() {
		bt.root = bt.bTreeIf.newNode(bt.t)
		bt.root.isLeaf = false
		n.p = bt.root
		n = n.splitAndGetChild(bt.bTreeIf.newNode(bt.t), key)
		bt.height++
	}

	for !n.isLeaf {
		c, keyValue := n.hitOrGetChild(key)
		if keyValue != nil {
			keyValue.value = value
			return
		}
		//n is full
		if c.isFull() {
			n = c.splitAndGetChild(bt.bTreeIf.newNode(bt.t), key)
		} else {
			n = c
		}
	}

	if _, keyValue := n.hitOrGetChild(key); keyValue != nil {
		keyValue.value = value
	} else {
		n.addKeyValue(key, value)
	}
}

func (bt *bTree) preOrderWalk(node *bTreeNode, callback func(*bTree, *bTreeNode) (bool)) (bool) {
	if node == nil {
		return false
	}
	if stop := callback(bt, node); stop {
		return true
	}
	for i := range node.c {
		if stop := bt.preOrderWalk(node.c[i], callback); stop {
			return true
		}
	}
	return false
}

func (bt *bTree) inOrderWalk(node *bTreeNode, callback func(*bTree, *bTreeNode, int) (bool)) (bool) {
	if node == nil {
		return false
	}
	for i := range node.c {
		if stop := bt.inOrderWalk(node.c[i], callback); stop {
			return true
		}
		if i < node.Len() {
			if stop := callback(bt, node, i); stop {
				return true
			}
		}
	}

	return false
}
