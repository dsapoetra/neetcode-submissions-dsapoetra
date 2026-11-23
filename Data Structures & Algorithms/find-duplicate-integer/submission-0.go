// Idea: the numbers are between 1 and n, so we can use a list of
// true/false flags: seen[x] = true once we have met the number x.
// The first number we meet twice is the answer.
// Time: O(n), extra space: O(n).
//
// (The follow-up asks for no extra memory. That needs the "slow and fast
// pointer" cycle trick, which is clever but much harder to understand.)
func findDuplicate(nums []int) int {
	seen := make([]bool, len(nums))

	for _, num := range nums {
		if seen[num] {
			return num
		}
		seen[num] = true
	}
	return -1 // never happens: the problem promises a duplicate
}
