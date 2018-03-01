package dp

import (
	"testing"
	"fmt"
	"reflect"
)

func TestLevenshteinDistance(t *testing.T)  {
	ldc := newLdc("algorithm","altruistic")
	ldc.addOp(newTwiddle())
	ldc.addOp(newCopy())
	ldc.addOp(newReplace())
	dist,opSeq := ldc.levenshteinDistance()
	if dist != 24 {
		t.Log(fmt.Sprintf("dist expect 24, but get %d", dist))
		t.Fail()
	}
	if !reflect.DeepEqual(opSeq, []string{"copy","copy","delete","replace","copy","insert", "insert","insert", "twiddle", "insert", "kill"}) {
		t.Log("opSeq wrong!\n")
		t.Log(opSeq)
		t.Fail()
	}
}
