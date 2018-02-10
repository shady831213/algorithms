# 关于BST的非递归非栈的前中后序遍历
## 网上很多都不是真正的非递归非栈遍历
----------------------------
完整的代码见[binarySearchTree.go](https://github.com/shady831213/algorithms/blob/master/tree/binaryTree/binarySearchTree/binarySearchTree.go)
test见[binarySearchTree_test.go](https://github.com/shady831213/algorithms/blob/master/tree/binaryTree/binarySearchTree/binarySearchTree_test.go)
test 命令：
```
go test
```
### 中序遍历
最简单，就是从最小节点循环找后继节点
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
-----
### 后续遍历
起始节点一定是最左边最底层的叶子节点，分两种情况：
  - 最小节点
  - 最小节点有右孩子，则沿着右边找到最小节点，如果最小节点还有右孩子，一直找下去
 确定起始节点后，如果节点是左孩子，则按找起始节点的方法，找到其父亲右面的最左边最底层的叶子节点；如果是右孩子，则返回父亲节点。如此一直循环
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
 
 ###前序遍历
 最麻烦。。。
 首先其实节点一定是根节点
 整个遍历过程分两个方向，向下和向上
 一开始向下遍历，如果有左孩子就一直往左走，如果没有左孩子有右孩子就往右，总之是有孩子节点就一直往下，左面优先级高，一直走到叶子节点，方向转成向上
 整个向上的过程分成两种情况，不分叶子节点还是非叶子节点：
 -如果节点是父节点的左节点，就向上遍历直到：
    - 向上的方向改变，即一直是左节点，左节点。。。突然一个节点是父节点的右节点，这时候停
    - 或者发现有一个节点有右孩子
 -如果节点是父节点的右节点，把父节点的右孩子置为空，找到父节点的后继节点，然后恢复父节点的右孩子，这个后继节点要么是根节点，要么是一个左节点
 向上转为向下的情况有两种：
 - 节点为根节点
 - 节点有右孩子
 此时，方向转为向下，并指向其右孩子，如果右孩子为空，即循环停止。
 
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
				if curNode.right != nil {
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
  
--------
### 时间复杂度
O(n)，n个节点的bst有n-1条边，这三个遍历每条边最多走两次，一次定位一次回溯，所以复杂度为O(2(n-1))=O(n)
### 空间复杂度
真正的非递归非栈，O(1)

