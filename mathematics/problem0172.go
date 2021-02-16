package mathematics

func trailingZeroes(n int) int {
	currentBase := 5
	result := 0
	temp := n / currentBase
	for temp > 0 {
		result += temp
		currentBase *= 5
		temp = n / currentBase
	}
	return result
}
