package dp

import (
	"math"
	"sort"
)

type point struct {
	x, y float64
}

func (p *point) dist(p1 *point) (float64) {
	return math.Sqrt(math.Pow(p.x-p1.x, 2) + math.Pow(p.y-p1.y, 2))
}

type pointSlice []*point

func (p *pointSlice) Less(i, j int) (bool) {
	return (*p)[i].x < (*p)[j].x
}

func (p *pointSlice) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func bitonicTSP(points []*point) (float64) {
	if len(points) < 3 {
		panic("There must be more than 3 points")
	}

	sortedPoints := pointSlice(points)
	sort.Slice(points, sortedPoints.Less)
	dist := make([][]float64, len(sortedPoints), cap(sortedPoints))
	for i := range dist {
		dist[i] = make([]float64, len(sortedPoints), cap(sortedPoints))
	}
	for numPoints := 1; numPoints < len(sortedPoints); numPoints++ {
		for idxPoint := 0; idxPoint < numPoints-1; idxPoint++ {
			dist[numPoints][idxPoint] = dist[numPoints-1][idxPoint] + sortedPoints[numPoints-1].dist(sortedPoints[numPoints])
		}
		dist[numPoints][numPoints-1] = dist[numPoints-1][0] + sortedPoints[numPoints].dist(sortedPoints[0])
		for idxPreviousPoint := 0; idxPreviousPoint < numPoints-1; idxPreviousPoint++ {
			temp := dist[numPoints-1][idxPreviousPoint] + sortedPoints[numPoints].dist(sortedPoints[idxPreviousPoint])
			if temp < dist[numPoints][numPoints-1] {
				dist[numPoints][numPoints-1] = temp
			}
		}
		dist[numPoints][numPoints] = dist[numPoints][numPoints-1] + sortedPoints[numPoints-1].dist(sortedPoints[numPoints])
	}
	return dist[len(sortedPoints)-1][len(sortedPoints)-1]
}
