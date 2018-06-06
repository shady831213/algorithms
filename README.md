# algorithms
[![Go Report Card](https://goreportcard.com/badge/github.com/shady831213/algorithms)](https://goreportcard.com/report/github.com/shady831213/algorithms)[![Build Status](https://travis-ci.org/shady831213/algorithms.svg?branch=master)](https://travis-ci.org/shady831213/algorithms)[![Maintainability](https://api.codeclimate.com/v1/badges/87b7c7f1222dfb1db63e/maintainability)](https://codeclimate.com/github/shady831213/algorithms/maintainability)[![Test Coverage](https://api.codeclimate.com/v1/badges/87b7c7f1222dfb1db63e/test_coverage)](https://codeclimate.com/github/shady831213/algorithms/test_coverage)

CLRS study. Codes are written with golang.

----------------

go version: 1.10.x

----------------

- [Sort](https://github.com/shady831213/algorithms/tree/master/sort)
  - [CountingSort](https://github.com/shady831213/algorithms/blob/master/sort/countingSort.go)
  - [HeapSort](https://github.com/shady831213/algorithms/blob/master/sort/heapSort.go)
  - [InsertSort](https://github.com/shady831213/algorithms/blob/master/sort/insertionSort.go)
  - [MergeSort](https://github.com/shady831213/algorithms/blob/master/sort/mergeSort.go)
  - [QuickSort](https://github.com/shady831213/algorithms/blob/master/sort/quickSort.go)
  
- [Heap](https://github.com/shady831213/algorithms/tree/master/heap)
  - [BinaryHeap on array](https://github.com/shady831213/algorithms/blob/master/heap/arrayHeap.go)
  - [BinaryHeap on linkedlist](https://github.com/shady831213/algorithms/blob/master/heap/linkedHeap.go)
  - [LeftistHeap](https://github.com/shady831213/algorithms/blob/master/heap/leftistHeap.go)
  - [FibonacciHeap](https://github.com/shady831213/algorithms/blob/master/heap/fibHeap.go)
  
- [Tree](https://github.com/shady831213/algorithms/tree/master/tree)
  - [binaryTree](https://github.com/shady831213/algorithms/tree/master/tree/binaryTree)
    - [BST](https://github.com/shady831213/algorithms/blob/master/tree/binaryTree/binarySearchTree.go)
    - [RedBlackTree](https://github.com/shady831213/algorithms/blob/master/tree/binaryTree/rbTree.go)
  - [B-Tree](https://github.com/shady831213/algorithms/tree/master/tree/bTree)
  - [RS-vEB-Tree](https://github.com/shady831213/algorithms/tree/master/tree/vEBTree)(Support single key multi value.Lazy hashtable is used to instead of array to reduce space complexity.Including Go Mixin design pattern)
  - [Disjoint-Set-Tree](https://github.com/shady831213/algorithms/tree/master/tree/disjointSetTree)
  
- [Graph](https://github.com/shady831213/algorithms/tree/master/graph) (including linkedMap, iterator)
  - [graph](https://github.com/shady831213/algorithms/blob/master/graph/graph.go)
  - [BFS](https://github.com/shady831213/algorithms/blob/master/graph/bfs.go)
  - [DFS](https://github.com/shady831213/algorithms/blob/master/graph/dfs.go)(use stack)
  - [StronglyConnectedComponents](https://github.com/shady831213/algorithms/blob/master/graph/stronglyConnectedComp.go)
  - [BioConnectedComponents](https://github.com/shady831213/algorithms/blob/master/graph/bioConnectedComp.go)(vertex bcc & edge bcc, use stack)  
  - [eulerCircuit](https://github.com/shady831213/algorithms/blob/master/graph/eulerCircuit.go)  
  - [mst](https://github.com/shady831213/algorithms/blob/master/graph/mst.go)(including Kruskal([disjointSet](https://github.com/shady831213/algorithms/tree/master/tree/disjointSetTree)) , Prim([fibonacci heap](https://github.com/shady831213/algorithms/blob/master/heap/fibHeap.go)), secondaryMst, mst reduce for Prim, linear time bottleneck spanning tree)
  - [Single-Source Shortest Path](https://github.com/shady831213/algorithms/blob/master/graph/sssp.go) (including bellmanFord, SPFA, Dijkstra, Gabow )
  - [All-Pairs Shortest Path](https://github.com/shady831213/algorithms/blob/master/graph/apsp.go) (including FloydWarshall, Johnson)
  - [Max Flow](https://github.com/shady831213/algorithms/blob/master/graph/flowGraph.go) (including flowGraph , preFlowGraph and allowedGraph data structure, Edmondes Karp, Push Relabel, Relabel to Front, Bipartite Graph Max Match and Hopcraft-Karp)
  
- [HashMap](https://github.com/shady831213/algorithms/tree/master/hashMap)(Support UpScale and DownScale)
  - [OpenAddressHashMap](https://github.com/shady831213/algorithms/blob/master/hashMap/openHashMap.go)
  - [LinkedListHashMap](https://github.com/shady831213/algorithms/blob/master/hashMap/chainedHashMap.go)
  
- [DynamicProgramming](https://github.com/shady831213/algorithms/tree/master/dp) (Including OOP pattern of golang)
  - [Double adjustable Euclidean traveling salesman](https://github.com/shady831213/algorithms/blob/master/dp/bitonicTSP.go)
  - [Pretty Print](https://github.com/shady831213/algorithms/blob/master/dp/prettyPrint.go)
  - [Levenshtein Distance](https://github.com/shady831213/algorithms/blob/master/dp/levenshteinDistance.go)
  - [Chess Game](https://github.com/shady831213/algorithms/blob/master/dp/chessGame.go)
  
- [GreedyAlgorithm](https://github.com/shady831213/algorithms/tree/master/greedy)
  - [Minimum average end time scheduling](https://github.com/shady831213/algorithms/blob/master/greedy/minAvgCompletedTimeSch.go)
