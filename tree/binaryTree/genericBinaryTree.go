package binaryTree

type GBTElement struct {
	BinaryTreeElement
	Parent, Left, Right *GBTElement
	SideValue           interface{}
}

type GBT struct {
	NilNode *GBTElement //nil node, Left, Right point to the root, Parent point to it self, empty pointer point to th NilNode
	Object  BinaryTreeIf
}

func (t *GBT) Init() {
	t.NilNode = new(GBTElement)
	t.NilNode.Left = t.NilNode
	t.NilNode.Right = t.NilNode
	t.NilNode.Parent = t.NilNode
	t.Object = t
}

func (t *GBT) IsNil(n interface{}) bool {
	return n.(*GBTElement) == t.NilNode
}

func (t *GBT) Root() interface{} {
	return t.NilNode.Left
}

func (t *GBT) Search(key uint32) interface{} {
	for cur := t.Root().(*GBTElement); !t.IsNil(cur); {
		if cur.Key == key {
			return cur
		} else if key < cur.Key {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	return t.NilNode
}

func (t *GBT) Insert(node interface{}) interface{} {
	target := t.Root().(*GBTElement)
	n, isNode := node.(*GBTElement)
	if !isNode {
		n = new(GBTElement)
		n.Key = node.(uint32)
		n.Left = t.NilNode
		n.Right = t.NilNode
		n.Parent = t.NilNode
	}
	for cur := t.Root().(*GBTElement); !t.IsNil(cur); {
		target = cur
		if n.Key < cur.Key {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	n.Parent = target
	if t.IsNil(target) {
		t.NilNode.Left = n
		t.NilNode.Right = n
	} else if n.Key < target.Key {
		target.Left = n
	} else {
		target.Right = n
	}

	return n
}

func (t *GBT) Delete(key uint32) interface{} {
	deleteNonCompletedNode := func(node *GBTElement) {
		var reConnectedNode *GBTElement
		if t.IsNil(node.Left) {
			reConnectedNode = node.Right
		} else {
			reConnectedNode = node.Left
		}
		if !t.IsNil(reConnectedNode) {
			reConnectedNode.Parent = node.Parent
		}
		if t.IsNil(node.Parent) {
			t.NilNode.Left = reConnectedNode
			t.NilNode.Right = reConnectedNode
		} else if node.Parent.Right == node {
			node.Parent.Right = reConnectedNode
		} else {
			node.Parent.Left = reConnectedNode
		}
	}
	node := t.Search(key).(*GBTElement)
	if t.IsNil(node) {
		return node
	}
	if t.IsNil(node.Left) || t.IsNil(node.Right) {
		deleteNonCompletedNode(node)
	} else {
		successor := t.Successor(node, t.Root()).(*GBTElement)
		_key, _value := successor.Key, successor.Value
		node.Key, node.Value = _key, _value
		deleteNonCompletedNode(successor)
	}
	return node
}

func (t *GBT) Min(node interface{}) interface{} {
	cur := node.(*GBTElement)
	for !t.IsNil(cur.Left) {
		cur = cur.Left
	}
	return cur
}

func (t *GBT) Max(node interface{}) interface{} {
	cur := node.(*GBTElement)
	for !t.IsNil(cur.Right) {
		cur = cur.Right
	}
	return cur
}

func (t *GBT) Predecesor(node interface{}, root interface{}) interface{} {
	r := root.(*GBTElement)
	if r == nil {
		r = t.Root().(*GBTElement)
	}
	n := node.(*GBTElement)
	if t.IsNil(n) {
		return t.NilNode
	}
	if !t.IsNil(n.Left) {
		return t.Max(n.Left)
	} else {
		cur := n
		for cur != r && cur.Parent.Right != cur {
			cur = cur.Parent
		}
		if cur == r {
			return t.NilNode
		}
		return cur.Parent
	}
}

func (t *GBT) Successor(node interface{}, root interface{}) interface{} {
	r := root.(*GBTElement)
	if r == nil {
		r = t.Root().(*GBTElement)
	}
	n := node.(*GBTElement)
	if t.IsNil(n) {
		return t.NilNode
	}
	if !t.IsNil(n.Right) {
		return t.Min(n.Right)
	} else {
		cur := n
		for cur != r && cur.Parent.Left != cur {
			cur = cur.Parent
		}
		if cur == r {
			return t.NilNode
		}
		return cur.Parent
	}
}

func (t *GBT) LeftRotate(node interface{}) interface{} {
	n := node.(*GBTElement)
	if t.IsNil(n.Right) {
		return t.NilNode
	}
	newNode := n.Right
	if n.Parent.Left == n {
		n.Parent.Left = newNode
	}
	if n.Parent.Right == n {
		n.Parent.Right = newNode
	}
	n.Parent, newNode.Parent = newNode, n.Parent
	newNode.Left, n.Right = n, newNode.Left
	if !t.IsNil(n.Right) {
		n.Right.Parent = n
	}
	return newNode
}

func (t *GBT) RightRotate(node interface{}) interface{} {
	n := node.(*GBTElement)
	if t.IsNil(n.Left) {
		return t.NilNode
	}
	newNode := n.Left
	if n.Parent.Right == n {
		n.Parent.Right = newNode
	}
	if n.Parent.Left == n {
		n.Parent.Left = newNode
	}
	n.Parent, newNode.Parent = newNode, n.Parent
	newNode.Right, n.Left = n, newNode.Right
	if !t.IsNil(n.Left) {
		n.Left.Parent = n
	}
	return newNode
}

//next node should always be successor node
//O(n), all the connections(n-1) are accessed less than or equal to 2 times
func (t *GBT) InOrderWalk(node interface{}, callback func(BinaryTreeIf, interface{}) bool) bool {
	n := node.(*GBTElement)
	for curNode := t.Min(n).(*GBTElement); !t.IsNil(curNode); {
		stop := callback(t.Object, curNode)
		if stop {
			return true
		}
		curNode = t.Successor(curNode, n).(*GBTElement)
	}
	return false
}

//GoDown: if the node has Left child, go through Left, otherwise if the node has Right child, go through Right. After it gets leaf node, go up
//GoUp : Left node: find a completed node or a Right node like this:
//                                                                 3
//                                                                  \
//                                                                 cur(5)
//                                                                 /
//                                                                4
// Right node: remove itself from Parent, then find the successor of Parent, then recover Parent
//During going up, when it gets root or a node has Right child , go down
//O(n), all the connections(n-1) are accessed less than or equal to 2 times
func (t *GBT) PreOrderWalk(node interface{}, callback func(BinaryTreeIf, interface{}) bool) bool {
	root := node.(*GBTElement)

	goDown := func(curNode *GBTElement) (*GBTElement, bool) {
		if !t.IsNil(curNode.Left) {
			return curNode.Left, true
		} else if !t.IsNil(curNode.Right) {
			return curNode.Right, true
		}
		return curNode, false
	}

	goUp := func(curNode *GBTElement) (*GBTElement, bool) {
		if curNode == root || !t.IsNil(curNode.Right) {
			return curNode.Right, true
		} else if curNode == curNode.Parent.Left {
			for curNode == curNode.Parent.Left {
				curNode = curNode.Parent
				if curNode == root || !t.IsNil(curNode.Right) {
					return curNode.Right, true
				}
			}
		} else {
			parentNode := curNode.Parent
			parentRightNode := parentNode.Right
			parentNode.Right = t.NilNode
			curNode = t.Successor(parentNode, root).(*GBTElement)
			parentNode.Right = parentRightNode
		}
		return curNode, false
	}

	down := true
	for curNode := root; !t.IsNil(curNode); {
		if down {
			stop := callback(t.Object, curNode)
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

//start from the leftist node, which must be the min node or leftist leaf node of Right sub tree of min node
//if the node is Left leaf node, find the leftist node of Right sub tree of Parent
//if the node is Right leaf node, go bach to Parent
//O(n), all the connections(n-1) are accessed less than or equal to 2 times

func (t *GBT) PostOrderWalk(node interface{}, callback func(BinaryTreeIf, interface{}) bool) bool {
	n := node.(*GBTElement)

	leftistNode := func(curNode *GBTElement) (nextNode *GBTElement) {
		nextNode = curNode
		for !t.IsNil(nextNode.Right) {
			nextNode = t.Min(nextNode.Right).(*GBTElement)
		}
		return
	}

	for curNode := leftistNode(t.Min(n).(*GBTElement)); curNode != n; {
		stop := callback(t.Object, curNode)
		if stop {
			return true
		}
		parentNode := curNode.Parent
		if curNode == parentNode.Left {
			curNode = leftistNode(parentNode)
		} else {
			curNode = parentNode
		}

	}
	return callback(t, n)
}

func NewGBT() BinaryTreeIf {
	t := new(GBT)
	t.Init()
	return t
}
