# BST
-----------------

CLRS Sec12 

-----------------

no stack no recursive inOrder, preOrder and posrOrder walk

inOrder walk

```go
func (t *BstIterative) InOrderWalk(node interface{}, callback func(interface{}) (bool)) (bool) {
	n := node.(*BstElement)
	for curNode := t.Min(n).(*BstElement); curNode != nil; {
		stop := callback(curNode)
		if stop {
			return true
		}
		curNode = t.Successor(curNode).(*BstElement)
	}
	return false
}
```

postOrder walk

 ```go
func (t *BstIterative) PostOrderWalk(node interface{}, callback func(interface{}) (bool)) (bool) {
	n := node.(*BstElement)

	leftistNode := func(curNode *BstElement) (nextNode *BstElement) {
		nextNode = curNode
		for nextNode.right != nil {
			nextNode = t.Min(nextNode.right).(*BstElement)
		}
		return
	}

	for curNode := leftistNode(t.Min(n).(*BstElement)); curNode != n; {
		stop := callback(curNode)
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
	return callback(n)
}
 ```
 preOrder walk
 
  ```go
func (t *BstIterative) PreOrderWalk(node interface{}, callback func(interface{}) (bool)) (bool) {
	root := node.(*BstElement)

	goDown := func(curNode *BstElement) (*BstElement, bool) {
		if curNode.left != nil {
			return curNode.left, true
		} else if curNode.right != nil {
			return curNode.right, true
		}
		return curNode, false
	}

	goUp := func(curNode *BstElement) (*BstElement, bool) {
		if curNode == root || curNode.right != nil {
			return curNode.right, true
		} else if curNode == curNode.parent.left {
			for curNode == curNode.parent.left {
				curNode = curNode.parent
				if curNode == root || curNode.right != nil {
					return curNode.right, true
				}
			}
		} else {
			parentNode := curNode.parent
			parentRightNode := parentNode.right
			parentNode.right = nil
			curNode = t.Successor(parentNode).(*BstElement)
			parentNode.right = parentRightNode
		}
		return curNode, false
	}

	down := true
	for curNode := root; curNode != nil; {
		if down {
			stop := callback(curNode)
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
 ```
 
 [Code](https://github.com/shady831213/algorithms/blob/master/tree/binaryTree/binarySearchTree.go)
 
 [Test](https://github.com/shady831213/algorithms/blob/master/tree/binaryTree/binarySearchTree_test.go)

# Red-Black Tree
-----------------

CLRS Sec13 

-----------------

 [Code](https://github.com/shady831213/algorithms/blob/master/tree/binaryTree/rbTree.go)
 
 [Test](https://github.com/shady831213/algorithms/blob/master/tree/binaryTree/rbTree_test.go)
