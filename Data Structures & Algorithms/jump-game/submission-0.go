// Idea (greedy): keep track of the farthest index we can reach so far.
// Walk through the array; if we are standing on an index that is beyond
// the farthest reach, we are stuck.
// Time: O(n).
func canJump(nums []int) bool {
	farthest := 0
	for i := 0; i < len(nums); i++ {
		// we cannot even get to index i
		if i > farthest {
			return false
		}
		// from i we can reach up to i + nums[i]
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}
	}
	return true
}
