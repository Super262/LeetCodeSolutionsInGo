package mathematics

func myPow(x float64, n int) float64 {
	result := 1.0
	isNegPower := n < 0
	if isNegPower {
		n = -n
	}
	for n != 0 {
		if n&1 != 0 {
			result *= x
		}
		x *= x
		n >>= 1
	}
	if isNegPower {
		result = 1 / result
	}
	return result
}
