package sorting

func kClosest(points [][]int, k int) [][]int {
	pointsCopied := make([][]int, len(points), len(points))
	for i := range pointsCopied {
		pointsCopied[i] = make([]int, 2, 2)
		copy(pointsCopied[i], points[i])
	}
	findKthSmallest0973(&pointsCopied, 0, len(pointsCopied)-1, k)
	return pointsCopied[0:k]
}

func findKthSmallest0973(points *[][]int, startIndex, endIndex, k int) {
	if startIndex >= endIndex {
		return
	}
	i := startIndex
	j := endIndex
	pivotDis := getDistance0973(&(*points)[startIndex+(endIndex-startIndex)/2])
	for i <= j {
		for i <= j && getDistance0973(&(*points)[i]) < pivotDis {
			i++
		}
		for i <= j && getDistance0973(&(*points)[j]) > pivotDis {
			j--
		}
		if i <= j {
			tmp := (*points)[j]
			(*points)[j] = (*points)[i]
			(*points)[i] = tmp
			i++
			j--
		}
	}
	if startIndex+k-1 <= j {
		findKthSmallest0973(points, startIndex, j, k)
		return
	}
	if startIndex+k-1 >= i {
		findKthSmallest0973(points, i, endIndex, k-(i-startIndex))
	}
}

func getDistance0973(p *[]int) int {
	return (*p)[0]*(*p)[0] + (*p)[1]*(*p)[1]
}
