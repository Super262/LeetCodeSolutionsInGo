package sorting

func wiggleSort(nums []int) {
	if nums == nil || len(nums) == 0 {
		return
	}
	for i := 1; i < len(nums); i++ {
		if (i%2 == 0 && nums[i-1] < nums[i]) || (i%2 == 1 && nums[i-1] > nums[i]) {
			swap0280(&nums, i-1, i)
		}
	}
}

func swap0280(nums *[]int, a, b int) {
	if nums == nil || len(*nums) == 0 {
		return
	}
	tmp := (*nums)[a]
	(*nums)[a] = (*nums)[b]
	(*nums)[b] = tmp
}
