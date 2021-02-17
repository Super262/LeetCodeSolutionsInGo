package stockexchange

func getMaxInt0714(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxProfit2(prices []int, fee int) int {
	if prices == nil {
		return 0
	}
	pricesLen := len(prices)
	if pricesLen == 0 {
		return 0
	}
	prev0 := -prices[0]
	prev1 := 0
	var cur0, cur1 int
	for i := 1; i < pricesLen; i++ {
		cur0 = getMaxInt0714(prev1-prices[i], prev0)
		cur1 = getMaxInt0714(prev0+prices[i]-fee, prev1)
		prev0 = cur0
		prev1 = cur1
	}
	return prev1
}
