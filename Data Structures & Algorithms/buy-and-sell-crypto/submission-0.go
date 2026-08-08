func maxProfit(prices []int) int {
	minPrice := 1000000000
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else {
			profit := price - minPrice
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}
