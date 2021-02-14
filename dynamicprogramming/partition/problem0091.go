package partition

func toInt0091(s *[]rune, first int, second int) int {
	if first >= second || (*s)[first] == '0' {
		return -1
	}
	num := int((*s)[first]-'0')*10 + int((*s)[second]-'0')
	if num > 26 {
		return -1
	}
	return num
}

func numDecodings(s string) int {
	sArray := []rune(s)
	sLen := len(sArray)
	f0 := 1
	f1 := 0
	if sArray[0] != '0' {
		f1 = 1
	}
	f2 := 0
	temp := 0
	for l := 2; l <= sLen; l++ {
		temp = toInt0091(&sArray, l-2, l-1)
		if sArray[l-1] != '0' {
			f2 += f1
		}
		if temp != -1 {
			f2 += f0
		}
		f0 = f1
		f1 = f2
		f2 = 0
	}
	return f1
}
