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

func (bt *bTree) leftNode(n *bTreeNode) (*bTreeNode, int) {
	if n.p == nil {
		return nil, -1
	}

	i := n.p.searchKeyIdx(n.keyValue[0].key)

	if n.p.LessByKey(n.keyValue[0].key, n.p.keyValue[i].key) {
		if i <= 0 {
			return nil, -1
		}
		return n.p.c[i-1], i - 1
	} else {
		return n.p.c[i], i
	}
}

func (bt *bTree) rightNode(n *bTreeNode) (*bTreeNode, int) {
	if n.p == nil {
		return nil, -1
	}

	i := n.p.searchKeyIdx(n.keyValue[0].key)

	if n.p.LessByKey(n.p.keyValue[i].key, n.keyValue[0].key) {
		if i >= n.p.Len()-1 {
			return nil, -1
		}
		return n.p.c[i+2], i + 1
	} else {
		return n.p.c[i+1], i
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
			leftNode, leftKeyValueIdx := bt.leftNode(node)
			rightNode, rightKeyValueIdx := bt.rightNode(node)
			if leftNode != nil && !leftNode.isEmpty() {
				node.keyValue = append([]*keyValue{n.keyValue[leftKeyValueIdx]}, node.keyValue...)
				node.c = append([]*bTreeNode{leftNode.c[leftNode.Len()]}, node.c...)
				if !node.isLeaf {
					node.c[0].p = node
				}
				n.keyValue[leftKeyValueIdx] = leftNode.keyValue[leftNode.Len()-1]
				leftNode.c = leftNode.c[:leftNode.Len()]
				leftNode.keyValue = leftNode.keyValue[:leftNode.Len()-1]
				n = node
			} else if rightNode != nil && !rightNode.isEmpty() {
				node.keyValue = append(node.keyValue, n.keyValue[rightKeyValueIdx])
				node.c = append(node.c, rightNode.c[0])
				if !node.isLeaf {
					node.c[node.Len()].p = node
				}
				n.keyValue[rightKeyValueIdx] = rightNode.keyValue[0]
				rightNode.keyValue = rightNode.keyValue[1:]
				rightNode.c = rightNode.c[1:]
				n = node
			} else if leftNode != nil {
				node.keyValue = append([]*keyValue{n.keyValue[leftKeyValueIdx]}, node.keyValue...)
				node.keyValue = append(leftNode.keyValue, node.keyValue...)
				for _, v := range leftNode.c {
					if v != nil {
						v.p = node
					}
				}
				node.c = append(leftNode.c, node.c...)
				n.removeKeyValue(n.keyValue[leftKeyValueIdx].key)
				if n.Len() == 0 {
					node.p = n.p
					bt.root = node
					bt.height --
				}
				n = node
			} else if rightNode != nil {
				rightNode.keyValue = append([]*keyValue{n.keyValue[rightKeyValueIdx]}, rightNode.keyValue...)
				rightNode.keyValue = append(node.keyValue, rightNode.keyValue...)
				for _, v := range node.c {
					if v != nil {
						v.p = rightNode
					}
				}
				rightNode.c = append(node.c, rightNode.c...)
				n.removeKeyValue(n.keyValue[rightKeyValueIdx].key)
				if n.Len() == 0 {
					rightNode.p = n.p
					bt.root = rightNode
					bt.height --
				}
				n = rightNode
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
