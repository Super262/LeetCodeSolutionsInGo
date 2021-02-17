package fibonaccisequence

func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	f1 := 1
	f2 := 2
	fn := 0
	n -= 2
	for n > 0 {
		fn = f1 + f2
		f1 = f2
		f2 = fn
		n--
	}
	return fn
}
