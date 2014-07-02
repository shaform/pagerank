package pagerank

import "math"

type Graph struct {
	InEdges    [][]int
	OutDegree  []int
	EmptyNodes []int
	NNode      int
}

func NewGraph(size int) *Graph {
	return &Graph{make([][]int, size), make([]int, size), []int{}, size}
}

func Pagerank(g *Graph, d float64, eps float64) []float64 {
	glen := g.NNode
	pg1, pg2 := make([]float64, glen), make([]float64, glen)
	weight1, weight2 := make([]float64, glen), make([]float64, glen)

	// initialize pagerank
	for i := range pg1 {
		pg1[i] = float64(1.0)
		if g.OutDegree[i] != 0 {
			weight1[i] = float64(1.0) / float64(g.OutDegree[i])
		}
	}

	// power iteration
	for {
		totalE := float64(0.0)
		for _, idx := range g.EmptyNodes {
			totalE += pg1[idx]
		}
		totalE /= float64(glen)

		diff := float64(0.0)
		for i := range pg2 {
			w := totalE
			for _, idx := range g.InEdges[i] {
				w += weight1[idx]
			}
			pg2[i] = (1.0 - d) + d*w
			if g.OutDegree[i] != 0 {
				weight2[i] = pg2[i] / float64(g.OutDegree[i])
			}

			diff += math.Pow(pg1[i]-pg2[i], 2)
		}

		if math.Sqrt(diff) < eps {
			break
		}
		pg1, pg2 = pg2, pg1
		weight1, weight2 = weight2, weight1
	}

	return pg2
}
