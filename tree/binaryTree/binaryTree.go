package binaryTree

type BinaryTreeElement struct {
	Key   uint32
	Value interface{}
}

type BinaryTreeIf interface {
	IsNil(interface{}) bool
	Root() interface{}
	Search(uint32) interface{}
	Insert(interface{}) interface{}
	Delete(uint32) interface{}
	Predecessor(interface{}, interface{}) interface{}
	Successor(interface{}, interface{}) interface{}
	LeftRotate(interface{}) interface{}
	RightRotate(interface{}) interface{}
	Min(interface{}) interface{}
	Max(interface{}) interface{}
	InOrderWalk(interface{}, func(BinaryTreeIf, interface{}) bool) bool
	PreOrderWalk(interface{}, func(BinaryTreeIf, interface{}) bool) bool
	PostOrderWalk(interface{}, func(BinaryTreeIf, interface{}) bool) bool
}
