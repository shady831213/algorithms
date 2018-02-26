package dp

import (
	"testing"
	"fmt"
	"math"
)

func TestBitonicTSP(t *testing.T)  {
	points := []*point{&point{0,6}, &point{1,0},&point{2,3},
	&point{5,4},&point{6,1},&point{7,5},&point{8,2}}
	result := bitonicTSP(points)
	if math.Abs(result - 25.584025) >= 0.0001 {
		t.Log(fmt.Sprintf("expect 25.584025, but get %f", result))
		t.Fail()
	}
}