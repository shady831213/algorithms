# 双调欧几里得旅行商问题
--------
## 问题：
平面上n个点，确定一条连接各点的最短闭合旅程。这个解的一般形式为NP的（在多项式时间内可以求出）。

J.L. Bentley 建议通过只考虑双调旅程(bitonic tours)来简化问题,这种旅程即为从最左点开始，严格地从左到右直至最右点，然后严格地从右到左直至出发点。下图(b)显示了同样的7个点的最短双调路线。在这种情况下，多项式的算法是可能的。事实上，存在确定的最优双调路线的O(n*n)时间的算法。

闭合巡游路线
![](https://github.com/shady831213/algorithms/blob/master/dp/static/dp1.PNG)

（a）图是最短的闭合旅程，长度为24.89。（b）图是问题经简化后，同样的点集的最短双调闭合旅程，长度为25.58。

--------
[代码](https://github.com/shady831213/algorithms/blob/master/dp/bitonicTSP.go)

[测试](https://github.com/shady831213/algorithms/blob/master/dp/bitonicTSP_test.go)

--------
## 思路：
建立一个二维数组，维度一代表总共的点数，维度二代表总共这么多点时，每个点到最后的点的最小距离，有最佳子结构：
i < numPoints - 1 :
D[numPoints][i] = D[numPoints-1][i] + d(p[numPoints] - p[numPoints-1])
这时，新的距离等于以前的距离加上新的点和倒数第二个点构成的线段

i == numPoints - 1:
D[numPoints][numPoints - 1] = min(D[numPoints-1][k] + d(p[k] - p[numPoints-1])) 0<=k<numPoints-1
这时，由于最新的点不能直接连倒数第二个点，必须绕一圈。要把最新的点连到除去倒数第二个点的某个点，取最小距离

i == numPoints:
D[numPoints][numPoints] = D[numPoints][numPoints-1] + d(p[numPoints-1] - p[numPoints]))
如果i和numPoints相同，表示新的点绕了一圈回到自己，也就是从倒数第二个点绕一圈到最新的点再加上二者构成的线段

## 算法步骤：
![](https://github.com/shady831213/algorithms/blob/master/dp/static/ph2.PNG)

![](https://github.com/shady831213/algorithms/blob/master/dp/static/ph3.PNG)

![](https://github.com/shady831213/algorithms/blob/master/dp/static/ph4.PNG)

![](https://github.com/shady831213/algorithms/blob/master/dp/static/ph5.PNG)


--------
# 整齐打印问题
--------
## 问题：
考虑在一个打印机上整齐地打印一段文章的问题。输入的正文是n个长度分别为L1、L2、……、Ln（以字符个数度量）的单词构成的序列。我们希望将这个段落在一些行上整齐地打印出来，每行至多M个字符。“整齐度”的标准如下：如果某一行包含从i到j的单词（i<j），且单词之间只留一个空格，则在行末多余的空格字符个数为 M - (j-i) - (Li+ …… + Lj)，它必须是非负值才能让该行容纳这些单词。我们希望所有行（除最后一行）的行末多余空格字符个数的立方和最小。请给出一个动态规划的算法，来在打印机整齐地打印一段又n个单词的文章。分析所给算法的执行时间和空间需求。

--------
[代码](https://github.com/shady831213/algorithms/blob/master/dp/prettyPrint.go)

[测试](https://github.com/shady831213/algorithms/blob/master/dp/prettyPrint_test.go)

--------
## 思路：
该问题和矩阵乘法链问题相似，从i到j的单词，如果超过1行的容量，分裂为子问题，1行以内直接求解：
M - (j-i) - (Li+...+Lj) >= 0 :
alignedIdx[i][j] = (M - (j-i) - (Li+...+Lj))^3

M - (j-i) - (Li+...+Lj) < 0 :
alignedIdx[i][j] = min(alignedIdx[i][k], alignedIdx[k+1][j]) i<=k<j

需要注意，最外层循环要循环长度，即j-i，这样才能从长度0，1...max自底向上构建自问题

