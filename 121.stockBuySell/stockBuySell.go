package stockbuysell

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	minValue, profit := prices[0], 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < minValue {
			minValue = prices[i]
			continue
		}

		if prices[i]-minValue > profit {
			profit = prices[i] - minValue
		}
	}

	return profit
}
