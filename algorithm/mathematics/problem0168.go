package mathematics

func convertToTitle(n int) string {
	if n == 0 {
		return ""
	}
	result := make([]rune, 0)
	for n > 0 {
		n--
		result = append(result, rune('A'+n%26))
		n /= 26
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}
