package stockexchange

func getMaxInt0123(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxProfit3(prices []int) int {
	if prices == nil {
		return 0
	}
	pricesLen := len(prices)
	if pricesLen == 0 {
		return 0
	}
	firstBuy := -prices[0]
	firstSell := 0
	secondBuy := -prices[0]
	secondSell := 0
	for i := 1; i < pricesLen; i++ {
		firstBuy = getMaxInt0123(firstBuy, -prices[i])
		firstSell = getMaxInt0123(firstSell, firstBuy+prices[i])
		secondBuy = getMaxInt0123(secondBuy, firstSell-prices[i])
		secondSell = getMaxInt0123(secondSell, secondBuy+prices[i])
	}
	return secondSell
}
