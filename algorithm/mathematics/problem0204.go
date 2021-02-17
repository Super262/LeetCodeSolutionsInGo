package mathematics

func countPrimes(n int) int {
	isPrime := make([]bool, n, n)
	for i := range isPrime {
		isPrime[i] = true
	}
	result := 0
	var n64, num64, temp, j int64
	n64 = int64(n)
	for num := 2; num < n; num++ {
		if isPrime[num] {
			result++
			num64 = int64(num)
			temp = num64 * num64
			if temp < n64 {
				for j = temp; j < n64; j += num64 {
					isPrime[int(j)] = false
				}
			}
		}
	}
	return result
}
