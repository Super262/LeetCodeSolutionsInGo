package binarysearch

import "sort"

func findClosestElements(arr []int, k int, x int) []int {
	result := make([]int, 0, k)
	if arr == nil || len(arr) == 0 {
		return result
	}
	rightIndex := findClosestUpper0658(&arr, x)
	leftIndex := rightIndex - 1
	for i := 0; i < k; i++ {
		if isLeftCloser0658(&arr, x, leftIndex, rightIndex) {
			result = append(result, arr[leftIndex])
			leftIndex--
		} else {
			result = append(result, arr[rightIndex])
			rightIndex++
		}
	}
	sort.Ints(result)
	return result
}

func findClosestUpper0658(arr *[]int, x int) int {
	start := 0
	end := len(*arr) - 1
	mid := 0
	for start+1 < end {
		mid = start + (end-start)/2
		if (*arr)[mid] >= x {
			end = mid
		} else {
			start = mid
		}
	}
	if (*arr)[start] >= x {
		return start
	}
	if (*arr)[end] >= x {
		return end
	}
	return len(*arr)
}

func isLeftCloser0658(arr *[]int, x int, leftIndex int, rightIndex int) bool {
	if leftIndex < 0 {
		return false
	}
	if rightIndex >= len(*arr) {
		return true
	}
	dl := x - (*arr)[leftIndex]
	dr := (*arr)[rightIndex] - x
	if dl == dr {
		return (*arr)[leftIndex] <= (*arr)[rightIndex]
	}
	return dl <= dr
}
