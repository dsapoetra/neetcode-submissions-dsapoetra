// Idea: i/2 is i with its last bit removed. So:
//   ones in i = ones in i/2 + the last bit of i (i % 2)
// We already know the answer for i/2 because it is smaller than i.
// Example: 6 = 110 -> ones(3 = 11) + 0 = 2
// Time: O(n).
func countBits(n int) []int {
	bits := make([]int, n+1) // bits[0] is 0
	for i := 1; i <= n; i++ {
		bits[i] = bits[i/2] + i%2
	}
	return bits
}
