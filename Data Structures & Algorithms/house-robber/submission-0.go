// For each house: either skip it, or rob it plus the best from two houses back.
func rob(nums []int) int {
	twoBack, oneBack := 0, 0
	for _, n := range nums {
		best := oneBack
		if twoBack+n > best {
			best = twoBack + n
		}
		twoBack = oneBack
		oneBack = best
	}
	return oneBack
}
