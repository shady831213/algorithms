# algorithms[![Go Report Card](https://goreportcard.com/badge/github.com/shady831213/algorithms)](https://goreportcard.com/report/github.com/shady831213/algorithms)[![Build Status](https://travis-ci.org/shady831213/algorithms.svg?branch=master)](https://travis-ci.org/shady831213/algorithms)
CLRS study. Codes are written with golang.

----------------

go version: >= 1.9.3

----------------

- [Sort](https://github.com/shady831213/algorithms/tree/master/sort) ![cover.run go](https://cover.run/go/github.com/shady831213/algorithms/sort.svg?tag=golang-1.10)
  - [CountingSort](https://github.com/shady831213/algorithms/blob/master/sort/countingSort.go)
  - [HeapSort](https://github.com/shady831213/algorithms/blob/master/sort/heapSort.go)
  - [InsertSort](https://github.com/shady831213/algorithms/blob/master/sort/insertionSort.go)
  - [MergeSort](https://github.com/shady831213/algorithms/blob/master/sort/mergeSort.go)
  - [QuickSort](https://github.com/shady831213/algorithms/blob/master/sort/quickSort.go)
  
- [Heap](https://github.com/shady831213/algorithms/tree/master/heap)
![cover.run go](https://cover.run/go/github.com/shady831213/algorithms/heap.svg?tag=golang-1.10)
  - [BinaryHeap on array](https://github.com/shady831213/algorithms/blob/master/heap/arrayHeap.go)
  - [BinaryHeap on linkedlist](https://github.com/shady831213/algorithms/blob/master/heap/linkedHeap.go)
  - [LeftistHeap](https://github.com/shady831213/algorithms/blob/master/heap/leftistHeap.go)
  - [FibonacciHeap](https://github.com/shady831213/algorithms/blob/master/heap/fibHeap.go)
  
- [Tree](https://github.com/shady831213/algorithms/tree/master/tree)
  - [binaryTree](https://github.com/shady831213/algorithms/tree/master/tree/binaryTree)
  ![cover.run go](https://cover.run/go/github.com/shady831213/algorithms/tree/binaryTree.svg?tag=golang-1.10)
    - [BST](https://github.com/shady831213/algorithms/blob/master/tree/binaryTree/binarySearchTree.go)
    - [RedBlackTree](https://github.com/shady831213/algorithms/blob/master/tree/binaryTree/rbTree.go)
  - [B-Tree](https://github.com/shady831213/algorithms/tree/master/tree/bTree)
  ![cover.run go](https://cover.run/go/github.com/shady831213/algorithms/tree/bTree.svg?tag=golang-1.10)
  - [RS-vEB-Tree](https://github.com/shady831213/algorithms/tree/master/tree/vEBTree)(Support single key multi value.Lazy hashtable is used to instead of array to reduce space complexity.Including Go Mixin design pattern)
  ![cover.run go](https://cover.run/go/github.com/shady831213/algorithms/tree/vEBTree.svg?tag=golang-1.10)
  - [Disjoint-Set-Tree](https://github.com/shady831213/algorithms/tree/master/tree/disjointSetTree)
  ![cover.run go](https://cover.run/go/github.com/shady831213/algorithms/tree/disjointSetTree.svg?tag=golang-1.10)
  
- [Graph](https://github.com/shady831213/algorithms/tree/master/graph) (including linkedMap, iterator)
![cover.run go](https://cover.run/go/github.com/shady831213/algorithms/graph.svg?tag=golang-1.10)
  - [AdjacencyMatrix and AdjacencyList](https://github.com/shady831213/algorithms/blob/master/graph/graph.go)
  - [BFS](https://github.com/shady831213/algorithms/blob/master/graph/bfs.go)
  - [DFS](https://github.com/shady831213/algorithms/blob/master/graph/dfs.go)(use stack)
  - [StronglyConnectedComponents](https://github.com/shady831213/algorithms/blob/master/graph/stronglyConnectedComp.go)
  - [BioConnectedComponents](https://github.com/shady831213/algorithms/blob/master/graph/bioConnectedComp.go)(vertex bcc & edge bcc, use stack)  
  - [eulerCircuit](https://github.com/shady831213/algorithms/blob/master/graph/eulerCircuit.go)  
  - [mst](https://github.com/shady831213/algorithms/blob/master/graph/mst.go)(including Kruskal([disjointSet](https://github.com/shady831213/algorithms/tree/master/tree/disjointSetTree)) , Prim([fibonacci heap](https://github.com/shady831213/algorithms/blob/master/heap/fibHeap.go)), secondaryMst, mst reduce for Prim, linear time bottleneck spanning tree)
  - [Single-Source Shortest Path](https://github.com/shady831213/algorithms/blob/master/graph/sssp.go) (including bellmanFord, SPFA, Dijkstra, Gabow )
  
- [HashMap](https://github.com/shady831213/algorithms/tree/master/hashMap)(Support UpScale and DownScale)
![cover.run go](https://cover.run/go/github.com/shady831213/algorithms/hashMap.svg?tag=golang-1.10)
  - [OpenAddressHashMap](https://github.com/shady831213/algorithms/blob/master/hashMap/openHashMap.go)
  - [LinkedListHashMap](https://github.com/shady831213/algorithms/blob/master/hashMap/chainedHashMap.go)
  
- [DynamicProgramming](https://github.com/shady831213/algorithms/tree/master/dp) (Including OOP pattern of golang)
![cover.run go](https://cover.run/go/github.com/shady831213/algorithms/dp.svg?tag=golang-1.10)
  - [Double adjustable Euclidean traveling salesman](https://github.com/shady831213/algorithms/blob/master/dp/bitonicTSP.go)
  - [Pretty Print](https://github.com/shady831213/algorithms/blob/master/dp/prettyPrint.go)
  - [Levenshtein Distance](https://github.com/shady831213/algorithms/blob/master/dp/levenshteinDistance.go)
  - [Chess Game](https://github.com/shady831213/algorithms/blob/master/dp/chessGame.go)
  
- [GreedyAlgorithm](https://github.com/shady831213/algorithms/tree/master/greedy)
![cover.run go](https://cover.run/go/github.com/shady831213/algorithms/greedy.svg?tag=golang-1.10)
  - [Minimum average end time scheduling](https://github.com/shady831213/algorithms/blob/master/greedy/minAvgCompletedTimeSch.go)
