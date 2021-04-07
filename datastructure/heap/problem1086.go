package heap

import (
	"container/heap"
	"sort"
)

func highFive(items [][]int) [][]int {
	idToHeap := make(map[int]*MinHeap1086)
	for _, stu := range items {
		_, existed := idToHeap[stu[0]]
		if !existed {
			idToHeap[stu[0]] = &MinHeap1086{}
		}
		heap.Push(idToHeap[stu[0]], stu[1])
		if len(*idToHeap[stu[0]]) > 5 {
			heap.Pop(idToHeap[stu[0]])
		}
	}
	result := make([][]int, 0)
	for id, scoreList := range idToHeap {
		result = append(result, []int{id, scoreList.Sum().(int) / 5})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i][0] <= result[j][0]
	})
	return result
}

type MinHeap1086 []int // 定义一个类型

func (minHeap MinHeap1086) Len() int { return len(minHeap) } // 绑定len方法,返回长度
func (minHeap MinHeap1086) Less(i, j int) bool { // 绑定less方法
	return minHeap[i] < minHeap[j] // 如果h[i]<minHeap[j]生成的就是小根堆，如果h[i]>minHeap[j]生成的就是大根堆
}
func (minHeap MinHeap1086) Swap(i, j int) { // 绑定swap方法，交换两个元素位置
	minHeap[i], minHeap[j] = minHeap[j], minHeap[i]
}

func (minHeap *MinHeap1086) Pop() interface{} { // 绑定pop方法，从最后拿出一个元素并返回
	old := *minHeap
	n := len(old)
	x := old[n-1]
	*minHeap = old[0 : n-1]
	return x
}

func (minHeap *MinHeap1086) Push(x interface{}) { // 绑定push方法，插入新元素
	*minHeap = append(*minHeap, x.(int))
}

func (minHeap *MinHeap1086) Peek() interface{} {
	return (*minHeap)[0]
}

func (minHeap *MinHeap1086) Sum() interface{} {
	sum := 0
	for _, v := range *minHeap {
		sum += v
	}
	return sum
}
