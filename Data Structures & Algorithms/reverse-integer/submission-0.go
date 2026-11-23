// Idea: take the digits off the end of x one by one (x % 10) and push
// them onto the end of the result (result * 10 + digit).
// In Go, % keeps the sign, so negative numbers work without extra code.
// Before pushing a digit, check that the result would still fit in a
// 32-bit integer; if not, the answer is 0.
// Time: O(number of digits).
func reverse(x int) int {
	const maxInt32 = 2147483647
	const minInt32 = -2147483648

	result := 0
	for x != 0 {
		digit := x % 10
		x /= 10

		// would result*10 + digit leave the 32-bit range?
		if result > maxInt32/10 || (result == maxInt32/10 && digit > 7) {
			return 0
		}
		if result < minInt32/10 || (result == minInt32/10 && digit < -8) {
			return 0
		}

		result = result*10 + digit
	}
	return result
}
