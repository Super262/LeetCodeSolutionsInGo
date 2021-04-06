package problem0380

import (
	"math/rand"
	"time"
)

type RandomizedSet struct {
	data        []int
	value2Index map[int]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{data: make([]int, 0), value2Index: make(map[int]int)}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (randSet *RandomizedSet) Insert(val int) bool {
	_, existed := randSet.value2Index[val]
	if existed {
		return false
	}
	randSet.data = append(randSet.data, val)
	randSet.value2Index[val] = len(randSet.data) - 1
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (randSet *RandomizedSet) Remove(val int) bool {
	_, existed := randSet.value2Index[val]
	if !existed {
		return false
	}
	currentIndex := randSet.value2Index[val]
	lastValue := randSet.data[len(randSet.data)-1]
	randSet.data[currentIndex] = lastValue
	randSet.value2Index[lastValue] = currentIndex
	delete(randSet.value2Index, val)
	randSet.data = randSet.data[0 : len(randSet.data)-1]
	return true
}

/** Get a random element from the set. */
func (randSet *RandomizedSet) GetRandom() int {
	rand.Seed(time.Now().UnixNano())
	return randSet.data[rand.Intn(len(randSet.data))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
