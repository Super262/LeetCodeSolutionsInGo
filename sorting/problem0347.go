package sorting

func topKFrequent(nums []int, k int) []int {
	numToFreq := make(map[int]int)
	for _, num := range nums {
		numToFreq[num] = numToFreq[num] + 1
	}
	numsLen := len(nums)
	freqToNum := make([][]int, numsLen+1, numsLen+1)
	for num, freq := range numToFreq {
		if freqToNum[freq] == nil {
			freqToNum[freq] = make([]int, 0)
		}
		freqToNum[freq] = append(freqToNum[freq], num)
	}
	result := make([]int, 0)
	for i := numsLen; i >= 0; i-- {
		if freqToNum[i] != nil {
			resultLen := len(result)
			if resultLen < k {
				dis := k - resultLen
				if len(freqToNum[i]) <= dis {
					result = append(result, freqToNum[i]...)
				} else {
					result = append(result, freqToNum[i][0:dis]...)
				}
			} else {
				break
			}
		}
	}
	return result
}
