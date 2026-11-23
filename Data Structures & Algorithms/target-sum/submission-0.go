// Idea: go through the numbers one by one and remember every total we
// can have so far, together with HOW MANY ways lead to that total.
// For the next number, each total splits in two: total + num and
// total - num.
// At the end, look up how many ways gave exactly the target.
// Time: O(n * number of different totals).
func findTargetSumWays(nums []int, target int) int {
	ways := map[int]int{0: 1} // before any number: total 0, one way

	for _, num := range nums {
		next := map[int]int{}
		for total, count := range ways {
			next[total+num] += count // put a "+" in front of num
			next[total-num] += count // put a "-" in front of num
		}
		ways = next
	}
	return ways[target]
}
