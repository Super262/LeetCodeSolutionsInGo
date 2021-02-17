package greedystrategy

import "strings"

func isSubsequence(s string, t string) bool {
	if strings.Compare(s, "") == 0 {
		return true
	}
	sArray := []rune(s)
	tArray := []rune(t)
	sLen := len(sArray)
	tLen := len(tArray)
	if sLen > tLen {
		return false
	}
	sP := 0
	tP := 0
	for sP < sLen && tP < tLen {
		if sArray[sP] == tArray[tP] {
			sP++
		}
		tP++
	}
	if sP < sLen {
		return false
	} else {
		return true
	}
}
