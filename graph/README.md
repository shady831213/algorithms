# graph.go
  
  graph data structure. Adjacency Matrix implemented by linked map
  
  --------------
  
  CLRS 22.1
  
  --------------
  
  [Code](https://github.com/shady831213/algorithms/blob/master/graph/graph.go)
  
  [Test](https://github.com/shady831213/algorithms/blob/master/graph/graph_test.go)
  
# bfs
  --------------
  
  CLRS 22.2
  
  --------------
  
  [Code](https://github.com/shady831213/algorithms/blob/master/graph/bfs.go)
  
  [Test](https://github.com/shady831213/algorithms/blob/master/graph/bfs_test.go)
  
# dfs
  --------------
  
  CLRS 22.3
  
  --------------
  Use stack to insead of recursive function
    
  [Code](https://github.com/shady831213/algorithms/blob/master/graph/dfs.go)
  
  [Test](https://github.com/shady831213/algorithms/blob/master/graph/dfs_test.go)
  
# StronglyConnectedComponents
  --------------
  
  CLRS 22.5
  
  --------------

  [Code](https://github.com/shady831213/algorithms/blob/master/graph/stronglyConnectedComp.go)
  
  [Test](https://github.com/shady831213/algorithms/blob/master/graph/stronglyConnectedComp_test.go)
  
  
# BioConnectedComponents
  --------------
  
  CLRS 22-2
  
  --------------
  Including VertexBCC & EdgeBCC
  
  [Code](https://github.com/shady831213/algorithms/blob/master/graph/bioConnectedComp.go)
  
  [Test](https://github.com/shady831213/algorithms/blob/master/graph/bioConnectedComp_test.go)

# eulerCircuit
  --------------
  
  CLRS 22-3
  
  --------------

  [Code](https://github.com/shady831213/algorithms/blob/master/graph/eulerCircuit.go)
  
  [Test](https://github.com/shady831213/algorithms/blob/master/graph/eulerCircuit_test.go)

# mst
  --------------
  
  CLRS Sec23
  
  --------------
  --------------
  ## Including mstKruskal and mstPrim.
  
   Dependencies:
  
   [disjointSet](https://github.com/shady831213/algorithms/blob/master/tree/disjointSetTree/disjointSetTree.go)  in [github.com/shady831213/algorithms/tree/disjointSetTree](https://github.com/shady831213/algorithms/tree/master/tree/disjointSetTree)
  
   [fibHeap](https://github.com/shady831213/algorithms/blob/master/heap/fibHeap.go) in [github.com/shady831213/algorithms/heap](https://github.com/shady831213/algorithms/tree/master/heap)
    
  
   [Code-mstKruskal&mstPrim](https://github.com/shady831213/algorithms/blob/master/graph/mst.go)
  
   [Test-mstKruskal&mstPrim](https://github.com/shady831213/algorithms/blob/master/graph/mst_test.go)
   
  ----------------
  ## Including secondaryMst(CLRS 32-1)
  
   [Code-secondaryMst](https://github.com/shady831213/algorithms/blob/master/graph/mst.go)
  
   [Test-secondaryMst](https://github.com/shady831213/algorithms/blob/master/graph/mst_test.go)
   
  -----------------
  ## Including mst reduce for Prim (CLRS 32-2)
  
   [Code-mstReducedPrim](https://github.com/shady831213/algorithms/blob/master/graph/mst.go)
  
   [Test-mstReducedPrim](https://github.com/shady831213/algorithms/blob/master/graph/mst_test.go)
   
   -----------------
  ## Including Linear time bottleneck spanning tree (CLRS 32-3)
  
  https://stackoverflow.com/questions/22875799/how-to-compute-a-minimum-bottleneck-spanning-tree-in-linear-time
  
   [Code-bottleNeckSpanningTree](https://github.com/shady831213/algorithms/blob/master/graph/mst.go)
  
   [Test-bottleNeckSpanningTree](https://github.com/shady831213/algorithms/blob/master/graph/mst_test.go)
  

# Single-Source Shortest Path
  --------------
  
  CLRS Sec24
  
  --------------
  ## Including Bellman Ford and Dijkstra.
   Dependencies:
  
   [fibHeap](https://github.com/shady831213/algorithms/blob/master/heap/fibHeap.go) in [github.com/shady831213/algorithms/heap](https://github.com/shady831213/algorithms/tree/master/heap)
   
   [Code-bellmanFord&dijkstra](https://github.com/shady831213/algorithms/blob/master/graph/sssp.go)
  
   [Test-bellmanFord&dijkstra](https://github.com/shady831213/algorithms/blob/master/graph/sssp_test.go)
   
  ## Including SPFA.
   [Shortest_Path_Faster_Algorithm](https://en.wikipedia.org/wiki/Shortest_Path_Faster_Algorithm)
   
   [Code-spfa](https://github.com/shady831213/algorithms/blob/master/graph/sssp.go)
  
   [Test-spfa](https://github.com/shady831213/algorithms/blob/master/graph/sssp_test.go)
   
  ## Including Gabow (CLRS 24-4)
   
   [Code-gabow](https://github.com/shady831213/algorithms/blob/master/graph/sssp.go)
  
   [Test-babow](https://github.com/shady831213/algorithms/blob/master/graph/sssp_test.go)
   
  ## Including Karp (CLRS 24-5)
   Test Vector comes from : http://www.columbia.edu/~cs2035/courses/ieor6614.S16/mmc.pdf
   
   [Code-karp](https://github.com/shady831213/algorithms/blob/master/graph/sssp.go)
  
   [Test-karp](https://github.com/shady831213/algorithms/blob/master/graph/sssp_test.go)
   
  ## Including Nested Boxes Problem (CLRS 24-2)
  
   [Code-nestedBoxes](https://github.com/shady831213/algorithms/blob/master/graph/sssp.go)
  
   [Test-nestedBoxes](https://github.com/shady831213/algorithms/blob/master/graph/sssp_test.go)
   
  
# All-Pairs Shortest Path
  --------------
  
  CLRS Sec25
  
  --------------
  ## Including FloydWarshall and Johnson.
   [Code](https://github.com/shady831213/algorithms/blob/master/graph/apsp.go)
  
   [Test](https://github.com/shady831213/algorithms/blob/master/graph/apsp_test.go)
