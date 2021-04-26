package mathematics

import "strings"

func strStr(haystack string, needle string) int {
	if strings.Compare(needle, "") == 0 {
		return 0
	}
	hChIndex := 0
	nChIndex := 0
	next := buildNext(&needle)
	for hChIndex < len(haystack) {
		if haystack[hChIndex] == needle[nChIndex] {
			hChIndex++
			nChIndex++
		} else if nChIndex != 0 {
			nChIndex = (*next)[nChIndex-1]
		} else {
			hChIndex++
		}
		if nChIndex == len(needle) {
			return hChIndex - nChIndex
		}
	}
	return -1
}

func buildNext(needle *string) *[]int {
	next := make([]int, len(*needle), len(*needle))
	now := 0
	x := 1
	for x < len(next) {
		if (*needle)[x] == (*needle)[now] {
			now++
			next[x] = now
			x++
		} else if now != 0 {
			now = next[now-1]
		} else {
			next[x] = 0
			x++
		}
	}
	return &next
}
