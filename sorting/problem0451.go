package sorting

func frequencySort(s string) string {
	charToFreq := make(map[rune]int)
	sCharArray := []rune(s)
	for _, ch := range sCharArray {
		charToFreq[ch] = charToFreq[ch] + 1
	}
	sLen := len(sCharArray)
	freqBuckets := make([][]rune, sLen+1, sLen+1)
	for ch, freq := range charToFreq {
		if freqBuckets[freq] == nil {
			freqBuckets[freq] = make([]rune, 0)
		}
		freqBuckets[freq] = append(freqBuckets[freq], ch)
	}
	result := make([]rune, 0, sLen)
	for i := sLen; i >= 0; i-- {
		if freqBuckets[i] != nil {
			freq := i
			for _, ch := range freqBuckets[i] {
				for j := 0; j < freq; j++ {
					result = append(result, ch)
				}
			}
		}
	}
	return string(result)
}
