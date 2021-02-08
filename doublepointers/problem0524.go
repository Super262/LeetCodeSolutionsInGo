package doublepointers

import "strings"

func findLongestWord(s string, d []string) string {
	if strings.Compare(s, "") == 0 || d == nil || len(d) == 0 {
		return ""
	}
	sArray := []rune(s)
	sLen := len(sArray)
	result := ""
	resultLen := 0
	sP := 0
	wP := 0
	wordLen := 0
	var wordArray []rune
	for _, word := range d {
		wordArray = []rune(word)
		wordLen = len(wordArray)
		sP = 0
		wP = 0
		for sP < sLen && wP < wordLen {
			if sArray[sP] == wordArray[wP] {
				wP++
			}
			sP++
		}
		if wP >= wordLen {
			if resultLen == wordLen {
				if strings.Compare(result, word) > 0 {
					result = word
				}
			} else if resultLen < wordLen {
				result = word
				resultLen = wordLen
			}
		}
	}
	return result
}
