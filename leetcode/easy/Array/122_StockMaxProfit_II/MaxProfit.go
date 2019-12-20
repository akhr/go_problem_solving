package maxprofit

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	buyPrice := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if profit := prices[i] - buyPrice; profit > 0 {
			maxProfit += profit
		}
		buyPrice = prices[i]
	}
	return maxProfit
}
