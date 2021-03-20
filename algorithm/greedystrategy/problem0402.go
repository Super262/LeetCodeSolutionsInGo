package greedystrategy

func removeKdigits(num string, k int) string {
	if num == "" || k >= len(num) {
		return "0"
	}
	deletedItemsCount := 0
	numArray := []rune(num)
	stack := make([]rune, len(numArray), len(numArray))
	stackTop := -1
	digit := '\u0000'
	for i := 0; i < len(numArray); i++ {
		digit = numArray[i]
		for stackTop >= 0 && deletedItemsCount < k && stack[stackTop] > numArray[i] {
			stackTop--
			deletedItemsCount++
		}
		if numArray[i] == '0' && stackTop == -1 {
			continue
		}
		stackTop++
		stack[stackTop] = digit
	}
	for deletedItemsCount < k {
		deletedItemsCount++
		stackTop--
	}
	if stackTop == -1 {
		return "0"
	}
	return string(stack[0 : stackTop+1])
}
