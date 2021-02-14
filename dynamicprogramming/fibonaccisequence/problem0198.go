package fibonaccisequence

func rob0198(nums []int) int {
	if nums == nil {
		return 0
	}
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}
	f0 := 0
	f1 := nums[0]
	fn := f1 + f0
	tempSum := 0
	for l := 2; l <= numsLen; l++ {
		tempSum = nums[l-1] + f0
		if tempSum > f1 {
			fn = tempSum
		} else {
			fn = f1
		}
		f0 = f1
		f1 = fn
	}
	return fn
}
