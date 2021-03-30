package heap

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

func (finder *MedianFinder) AddNum(num int) {
	if finder.maxHeap.Len() == 0 || num <= finder.maxHeap.Peek().(int) {
		heap.Push(finder.maxHeap, num)
	} else {
		heap.Push(finder.minHeap, num)
	}
	finder.Balance()
}

func (finder *MedianFinder) FindMedian() float64 {
	if finder.maxHeap.Len() == finder.minHeap.Len() {
		return float64(finder.maxHeap.Peek().(int)+finder.minHeap.Peek().(int)) / 2.0
	} else {
		return float64(finder.maxHeap.Peek().(int))
	}
}

func (finder *MedianFinder) Balance() {
	for finder.maxHeap.Len() < finder.minHeap.Len() {
		heap.Push(finder.maxHeap, heap.Pop(finder.minHeap))
	}
	for finder.maxHeap.Len() > finder.minHeap.Len()+1 {
		heap.Push(finder.minHeap, heap.Pop(finder.maxHeap))
	}
}

type MaxHeap0295 []int // 定义一个类型

func (maxHeap MaxHeap0295) Len() int { return len(maxHeap) } // 绑定len方法,返回长度
func (maxHeap MaxHeap0295) Less(i, j int) bool { // 绑定less方法
	return maxHeap[i] > maxHeap[j] // 如果h[i]<maxHeap[j]生成的就是小根堆，如果h[i]>maxHeap[j]生成的就是大根堆
}
func (maxHeap MaxHeap0295) Swap(i, j int) { // 绑定swap方法，交换两个元素位置
	maxHeap[i], maxHeap[j] = maxHeap[j], maxHeap[i]
}

func (maxHeap *MaxHeap0295) Pop() interface{} { // 绑定pop方法，从最后拿出一个元素并返回
	old := *maxHeap
	n := len(old)
	x := old[n-1]
	*maxHeap = old[0 : n-1]
	return x
}

func (maxHeap *MaxHeap0295) Push(x interface{}) { // 绑定push方法，插入新元素
	*maxHeap = append(*maxHeap, x.(int))
}

func (maxHeap *MaxHeap0295) Peek() interface{} {
	return (*maxHeap)[0]
}

type MinHeap0295 []int // 定义一个类型

func (minHeap MinHeap0295) Len() int { return len(minHeap) } // 绑定len方法,返回长度
func (minHeap MinHeap0295) Less(i, j int) bool { // 绑定less方法
	return minHeap[i] < minHeap[j] // 如果h[i]<minHeap[j]生成的就是小根堆，如果h[i]>minHeap[j]生成的就是大根堆
}
func (minHeap MinHeap0295) Swap(i, j int) { // 绑定swap方法，交换两个元素位置
	minHeap[i], minHeap[j] = minHeap[j], minHeap[i]
}

func (minHeap *MinHeap0295) Pop() interface{} { // 绑定pop方法，从最后拿出一个元素并返回
	old := *minHeap
	n := len(old)
	x := old[n-1]
	*minHeap = old[0 : n-1]
	return x
}

func (minHeap *MinHeap0295) Push(x interface{}) { // 绑定push方法，插入新元素
	*minHeap = append(*minHeap, x.(int))
}

func (minHeap *MinHeap0295) Peek() interface{} {
	return (*minHeap)[0]
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
