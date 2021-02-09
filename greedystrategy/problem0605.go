package greedystrategy

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	if flowerbed == nil {
		return false
	}
	fLen := len(flowerbed)
	if fLen == 0 {
		return false
	}
	result := 0
	for i := 0; i < fLen; i++ {
		if flowerbed[i] == 0 {
			if (i == 0 || flowerbed[i-1] == 0) && (i == fLen-1 || flowerbed[i+1] == 0) {
				result++
				flowerbed[i] = 1
				if result == n {
					return true
				}
			}
		}
	}
	return false
}
