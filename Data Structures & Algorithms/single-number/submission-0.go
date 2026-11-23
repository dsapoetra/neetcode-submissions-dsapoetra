// Idea (XOR): the ^ operator has two handy rules:
//   x ^ x = 0   (a number cancels itself)
//   x ^ 0 = x
// XOR all the numbers together: every pair cancels out and only the
// number without a partner is left.
// Time: O(n), no extra memory.
func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}
