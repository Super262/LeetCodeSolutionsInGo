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

func (this *DualHeap0480) Add(num int) {
	if this.smallerSize == 0 || this.smallerNums.Len() == 0 || num <= this.smallerNums.Peek().(int) {
		heap.Push(&this.smallerNums, num)
		this.smallerSize++
	} else {
		heap.Push(&this.largerNums, num)
		this.largerSize++
	}
	this.BalanceTwoParts()
}

func (this *DualHeap0480) Remove(num int) {
	this.PushDeletion(num)
	if this.smallerSize > 0 && this.smallerNums.Len() > 0 && num <= this.smallerNums.Peek().(int) {
		this.smallerSize--
		this.CleanAndAlignSmaller()
	} else if this.largerSize > 0 {
		this.largerSize--
		this.CleanAndAlignLarger()
	}
	this.BalanceTwoParts()
}

func (this *DualHeap0480) GetMedian() float64 {
	if this.smallerSize == this.largerSize && this.smallerNums.Len() > 0 && this.largerNums.Len() > 0 {
		return (float64(this.smallerNums.Peek().(int)) + float64(this.largerNums.Peek().(int))) / 2.0
	} else if this.smallerNums.Len() > 0 {
		return float64(this.smallerNums.Peek().(int))
	}
	return math.MinInt32
}

func (this *DualHeap0480) BalanceTwoParts() {
	for this.largerNums.Len() > 0 && this.smallerSize < this.largerSize {
		heap.Push(&this.smallerNums, heap.Pop(&this.largerNums))
		this.smallerSize++
		this.largerSize--
		this.CleanAndAlignLarger()
	}
	for this.smallerNums.Len() > 0 && this.smallerSize > this.largerSize+1 {
		heap.Push(&this.largerNums, heap.Pop(&this.smallerNums))
		this.largerSize++
		this.smallerSize--
		this.CleanAndAlignSmaller()
	}
}

func (this *DualHeap0480) CleanAndAlignLarger() {
	peekValue := 0
	for this.largerNums.Len() > 0 {
		peekValue = this.largerNums.Peek().(int)
		if !this.PollDeletion(peekValue) {
			break
		}
		heap.Pop(&this.largerNums)
	}
}

func (this *DualHeap0480) CleanAndAlignSmaller() {
	peekValue := 0
	for this.smallerNums.Len() > 0 {
		peekValue = this.smallerNums.Peek().(int)
		if !this.PollDeletion(peekValue) {
			break
		}
		heap.Pop(&this.smallerNums)
	}
}

func (this *DualHeap0480) PollDeletion(key int) bool {
	remainsCount, existed := (this.delayedRemovals)[key]
	if !existed {
		return false
	}
	if remainsCount == 1 {
		delete(this.delayedRemovals, key)
	} else {
		(this.delayedRemovals)[key] = remainsCount - 1
	}
	return true
}

func (this *DualHeap0480) PushDeletion(key int) {
	remainsCount, existed := (this.delayedRemovals)[key]
	if existed {
		(this.delayedRemovals)[key] = remainsCount + 1
	} else {
		(this.delayedRemovals)[key] = 1
	}
}

type MaxHeap0480 []int // 定义一个类型

func (h MaxHeap0480) Len() int {
	return len(h) // 绑定len方法,返回长度
}
func (h MaxHeap0480) Less(i, j int) bool { // 绑定less方法
	return h[i] > h[j] // 如果h[i]<h[j]生成的就是小根堆，如果h[i]>h[j]生成的就是大根堆
}
func (h MaxHeap0480) Swap(i, j int) { // 绑定swap方法，交换两个元素位置
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap0480) Pop() interface{} { // 绑定pop方法，从最后拿出一个元素并返回
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MaxHeap0480) Push(x interface{}) { // 绑定push方法，插入新元素
	*h = append(*h, x.(int))
}

func (this *MaxHeap0480) Peek() interface{} {
	return (*this)[0]
}

type MinHeap0480 []int // 定义一个类型

func (h MinHeap0480) Len() int {
	return len(h) // 绑定len方法,返回长度
}
func (h MinHeap0480) Less(i, j int) bool { // 绑定less方法
	return h[i] < h[j] // 如果h[i]<h[j]生成的就是小根堆，如果h[i]>h[j]生成的就是大根堆
}
func (h MinHeap0480) Swap(i, j int) { // 绑定swap方法，交换两个元素位置
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap0480) Pop() interface{} { // 绑定pop方法，从最后拿出一个元素并返回
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MinHeap0480) Push(x interface{}) { // 绑定push方法，插入新元素
	*h = append(*h, x.(int))
}

func (this *MinHeap0480) Peek() interface{} {
	return (*this)[0]
}
