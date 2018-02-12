package binaryTree
type BinaryTreeElement struct {
	Key uint32
	Value interface{}
}

type BinaryTreeIf interface {
	IsNil(interface{})(bool)
	Root()(interface{})
	Search(uint32)(interface{})
	Insert(interface{})
	Delete(uint32)
	Predecesor(interface{})(interface{})
	Successor(interface{})(interface{})
	LeftRotate(interface{})
	RightRotate(interface{})
	Min(interface{})(interface{})
	Max(interface{})(interface{})
	InOrderWalk(interface{}, func(BinaryTreeIf, interface{})(bool))(bool)
	PreOrderWalk(interface{}, func(BinaryTreeIf, interface{})(bool))(bool)
	PostOrderWalk(interface{}, func(BinaryTreeIf, interface{})(bool))(bool)
}
