package sorting

func topKFrequent(nums []int, k int) []int {
	numToFreq := make(map[int]int)
	for _, num := range nums {
		numToFreq[num] = numToFreq[num] + 1
	}
	numsLen := len(nums)
	freqBuckets := make([][]int, numsLen+1, numsLen+1)
	for num, freq := range numToFreq {
		if freqBuckets[freq] == nil {
			freqBuckets[freq] = make([]int, 0)
		}
		freqBuckets[freq] = append(freqBuckets[freq], num)
	}
	result := make([]int, 0)
	for i := numsLen; i >= 0; i-- {
		if freqBuckets[i] != nil {
			resultLen := len(result)
			if resultLen < k {
				dis := k - resultLen
				if len(freqBuckets[i]) <= dis {
					result = append(result, freqBuckets[i]...)
				} else {
					result = append(result, freqBuckets[i][0:dis]...)
				}
			} else {
				break
			}
		}
	}
	return result
}
