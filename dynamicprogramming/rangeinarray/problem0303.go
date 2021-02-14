package rangeinarray

type NumArray struct {
	preSum []int
}

func Constructor(nums []int) NumArray {
	if nums == nil {
		na := NumArray{nil}
		return na
	}
	numsLen := len(nums)
	if numsLen == 0 {
		na := NumArray{nil}
		return na
	}
	na := NumArray{make([]int, numsLen+1, numsLen+1)}
	for i := range nums {
		na.preSum[i+1] = na.preSum[i] + nums[i]
	}
	return na
}

func (na *NumArray) SumRange(i int, j int) int {
	if na.preSum == nil {
		return 0
	} else {
		return na.preSum[j+1] - na.preSum[i]
	}
}
