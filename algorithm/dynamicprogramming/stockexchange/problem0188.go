package stockexchange

func getMaxInt0188(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxProfit4(k int, prices []int) int {
	if k == 0 || prices == nil {
		return 0
	}
	pricesLen := len(prices)
	if pricesLen == 0 {
		return 0
	}
	buy := make([]int, k, k)
	sell := make([]int, k, k)
	for i := range buy {
		buy[i] = -prices[0]
	}
	for i := 1; i < pricesLen; i++ {
		buy[0] = getMaxInt0188(buy[0], -prices[i])
		sell[0] = getMaxInt0188(sell[0], buy[0]+prices[i])
		for j := 1; j < k; j++ {
			buy[j] = getMaxInt0188(buy[j], sell[j-1]-prices[i])
			sell[j] = getMaxInt0188(buy[j]+prices[i], sell[j])
		}
	}
	return sell[k-1]
}
