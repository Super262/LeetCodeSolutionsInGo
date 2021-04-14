package sorting

func wiggleSortII(nums []int) {
	if nums == nil || len(nums) == 0 {
		return
	}
	n := len(nums)
	answer := make([]int, n, n)
	smallerMid := partitionBySmallerMid0324(&nums, 0, n-1, (n-1)/2+1)
	for i := range answer {
		answer[i] = smallerMid
	}
	left := 1
	right := 0
	if n%2 == 0 {
		right = n - 2
	} else {
		right = n - 1
	}
	for _, num := range nums {
		if num > smallerMid {
			answer[left] = num
			left += 2
		} else if num < smallerMid {
			answer[right] = num
			right -= 2
		}
	}
	copy(nums, answer)
}

func partitionBySmallerMid0324(nums *[]int, start, end, k int) int {
	if start >= end {
		return (*nums)[start]
	}
	pivot := (*nums)[start+(end-start)/2]
	i := start
	j := end
	for i <= j {
		for i <= j && (*nums)[i] < pivot {
			i++
		}
		for i <= j && (*nums)[j] > pivot {
			j--
		}
		if i <= j {
			tmp := (*nums)[i]
			(*nums)[i] = (*nums)[j]
			(*nums)[j] = tmp
			i++
			j--
		}
	}
	if start+k-1 <= j {
		return partitionBySmallerMid0324(nums, start, j, k)
	}
	if start+k-1 >= i {
		return partitionBySmallerMid0324(nums, i, end, k-(i-start))
	}
	return (*nums)[j+1]
}
