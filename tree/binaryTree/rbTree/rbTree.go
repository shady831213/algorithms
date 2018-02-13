package rbTree

import (
	"algorithms/tree/binaryTree/genericBinaryTree"
)

const (
	black = true
	red   = false
)

type RBT struct {
	genericBinaryTree.GBT
}

func (t *RBT) setColor(node *genericBinaryTree.GBTElement, color bool) {
	node.SideValue = color
}

func (t *RBT) color(node *genericBinaryTree.GBTElement) (black bool) {
	return t.IsNil(node) || node.SideValue.(bool)
}

func (t *RBT) Insert(node interface{}) (interface{}) {
	n := t.GBT.Insert(node).(*genericBinaryTree.GBTElement)
	t.setColor(n, red)
	t.insertFix(n)
	return n
}

func (t *RBT) insertFix(node interface{})() {
	n := node.(*genericBinaryTree.GBTElement)
	//only can violate property 3: both left and right children of red node must be black
	for !t.color(n.Parent) && !t.color(n) {
		grandNode := n.Parent.Parent//must be black
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
			t.setColor(grandNode,red)
			if n.Parent == grandNode.Left {
				//case 2 n is right child of parent
				//no matter which side is n, case3 rotation will not violate red black tree propert.
				//but the reason why do left rotation is for the BALANCE!!
				if n == n.Parent.Right {
					t.LeftRotate(n.Parent)
				}
				//case 3 n is left child of parent
				t.setColor(grandNode.Left,black)
				t.RightRotate(grandNode)
			} else {
				if n == n.Parent.Left {
					t.RightRotate(n.Parent)
				}
				t.setColor(grandNode.Right,black)
				t.LeftRotate(grandNode)
			}
		}
	}
	t.setColor(t.Root().(*genericBinaryTree.GBTElement), black)
}

func New() *RBT {
	t := new(RBT)
	t.Init()
	t.GBT.Object = t
	return t
}
