package dp

import (
	"math"
)

type lDOperation interface {
	name() string
	cost(int, int, *lDComputor) (int)
	preOpIdx(int, int, *lDComputor) (int, int)
}

type copy struct {
}

func (c *copy) name() (string) {
	return "copy"
}

func (c *copy) cost(i int, j int, ldc *lDComputor) (int) {
	if ldc.seq0[i] == ldc.seq1[j] {
		return ldc.cost[i-1][j-1] + 1
	}
	return math.MaxInt32
}
func (c *copy) preOpIdx(i int, j int, ldc *lDComputor) (int, int) {
	return i-1, j-1
}

type replace struct {
}

func (c *replace) name() (string) {
	return "replace"
}

func (r *replace) cost(i int, j int, ldc *lDComputor) (int) {
	if ldc.seq0[i] != ldc.seq1[j] {
		return ldc.cost[i-1][j-1] + 4
	}
	return math.MaxInt32
}

func (r *replace) preOpIdx(i int, j int, ldc *lDComputor) (int, int) {
	return i-1, j-1
}

type insert struct {
}

func (is *insert) name() (string) {
	return "insert"
}

func (is *insert) cost(i int, j int, ldc *lDComputor) (int) {
	return ldc.cost[i][j-1] + 3
}

func (is *insert) preOpIdx(i int, j int, ldc *lDComputor) (int, int) {
	return i, j-1
}

type delete struct {
}

func (c *delete) name() (string) {
	return "delete"
}

func (d *delete) cost(i int, j int, ldc *lDComputor) (int) {
	return ldc.cost[i-1][j] + 2
}

func (d *delete) preOpIdx(i int, j int, ldc *lDComputor) (int, int) {
	return i-1, j
}

type twiddle struct {
}

func (c *twiddle) name() (string) {
	return "twiddle"
}

func (t *twiddle) cost(i int, j int, ldc *lDComputor) (int) {
	if i > 1 && j > 1 && ldc.seq0[i] == ldc.seq1[j-1] && ldc.seq0[i-1] == ldc.seq1[j] {
		return ldc.cost[i-2][j-2] + 2
	}
	return math.MaxInt32
}

func (t *twiddle) preOpIdx(i int, j int, ldc *lDComputor) (int, int) {
	return i - 2, j - 2
}

type kill struct {
}

func (k *kill) name() (string) {
	return "kill"
}

func (t *kill) cost(i int, j int, ldc *lDComputor) (int) {
	return 1
}

func (t *kill) preOpIdx(i int, j int, ldc *lDComputor) (int, int) {
	return 0,0
}

type lDComputor struct {
	cost       [][]int
	opSeq      [][]lDOperation
	seq0, seq1 []byte
	ops        []lDOperation
	delete *delete
	insert *insert
	kill *kill
}

func (ldc *lDComputor) Init(word0, word1 string) (*lDComputor) {
	ldc.seq0 = append([]byte{0}, ([]byte)(word0)...)
	ldc.seq1 = append([]byte{0}, ([]byte)(word1)...)
	ldc.ops = make([]lDOperation,0,0)
	ldc.delete = new(delete)
	ldc.insert = new(insert)
	ldc.kill = new(kill)
	ldc.ops = append(ldc.ops,[]lDOperation{ldc.delete,ldc.insert}...)
	ldc.cost = make([][]int, len(ldc.seq0), len(ldc.seq0))
	ldc.opSeq = make([][]lDOperation, len(ldc.seq0), len(ldc.seq0))
	for i := range ldc.cost {
		ldc.cost[i] = make([]int, len(ldc.seq1), len(ldc.seq1))
		ldc.opSeq[i] = make([]lDOperation, len(ldc.seq1), len(ldc.seq1))
		if i > 0 {
			ldc.cost[i][0] = ldc.delete.cost(i,0,ldc)
			ldc.opSeq[i][0] = ldc.delete
		}
		for j := range ldc.cost[i][1:] {
			ldc.cost[i][j+1] = ldc.insert.cost(0,j+1,ldc)
			ldc.opSeq[i][j+1] = ldc.insert
		}
	}
	return ldc
}

func (ldc *lDComputor) addOp(op lDOperation) {
	ldc.ops = append([]lDOperation{op}, ldc.ops...)
}

func (ldc *lDComputor) levenshteinDistance() (int, []string) {
	for i := 1;i < len(ldc.cost);i++{
		for j := 1;j < len(ldc.cost[i]);j++ {
			ldc.cost[i][j] = math.MaxInt32
			for op := range ldc.ops {
				temp := ldc.ops[op].cost(i, j, ldc)
				if temp < ldc.cost[i][j] {
					ldc.cost[i][j] = temp
					ldc.opSeq[i][j] = ldc.ops[op]
				}
			}
		}
	}

	opSeq := make([]string, 0, 0)
	minDist := ldc.cost[len(ldc.seq0)-1][len(ldc.seq1)-1]
	minI := len(ldc.seq0)-1
	for i := 1; i < len(ldc.seq0) - 1; i++ {
		temp := ldc.cost[i][len(ldc.seq1)-1] + ldc.kill.cost(0,0,ldc)
		if  temp < minDist {
			minDist = temp
			minI = i
			opSeq = []string{ldc.kill.name()}
		}
	}

	for i, j := minI, len(ldc.seq1)-1; i != 0 || j != 0; {
		op := ldc.opSeq[i][j]
		opSeq = append([]string{op.name()}, opSeq...)
		i, j = op.preOpIdx(i, j, ldc)
	}
	return minDist, opSeq
}

func newLdc(word0, word1 string) (*lDComputor) {
	return new(lDComputor).Init(word0, word1)
}
