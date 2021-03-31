package twopointers

func minWindow(s string, t string) string {
	if s == "" || t == "" {
		return ""
	}
	tArray := []rune(t)
	countT := make([]int, 256, 256)
	uniqueT := 0
	for i := range tArray {
		countT[int(tArray[i])]++
		if countT[int(tArray[i])] == 1 {
			uniqueT++
		}
	}
	sArray := []rune(s)
	uniqueS := 0
	countS := make([]int, 256, 256)
	left := 0
	right := 0
	ansL := -1
	ansR := -1
	for left < len(sArray) {
		for right < len(sArray) && uniqueS < uniqueT {
			countS[int(sArray[right])]++
			if countS[int(sArray[right])] == countT[int(sArray[right])] {
				uniqueS++
			}
			right++
		}
		if uniqueS == uniqueT {
			if ansL == -1 || right-left < ansR-ansL {
				ansL = left
				ansR = right
			}
		}
		countS[int(sArray[left])]--
		if countS[int(sArray[left])] == countT[int(sArray[left])]-1 {
			uniqueS--
		}
		left++
	}
	if ansL == -1 {
		return ""
	}
	return string(sArray[ansL:ansR])
}
