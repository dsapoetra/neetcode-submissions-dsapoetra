// Idea (backtracking): fill the permutation position by position. For each
// position try every number that is not used yet, go deeper, then take
// the number back out and try the next one.
// Time: O(n * n!), there are n! permutations.
func permute(nums []int) [][]int {
	res := [][]int{}
	current := []int{}
	used := make([]bool, len(nums))

	var build func()
	build = func() {
		if len(current) == len(nums) {
			// full permutation: save a COPY, because current keeps changing
			res = append(res, append([]int{}, current...))
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true // choose
			current = append(current, nums[i])

			build() // fill the next position

			current = current[:len(current)-1] // undo the choice
			used[i] = false
		}
	}

	build()
	return res
}
