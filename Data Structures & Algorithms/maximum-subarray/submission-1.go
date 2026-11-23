// Idea (Kadane's algorithm): keep a running sum. If the running sum
// becomes negative it can only hurt us, so throw it away and start again.
// Time: O(n).
func maxSubArray(nums []int) int {
	best := nums[0] // best sum found so far
	cur := 0        // sum of the subarray we are building now

	for _, n := range nums {
		// a negative running sum is useless, restart from here
		if cur < 0 {
			cur = 0
		}
		cur += n

		if cur > best {
			best = cur
		}
	}
	return best
}
