package dp

import (
	"math"
)
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
	lDOperation
}

func (c *copy)init(cost int)(*lDOperation)  {
	op:=c.lDOperation.init("copy", cost)
	op.lDCompute = c
	return op
}

func (c *copy) updateCost(i int, j int, ldc *lDComputor) (int) {
	if ldc.seq0[i] == ldc.seq1[j] {
		return ldc.cost[i-1][j-1] + c.cost//1
	}
	return math.MaxInt32
}
func (c *copy) preOpIdx(i int, j int, ldc *lDComputor) (int, int) {
	return i-1, j-1
}

func newCopy(cost int)(*lDOperation) {
	return new(copy).init(cost)
}

type replace struct {
	lDOperation
}

func (r *replace)init(cost int)(*lDOperation)  {
	op:=r.lDOperation.init("replace", cost)
	op.lDCompute = r
	return op
}

func (r *replace) updateCost(i int, j int, ldc *lDComputor) (int) {
	if ldc.seq0[i] != ldc.seq1[j] {
		return ldc.cost[i-1][j-1] + r.cost//4
	}
	return math.MaxInt32
}

func newReplace(cost int)(*lDOperation) {
	return new(replace).init(cost)
}

func (r *replace) preOpIdx(i int, j int, ldc *lDComputor) (int, int) {
	return i-1, j-1
}

type insert struct {
	lDOperation
}

func (is *insert)init(cost int)(*lDOperation)  {
	op:=is.lDOperation.init("insert", cost)
	op.lDCompute = is
	return op
}

func (is *insert) updateCost(i int, j int, ldc *lDComputor) (int) {
	return ldc.cost[i][j-1] + is.cost//3
}

func (is *insert) preOpIdx(i int, j int, ldc *lDComputor) (int, int) {
	return i, j-1
}

func newInsert(cost int)(*lDOperation) {
	return new(insert).init(cost)
}


type delete struct {
	lDOperation
}

func (d *delete)init(cost int)(*lDOperation)  {
	op:=d.lDOperation.init("delete", cost)
	op.lDCompute = d
	return op
}

func (d *delete) updateCost(i int, j int, ldc *lDComputor) (int) {
	return ldc.cost[i-1][j] + d.cost//2
}

func (d *delete) preOpIdx(i int, j int, ldc *lDComputor) (int, int) {
	return i-1, j
}

func newDelete(cost int)(*lDOperation) {
	return new(delete).init(cost)
}

type twiddle struct {
	lDOperation
}

func (t *twiddle)init(cost int)(*lDOperation)  {
	op:=t.lDOperation.init("twiddle", cost)
	op.lDCompute = t
	return op
}

func (t *twiddle) updateCost(i int, j int, ldc *lDComputor) (int) {
	if i > 1 && j > 1 && ldc.seq0[i] == ldc.seq1[j-1] && ldc.seq0[i-1] == ldc.seq1[j] {
		return ldc.cost[i-2][j-2] + t.cost//2
	}
	return math.MaxInt32
}

func (t *twiddle) preOpIdx(i int, j int, ldc *lDComputor) (int, int) {
	return i - 2, j - 2
}

func newTwiddle(cost int)(*lDOperation) {
	return new(twiddle).init(cost)
}

type kill struct {
	lDOperation
}

func (k *kill)init(cost int)(*lDOperation)  {
	op:=k.lDOperation.init("kill", cost)
	op.lDCompute = k
	return op
}

func (t *kill) updateCost(i int, j int, ldc *lDComputor) (int) {
	return t.cost
}

func (t *kill) preOpIdx(i int, j int, ldc *lDComputor) (int, int) {
	return 0,0
}

func newKill(cost int)(*lDOperation) {
	return new(kill).init(cost)
}


//levenshteinDistance computor
type lDComputor struct {
	cost       [][]int//side data, store cost
	opSeq      [][]*lDOperation// side data, store operation
	seq0, seq1 []byte// string to array
	ops        []*lDOperation//operations set
	delete *lDOperation
	insert *lDOperation
	kill *lDOperation
}

//init
func (ldc *lDComputor) init(word0, word1 string, delete, insert, kill*lDOperation) (*lDComputor) {
	//string to array, and 1 to deal with edge logic
	ldc.seq0 = append([]byte{0}, ([]byte)(word0)...)
	ldc.seq1 = append([]byte{0}, ([]byte)(word1)...)
	ldc.ops = make([]*lDOperation,0,0)
	ldc.delete = delete
	ldc.insert = insert
	ldc.kill = kill
	ldc.ops = append(ldc.ops,[]*lDOperation{ldc.delete,ldc.insert}...)
	//initialize side datas
	ldc.cost = make([][]int, len(ldc.seq0), len(ldc.seq0))
	ldc.opSeq = make([][]*lDOperation, len(ldc.seq0), len(ldc.seq0))
	for i := range ldc.cost {
		ldc.cost[i] = make([]int, len(ldc.seq1), len(ldc.seq1))
		ldc.opSeq[i] = make([]*lDOperation, len(ldc.seq1), len(ldc.seq1))
		//j = 0 , use delete to init, means delete all chars in word0
		if i > 0 {
			ldc.cost[i][0] = ldc.delete.updateCost(i,0,ldc)
			ldc.opSeq[i][0] = ldc.delete
		}
		//i = 0, use insert to init, means insert all chars in word1
		for j := range ldc.cost[i][1:] {
			ldc.cost[i][j+1] = ldc.insert.updateCost(0,j+1,ldc)
			ldc.opSeq[i][j+1] = ldc.insert
		}
	}
	return ldc
}

func (ldc *lDComputor) addOp(op *lDOperation) {
	ldc.ops = append([]*lDOperation{op}, ldc.ops...)
}

func (ldc *lDComputor) levenshteinDistance() (int, []string) {
	//dynamic solve
	for i := 1;i < len(ldc.cost);i++{
		for j := 1;j < len(ldc.cost[i]);j++ {
			ldc.cost[i][j] = math.MaxInt32
			for op := range ldc.ops {
				temp := ldc.ops[op].updateCost(i, j, ldc)
				if temp < ldc.cost[i][j] {
					ldc.cost[i][j] = temp
					ldc.opSeq[i][j] = ldc.ops[op]
				}
			}
		}
	}

	//deal with kill, pick up the min value in len(ldc.seq1)-1 vol
	opSeq := make([]string, 0, 0)
	minDist := ldc.cost[len(ldc.seq0)-1][len(ldc.seq1)-1]
	minI := len(ldc.seq0)-1
	for i := 1; i < len(ldc.seq0) - 1; i++ {
		temp := ldc.cost[i][len(ldc.seq1)-1] + ldc.kill.updateCost(0,0,ldc)
		if  temp < minDist {
			minDist = temp
			minI = i
			opSeq = []string{ldc.kill.name}
		}
	}

	//track the seq path
	for i, j := minI, len(ldc.seq1)-1; i != 0 || j != 0; {
		op := ldc.opSeq[i][j]
		opSeq = append([]string{op.name}, opSeq...)
		i, j = op.preOpIdx(i, j, ldc)
	}
	return minDist, opSeq
}

func newLdc(word0, word1 string, delete, insert, kill*lDOperation) (*lDComputor) {
	return new(lDComputor).init(word0, word1,  delete, insert, kill)
}
