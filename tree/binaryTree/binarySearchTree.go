package binaryTree

type BstElement struct {
	parent, left, right *BstElement
	key uint32
	value interface{}
}

type Bst struct {
	root *BstElement
}

type BstIf interface {
	Insert(uint32, interface{})
	Delete(uint32)
	Search(uint32)(*BstElement)
	Predecesor(uint32)(*BstElement)
	Successor(uint32)(*BstElement)
	Min(*BstElement)(*BstElement)
	Max(*BstElement)(*BstElement)
	InOrderWalk(*BstElement, func(*BstElement))
	InOrderNext(*BstElement)(*BstElement)
	PreOrderWalk(*BstElement, func(*BstElement))
	PreOrderNext(*BstElement)(*BstElement)
	PostOrderWalk(*BstElement, func(*BstElement))
	PostOrderNext(*BstElement)(*BstElement)
}