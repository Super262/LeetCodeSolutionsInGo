package heap

import "container/heap"

func mergeKLists(lists []*ListNode) *ListNode {
	minHeap := &MinHeap0023{}
	for _, l := range lists {
		if l == nil {
			continue
		}
		heap.Push(minHeap, l)
	}
	resultHead := &ListNode{}
	resultTop := resultHead
	for minHeap.Len() > 0 {
		currentNode := heap.Pop(minHeap).(*ListNode)
		resultTop.Next = currentNode
		if currentNode.Next != nil {
			heap.Push(minHeap, currentNode.Next)
		}
		currentNode = currentNode.Next
		resultTop = resultTop.Next
	}
	return resultHead.Next
}

type MinHeap0023 []*ListNode // 定义一个类型

func (minHeap MinHeap0023) Len() int {
	return len(minHeap) // 绑定len方法,返回长度
}
func (minHeap MinHeap0023) Less(i, j int) bool { // 绑定less方法
	return minHeap[i].Val < minHeap[j].Val // 如果h[i]<minHeap[j]生成的就是小根堆，如果h[i]>minHeap[j]生成的就是大根堆
}
func (minHeap MinHeap0023) Swap(i, j int) { // 绑定swap方法，交换两个元素位置
	minHeap[i], minHeap[j] = minHeap[j], minHeap[i]
}

func (minHeap *MinHeap0023) Pop() interface{} { // 绑定pop方法，从最后拿出一个元素并返回
	old := *minHeap
	n := len(old)
	x := old[n-1]
	*minHeap = old[0 : n-1]
	return x
}

func (minHeap *MinHeap0023) Push(x interface{}) { // 绑定push方法，插入新元素
	*minHeap = append(*minHeap, x.(*ListNode))
}

func (minHeap *MinHeap0023) Peek() interface{} {
	return (*minHeap)[0]
}
