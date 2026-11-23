// Idea: keep replacing the number with the sum of the squares of its
// digits. Either we reach 1 (happy), or we come back to a number we have
// seen before. Then we are in a loop and will never reach 1.
// A set of the numbers we have seen detects the loop.
func isHappy(n int) bool {
	seen := map[int]bool{}

	for n != 1 {
		if seen[n] {
			return false // we were here before: a loop
		}
		seen[n] = true
		n = sumOfDigitSquares(n)
	}
	return true
}

// sumOfDigitSquares returns, for example, 1*1 + 9*9 = 82 for 19.
func sumOfDigitSquares(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	return sum
}
