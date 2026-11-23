// Idea (backtracking): sort the numbers, then build combinations one
// number at a time. Each number can be used only once, so after choosing
// index i we continue from i+1.
// To avoid the same combination twice: at one position of the combination
// never try the same VALUE twice (skip equal neighbors in the sorted list).
// Time: O(2^n) in the worst case.
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := [][]int{}
	current := []int{}

	// build tries every number from index start on as the next number.
	var build func(start int, remaining int)
	build = func(start int, remaining int) {
		if remaining == 0 {
			// found one: save a COPY, because current keeps changing
			res = append(res, append([]int{}, current...))
			return
		}

		for i := start; i < len(candidates); i++ {
			// same value as the previous try at this position: skip
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			// sorted list: everything after this is too big as well
			if candidates[i] > remaining {
				break
			}

			current = append(current, candidates[i]) // choose
			build(i+1, remaining-candidates[i])      // go deeper
			current = current[:len(current)-1]       // undo the choice
		}
	}

	build(0, target)
	return res
}
