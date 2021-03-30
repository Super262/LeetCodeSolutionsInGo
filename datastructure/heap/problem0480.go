package heap

import (
	"container/heap"
	"math"
)

func medianSlidingWindow(nums []int, k int) []float64 {
	if nums == nil || len(nums) < k || k <= 0 {
		return []float64{}
	}
	dualHeap := Constructor0480()
	result := make([]float64, len(nums)-k+1, len(nums)-k+1)
	resultTop := 0
	for i := 0; i < k; i++ {
		dualHeap.Add(nums[i])
	}
	result[resultTop] = dualHeap.GetMedian()
	resultTop++
	for start := 1; start < len(nums)-k+1; start++ {
		dualHeap.Remove(nums[start-1])
		dualHeap.Add(nums[start+k-1])
		result[resultTop] = dualHeap.GetMedian()
		resultTop++
	}
	return result
}

type DualHeap0480 struct {
	largerNums      MinHeap0480
	smallerNums     MaxHeap0480
	delayedRemovals map[int]int
	smallerSize     int
	largerSize      int
}

func Constructor0480() *DualHeap0480 {
	dualHeap := DualHeap0480{largerNums: MinHeap0480{}, smallerNums: MaxHeap0480{}, delayedRemovals: make(map[int]int), smallerSize: 0, largerSize: 0}
	return &dualHeap
}

func (dualHeap *DualHeap0480) Add(num int) {
	if dualHeap.smallerSize == 0 || dualHeap.smallerNums.Len() == 0 || num <= dualHeap.smallerNums.Peek().(int) {
		heap.Push(&dualHeap.smallerNums, num)
		dualHeap.smallerSize++
	} else {
		heap.Push(&dualHeap.largerNums, num)
		dualHeap.largerSize++
	}
	dualHeap.BalanceTwoParts()
}

func (dualHeap *DualHeap0480) Remove(num int) {
	dualHeap.PushDeletion(num)
	if dualHeap.smallerSize > 0 && dualHeap.smallerNums.Len() > 0 && num <= dualHeap.smallerNums.Peek().(int) {
		dualHeap.smallerSize--
		dualHeap.CleanAndAlignSmaller()
	} else if dualHeap.largerSize > 0 {
		dualHeap.largerSize--
		dualHeap.CleanAndAlignLarger()
	}
	dualHeap.BalanceTwoParts()
}

func (dualHeap *DualHeap0480) GetMedian() float64 {
	if dualHeap.smallerSize == dualHeap.largerSize && dualHeap.smallerNums.Len() > 0 && dualHeap.largerNums.Len() > 0 {
		return (float64(dualHeap.smallerNums.Peek().(int)) + float64(dualHeap.largerNums.Peek().(int))) / 2.0
	} else if dualHeap.smallerNums.Len() > 0 {
		return float64(dualHeap.smallerNums.Peek().(int))
	}
	return math.MinInt32
}

func (dualHeap *DualHeap0480) BalanceTwoParts() {
	for dualHeap.largerNums.Len() > 0 && dualHeap.smallerSize < dualHeap.largerSize {
		heap.Push(&dualHeap.smallerNums, heap.Pop(&dualHeap.largerNums))
		dualHeap.smallerSize++
		dualHeap.largerSize--
		dualHeap.CleanAndAlignLarger()
	}
	for dualHeap.smallerNums.Len() > 0 && dualHeap.smallerSize > dualHeap.largerSize+1 {
		heap.Push(&dualHeap.largerNums, heap.Pop(&dualHeap.smallerNums))
		dualHeap.largerSize++
		dualHeap.smallerSize--
		dualHeap.CleanAndAlignSmaller()
	}
}

func (dualHeap *DualHeap0480) CleanAndAlignLarger() {
	peekValue := 0
	for dualHeap.largerNums.Len() > 0 {
		peekValue = dualHeap.largerNums.Peek().(int)
		if !dualHeap.PollDeletion(peekValue) {
			break
		}
		heap.Pop(&dualHeap.largerNums)
	}
}

func (dualHeap *DualHeap0480) CleanAndAlignSmaller() {
	peekValue := 0
	for dualHeap.smallerNums.Len() > 0 {
		peekValue = dualHeap.smallerNums.Peek().(int)
		if !dualHeap.PollDeletion(peekValue) {
			break
		}
		heap.Pop(&dualHeap.smallerNums)
	}
}

func (dualHeap *DualHeap0480) PollDeletion(key int) bool {
	remainsCount, existed := (dualHeap.delayedRemovals)[key]
	if !existed {
		return false
	}
	if remainsCount == 1 {
		delete(dualHeap.delayedRemovals, key)
	} else {
		(dualHeap.delayedRemovals)[key] = remainsCount - 1
	}
	return true
}

func (dualHeap *DualHeap0480) PushDeletion(key int) {
	remainsCount, existed := (dualHeap.delayedRemovals)[key]
	if existed {
		(dualHeap.delayedRemovals)[key] = remainsCount + 1
	} else {
		(dualHeap.delayedRemovals)[key] = 1
	}
}

type MaxHeap0480 []int // 定义一个类型

func (maxHeap MaxHeap0480) Len() int {
	return len(maxHeap) // 绑定len方法,返回长度
}
func (maxHeap MaxHeap0480) Less(i, j int) bool { // 绑定less方法
	return maxHeap[i] > maxHeap[j] // 如果h[i]<maxHeap[j]生成的就是小根堆，如果h[i]>maxHeap[j]生成的就是大根堆
}
func (maxHeap MaxHeap0480) Swap(i, j int) { // 绑定swap方法，交换两个元素位置
	maxHeap[i], maxHeap[j] = maxHeap[j], maxHeap[i]
}

func (maxHeap *MaxHeap0480) Pop() interface{} { // 绑定pop方法，从最后拿出一个元素并返回
	old := *maxHeap
	n := len(old)
	x := old[n-1]
	*maxHeap = old[0 : n-1]
	return x
}

func (maxHeap *MaxHeap0480) Push(x interface{}) { // 绑定push方法，插入新元素
	*maxHeap = append(*maxHeap, x.(int))
}

func (maxHeap *MaxHeap0480) Peek() interface{} {
	return (*maxHeap)[0]
}

type MinHeap0480 []int // 定义一个类型

func (minHeap MinHeap0480) Len() int {
	return len(minHeap) // 绑定len方法,返回长度
}
func (minHeap MinHeap0480) Less(i, j int) bool { // 绑定less方法
	return minHeap[i] < minHeap[j] // 如果h[i]<minHeap[j]生成的就是小根堆，如果h[i]>minHeap[j]生成的就是大根堆
}
func (minHeap MinHeap0480) Swap(i, j int) { // 绑定swap方法，交换两个元素位置
	minHeap[i], minHeap[j] = minHeap[j], minHeap[i]
}

func (minHeap *MinHeap0480) Pop() interface{} { // 绑定pop方法，从最后拿出一个元素并返回
	old := *minHeap
	n := len(old)
	x := old[n-1]
	*minHeap = old[0 : n-1]
	return x
}

func (minHeap *MinHeap0480) Push(x interface{}) { // 绑定push方法，插入新元素
	*minHeap = append(*minHeap, x.(int))
}

func (minHeap *MinHeap0480) Peek() interface{} {
	return (*minHeap)[0]
}
