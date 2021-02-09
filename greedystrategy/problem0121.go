package greedystrategy

func maxProfit(prices []int) int {
	if prices == nil {
		return 0
	}
	pLen := len(prices)
	if pLen == 0 {
		return 0
	}
	preMin := prices[0]
	profit := 0
	tempDis := 0
	for i := 1; i < pLen; i++ {
		if preMin < prices[i] {
			tempDis = prices[i] - preMin
			if tempDis > profit {
				profit = tempDis
			}
		} else {
			preMin = prices[i]
		}
	}
	return profit
}
