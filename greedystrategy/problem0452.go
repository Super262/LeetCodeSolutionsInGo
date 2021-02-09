package greedystrategy

import "sort"

type points0452 [][]int

func (p points0452) Len() int {
	return len(p)
}
func (p points0452) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p points0452) Less(i, j int) bool {
	return p[i][1] < p[j][1]
}

func findMinArrowShots(points [][]int) int {
	if points == nil {
		return 0
	}
	pointsLen := len(points)
	if pointsLen == 0 {
		return 0
	}
	sort.Sort(points0452(points))
	result := 1
	prev := points[0]
	for i := 1; i < pointsLen; i++ {
		if points[i][0] > prev[1] {
			result++
			prev = points[i]
		}
	}
	return result
}
