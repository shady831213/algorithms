# Double adjustable Euclidean traveling salesman
--------
[Code](https://github.com/shady831213/algorithms/blob/master/dp/bitonicTSP.go)

[Test](https://github.com/shady831213/algorithms/blob/master/dp/bitonicTSP_test.go)

--------
## Solution：
![](https://github.com/shady831213/algorithms/blob/master/dp/static/ph2.PNG)

![](https://github.com/shady831213/algorithms/blob/master/dp/static/ph3.PNG)

![](https://github.com/shady831213/algorithms/blob/master/dp/static/ph4.PNG)

![](https://github.com/shady831213/algorithms/blob/master/dp/static/ph5.PNG)


--------
# Pretty Print
--------
[Code](https://github.com/shady831213/algorithms/blob/master/dp/prettyPrint.go)

[Test](https://github.com/shady831213/algorithms/blob/master/dp/prettyPrint_test.go)

--------
## Solution：
M - (j-i) - (Li+...+Lj) >= 0 :
alignedIdx[i][j] = (M - (j-i) - (Li+...+Lj))^3

M - (j-i) - (Li+...+Lj) < 0 :
alignedIdx[i][j] = min(alignedIdx[i][k], alignedIdx[k+1][j]) i<=k<j



--------
# Levenshtein Distance
--------
[Code](https://github.com/shady831213/algorithms/blob/master/dp/levenshteinDistance.go)

[Test](https://github.com/shady831213/algorithms/blob/master/dp/levenshteinDistance_test.go)

--------
## Solution：
![](https://github.com/shady831213/algorithms/blob/master/dp/static/编辑距离4.gif)
c[i][j] = MIN(c[m,n], MIN(c[i,n]+cost(kill)))， 0<=i

Gene alignment
copy : -1
replace : 1
insert : 2
delete ： 2
twiddle : Max
kill ： Max


## OOP pattern：
```go
//base method and data
type lDCompute interface {
	updateCost(int, int, *lDComputor) (int)
	preOpIdx(int, int, *lDComputor) (int, int)
}

type lDOperation struct {
	name string
	cost int
	lDCompute
}

func (op *lDOperation)init(name string, cost int)(*lDOperation)  {
	op.name = name
	op.cost = cost
	return op
}

//operations
type copy struct {
	lDOperation//base
}

func (c *copy)init(cost int)(*copy)  {
	op:=c.lDOperation.init("copy", cost)
	op.lDCompute = c
	return c
}

func (c *copy) updateCost(i int, j int, ldc *lDComputor) (int) {
	//...
}
func (c *copy) preOpIdx(i int, j int, ldc *lDComputor) (int, int) {
	//...
}

func newCopy(cost int)(*copy) {
	return new(copy).init(cost)
}

//object
copyObj := newCopy(1)

//base type pointer
var op *lDOperation
op = &newCopy(1).lDOperation
```


--------
# Chess Game
--------
[Code](https://github.com/shady831213/algorithms/blob/master/dp/chessGame.go)

[Test](https://github.com/shady831213/algorithms/blob/master/dp/chessGame_test.go)

--------
## Solution：

score[i][j] = max(score[i-1][j-1],score[i-1][j],score[i-1][j+1]) + board[i][j]


![](https://github.com/shady831213/algorithms/blob/master/dp/static/棋.PNG)
