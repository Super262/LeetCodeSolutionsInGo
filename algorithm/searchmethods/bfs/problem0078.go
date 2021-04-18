package bfs

func subsets(nums []int) [][]int {
	queue := make([][]int, 0)
	if nums == nil || len(nums) == 0 {
		return queue
	}
	queue = append(queue, make([]int, 0))
	index := 0
	for index < len(queue) {
		baseSet := queue[index]
		for _, num := range nums {
			if len(baseSet) > 0 && baseSet[len(baseSet)-1] >= num {
				continue
			}
			tmp := make([]int, len(baseSet), len(baseSet))
			copy(tmp, baseSet)
			tmp = append(tmp, num)
			queue = append(queue, tmp)
		}
		index++
	}
	return queue
}
