// ways(n) = ways(n-1) + ways(n-2), same as Fibonacci.
func climbStairs(n int) int {
	prev, cur := 1, 1
	for i := 2; i <= n; i++ {
		prev, cur = cur, prev+cur
	}
	return cur
}
