package binaryTree

type bstElement struct {
	binaryTreeElement
	parent, left, right *bstElement
}

type bst struct {
	root *bstElement
}

func (t *bst) IsNil(n interface{}) bool {
	return n == nil
}

func (t *bst) Root() interface{} {
	return t.root
}

func (t *bst) Search(key uint32) interface{} {
	for cur := t.root; cur != nil; {
		if cur.Key == key {
			return cur
		} else if key < cur.Key {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	return nil
}

func (t *bst) Insert(node interface{}) interface{} {
	var target *bstElement
	n, isNode := node.(*bstElement)
	if !isNode {
		n = new(bstElement)
		n.Key = node.(uint32)
	}
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
	return n
}

func (t *bst) Delete(key uint32) interface{} {
	deleteNonCompletedNode := func(node *bstElement) {
		var reConnectedNode *bstElement
		if node.left == nil {
			reConnectedNode = node.right
		} else {
			reConnectedNode = node.left
		}
		if reConnectedNode != nil {
			reConnectedNode.parent = node.parent
		}
		if node.parent == nil {
			t.root = reConnectedNode
		} else if node.parent.right == node {
			node.parent.right = reConnectedNode
		} else {
			node.parent.left = reConnectedNode
		}
		node = nil
	}
	node := t.Search(key).(*bstElement)
	if node == nil {
		return node
	}
	if node.left == nil || node.right == nil {
		deleteNonCompletedNode(node)
	} else {
		successor := t.Successor(node, t.Root()).(*bstElement)
		_key, _value := successor.Key, successor.Value
		deleteNonCompletedNode(successor)
		node.Key, node.Value = _key, _value
	}
	return node
}

func (t *bst) Min(node interface{}) interface{} {
	cur := node.(*bstElement)
	for cur.left != nil {
		cur = cur.left
	}
	return cur
}

func (t *bst) Max(node interface{}) interface{} {
	cur := node.(*bstElement)
	for cur.right != nil {
		cur = cur.right
	}
	return cur
}

func (t *bst) Predecessor(node interface{}, root interface{}) interface{} {
	n := node.(*bstElement)
	if n == nil {
		return nil
	}
	if n.left != nil {
		return t.Max(n.left)
	}
	cur := n
	for cur.parent != nil && cur.parent.right != cur {
		cur = cur.parent
	}
	return cur.parent

}

func (t *bst) Successor(node interface{}, root interface{}) interface{} {
	n := node.(*bstElement)
	if n == nil {
		return nil
	}
	if n.right != nil {
		return t.Min(n.right)
	}
	cur := n
	for cur.parent != nil && cur.parent.left != cur {
		cur = cur.parent
	}
	return cur.parent

}

func (t *bst) LeftRotate(node interface{}) interface{} {
	panic("not implement in bst!")
}

func (t *bst) RightRotate(node interface{}) interface{} {
	panic("not implement in bst!")
}

type bstRecrusive struct {
	bst
}

func (t *bstRecrusive) InOrderWalk(node interface{}, callback func(binaryTreeIf, interface{}) bool) bool {
	n := node.(*bstElement)
	if n != nil {
		stop := t.InOrderWalk(n.left, callback)
		if stop {
			return true
		}
		stop = callback(t, n)
		if stop {
			return true
		}
		stop = t.InOrderWalk(n.right, callback)
		return stop
	}
	return false
}

func (t *bstRecrusive) PreOrderWalk(node interface{}, callback func(binaryTreeIf, interface{}) bool) bool {
	n := node.(*bstElement)
	if n != nil {
		stop := callback(t, n)
		if stop {
			return true
		}
		stop = t.PreOrderWalk(n.left, callback)
		if stop {
			return true
		}
		stop = t.PreOrderWalk(n.right, callback)
		return stop
	}
	return false
}

func (t *bstRecrusive) PostOrderWalk(node interface{}, callback func(binaryTreeIf, interface{}) bool) bool {
	n := node.(*bstElement)
	if n != nil {
		stop := t.PostOrderWalk(n.left, callback)
		if stop {
			return true
		}
		stop = t.PostOrderWalk(n.right, callback)
		if stop {
			return true
		}
		stop = callback(t, n)
		return stop
	}
	return false
}

func newBstRecrusive() *bstRecrusive {
	return new(bstRecrusive)
}

type bstIterative struct {
	bst
}

//next node should always be successor node
//O(n), all the connections(n-1) are accessed less than or equal to 2 times
func (t *bstIterative) InOrderWalk(node interface{}, callback func(binaryTreeIf, interface{}) bool) bool {
	n := node.(*bstElement)
	for curNode := t.Min(n).(*bstElement); curNode != nil; {
		stop := callback(t, curNode)
		if stop {
			return true
		}
		curNode = t.Successor(curNode, n).(*bstElement)
	}
	return false
}

//GoDown: if the node has left child, go through left, otherwise if the node has right child, go through right. After it gets leaf node, go up
//GoUp : left node: find a completed node or a right node like this:
//                                                                 3
//                                                                  \
//                                                                 cur(5)
//                                                                 /
//                                                                4
// right node: remove itself from parent, then find the successor of parent, then recover parent
//During going up, when it gets root or a node has right child , go down
//O(n), all the connections(n-1) are accessed less than or equal to 2 times
func (t *bstIterative) PreOrderWalk(node interface{}, callback func(binaryTreeIf, interface{}) bool) bool {
	root := node.(*bstElement)

	goDown := func(curNode *bstElement) (*bstElement, bool) {
		if curNode.left != nil {
			return curNode.left, true
		} else if curNode.right != nil {
			return curNode.right, true
		}
		return curNode, false
	}

	goUp := func(curNode *bstElement) (*bstElement, bool) {
		if curNode == root || curNode.right != nil {
			return curNode.right, true
		} else if curNode == curNode.parent.left {
			for curNode == curNode.parent.left {
				curNode = curNode.parent
				if curNode == root || curNode.right != nil {
					return curNode.right, true
				}
			}
		} else {
			parentNode := curNode.parent
			parentRightNode := parentNode.right
			parentNode.right = nil
			curNode = t.Successor(parentNode, root).(*bstElement)
			parentNode.right = parentRightNode
		}
		return curNode, false
	}

	down := true
	for curNode := root; curNode != nil; {
		if down {
			stop := callback(t, curNode)
			if stop {
				return true
			}
			curNode, down = goDown(curNode)
		} else {
			curNode, down = goUp(curNode)
		}
	}
	return false
}

//start from the leftist node, which must be the min node or leftist leaf node of right sub tree of min node
//if the node is left leaf node, find the leftist node of right sub tree of parent
//if the node is right leaf node, go bach to parent
//O(n), all the connections(n-1) are accessed less than or equal to 2 times

func (t *bstIterative) PostOrderWalk(node interface{}, callback func(binaryTreeIf, interface{}) bool) bool {
	n := node.(*bstElement)

	leftistNode := func(curNode *bstElement) (nextNode *bstElement) {
		nextNode = curNode
		for nextNode.right != nil {
			nextNode = t.Min(nextNode.right).(*bstElement)
		}
		return
	}

	for curNode := leftistNode(t.Min(n).(*bstElement)); curNode != n; {
		stop := callback(t, curNode)
		if stop {
			return true
		}
		parentNode := curNode.parent
		if curNode == parentNode.left {
			curNode = leftistNode(parentNode)
		} else {
			curNode = parentNode
		}

	}
	return callback(t, n)
}

func newBstIterative() *bstIterative {
	return new(bstIterative)
}
