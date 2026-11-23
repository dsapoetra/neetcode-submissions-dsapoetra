// Idea: walk through the prices once. Remember the cheapest price seen
// so far (the best day to buy). For every day, check the profit if we
// sold today, and keep the biggest profit.
// Time: O(n).
func maxProfit(prices []int) int {
	minPrice := prices[0] // cheapest price so far
	best := 0             // best profit so far (0 = do not trade)

	for _, price := range prices {
		if price < minPrice {
			minPrice = price // found a cheaper day to buy
		}
		profit := price - minPrice // buy at the cheapest, sell today
		if profit > best {
			best = profit
		}
	}
	return best
}
