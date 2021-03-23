package doublepointers

import "sort"

func fourSum(nums []int, target int) [][]int {
	results := make([][]int, 0, 0)
	if nums == nil {
		return results
	}

	boundD := len(nums)
	if boundD < 4 {
		return results
	}
	boundB := boundD - 2
	boundA := boundB - 1
	sort.Ints(nums)
	for a := 0; a < boundA; a++ {
		if a > 0 && nums[a-1] == nums[a] {
			continue
		}
		for b := a + 1; b < boundB; b++ {
			if b > a+1 && nums[b-1] == nums[b] {
				continue
			}
			twoSum0018(&nums, target-nums[a]-nums[b], b+1, boundD, nums[a], nums[b], &results)
		}
	}
	return results
}

func twoSum0018(nums *[]int, target int, start int, end int, prefix1 int, prefix2 int, results *[][]int) {
	if start >= end-1 {
		return
	}
	left := start
	right := end - 1
	tempSum := 0
	for left < right {
		tempSum = (*nums)[left] + (*nums)[right]
		if tempSum == target {
			*results = append(*results, []int{prefix1, prefix2, (*nums)[left], (*nums)[right]})
			right--
			left++
			for left > 0 && left < right && (*nums)[left-1] == (*nums)[left] {
				left++
			}
		} else if tempSum > target {
			right--
		} else {
			left++
		}
	}
}
