func maxProfit(prices []int) int {
	buy := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < buy {
			buy = prices[i]
		}
		if prices[i]-buy > maxProfit {
			maxProfit = prices[i] - buy
		}
	}

	return maxProfit
}
