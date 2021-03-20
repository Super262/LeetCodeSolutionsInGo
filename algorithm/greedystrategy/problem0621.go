package greedystrategy

import "sort"

func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}
	taskFrequency := make([]int, 26, 26)
	for _, t := range tasks {
		taskFrequency[int(t-'A')]++
	}
	sort.Ints(taskFrequency)
	lastPartLen := 0
	for i := 0; i < len(taskFrequency); i++ {
		if taskFrequency[i] == taskFrequency[len(taskFrequency)-1] {
			lastPartLen++
		}
	}
	temp1 := len(tasks)
	temp2 := (taskFrequency[len(taskFrequency)-1]-1)*(n+1) + lastPartLen
	if temp1 > temp2 {
		return temp1
	}
	return temp2
}
