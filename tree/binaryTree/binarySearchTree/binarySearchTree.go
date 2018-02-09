package binarySearchTree

import (
	"algorithms/tree/binaryTree"
)

type BstElement struct {
	binaryTree.BinaryTreeElement
	parent, left, right *BstElement
}

type Bst struct {
	root *BstElement
}

func (t *Bst) Search(key uint32) (interface{}) {
	for cur := t.root; cur != nil; {
		if cur.Key == key {
			return cur
		}else if key < cur.Key {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	return nil
}

func (t *Bst) Insert(node interface{}) {
	var target *BstElement
	n := node.(*BstElement)
	for cur := t.root; cur != nil; {
		target = cur
		if n.Key < cur.Key {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	n.parent = target
	if target == nil {
		t.root = n
	} else if n.Key < target.Key {
		target.left = n
	} else {
		target.right = n
	}
}

func (t *Bst) Delete(key uint32) {
	deleteNonCompletedNode := func (node *BstElement) {
		var reConnectedNode  *BstElement
		if node.left == nil {
			reConnectedNode = node.right
		} else {
			reConnectedNode = node.left
		}
		if reConnectedNode != nil {
			reConnectedNode.parent = node.parent
		}
		if node.parent.right == node {
			node.parent.right = reConnectedNode
		} else {
			node.parent.left = reConnectedNode
		}
		node = nil
	}
	node := t.Search(key).(*BstElement)
	if node == nil {
		return
	}
	if node.left == nil || node.right == nil {
		deleteNonCompletedNode(node)
	} else {
		successor := t.Successor(key).(*BstElement)
		_key, _value := successor.Key, successor.Value
		deleteNonCompletedNode(successor)
		node.Key, node.Value = _key, _value
	}
}

func (t *Bst) Min(node interface{}) (interface{}) {
	cur:=node.(*BstElement)
	for cur.left != nil{
		cur = cur.left
	}
	return cur
}

func (t *Bst) Max(node interface{}) (interface{}) {
	cur:=node.(*BstElement)
	for cur.right != nil{
		cur = cur.right
	}
	return cur
}

func (t *Bst) Predecesor(node interface{}) (interface{}) {
	n := node.(*BstElement)
	if n == nil {
		return nil
	}
	if n.left != nil {
		return t.Max(n.left)
	} else {
		cur := n
		for cur.parent == nil || cur.parent.right != cur {
			cur = cur.parent
		}
		return cur.parent
	}
}

func (t *Bst) Successor(node interface{}) (interface{}) {
	n := node.(*BstElement)
	if n == nil {
		return nil
	}
	if n.right != nil {
		return t.Min(n.right)
	} else {
		cur := n
		for cur.parent == nil || cur.parent.left != cur {
			cur = cur.parent
		}
		return cur.parent
	}
}

type BstRecrusive struct {
	Bst
}

func (t *BstRecrusive) InOrderWalk(node interface{}, callback func(interface{}, ...interface{}), args ...interface{}) {
	if node != nil {
		n := node.(*BstElement)
		t.InOrderWalk(n.left, callback, args)
		callback(n, args)
		t.InOrderWalk(n.right, callback, args)
	}
}

func NewBstRecrusive() binaryTree.BinaryTreeIf {
	return new(BstRecrusive)
}
