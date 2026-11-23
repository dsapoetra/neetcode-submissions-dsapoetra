// Idea (dynamic programming): dp[i] = length of the longest increasing
// subsequence that ENDS at index i.
// For each i, look at every earlier j with nums[j] < nums[i]:
// we can put nums[i] after it, so dp[i] = dp[j] + 1.
// Time: O(n^2).
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	best := 0

	for i := 0; i < len(nums); i++ {
		dp[i] = 1 // the number alone is a subsequence of length 1

		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}

		if dp[i] > best {
			best = dp[i]
		}
	}
	return best
}
