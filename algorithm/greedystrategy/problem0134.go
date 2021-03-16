package greedystrategy

func canCompleteCircuit(gas []int, cost []int) int {
	circuitLen := len(gas)
	firstPosition := 0
	currentPosition := 0
	var pathLen int
	var sumOfCost int
	var sumOfGas int
	for firstPosition < circuitLen {
		pathLen = 0
		sumOfCost = 0
		sumOfGas = 0
		for pathLen < circuitLen {
			currentPosition = (firstPosition + pathLen) % circuitLen
			sumOfGas += gas[currentPosition]
			sumOfCost += cost[currentPosition]
			if sumOfCost > sumOfGas {
				break
			} else {
				pathLen++
			}
		}
		if pathLen == circuitLen {
			return firstPosition
		} else {
			firstPosition += pathLen + 1
		}
	}
	return -1
}
