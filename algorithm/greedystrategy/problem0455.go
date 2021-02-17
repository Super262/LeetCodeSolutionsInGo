package greedystrategy

import "sort"

func findContentChildren(g []int, s []int) int {
	if g == nil || s == nil {
		return 0
	}
	gLen := len(g)
	sLen := len(s)
	if gLen == 0 || sLen == 0 {
		return 0
	}
	sort.Ints(g)
	sort.Ints(s)
	result := 0
	gP := 0
	sP := 0
	for gP < gLen && sP < sLen {
		if g[gP] <= s[sP] {
			result++
			gP++
		}
		sP++
	}
	return result
}
