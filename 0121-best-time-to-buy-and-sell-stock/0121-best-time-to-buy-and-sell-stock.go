func maxProfit(prices []int) int {
	buy, profit := prices[0], 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < buy {
			buy = prices[i]
		} else {
			sell := prices[i] - buy
			if sell > profit {
				profit = sell
			}
		}
	}
	return profit
}