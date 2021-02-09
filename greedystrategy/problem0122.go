package greedystrategy

func maxProfit2(prices []int) int {
	result := 0
	if prices == nil {
		return 0
	}
	pLen := len(prices)
	if pLen == 0 {
		return 0
	}
	for i := 1; i < pLen; i++ {
		if prices[i] > prices[i-1] {
			result += prices[i] - prices[i-1]
		}
	}
	return result
}
