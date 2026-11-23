// Idea: at the end of every day we are in one of three situations.
// Keep the best profit for each of them:
//   holding = we own a coin
//   sold    = we sold today (tomorrow is the cooldown day)
//   resting = we own nothing and did not sell today (free to buy tomorrow)
// Each day the new values come from yesterday's values:
//   holding: keep holding, OR buy today (only allowed after resting)
//   sold:    we were holding and sell today
//   resting: keep resting, OR yesterday we sold (cooldown is over)
// Time: O(n), no extra memory.
func maxProfit(prices []int) int {
	holding := -prices[0] // bought on the first day
	sold := 0
	resting := 0

	for i := 1; i < len(prices); i++ {
		price := prices[i]

		newHolding := holding
		if resting-price > newHolding {
			newHolding = resting - price // buy today
		}

		newSold := holding + price // sell today

		newResting := resting
		if sold > newResting {
			newResting = sold // yesterday's sale, cooldown done
		}

		holding, sold, resting = newHolding, newSold, newResting
	}

	// the best result is never "still holding a coin"
	if sold > resting {
		return sold
	}
	return resting
}
