package binaryTree

type binaryTreeElement struct {
	Key   uint32
	Value interface{}
}

type binaryTreeIf interface {
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
	InOrderWalk(interface{}, func(binaryTreeIf, interface{}) bool) bool
	PreOrderWalk(interface{}, func(binaryTreeIf, interface{}) bool) bool
	PostOrderWalk(interface{}, func(binaryTreeIf, interface{}) bool) bool
}
