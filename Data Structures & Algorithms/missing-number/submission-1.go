// Idea: the numbers should be 0, 1, 2, ..., n. Add up 0..n and subtract
// every number we actually have. What is left is the missing number.
// Time: O(n).
func missingNumber(nums []int) int {
	missing := 0
	for i := 0; i < len(nums); i++ {
		missing += (i + 1) // add the expected number (1..n)
		missing -= nums[i] // remove the number we have
	}
	return missing
}
