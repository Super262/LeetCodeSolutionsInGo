package mathematics

func majorityElement2(nums []int) []int {
	numsLen := len(nums)
	vote1 := 0
	vote2 := 0
	candidate1 := nums[0]
	candidate2 := nums[0]
	for i := 0; i < numsLen; i++ {
		if nums[i] == candidate1 {
			vote1++
			continue
		}
		if nums[i] == candidate2 {
			vote2++
			continue
		}
		if vote1 != 0 && vote2 != 0 {
			vote1--
			vote2--
			continue
		}
		if vote1 == 0 {
			candidate1 = nums[i]
			vote1 = 1
			continue
		}
		if vote2 == 0 {
			candidate2 = nums[i]
			vote2 = 1
			continue
		}
	}
	result := make([]int, 0, 2)
	vote1 = 0
	vote2 = 0
	for _, num := range nums {
		if num == candidate1 {
			vote1++
		} else if num == candidate2 {
			vote2++
		}
	}
	if vote1 > numsLen/3 {
		result = append(result, candidate1)
	}
	if vote2 > numsLen/3 {
		result = append(result, candidate2)
	}
	return result
}
