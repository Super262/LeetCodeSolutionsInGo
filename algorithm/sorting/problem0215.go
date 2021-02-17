package sorting

func doPartition(nums *[]int, start int, end int) int {
	if start == end {
		return start
	}
	pivot := (*nums)[start]
	i := start
	j := end
	for i < j {
		for i < j && (*nums)[j] <= pivot {
			j--
		}
		if i < j {
			(*nums)[i] = (*nums)[j]
		}
		for i < j && (*nums)[i] > pivot {
			i++
		}
		if i < j {
			(*nums)[j] = (*nums)[i]
		}
	}
	(*nums)[i] = pivot
	return i
}

func findKthLargest(nums []int, k int) int {
	k--
	high := len(nums) - 1
	low := 0
	pIndex := 0
	for low < high {
		pIndex = doPartition(&nums, low, high)
		if pIndex == k {
			return nums[k]
		} else if pIndex > k {
			high = pIndex - 1
		} else {
			low = pIndex + 1
		}
	}
	return nums[low]
}
