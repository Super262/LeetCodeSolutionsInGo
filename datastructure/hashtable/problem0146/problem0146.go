package problem0146

type LRUCache struct {
	capacity int
	data     map[int]*DLinkNode0146
	head     *DLinkNode0146
	tail     *DLinkNode0146
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{capacity: capacity, data: make(map[int]*DLinkNode0146), head: &DLinkNode0146{key: -1, value: -1, prev: nil, next: nil}, tail: &DLinkNode0146{key: -1, value: -1, prev: nil, next: nil}}
	cache.head.next = cache.tail
	cache.tail.prev = cache.head
	return cache
}

func (cache *LRUCache) Get(key int) int {
	currentNode, existed := cache.data[key]
	if !existed {
		return -1
	}
	currentNode = RemoveNode0146(currentNode)
	cache.AddNodeOnLeft0146(currentNode)
	return currentNode.value
}

func (cache *LRUCache) Put(key int, value int) {
	_, existed := cache.data[key]
	var currentNode *DLinkNode0146
	if existed {
		currentNode = RemoveNode0146(cache.data[key])
		currentNode.value = value
	} else {
		if len(cache.data) == cache.capacity {
			deadNode := cache.PopNodeOnRight0146()
			if deadNode != nil {
				delete(cache.data, deadNode.key)
			}
		}
		currentNode = &DLinkNode0146{key: key, value: value, prev: nil, next: nil}
		cache.data[key] = currentNode
	}
	cache.AddNodeOnLeft0146(currentNode)
}

func RemoveNode0146(target *DLinkNode0146) *DLinkNode0146 {
	if target == nil {
		return nil
	}
	target.prev.next = target.next
	target.next.prev = target.prev
	target.next = nil
	target.prev = nil
	return target
}

func (cache *LRUCache) AddNodeOnLeft0146(target *DLinkNode0146) {
	if target == nil {
		return
	}
	target.next = cache.head.next
	target.prev = cache.head
	target.next.prev = target
	cache.head.next = target
}

func (cache *LRUCache) PopNodeOnRight0146() *DLinkNode0146 {
	target := cache.tail.prev
	return RemoveNode0146(target)
}

type DLinkNode0146 struct {
	key   int
	value int
	prev  *DLinkNode0146
	next  *DLinkNode0146
}
