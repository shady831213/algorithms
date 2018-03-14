package bTree

import (
	"sort"
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
	return n.bTreeNodeSort.LessByKey(n.keyValue[i].key, n.keyValue[j].key)
}

func (n *bTreeNode) searchKeyIdx(key interface{}) (int) {
	//binary search
	i, j := 0, len(n.keyValue)-1
	for i != j {
		mid := (j-i)/2 + i
		if key == n.keyValue[mid].key {
			i = mid
			j = mid
		} else if n.LessByKey(key, n.keyValue[mid].key) {
			j = mid
		} else {
			i = mid + 1
		}
	}
	return i
}

func (n *bTreeNode) getChildOrKeyValue(key interface{}) (interface{}, int) {
	if len(n.keyValue) == 0 {
		return nil, -1
	}

	i := n.searchKeyIdx(key)

	if key == n.keyValue[i].key {
		return n.keyValue[i], i
	} else if n.LessByKey(key, n.keyValue[i].key) {
		return n.c[i], i
	} else {
		return n.c[i+1], i + 1
	}
}

func (n *bTreeNode) predecesor(key interface{}) (*keyValue, *bTreeNode, int) {
	if len(n.keyValue) == 0 {
		return nil, nil, -1
	}

	i := n.searchKeyIdx(key)

	if key == n.keyValue[i].key {
		if i <= 0 {
			return nil, nil, -1
		}
		return n.keyValue[i-1], n.c[i-1], i - 1
	} else if n.LessByKey(key, n.keyValue[i].key) {
		if i <= 0 {
			return nil, nil, -1
		}
		return n.keyValue[i-1], n.c[i-1], i - 1
	} else {
		return n.keyValue[i], n.c[i], i
	}
}

