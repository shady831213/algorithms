package genericBinaryTree
import (
	"algorithms/tree/binaryTree"
)

type GBTElement struct {
	binaryTree.BinaryTreeElement
	parent, left, right *GBTElement
	SideValue interface{}
}

type GBT struct {
	nilNode *GBTElement//nil node
}

func (t *GBT) init() {
	t.nilNode = new(GBTElement)
	t.nilNode.left = t.nilNode
	t.nilNode.right = t.nilNode
	t.nilNode.parent = t.nilNode
}


func (t *GBT) IsNil(n interface{}) (bool) {
	return n.(*GBTElement) == t.nilNode
}


func (t *GBT) Root() (interface{}) {
	return t.nilNode.left
}

func (t *GBT) Search(key uint32) (interface{}) {
	for cur := t.Root().(*GBTElement); !t.IsNil(cur); {
		if cur.Key == key {
			return cur
		} else if key < cur.Key {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	return t.nilNode
}

func (t *GBT) Insert(node interface{}) {
	target := t.Root().(*GBTElement)
	n, isNode := node.(*GBTElement)
	if !isNode {
		n = new(GBTElement)
		n.Key = node.(uint32)
		n.left = t.nilNode
		n.right = t.nilNode
		n.parent = t.nilNode
	}
	for cur := t.Root().(*GBTElement); !t.IsNil(cur); {
		target = cur
		if n.Key < cur.Key {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	n.parent = target
	if t.IsNil(target) {
		t.nilNode.left = n
	} else if n.Key < target.Key {
		target.left = n
	} else {
		target.right = n
	}
}

func (t *GBT) Delete(key uint32) {
	deleteNonCompletedNode := func(node *GBTElement) {
		var reConnectedNode *GBTElement
		if t.IsNil(node.left) {
			reConnectedNode = node.right
		} else {
			reConnectedNode = node.left
		}
		if !t.IsNil(reConnectedNode) {
			reConnectedNode.parent = node.parent
		}
		if t.IsNil(node.parent) {
			t.nilNode.left = reConnectedNode
		} else if node.parent.right == node {
			node.parent.right = reConnectedNode
		} else {
			node.parent.left = reConnectedNode
		}
		node = t.nilNode
	}
	node := t.Search(key).(*GBTElement)
	if t.IsNil(node) {
		return
	}
	if t.IsNil(node.left) || t.IsNil(node.right) {
		deleteNonCompletedNode(node)
	} else {
		successor := t.Successor(node).(*GBTElement)
		_key, _value := successor.Key, successor.Value
		deleteNonCompletedNode(successor)
		node.Key, node.Value = _key, _value
	}
}

func (t *GBT) Min(node interface{}) (interface{}) {
	cur := node.(*GBTElement)
	for !t.IsNil(cur.left) {
		cur = cur.left
	}
	return cur
}

func (t *GBT) Max(node interface{}) (interface{}) {
	cur := node.(*GBTElement)
	for !t.IsNil(cur.right) {
		cur = cur.right
	}
	return cur
}

func (t *GBT) Predecesor(node interface{}) (interface{}) {
	n := node.(*GBTElement)
	if t.IsNil(n) {
		return t.nilNode
	}
	if !t.IsNil(n.left) {
		return t.Max(n.left)
	} else {
		cur := n
		for !t.IsNil(cur.parent) && cur.parent.right != cur {
			cur = cur.parent
		}
		return cur.parent
	}
}

func (t *GBT) Successor(node interface{}) (interface{}) {
	n := node.(*GBTElement)
	if t.IsNil(n) {
		return t.nilNode
	}
	if !t.IsNil(n.right) {
		return t.Min(n.right)
	} else {
		cur := n
		for !t.IsNil(cur.parent) && cur.parent.left != cur {
			cur = cur.parent
		}
		return cur.parent
	}
}

//next node should always be successor node
//O(n), all the connections(n-1) are accessed less than or equal to 2 times
func (t *GBT) InOrderWalk(node interface{}, callback func(binaryTree.BinaryTreeIf, interface{}) (bool)) (bool) {
	n := node.(*GBTElement)
	for curNode := t.Min(n).(*GBTElement); !t.IsNil(curNode); {
		stop := callback(t, curNode)
		if stop {
			return true
		}
		curNode = t.Successor(curNode).(*GBTElement)
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
func (t *GBT) PreOrderWalk(node interface{}, callback func(binaryTree.BinaryTreeIf, interface{}) (bool)) (bool) {
	root := node.(*GBTElement)

	goDown := func(curNode *GBTElement) (*GBTElement, bool) {
		if !t.IsNil(curNode.left) {
			return curNode.left, true
		} else if !t.IsNil(curNode.right) {
			return curNode.right, true
		}
		return curNode, false
	}

	goUp := func(curNode *GBTElement) (*GBTElement, bool) {
		if curNode == root || !t.IsNil(curNode.right) {
			return curNode.right, true
		} else if curNode == curNode.parent.left {
			for curNode == curNode.parent.left {
				curNode = curNode.parent
				if curNode == root || !t.IsNil(curNode.right) {
					return curNode.right, true
				}
			}
		} else {
			parentNode := curNode.parent
			parentRightNode := parentNode.right
			parentNode.right = t.nilNode
			curNode = t.Successor(parentNode).(*GBTElement)
			parentNode.right = parentRightNode
		}
		return curNode, false
	}

	down := true
	for curNode := root; !t.IsNil(curNode); {
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

func (t *GBT) PostOrderWalk(node interface{}, callback func(binaryTree.BinaryTreeIf, interface{}) (bool)) (bool) {
	n := node.(*GBTElement)

	leftistNode := func(curNode *GBTElement) (nextNode *GBTElement) {
		nextNode = curNode
		for !t.IsNil(nextNode.right) {
			nextNode = t.Min(nextNode.right).(*GBTElement)
		}
		return
	}

	for curNode := leftistNode(t.Min(n).(*GBTElement)); curNode != n; {
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

func New() binaryTree.BinaryTreeIf {
	t := new(GBT)
	t.init()
	return t
}

