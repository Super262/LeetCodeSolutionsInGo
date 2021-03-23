package doublepointers

import "sort"

func threeSum(nums []int) [][]int {
	results := make([][]int, 0, 0)
	if nums == nil || len(nums) < 3 {
		return results
	}
	sort.Ints(nums)
	maxBound := len(nums) - 1
	for i := range nums {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		twoSum0015(&nums, -nums[i], i+1, maxBound, &results)
	}
	return results
}

func twoSum0015(nums *[]int, target int, start int, end int, results *[][]int) {
	if start >= end {
		return
	}
	left := start
	right := end
	tempSum := 0
	for left < right {
		tempSum = (*nums)[left] + (*nums)[right]
		if tempSum == target {
			*results = append(*results, []int{-target, (*nums)[left], (*nums)[right]})
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
