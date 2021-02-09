package greedystrategy

func maxProfit(prices []int) int {
	if prices == nil {
		return 0
	}
	preMin := prices[0]
	profit := 0
	tempDis := 0
	for _, p := range prices {
		if preMin < p {
			tempDis = p - preMin
			if tempDis > profit {
				profit = tempDis
			}
		} else {
			preMin = p
		}
	}
	return profit
}