func (n *bTreeNode) successor(key interface{}) (*keyValue, *bTreeNode, int) {
	if len(n.keyValue) == 0 {
		return nil, nil, -1
	}

	i := n.searchKeyIdx(key)

	if key == n.keyValue[i].key {
		if i >= len(n.keyValue)-1 {
			return nil, nil, -1
		}
		return n.keyValue[i+1], n.c[i+2], i + 1
	} else if n.LessByKey(n.keyValue[i].key, key) {
		if i >= len(n.keyValue)-1 {
			return nil, nil, -1
		}
		return n.keyValue[i+1], n.c[i+2], i + 1
	} else {
		return n.keyValue[i], n.c[i+1], i
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

func (n *bTreeNode) removeKeyValue(key interface{}) () {
	keyValueOrChild, keyValueIdx := n.getChildOrKeyValue(key)
	if _, ok := keyValueOrChild.(*keyValue); ok {
		n.keyValue = append(n.keyValue[:keyValueIdx], n.keyValue[keyValueIdx+1:]...)
		n.c = append(n.c[:keyValueIdx], n.c[keyValueIdx+1:]...)
	}
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

func (n *bTreeNode) isEmpty() (bool) {
	return len(n.keyValue) <= n.t-1
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

func (bt *bTree) split(n *bTreeNode) (int) {
	//because it is up-down, parent must be not full
	if !n.isFull() {
		panic("split when not full!")
	}
	n2 := bt.bTreeIf.newNode(bt.t)
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

func (bt *bTree) insertOrSet(n *bTreeNode, key, value interface{}) *bTreeNode {
	nodeOrKeyValue, _ := n.getChildOrKeyValue(key)
	//if node hit the key ,set value
	if keyValue, ok := nodeOrKeyValue.(*keyValue); ok {
		keyValue.value = value
		return nil
	}

	if n.isFull() {
		// if node is full, split
		if n.p == nil {
			//if node is root , increase hight
			n.p = bt.bTreeIf.newNode(bt.t)
			n.p.isLeaf = false
			bt.root = n.p
			bt.height ++
		}
		p := n.p
		i := bt.split(n)
		if p.keyValue[i].key == key {
			//if hit the middle key, set value
			p.keyValue[i].value = value
			return nil
		} else if p.LessByKey(key, p.keyValue[i].key) {
			//reture left part
			return p.c[i]
		} else {
			//reture right part
			return p.c[i+1]
		}
	} else if n.isLeaf{
		// if it is leaf, add key-value
		n.addKeyValue(key, value)
		return nil
	} else {
		// return child
		return nodeOrKeyValue.(*bTreeNode)
	}
}

func (bt *bTree) insert(key, value interface{}) {
	//empty tree
	if bt.root == nil {
		bt.root = bt.bTreeIf.newNode(bt.t)
		bt.height++
	}

	for n := bt.root;n != nil; {
		n = bt.insertOrSet(n, key, value)
	}

}

func (bt *bTree) remove(key interface{}) {
	n := bt.root
	k := key
	for !n.isLeaf {
		nodeOrKeyValue, keyValueIdx := n.getChildOrKeyValue(k)
		if _, ok := nodeOrKeyValue.(*keyValue); ok {
			preChild, postChild := n.c[keyValueIdx], n.c[keyValueIdx+1]
			if !preChild.isEmpty() {
				predecesorIdx := preChild.Len() - 1
				n.keyValue[keyValueIdx] = preChild.keyValue[predecesorIdx]
				n = preChild
				k = preChild.keyValue[predecesorIdx].key
			} else if !postChild.isEmpty() {
				successorIdx := 0
				n.keyValue[keyValueIdx] = postChild.keyValue[successorIdx]
				n = postChild
				k = postChild.keyValue[successorIdx].key
			} else {
				preChild.keyValue = append(preChild.keyValue, n.keyValue[keyValueIdx])
				postChild.keyValue = append(preChild.keyValue, postChild.keyValue...)
				for _, v := range preChild.c {
					if v != nil {
						v.p = postChild
					}
				}
				postChild.c = append(preChild.c, postChild.c...)
				n.removeKeyValue(k)
				if n.Len() == 0 {
					postChild.p = n.p
					bt.root = postChild
					bt.height --
				}
				n = postChild
				k = key
			}
		} else if node := nodeOrKeyValue.(*bTreeNode); node.isEmpty() {
			_, preNode, preKeyValueIdx := n.predecesor(node.keyValue[0].key)
			_, sucNode, sucKeyValueIdx := n.successor(node.keyValue[0].key)
			if preNode != nil && !preNode.isEmpty() {
				node.keyValue = append([]*keyValue{n.keyValue[preKeyValueIdx]}, node.keyValue...)
				node.c = append([]*bTreeNode{preNode.c[preNode.Len()]}, node.c...)
				if !node.isLeaf {
					node.c[0].p = node
				}
				n.keyValue[preKeyValueIdx] = preNode.keyValue[preNode.Len()-1]
				preNode.c = preNode.c[:preNode.Len()]
				preNode.keyValue = preNode.keyValue[:preNode.Len()-1]
				n = node
			} else if sucNode != nil && !sucNode.isEmpty() {
				node.keyValue = append(node.keyValue, n.keyValue[sucKeyValueIdx])
				node.c = append(node.c, sucNode.c[0])
				if !node.isLeaf {
					node.c[node.Len()].p = node
				}
				n.keyValue[sucKeyValueIdx] = sucNode.keyValue[0]
				sucNode.keyValue = sucNode.keyValue[1:]
				sucNode.c = sucNode.c[1:]
				n = node
			} else if preNode != nil {
				node.keyValue = append([]*keyValue{n.keyValue[preKeyValueIdx]}, node.keyValue...)
				node.keyValue = append(preNode.keyValue, node.keyValue...)
				for _, v := range preNode.c {
					if v != nil {
						v.p = node
					}
				}
				node.c = append(preNode.c, node.c...)
				n.removeKeyValue(n.keyValue[preKeyValueIdx].key)
				if n.Len() == 0 {
					node.p = n.p
					bt.root = node
					bt.height --
				}
				n = node
			} else if sucNode != nil {
				sucNode.keyValue = append([]*keyValue{n.keyValue[sucKeyValueIdx]}, sucNode.keyValue...)
				sucNode.keyValue = append(node.keyValue, sucNode.keyValue...)
				for _, v := range node.c {
					if v != nil {
						v.p = sucNode
					}
				}
				sucNode.c = append(node.c, sucNode.c...)
				n.removeKeyValue(n.keyValue[sucKeyValueIdx].key)
				if n.Len() == 0 {
					sucNode.p = n.p
					bt.root = sucNode
					bt.height --
				}
				n = sucNode
			}
		} else {
			n = node
		}

	}
	n.removeKeyValue(k)
	if n.Len() == 0 {
		bt.root = nil
		bt.height --
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
