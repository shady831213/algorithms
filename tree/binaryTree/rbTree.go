package binaryTree

const (
	black = true
	red   = false
	left  = true
)

type rbt struct {
	gbt
}

func (t *rbt) setColor(node *gbtElement, color bool) {
	node.SideValue = color
}

func (t *rbt) color(node *gbtElement) (black bool) {
	return t.IsNil(node) || node.SideValue.(bool)
}

func (t *rbt) otherSideNode(side bool, node *gbtElement) *gbtElement {
	if side == left {
		return node.Right
	}
	return node.Left
}
func (t *rbt) invDirRotation(side bool, node *gbtElement) interface{} {
	if side == left {
		return t.RightRotate(node)
	}
	return t.LeftRotate(node)
}
func (t *rbt) sameSideNode(side bool, node *gbtElement) *gbtElement {
	if side == left {
		return node.Left
	}
	return node.Right
}

func (t *rbt) sameDirRotation(side bool, node *gbtElement) interface{} {
	if side == left {
		return t.LeftRotate(node)
	}
	return t.RightRotate(node)
}

func (t *rbt) Insert(node interface{}) interface{} {
	n := t.gbt.Insert(node).(*gbtElement)
	t.setColor(n, red)
	t.insertFix(n)
	return n
}

func (t *rbt) insertFix(node interface{}) {
	n := node.(*gbtElement)
	//only can violate property 3: both left and right children of red node must be black
	for !t.color(n.Parent) && !t.color(n) {
		grandNode := n.Parent.Parent //must be black
		uncleNode := grandNode.Right
		if n.Parent == uncleNode {
			uncleNode = grandNode.Left
		}
		//case1: uncle node is red
		if !t.color(uncleNode) {
			t.setColor(grandNode, red)
			t.setColor(grandNode.Left, black)
			t.setColor(grandNode.Right, black)
			n = grandNode
			//case2&3: uncle node is black
		} else {
			side := n.Parent == grandNode.Left
			t.setColor(grandNode, red)
			//case 2 n is right child of parent
			if n == t.otherSideNode(side, n.Parent) {
				t.sameDirRotation(side, n.Parent)
			}
			//case 3 n is left child of parent
			t.setColor(t.sameSideNode(side, grandNode), black)
			t.invDirRotation(side, grandNode)
		}
	}
	t.setColor(t.Root().(*gbtElement), black)
}

func (t *rbt) Delete(key uint32) interface{} {
	deleteNonCompletedNode := func(node *gbtElement) (deletedNode *gbtElement, nextNode *gbtElement) {
		var reConnectedNode *gbtElement
		if t.IsNil(node.Left) {
			reConnectedNode = node.Right
		} else {
			reConnectedNode = node.Left
		}
		//mean's another black color
		reConnectedNode.Parent = node.Parent
		if t.IsNil(node.Parent) {
			t.NilNode.Left = reConnectedNode
			t.NilNode.Right = reConnectedNode
		} else if node.Parent.Right == node {
			node.Parent.Right = reConnectedNode
		} else {
			node.Parent.Left = reConnectedNode
		}
		return node, reConnectedNode
	}
	node := t.Search(key).(*gbtElement)
	if t.IsNil(node) {
		return node
	}
	var deletedNode, reConnectedNode *gbtElement
	if t.IsNil(node.Left) || t.IsNil(node.Right) {
		deletedNode, reConnectedNode = deleteNonCompletedNode(node)
	} else {
		successor := t.Successor(node, t.Root()).(*gbtElement)
		_key, _value := successor.Key, successor.Value
		node.Key, node.Value = _key, _value
		deletedNode, reConnectedNode = deleteNonCompletedNode(successor)
	}
	if t.color(deletedNode) {
		//Now, reConnectedNode is black-black or black-red
		t.deleteFix(reConnectedNode)
	}
	//recover NilNode
	t.NilNode.Parent = t.NilNode
	return node
}

func (t *rbt) deleteFix(node interface{}) {
	n := node.(*gbtElement)
	//n always points to the black-black or black-red node.The purpose is to remove the additional black color,
	//which means add a black color in the same side or reduce a black color in the other side
	for n != t.Root() && t.color(n) {
		side := n == n.Parent.Left
		brotherNode := t.otherSideNode(side, n.Parent)
		//case 1 brotherNode node is red, so parent must be black.Turn brotherNode node to a black one, convert to case 2,3,4
		if !t.color(brotherNode) {
			t.setColor(n.Parent, red)
			t.setColor(brotherNode, black)
			t.sameDirRotation(side, n.Parent)
			//case 2, 3, 4 brotherNode node is black
		} else {
			//case 2 move black-blcak or blcak-red node up
			if t.color(brotherNode.Left) && t.color(brotherNode.Right) {
				t.setColor(brotherNode, red)
				n = n.Parent
				//case 3 convert to case 4
			} else if t.color(t.otherSideNode(side, brotherNode)) {
				t.setColor(brotherNode, red)
				t.setColor(t.sameSideNode(side, brotherNode), black)
				t.invDirRotation(side, brotherNode)
				//case 4 add a black to left, turn black-black or black-red to black or red
			} else {
				t.setColor(brotherNode, t.color(n.Parent))
				t.setColor(n.Parent, black)
				t.setColor(t.otherSideNode(side, brotherNode), black)
				t.sameDirRotation(side, n.Parent)
				n = t.Root().(*gbtElement)
			}
		}

	}
	t.setColor(n, black)

}

func newRBT() *rbt {
	t := new(rbt)
	t.Init()
	t.gbt.Object = t
	return t
}
