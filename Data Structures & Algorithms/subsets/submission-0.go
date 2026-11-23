// Idea: start with only the empty subset. Then take the numbers one by
// one: for every subset we already have, make a copy that ALSO contains
// the new number. That doubles the list each time.
// Example with [1, 2]:
//   start:    []
//   add 1:    [], [1]
//   add 2:    [], [1], [2], [1,2]
// Time: O(n * 2^n), there are 2^n subsets.
func subsets(nums []int) [][]int {
	res := [][]int{{}} // just the empty subset

	for _, num := range nums {
		count := len(res) // only loop over the subsets that exist right now
		for i := 0; i < count; i++ {
			// copy the old subset and put the new number at the end
			withNum := make([]int, len(res[i]), len(res[i])+1)
			copy(withNum, res[i])
			withNum = append(withNum, num)
			res = append(res, withNum)
		}
	}
	return res
}
