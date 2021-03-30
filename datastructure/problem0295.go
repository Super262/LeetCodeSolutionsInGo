package datastructure

import "container/heap"

type MedianFinder struct {
	minHeap *MinHeap0295
	maxHeap *MaxHeap0295
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	finder := MedianFinder{minHeap: &MinHeap0295{}, maxHeap: &MaxHeap0295{}}
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

type MaxHeap0295 []int // 定义一个类型

func (h MaxHeap0295) Len() int { return len(h) } // 绑定len方法,返回长度
func (h MaxHeap0295) Less(i, j int) bool { // 绑定less方法
	return h[i] > h[j] // 如果h[i]<h[j]生成的就是小根堆，如果h[i]>h[j]生成的就是大根堆
}
func (h MaxHeap0295) Swap(i, j int) { // 绑定swap方法，交换两个元素位置
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap0295) Pop() interface{} { // 绑定pop方法，从最后拿出一个元素并返回
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MaxHeap0295) Push(x interface{}) { // 绑定push方法，插入新元素
	*h = append(*h, x.(int))
}

func (this *MaxHeap0295) Peek() interface{} {
	return (*this)[0]
}

type MinHeap0295 []int // 定义一个类型

func (h MinHeap0295) Len() int { return len(h) } // 绑定len方法,返回长度
func (h MinHeap0295) Less(i, j int) bool { // 绑定less方法
	return h[i] < h[j] // 如果h[i]<h[j]生成的就是小根堆，如果h[i]>h[j]生成的就是大根堆
}
func (h MinHeap0295) Swap(i, j int) { // 绑定swap方法，交换两个元素位置
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap0295) Pop() interface{} { // 绑定pop方法，从最后拿出一个元素并返回
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MinHeap0295) Push(x interface{}) { // 绑定push方法，插入新元素
	*h = append(*h, x.(int))
}

func (this *MinHeap0295) Peek() interface{} {
	return (*this)[0]
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
