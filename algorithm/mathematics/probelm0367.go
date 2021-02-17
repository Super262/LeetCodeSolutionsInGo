package mathematics

func isPerfectSquare(num int) bool {
	dis := 1
	for num > 0 {
		num -= dis
		dis += 2
	}
	return num == 0
}
