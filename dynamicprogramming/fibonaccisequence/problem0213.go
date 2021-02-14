package fibonaccisequence

func robPart0213(nums *[]int, start int, length int) int {
	f0 := 0
	f1 := (*nums)[start]
	fn := f1 + f0
	tempSum := 0
	for l := 2; l <= length; l++ {
		tempSum = (*nums)[start+l-1] + f0
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

func rob0213(nums []int) int {
	if nums == nil {
		return 0
	}
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}
	if numsLen == 1 {
		return nums[0]
	}
	result1 := robPart0213(&nums, 0, numsLen-1)
	result2 := robPart0213(&nums, 1, numsLen-1)
	if result1 > result2 {
		return result1
	} else {
		return result2
	}
}
