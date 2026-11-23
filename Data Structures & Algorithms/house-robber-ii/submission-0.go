// Idea: the houses are in a circle, so the first and last house are
// neighbors and we cannot rob both. So solve the normal (straight line)
// House Robber twice:
//   1) without the last house
//   2) without the first house
// and take the better result.
// Time: O(n).
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	withoutLast := robLine(nums[:len(nums)-1])
	withoutFirst := robLine(nums[1:])
	if withoutLast > withoutFirst {
		return withoutLast
	}
	return withoutFirst
}

// robLine solves House Robber for houses in a straight line.
func robLine(nums []int) int {
	twoBack := 0 // best money up to two houses back
	oneBack := 0 // best money up to the previous house
	for _, money := range nums {
		best := oneBack // skip this house
		if twoBack+money > best {
			best = twoBack + money // rob this house
		}
		twoBack = oneBack
		oneBack = best
	}
	return oneBack
}
