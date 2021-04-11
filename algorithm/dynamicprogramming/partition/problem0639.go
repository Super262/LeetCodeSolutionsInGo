package partition

func numDecodingsII(s string) int {
	charArray := []rune(s)
	n := len(charArray)
	if n == 0 {
		return 1
	}
	var f0, f1, f2, maxBound int64
	maxBound = 1000000007
	f0 = 1
	f1 = countOne0639(&charArray, 0) * f0
	for l := 2; l <= n; l++ {
		f2 = f1*countOne0639(&charArray, l-1)%maxBound + f0*countTwo0639(&charArray, l-2)%maxBound
		f2 %= maxBound
		f0 = f1
		f1 = f2
		f2 = 0
	}
	return int(f1)
}

func countOne0639(s *[]rune, start int) int64 {
	if (*s)[start] == '0' {
		return 0
	} else if (*s)[start] == '*' {
		return 9
	}
	return 1
}

func countTwo0639(s *[]rune, start int) int64 {
	if (*s)[start] == '*' {
		if (*s)[start+1] == '*' {
			return 15
		} else if '0' <= (*s)[start+1] && (*s)[start+1] <= '6' {
			return 2
		} else {
			return 1
		}
	} else if (*s)[start] == '1' {
		if (*s)[start+1] == '*' {
			return 9
		} else {
			return 1
		}
	} else if (*s)[start] == '2' {
		if (*s)[start+1] == '*' {
			return 6
		} else if '0' <= (*s)[start+1] && (*s)[start+1] <= '6' {
			return 1
		} else {
			return 0
		}
	} else {
		return 0
	}
}
