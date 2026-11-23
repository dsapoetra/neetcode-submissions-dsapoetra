// Remember the cheapest price so far, try selling at every price.
func maxProfit(prices []int) int {
	minPrice := prices[0]
	best := 0
	for _, p := range prices {
		if p < minPrice {
			minPrice = p
		}
		if p-minPrice > best {
			best = p - minPrice
		}
	}
	return best
}
