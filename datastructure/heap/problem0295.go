package heap

import "container/heap"

type MedianFinder struct {
	minHeap *MinHeap
	maxHeap *MaxHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	finder := MedianFinder{minHeap: &MinHeap{}, maxHeap: &MaxHeap{}}
	heap.Init(finder.minHeap)
	heap.Init(finder.maxHeap)
	return finder
}

func (this *MedianFinder) AddNum(num int) {
	if this.maxHeap.Len() == 0 || num <= this.maxHeap.Peek().(int) {
		heap.Push(this.maxHeap, num)
	} else {
		heap.Push(this.minHeap, num)
	}
	this.Balance()
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxHeap.Len() == this.minHeap.Len() {
		return float64(this.maxHeap.Peek().(int)+this.minHeap.Peek().(int)) / 2.0
	} else {
		return float64(this.maxHeap.Peek().(int))
	}
}

func (this *MedianFinder) Balance() {
	for this.maxHeap.Len() < this.minHeap.Len() {
		heap.Push(this.maxHeap, heap.Pop(this.minHeap))
	}
	for this.maxHeap.Len() > this.minHeap.Len()+1 {
		heap.Push(this.minHeap, heap.Pop(this.maxHeap))
	}
}

type MaxHeap []int // 定义一个类型

func (h MaxHeap) Len() int { return len(h) } // 绑定len方法,返回长度
func (h MaxHeap) Less(i, j int) bool { // 绑定less方法
	return h[i] > h[j] // 如果h[i]<h[j]生成的就是小根堆，如果h[i]>h[j]生成的就是大根堆
}
func (h MaxHeap) Swap(i, j int) { // 绑定swap方法，交换两个元素位置
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Pop() interface{} { // 绑定pop方法，从最后拿出一个元素并返回
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MaxHeap) Push(x interface{}) { // 绑定push方法，插入新元素
	*h = append(*h, x.(int))
}

func (this *MaxHeap) Peek() interface{} {
	return (*this)[0]
}

type MinHeap []int // 定义一个类型

func (h MinHeap) Len() int { return len(h) } // 绑定len方法,返回长度
func (h MinHeap) Less(i, j int) bool { // 绑定less方法
	return h[i] < h[j] // 如果h[i]<h[j]生成的就是小根堆，如果h[i]>h[j]生成的就是大根堆
}
func (h MinHeap) Swap(i, j int) { // 绑定swap方法，交换两个元素位置
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Pop() interface{} { // 绑定pop方法，从最后拿出一个元素并返回
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MinHeap) Push(x interface{}) { // 绑定push方法，插入新元素
	*h = append(*h, x.(int))
}

func (this *MinHeap) Peek() interface{} {
	return (*this)[0]
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
