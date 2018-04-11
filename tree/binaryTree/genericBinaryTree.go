package binaryTree

type gbtElement struct {
	binaryTreeElement
	Parent, Left, Right *gbtElement
	SideValue           interface{}
}

type gbt struct {
	NilNode *gbtElement //nil node, Left, Right point to the root, Parent point to it self, empty pointer point to th NilNode
	Object  binaryTreeIf
}

func (t *gbt) Init() {
	t.NilNode = new(gbtElement)
	t.NilNode.Left = t.NilNode
	t.NilNode.Right = t.NilNode
	t.NilNode.Parent = t.NilNode
	t.Object = t
}

func (t *gbt) IsNil(n interface{}) bool {
	return n.(*gbtElement) == t.NilNode
}

func (t *gbt) Root() interface{} {
	return t.NilNode.Left
}

func (t *gbt) Search(key uint32) interface{} {
	for cur := t.Root().(*gbtElement); !t.IsNil(cur); {
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

func (t *gbt) Insert(node interface{}) interface{} {
	target := t.Root().(*gbtElement)
	n, isNode := node.(*gbtElement)
	if !isNode {
		n = new(gbtElement)
		n.Key = node.(uint32)
		n.Left = t.NilNode
		n.Right = t.NilNode
		n.Parent = t.NilNode
	}
	for cur := t.Root().(*gbtElement); !t.IsNil(cur); {
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

func (t *gbt) Delete(key uint32) interface{} {
	deleteNonCompletedNode := func(node *gbtElement) {
		var reConnectedNode *gbtElement
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
	node := t.Search(key).(*gbtElement)
	if t.IsNil(node) {
		return node
	}
	if t.IsNil(node.Left) || t.IsNil(node.Right) {
		deleteNonCompletedNode(node)
	} else {
		successor := t.Successor(node, t.Root()).(*gbtElement)
		_key, _value := successor.Key, successor.Value
		node.Key, node.Value = _key, _value
		deleteNonCompletedNode(successor)
	}
	return node
}

func (t *gbt) Min(node interface{}) interface{} {
	cur := node.(*gbtElement)
	for !t.IsNil(cur.Left) {
		cur = cur.Left
	}
	return cur
}

func (t *gbt) Max(node interface{}) interface{} {
	cur := node.(*gbtElement)
	for !t.IsNil(cur.Right) {
		cur = cur.Right
	}
	return cur
}

func (t *gbt) Predecessor(node interface{}, root interface{}) interface{} {
	r := root.(*gbtElement)
	if r == nil {
		r = t.Root().(*gbtElement)
	}
	n := node.(*gbtElement)
	if t.IsNil(n) {
		return t.NilNode
	}
	if !t.IsNil(n.Left) {
		return t.Max(n.Left)
	}
	cur := n
	for cur != r && cur.Parent.Right != cur {
		cur = cur.Parent
	}
	if cur == r {
		return t.NilNode
	}
	return cur.Parent
}

func (t *gbt) Successor(node interface{}, root interface{}) interface{} {
	r := root.(*gbtElement)
	if r == nil {
		r = t.Root().(*gbtElement)
	}
	n := node.(*gbtElement)
	if t.IsNil(n) {
		return t.NilNode
	}
	if !t.IsNil(n.Right) {
		return t.Min(n.Right)
	}
	cur := n
	for cur != r && cur.Parent.Left != cur {
		cur = cur.Parent
	}
	if cur == r {
		return t.NilNode
	}
	return cur.Parent
}

func (t *gbt) LeftRotate(node interface{}) interface{} {
	n := node.(*gbtElement)
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

func (t *gbt) RightRotate(node interface{}) interface{} {
	n := node.(*gbtElement)
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
func (t *gbt) InOrderWalk(node interface{}, callback func(binaryTreeIf, interface{}) bool) bool {
	n := node.(*gbtElement)
	for curNode := t.Min(n).(*gbtElement); !t.IsNil(curNode); {
		stop := callback(t.Object, curNode)
		if stop {
			return true
		}
		curNode = t.Successor(curNode, n).(*gbtElement)
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
func (t *gbt) PreOrderWalk(node interface{}, callback func(binaryTreeIf, interface{}) bool) bool {
	root := node.(*gbtElement)

	goDown := func(curNode *gbtElement) (*gbtElement, bool) {
		if !t.IsNil(curNode.Left) {
			return curNode.Left, true
		} else if !t.IsNil(curNode.Right) {
			return curNode.Right, true
		}
		return curNode, false
	}

	goUp := func(curNode *gbtElement) (*gbtElement, bool) {
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
			curNode = t.Successor(parentNode, root).(*gbtElement)
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

func (t *gbt) PostOrderWalk(node interface{}, callback func(binaryTreeIf, interface{}) bool) bool {
	n := node.(*gbtElement)

	leftistNode := func(curNode *gbtElement) (nextNode *gbtElement) {
		nextNode = curNode
		for !t.IsNil(nextNode.Right) {
			nextNode = t.Min(nextNode.Right).(*gbtElement)
		}
		return
	}

	for curNode := leftistNode(t.Min(n).(*gbtElement)); curNode != n; {
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

func newGBT() binaryTreeIf {
	t := new(gbt)
	t.Init()
	return t
}
