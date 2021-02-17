package mathematics

func majorityElement(nums []int) int {
	votes := 0
	result := 0
	for _, num := range nums {
		if votes == 0 {
			result = num
			votes++
		} else {
			if num == result {
				votes++
			} else {
				votes--
			}
		}
	}
	return result
}
