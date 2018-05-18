# graph.go
  
  使用linkMap实现的邻接矩阵表示
 
# bfs.go
  
  返回一个广度优先前趋子图， 前趋子图是邻接表还是邻接矩阵依赖于输入的图是邻接表还是邻接矩阵
  
# dfs.go
  返回深度优先前森林，反向边图，正向边图和交叉边图， 用栈替换了递归。用自定义linkedMap 代替map, 以解决key的顺序随机的问题。
  
  
  树边：
    白色时发现的边
    
    
  反向边：
    灰色时发现的边
    
    
  交叉边：
    黑色时发现的边，起点比终点发现的晚
    
    
  正向边：
    黑色时发现的边，终点比起点发现的晚，且时间差大于1
    

# 强连通组件

  收缩节点后，SCC留下的便为所有交叉边
  
# 双向连通组件
  点双连通和边双连通，用栈
  https://blog.csdn.net/STILLxjy/article/details/70176689

# 欧拉回路
  https://blog.csdn.net/qq_35649707/article/details/75578102#uoj117%E6%B1%82%E7%BB%99%E5%AE%9A%E5%9B%BE%E7%9A%84%E6%AC%A7%E6%8B%89%E5%9B%9E%E8%B7%AF


# 23-1 次优最小生成树
  对最小生成树每个顶点用bfs遍历，构造二维矩阵表明每两点间的最大权值，再遍历原图的所有边(出去最小生成树的边)，对最小生成树添加边，去掉边点间最大的权值的边，计算总权值，遍历后取最小的解决方案。

# 23-2 稀疏图的最小生成树
  https://blog.csdn.net/zilingxiyue/article/details/44730607
  
  但是原文说书中算法更新T不用origin是不对的，CLRS中的算法没有问题，因为mst-reduce需要多次调用，origin中的引用边在第一次之后的调用和并查集中指的边很可能不一样。
  最终的mstReducedPrim这样：
  ```go
  func mstReducedPrim(g graphWeightily, k int) graphWeightily {

	t := createGraphByType(g).(graphWeightily)

	origin := make(map[edge]edge)
	for _, e := range g.AllEdges() {
		origin[e] = e
	}

	newG := g

	for i := 0; i < k; i++ {
		newG, origin = mstReduceOnce(newG, t, origin)
	}

	newT := mstPrim(newG)

	for _, e := range newT.AllEdges() {
		t.AddEdgeWithWeight(origin[e], newT.Weight(origin[e]))
	}
	return t
}
  ```

# 23-3 瓶颈生成树
  https://stackoverflow.com/questions/22875799/how-to-compute-a-minimum-bottleneck-spanning-tree-in-linear-time
  
  线性时间中位数分组类似快排
  
  对较小部分的边+所有的点dfs后，连通分量为1，表示有冗余的边，递归所有较小部分；如果连通分量比1多， 将所有树边加入到生成树中， 用并查集收缩所有的连通分量， 对连接连通分量的边，收缩到权值最小的边，更新origin属性指向原边， 递归收缩后的图。
 
 # 单源最短路径
   Bellman Ford是用的动态规划的思想，Dijkstra用的贪心的思想
   
   SPFA是Bellman Ford的队列实现，类似于BFS
   
   Gabow是按位处理，每轮计算后将结果左移一位
   
   Karp是在原有的单元最短路径算法上加上动态规划，维护一个每条边的查找表，的测试向量：
   
   http://www.columbia.edu/~cs2035/courses/ieor6614.S16/mmc.pdf
   
 # 顶点对最短路径
   Floyd求解最短路径矩阵的矩阵表示意义是，pi[i][j]->j,即pi[i][j]为start, j为end。pi[i]向量为以i为起点的单源最短路径树。
   
 # 最大流
   残留网络中的边为cap - flow不为0的边，在residualGraph的data structure中，在更新flow时会check该值，如果为0， 删除该边。
   饱和顶点不包括s,t
   Relabel to Front的邻接表内容不能改变
