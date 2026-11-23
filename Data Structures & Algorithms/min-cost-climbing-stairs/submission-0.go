// Idea (dynamic programming): reach[i] = the cheapest total cost to stand
// on step i (before paying for step i itself).
// We may start on step 0 or step 1 for free.
// To arrive at step i we come from step i-1 or step i-2, and we pay the
// cost of the step we leave.
// The "top" is one step past the last one, so the answer is reach[n].
// Time: O(n).
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	reach := make([]int, n+1) // reach[0] = reach[1] = 0: free starting points

	for i := 2; i <= n; i++ {
		fromOneBelow := reach[i-1] + cost[i-1]
		fromTwoBelow := reach[i-2] + cost[i-2]

		if fromOneBelow < fromTwoBelow {
			reach[i] = fromOneBelow
		} else {
			reach[i] = fromTwoBelow
		}
	}
	return reach[n]
}
