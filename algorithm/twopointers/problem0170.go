package twopointers

type TwoSum struct {
	count map[int]int
}

/** Initialize your data structure here. */
func Constructor() TwoSum {
	return TwoSum{count: make(map[int]int)}
}

/** Add the number to an internal data structure.. */
func (this *TwoSum) Add(number int) {
	c, ok := this.count[number]
	if ok {
		this.count[number] = c + 1
	} else {
		this.count[number] = 1
	}
}

/** Find if there exists any pair of numbers which sum is equal to the value. */
func (this *TwoSum) Find(value int) bool {
	num2 := 0
	c := 0
	ok := false
	for num1 := range this.count {
		num2 = value - num1
		c, ok = this.count[num2]
		if !ok {
			continue
		}
		if num1 == num2 && c > 1 {
			return true
		}
		if num1 != num2 && c > 0 {
			return true
		}
	}
	return false
}

/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */
