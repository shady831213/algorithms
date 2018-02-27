package dp

import (
	"testing"
	"fmt"
	"math"
	"reflect"
)

func TestBitonicTSP(t *testing.T)  {
	points := []*point{newPoint(0,6), newPoint(1,0),newPoint(2,3),
		newPoint(5,4),newPoint(6,1),newPoint(7,5),newPoint(8,2)}
	expPath := []*line{newLine(points[0], points[1]), newLine(points[0], points[2]), newLine(points[2], points[3]),
		newLine(points[1], points[4]), newLine(points[3], points[5]), newLine(points[4], points[6]), newLine(points[5], points[6])}
	result, path := bitonicTSP(points)
	if math.Abs(result - 25.584025) >= 0.0001 {
		t.Log(fmt.Sprintf("expect 25.584025, but get %f", result))
		t.Fail()
	}
	if !reflect.DeepEqual(path, expPath) {
		t.Log("Path not match!")
		t.Fail()
	}
}