// Idea: to reach step i we come from step i-1 or step i-2, so
// ways(i) = ways(i-1) + ways(i-2). That is the Fibonacci sequence.
// We only need the last two values, not a whole array.
// Time: O(n).
func climbStairs(n int) int {
	prev := 1 // ways to reach step i-2
	cur := 1  // ways to reach step i-1

	for i := 2; i <= n; i++ {
		next := prev + cur
		prev = cur
		cur = next
	}
	return cur
}
