// Idea (backtracking): build a combination one number at a time.
// At each step try every number from index "start" onwards:
//   pick it -> go deeper -> un-pick it (backtrack) and try the next one.
// We pass the same index again when going deeper because a number can
// be used many times. Starting from "start" avoids duplicates like
// [2,3] and [3,2].
func combinationSum(nums []int, target int) [][]int {
	res := [][]int{}
	current := []int{} // the combination we are building

	var backtrack func(start int, remaining int)
	backtrack = func(start int, remaining int) {
		// we hit the target exactly: save a COPY of current
		if remaining == 0 {
			saved := make([]int, len(current))
			copy(saved, current)
			res = append(res, saved)
			return
		}

		for i := start; i < len(nums); i++ {
			// this number is too big, skip it
			if nums[i] > remaining {
				continue
			}
			current = append(current, nums[i])    // pick
			backtrack(i, remaining-nums[i])       // go deeper (i again: reuse allowed)
			current = current[:len(current)-1]    // un-pick
		}
	}

	backtrack(0, target)
	return res
}
