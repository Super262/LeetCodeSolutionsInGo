package others

import (
	"container/heap"
	"math"
)

func kthSmallest(matrix [][]int, k int) int {
	if k <= 0 || matrix == nil {
		return math.MaxInt32
	}
	h := len(matrix)
	if h == 0 {
		return math.MaxInt32
	}
	w := len(matrix[0])
	if w == 0 {
		return math.MaxInt32
	}
	used := make([][]bool, h, h)
	for i := range used {
		used[i] = make([]bool, w, w)
	}
	minHeap := &Pair0378Heap{}
	heap.Init(minHeap)
	heap.Push(minHeap, Pair0378{matrix[0][0], 0, 0})
	used[0][0] = true
	curX := 0
	curY := 0
	for i := 0; i < k-1; i++ {
		currentPair := heap.Pop(minHeap).(Pair0378)
		curX = currentPair.x
		curY = currentPair.y
		if curX+1 < w && !used[curY][curX+1] {
			heap.Push(minHeap, Pair0378{matrix[curY][curX+1], curY, curX + 1})
			used[curY][curX+1] = true
		}
		if curY+1 < h && !used[curY+1][curX] {
			heap.Push(minHeap, Pair0378{matrix[curY+1][curX], curY + 1, curX})
			used[curY+1][curX] = true
		}
	}
	return (*minHeap)[0].value
}

//定义一个结构体
type Pair0378 struct {
	value int
	y     int
	x     int
}

// 定义一个堆结构体
type Pair0378Heap []Pair0378

// 实现heap.Interface接口
func (pairs Pair0378Heap) Len() int {
	return len(pairs)
}

// 实现sort.Iterface
func (pairs Pair0378Heap) Swap(i, j int) {
	pairs[i], pairs[j] = pairs[j], pairs[i]
}
func (pairs Pair0378Heap) Less(i, j int) bool {
	return pairs[i].value < pairs[j].value
}

// 实现heap.Interface接口定义的额外方法
func (pairs *Pair0378Heap) Push(h interface{}) {
	*pairs = append(*pairs, h.(Pair0378))
}
func (pairs *Pair0378Heap) Pop() (x interface{}) {
	n := len(*pairs)
	x = (*pairs)[n-1]          // 返回删除的元素
	*pairs = (*pairs)[0 : n-1] // [n:m]不包括下标为m的元素
	return x
}
