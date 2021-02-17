package mathematics

import "sort"

func minMoves2(nums []int) int {
	if nums == nil || len(nums) < 2 {
		return 0
	}
	sort.Ints(nums)
	median := nums[len(nums)/2]
	moves := 0
	for _, num := range nums {
		moves += getIntAbs0462(num - median)
	}
	return moves
}

func getIntAbs0462(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}
