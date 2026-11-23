// Idea (dynamic programming): dp[a] = fewest coins needed to make amount a.
// To make amount a, try every coin c as the last coin:
//   dp[a] = dp[a-c] + 1   (take the smallest over all coins)
// Time: O(amount * number of coins).
func coinChange(coins []int, amount int) int {
	// amount+1 is used as "impossible" because we never need more
	// than amount coins (the smallest coin is 1).
	impossible := amount + 1

	dp := make([]int, amount+1)
	for a := 1; a <= amount; a++ {
		dp[a] = impossible
	}
	// dp[0] is already 0: zero coins make amount 0

	for a := 1; a <= amount; a++ {
		for _, c := range coins {
			// coin c can be the last coin only if it is not bigger than a
			if c <= a && dp[a-c]+1 < dp[a] {
				dp[a] = dp[a-c] + 1
			}
		}
	}

	if dp[amount] == impossible {
		return -1
	}
	return dp[amount]
}
