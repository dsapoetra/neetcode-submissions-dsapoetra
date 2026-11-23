// Idea: add 1 like on paper, starting at the last digit.
//   - a digit smaller than 9 just goes up by one, and we are done
//   - a 9 becomes 0 and the "carry" moves on to the digit on its left
// If every digit was a 9 (like 999), we fall out of the loop and need a
// new 1 in front: 1000.
// Time: O(n).
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0 // 9 + 1 = 10: write 0, carry on
	}

	// all digits were 9
	return append([]int{1}, digits...)
}
