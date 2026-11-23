// Idea (backtracking): sort the numbers so equal numbers are next to each
// other. Every state of "current" is one subset, so save it at every step.
// Then try to add each later number.
// To avoid the same subset twice: at one position never try the same
// VALUE twice (skip equal neighbors).
// Time: O(n * 2^n).
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	current := []int{}

	var build func(start int)
	build = func(start int) {
		// current is a subset: save a copy
		res = append(res, append([]int{}, current...))

		for i := start; i < len(nums); i++ {
			// same value as the previous try at this position: skip
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			current = append(current, nums[i]) // choose
			build(i + 1)                       // go deeper
			current = current[:len(current)-1] // undo the choice
		}
	}

	build(0)
	return res
}
