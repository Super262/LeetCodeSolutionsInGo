package binarysearch

func minArea(image [][]byte, x int, y int) int {
	if image == nil || len(image) == 0 || image[0] == nil || len(image[0]) == 0 || image[x][y] == '0' {
		return 0
	}
	left := findFirst0302(&image, 0, y, false)
	right := findLast0302(&image, y, len(image[0])-1, false)
	up := findFirst0302(&image, 0, x, true)
	down := findLast0302(&image, x, len(image)-1, true)
	return (right - left + 1) * (down - up + 1)
}

func findLast0302(image *[][]byte, start, end int, needCheckRow bool) int {
	for start+1 < end {
		mid := start + (end-start)/2
		if (needCheckRow && checkRow0302(image, mid)) || (!needCheckRow && checkCol0302(image, mid)) {
			start = mid
		} else {
			end = mid
		}
	}
	if (needCheckRow && checkRow0302(image, end)) || (!needCheckRow && checkCol0302(image, end)) {
		return end
	}
	return start
}

func findFirst0302(image *[][]byte, start, end int, needCheckRow bool) int {
	for start+1 < end {
		mid := start + (end-start)/2
		if (needCheckRow && checkRow0302(image, mid)) || (!needCheckRow && checkCol0302(image, mid)) {
			end = mid
		} else {
			start = mid
		}
	}
	if (needCheckRow && checkRow0302(image, start)) || (!needCheckRow && checkCol0302(image, start)) {
		return start
	}
	return end
}

func checkRow0302(image *[][]byte, row int) bool {
	for col := 0; col < len((*image)[0]); col++ {
		if (*image)[row][col] == '1' {
			return true
		}
	}
	return false
}

func checkCol0302(image *[][]byte, col int) bool {
	for row := 0; row < len((*image)); row++ {
		if (*image)[row][col] == '1' {
			return true
		}
	}
	return false
}
