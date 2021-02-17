package mathematics

func productExceptSelf(nums []int) []int {
	numsLen := len(nums)
	product := make([]int, numsLen, numsLen)
	for i := range product {
		product[i] = 1
	}
	left := 1
	for i := 1; i < numsLen; i++ {
		left *= nums[i-1]
		product[i] *= left
	}
	right := 1
	for i := numsLen - 2; i >= 0; i-- {
		right *= nums[i+1]
		product[i] *= right
	}
	return product
}
