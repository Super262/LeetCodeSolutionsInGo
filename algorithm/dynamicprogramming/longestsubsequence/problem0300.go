package longestsubsequence

func findPreMax0300(tails *[]int, currentLen int, key int) int {
	low := 0
	high := currentLen
	mid := 0
	for low < high {
		mid = low + (high-low)/2
		if (*tails)[mid] == key {
			return mid
		} else if (*tails)[mid] > key {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func lengthOfLIS(nums []int) int {
	if nums == nil {
		return 0
	}
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}
	tails := make([]int, numsLen, numsLen)
	tails[0] = nums[0]
	currentLen := 1
	preMaxIndex := 0
	for i := 1; i < numsLen; i++ {
		preMaxIndex = findPreMax0300(&tails, currentLen, nums[i])
		tails[preMaxIndex] = nums[i]
		if preMaxIndex == currentLen {
			currentLen++
		}
	}
	return currentLen
}
