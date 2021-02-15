package stockexchange

func getMaxInt0309(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxProfit(prices []int) int {
	if prices == nil {
		return 0
	}
	pricesLen := len(prices)
	if pricesLen == 0 {
		return 0
	}
	prev0 := -prices[0] // has some stock
	prev1 := 0          // cooling
	prev2 := 0          // cooled
	var cur0, cur1, cur2 int
	for i := 1; i < pricesLen; i++ {
		cur0 = getMaxInt0309(prev0, prev2-prices[i])
		cur1 = prev0 + prices[i]
		cur2 = getMaxInt0309(prev1, prev2)
		prev0 = cur0
		prev1 = cur1
		prev2 = cur2
	}
	return getMaxInt0309(prev1, prev2)
}
