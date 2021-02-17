package binarysearch

func nextGreatestLetter(letters []byte, target byte) byte {
	dicLen := len(letters)
	if target >= letters[dicLen-1] || target < letters[0] {
		return letters[0]
	}
	low := 0
	high := dicLen - 1
	mid := 0
	for low <= high {
		mid = low + (high-low)/2
		if letters[mid] <= target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return letters[low]
}
