package heap

import "container/heap"

func nthUglyNumber(n int) int {
	visited := make(map[int64]bool)
	minHeap := &MinHeap0264{}
	factors := []int64{2, 3, 5}
	heap.Push(minHeap, int64(1))
	visited[1] = true
	var currentVal int64
	for i := 0; i < n; i++ {
		currentVal = heap.Pop(minHeap).(int64)
		for _, f := range factors {
			nextVal := f * currentVal
			_, existed := visited[nextVal]
			if existed {
				continue
			}
			visited[nextVal] = true
			heap.Push(minHeap, nextVal)
		}
	}
	return int(currentVal)
}

type MinHeap0264 []int64 // 定义一个类型

func (minHeap MinHeap0264) Len() int { return len(minHeap) } // 绑定len方法,返回长度
func (minHeap MinHeap0264) Less(i, j int) bool { // 绑定less方法
	return minHeap[i] < minHeap[j] // 如果h[i]<minHeap[j]生成的就是小根堆，如果h[i]>minHeap[j]生成的就是大根堆
}
func (minHeap MinHeap0264) Swap(i, j int) { // 绑定swap方法，交换两个元素位置
	minHeap[i], minHeap[j] = minHeap[j], minHeap[i]
}

func (minHeap *MinHeap0264) Pop() interface{} { // 绑定pop方法，从最后拿出一个元素并返回
	old := *minHeap
	n := len(old)
	x := old[n-1]
	*minHeap = old[0 : n-1]
	return x
}

func (minHeap *MinHeap0264) Push(x interface{}) { // 绑定push方法，插入新元素
	*minHeap = append(*minHeap, x.(int64))
}

func (minHeap *MinHeap0264) Peek() interface{} {
	return (*minHeap)[0]
}
