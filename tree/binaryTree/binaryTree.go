package binaryTree
type BinaryTreeElement struct {
	Key uint32
	Value interface{}
}

type BinaryTreeIf interface {
	Search(uint32)(interface{})
	Insert(interface{})
	Delete(uint32)
	Predecesor(interface{})(interface{})
	Successor(interface{})(interface{})
	Min(interface{})(interface{})
	Max(interface{})(interface{})
	InOrderWalk(interface{}, func(interface{})(bool))(bool)
	//PreOrderWalk(func(interface{}, ...interface{}), ...interface{})
	//PostOrderWalk(func(interface{}, ...interface{}), ...interface{})
}
