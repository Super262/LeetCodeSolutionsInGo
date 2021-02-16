package mathematics

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	dec2Hex := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	numU := uint32(num)
	result := make([]rune, 0)
	for numU != 0 {
		result = append(result, dec2Hex[numU&15])
		numU >>= 4
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}
