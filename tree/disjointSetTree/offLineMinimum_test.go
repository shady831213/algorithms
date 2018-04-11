package disjointSetTree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestOfflineMinimum(t *testing.T) {
	seq := []int{4, 8,
		offlineminimumExtract,
		3,
		offlineminimumExtract,
		9, 2, 6,
		offlineminimumExtract,
		offlineminimumExtract,
		offlineminimumExtract,
		1, 7,
		offlineminimumExtract,
		5}
	exp := []int{4, 3, 2, 6, 8, 1}
	extractSeq := offLineMinimum(seq)
	if !reflect.DeepEqual(extractSeq, exp) {
		t.Log(fmt.Sprintf("expct :%+v, actual:%+v", exp, extractSeq))
		t.Fail()
	}
}
