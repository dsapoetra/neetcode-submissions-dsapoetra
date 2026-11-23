// Idea (dynamic programming): ways[a] = number of ways to make amount a.
// Add the coins ONE BY ONE. With a new coin, every way to make
// (a - coin) becomes a way to make a, by adding this coin.
// Because the coin loop is the OUTER loop, each combination is counted
// once (1+2 and 2+1 are the same combination, not two).
// Time: O(coins * amount).
func change(amount int, coins []int) int {
	ways := make([]int, amount+1)
	ways[0] = 1 // one way to make 0: use no coins

	for _, coin := range coins {
		for a := coin; a <= amount; a++ {
			ways[a] += ways[a-coin]
		}
	}
	return ways[amount]
}
