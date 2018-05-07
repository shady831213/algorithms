package graph

import "math"

func floydWarShall(g weightedGraph, init func(weightedGraph) ([][][]int, []interface{}), handler func(*[][][]int, int, int, int),
	rebuild func([]interface{}, weightedGraph, [][]int)) weightedGraph {
	array, vertices := init(g)

	for k := range array[:len(array)-1] {
		for i := range array[k] {
			for j := range array[k][i] {
				handler(&array, k, i, j)
			}
		}
	}
	newG := createGraphByType(g).(weightedGraph)
	rebuild(vertices, newG, array[len(array)-1])
	return newG
}

func distFloydWarShall(g weightedGraph) weightedGraph {
	init := func(g weightedGraph) ([][][]int, []interface{}) {
		vertices := g.AllVertices()
		array := make([][][]int, len(vertices)+1, len(vertices)+1)

		for k := range array {
			array[k] = make([][]int, len(vertices), len(vertices))
			for i := range array[k] {
				array[k][i] = make([]int, len(vertices), len(vertices))
				if k == 0 {
					for j := range array[k][i] {
						currentEdge := edge{vertices[i], vertices[j]}
						if i == j {
							array[k][i][j] = 0
						} else if !g.CheckEdge(currentEdge) {
							array[k][i][j] = math.MaxInt32
						} else {
							array[k][i][j] = g.Weight(currentEdge)
						}
					}
				}
			}
		}

		return array, vertices
	}

	handler := func(array *[][][]int, k, i, j int) {
		(*array)[k+1][i][j] = (*array)[k][i][j]
		if (*array)[k][i][k]+(*array)[k][k][j] < (*array)[k+1][i][j] {
			(*array)[k+1][i][j] = (*array)[k][i][k] + (*array)[k][k][j]
		}
	}

	rebuild := func(vertices []interface{}, g weightedGraph, array [][]int) {
		for i := range vertices {
			for j := range vertices {
				if array[i][j] < math.MaxInt32 {
					g.AddEdgeWithWeight(edge{vertices[i], vertices[j]}, array[i][j])
				}
			}
		}
	}

	return floydWarShall(g, init, handler, rebuild)
}
