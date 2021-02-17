package sorting

func swap0075(nums []int, a int, b int) {
	t := nums[a]
	nums[a] = nums[b]
	nums[b] = t
}
func sortColors(nums []int) {
	zerosP := -1
	onesP := 0
	twosP := len(nums)
	for onesP < twosP {
		if nums[onesP] == 0 {
			zerosP++
			swap0075(nums, onesP, zerosP)
			onesP++
		} else if nums[onesP] == 2 {
			twosP--
			swap0075(nums, onesP, twosP)
		} else {
			onesP++
		}
	}
}
