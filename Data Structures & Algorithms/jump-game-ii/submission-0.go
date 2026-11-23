// Idea (greedy): think in "rounds". With 1 jump we can reach some range of
// positions, with 2 jumps a further range, and so on.
//   currentEnd = the last position we can reach with the jumps made so far
//   farthest   = the farthest position we could reach with one more jump
// Walk through the list. When we arrive at currentEnd we MUST jump again,
// and the new range ends at farthest.
// Time: O(n).
func jump(nums []int) int {
	jumps := 0
	currentEnd := 0
	farthest := 0

	// we stop before the last index: standing there needs no more jumps
	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}

		if i == currentEnd {
			jumps++ // one more jump
			currentEnd = farthest
		}
	}
	return jumps
}
