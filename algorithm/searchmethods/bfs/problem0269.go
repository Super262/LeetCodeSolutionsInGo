package bfs

import "container/heap"

func alienOrder(words []string) string {
	if words == nil || len(words) == 0 {
		return ""
	}
	if len(words) == 1 {
		return words[0]
	}
	charArrays := make([][]rune, len(words), len(words))
	for _, w := range words {
		charArrays = append(charArrays, []rune(w))
	}
	inDegree := make(map[rune]int)
	neighbors := make(map[rune]map[rune]int)
	for _, w := range charArrays {
		for _, ch := range w {
			inDegree[ch] = 0
			_, ok := neighbors[ch]
			if !ok {
				neighbors[ch] = make(map[rune]int)
				neighbors[ch][ch] = 1
			}
		}
	}
	for i := 1; i < len(charArrays); i++ {
		varianceIndex := getVariance(&charArrays[i-1], &charArrays[i])
		if varianceIndex == -1 {
			if len(charArrays[i-1]) > len(charArrays[i]) {
				return ""
			}
			continue
		}
		_, ok := neighbors[charArrays[i-1][varianceIndex]][charArrays[i][varianceIndex]]
		if !ok {
			neighbors[charArrays[i-1][varianceIndex]][charArrays[i][varianceIndex]] = 1
			inDegree[charArrays[i][varianceIndex]]++
		}
	}
	minHeap := &MinHeap0269{}
	for v, d := range inDegree {
		if d == 0 {
			heap.Push(minHeap, v)
		}
	}
	result := make([]rune, 0, len(inDegree))
	for minHeap.Len() > 0 {
		node := heap.Pop(minHeap).(rune)
		result = append(result, node)
		for n, _ := range neighbors[node] {
			inDegree[n]--
			if inDegree[n] == 0 {
				heap.Push(minHeap, n)
			}
		}
	}
	if len(result) == len(inDegree) {
		return string(result)
	}
	return ""
}

func getVariance(word1 *[]rune, word2 *[]rune) int {
	if word1 == nil || word2 == nil {
		return -1
	}
	i := 0
	for i < len(*word1) && i < len(*word2) {
		if (*word1)[i] != (*word2)[i] {
			return i
		}
		i++
	}
	return -1
}

type MinHeap0269 []rune // 定义一个类型

func (minHeap MinHeap0269) Len() int {
	return len(minHeap) // 绑定len方法,返回长度
}
func (minHeap MinHeap0269) Less(i, j int) bool { // 绑定less方法
	return minHeap[i] < minHeap[j] // 如果h[i]<minHeap[j]生成的就是小根堆，如果h[i]>minHeap[j]生成的就是大根堆
}
func (minHeap MinHeap0269) Swap(i, j int) { // 绑定swap方法，交换两个元素位置
	minHeap[i], minHeap[j] = minHeap[j], minHeap[i]
}

func (minHeap *MinHeap0269) Pop() interface{} { // 绑定pop方法，从最后拿出一个元素并返回
	old := *minHeap
	n := len(old)
	x := old[n-1]
	*minHeap = old[0 : n-1]
	return x
}

func (minHeap *MinHeap0269) Push(x interface{}) { // 绑定push方法，插入新元素
	*minHeap = append(*minHeap, x.(rune))
}

func (minHeap *MinHeap0269) Peek() interface{} {
	return (*minHeap)[0]
}
