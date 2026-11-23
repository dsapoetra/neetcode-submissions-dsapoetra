// Idea: at each house we have two choices:
//   - skip it  -> money stays the same as the previous house
//   - rob it   -> money = this house + best up to two houses back
// Take the bigger one. We only need the last two answers.
// Time: O(n).
func rob(nums []int) int {
	twoBack := 0 // best money up to two houses back
	oneBack := 0 // best money up to the previous house

	for _, money := range nums {
		best := oneBack // choice 1: skip this house
		if twoBack+money > best {
			best = twoBack + money // choice 2: rob this house
		}
		// slide forward by one house
		twoBack = oneBack
		oneBack = best
	}
	return oneBack
}
