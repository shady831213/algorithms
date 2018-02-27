package dp

import (
	"math"
	"sort"
	"fmt"
)

type point struct {
	x, y float64
}

func newPoint(x, y float64) *point {
	return &point{x, y}
}

type line struct {
	p0,p1 *point
}

func newLine(p0,p1 *point) *line {
	return &line{p0,p1}
}

func (l *line) dist() (float64) {
	return math.Sqrt(math.Pow(l.p0.x-l.p1.x, 2) + math.Pow(l.p0.y-l.p1.y, 2))
}

func (l *line) print() {
	fmt.Println("----")
	fmt.Println(l.p0)
	fmt.Println(l.p1)
	fmt.Println("----")
}

type pointSlice []*point

func (p *pointSlice) Less(i, j int) (bool) {
	return (*p)[i].x < (*p)[j].x
}

func (p *pointSlice) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func bitonicTSP(points []*point) (float64, []*line) {
	if len(points) < 2 {
		panic("There must be more than 2 points")
	}

	sortedPoints := pointSlice(points)
	sort.Slice(points, sortedPoints.Less)

	//side data structures
	dist := make([][]float64, len(sortedPoints), cap(sortedPoints))
	lines := make([][][]*line, len(sortedPoints), cap(sortedPoints))
	for i := range dist {
		dist[i] = make([]float64, len(sortedPoints), cap(sortedPoints))
		lines[i] = make([][]*line, len(sortedPoints), cap(sortedPoints))
	}

	for numPoints := 1; numPoints < len(sortedPoints); numPoints++ {
		//idxPoint < numPoints-1
		for idxPoint := 0; idxPoint < numPoints-1; idxPoint++ {
			line := newLine(sortedPoints[numPoints-1], sortedPoints[numPoints])
			dist[numPoints][idxPoint] = dist[numPoints-1][idxPoint] + line.dist()
			lines[numPoints][idxPoint] = append(lines[numPoints-1][idxPoint], line)
		}

		//idxPoint == numPoints-1
		line := newLine(sortedPoints[0], sortedPoints[numPoints])
		dist[numPoints][numPoints-1] = dist[numPoints-1][0] + line.dist()
		lines[numPoints][numPoints-1] = append(lines[numPoints-1][0],  line)
		for idxPreviousPoint := 0; idxPreviousPoint < numPoints-1; idxPreviousPoint++ {
			line := newLine(sortedPoints[idxPreviousPoint], sortedPoints[numPoints])
			temp := dist[numPoints-1][idxPreviousPoint] + line.dist()
			if temp < dist[numPoints][numPoints-1] {
				dist[numPoints][numPoints-1] = temp
				lines[numPoints][numPoints-1] = append(lines[numPoints-1][idxPreviousPoint],  line)
			}
		}

		//idxPoint == numPoints
		line = newLine(sortedPoints[numPoints-1], sortedPoints[numPoints])
		dist[numPoints][numPoints] = dist[numPoints][numPoints-1] + line.dist()
		lines[numPoints][numPoints] = append(lines[numPoints][numPoints-1],line)
	}

	return dist[len(sortedPoints)-1][len(sortedPoints)-1], lines[len(sortedPoints)-1][len(sortedPoints)-1]
}
