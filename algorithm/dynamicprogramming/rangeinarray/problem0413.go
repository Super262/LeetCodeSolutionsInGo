package rangeinarray

func numberOfArithmeticSlices(A []int) int {
	if A == nil {
		return 0
	}
	aLen := len(A)
	if aLen < 3 {
		return 0
	}
	dp := make([]int, aLen, aLen)
	for i := 2; i < aLen; i++ {
		if A[i]-A[i-1] == A[i-1]-A[i-2] {
			dp[i] = dp[i-1] + 1
		}
	}
	result := 0
	for _, r := range dp {
		result += r
	}
	return result
}
