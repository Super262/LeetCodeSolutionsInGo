package longestsubsequence

func wiggleMaxLength(nums []int) int {
	if nums == nil {
		return 0
	}
	numsLen := len(nums)
	if numsLen < 2 {
		return numsLen
	}
	up := 1
	down := 1
	for i := 1; i < numsLen; i++ {
		if nums[i] > nums[i-1] {
			up = down + 1
		} else if nums[i] < nums[i-1] {
			down = up + 1
		}
	}
	if up > down {
		return up
	} else {
		return down
	}
}
