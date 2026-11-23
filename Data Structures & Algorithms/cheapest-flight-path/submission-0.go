// Idea (Bellman-Ford with a limit): price[x] = the cheapest price we know
// to reach city x. At most k stops means at most k+1 flights.
// So do k+1 rounds; in each round every flight may be used ONCE.
// Important: a round must read the prices of the PREVIOUS round (a copy).
// Otherwise one round could chain several flights together.
// Time: O(k * flights).
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	const inf = 1000000000

	price := make([]int, n)
	for i := range price {
		price[i] = inf
	}
	price[src] = 0

	for round := 0; round <= k; round++ {
		// copy, so this round only extends routes from the round before
		next := make([]int, n)
		copy(next, price)

		for _, flight := range flights {
			from, to, cost := flight[0], flight[1], flight[2]
			if price[from] != inf && price[from]+cost < next[to] {
				next[to] = price[from] + cost
			}
		}
		price = next
	}

	if price[dst] == inf {
		return -1
	}
	return price[dst]
}
