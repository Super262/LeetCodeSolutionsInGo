package greedystrategy

import "sort"

type points0406 [][]int

func (p points0406) Len() int {
	return len(p)
}
func (p points0406) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p points0406) Less(i, j int) bool {
	if p[i][0] == p[j][0] {
		return p[i][1] < p[j][1]
	} else {
		return p[i][0] > p[j][0]
	}
}

func reconstructQueue(people [][]int) [][]int {
	sort.Sort(points0406(people))
	pLen := len(people)
	result := make([][]int, 0, pLen)
	for _, p := range people {
		rearPart := append([][]int{}, result[p[1]:]...)
		result = append(result[0:p[1]], p)
		result = append(result, rearPart...)
	}
	return result
}
