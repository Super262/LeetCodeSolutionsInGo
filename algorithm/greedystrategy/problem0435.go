package greedystrategy

import "sort"

type points0435 [][]int

func (p points0435) Len() int {
	return len(p)
}
func (p points0435) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p points0435) Less(i, j int) bool {
	return p[i][1] < p[j][1]
}

func eraseOverlapIntervals(intervals [][]int) int {
	if intervals == nil {
		return 0
	}
	inLen := len(intervals)
	if inLen == 0 {
		return 0
	}
	sort.Sort(points0435(intervals))
	result := 0
	prev := intervals[0]
	for i := 1; i < inLen; i++ {
		if intervals[i][0] < prev[1] {
			result++
		} else {
			prev = intervals[i]
		}
	}
	return result
}
