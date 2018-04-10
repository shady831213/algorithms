package dp

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestLevenshteinDistance(t *testing.T) {
	ldc := newLdc("algorithm", "altruistic", &newDelete(2).lDOperation, &newInsert(3).lDOperation, &newKill(1).lDOperation)
	ldc.addOp(&newTwiddle(2).lDOperation)
	ldc.addOp(&newCopy(1).lDOperation)
	ldc.addOp(&newReplace(4).lDOperation)
	dist, opSeq := ldc.levenshteinDistance()
	if dist != 24 {
		t.Log(fmt.Sprintf("dist expect 24, but get %d", dist))
		t.Fail()
	}
	if !reflect.DeepEqual(opSeq, []string{"copy", "copy", "delete", "replace", "copy", "insert", "insert", "insert", "twiddle", "insert", "kill"}) {
		t.Log("opSeq wrong!\n")
		t.Log(opSeq)
		t.Fail()
	}
}

func TestGeneAlign(t *testing.T) {
	ldc := newLdc("GATCGGCAT", "CAATGTGAATC", &newDelete(2).lDOperation, &newInsert(2).lDOperation, &newKill(math.MaxInt32).lDOperation)
	ldc.addOp(&newCopy(-1).lDOperation)
	ldc.addOp(&newReplace(1).lDOperation)
	dist, opSeq := ldc.levenshteinDistance()
	if dist != 3 {
		t.Log(fmt.Sprintf("dist expect 3, but get %d", dist))
		t.Fail()
	}
	if !reflect.DeepEqual(opSeq, []string{"insert", "replace", "copy", "copy", "replace", "replace", "copy", "replace", "copy", "copy", "insert"}) {
		t.Log("opSeq wrong!\n")
		t.Log(opSeq)
		t.Fail()
	}
}
