// Idea: a palindrome is a mirror around its center. Try every position
// as a center and grow outwards while the letters on both sides match.
// There are two kinds of center:
//   - one letter      (odd length, like "aba")
//   - between letters (even length, like "abba")
// Time: O(n^2).
func longestPalindrome(s string) string {
	bestStart := 0
	bestLen := 0

	for i := 0; i < len(s); i++ {
		// odd length: center is letter i
		left, right := growPalindrome(s, i, i)
		if right-left+1 > bestLen {
			bestStart = left
			bestLen = right - left + 1
		}

		// even length: center is between i and i+1
		left, right = growPalindrome(s, i, i+1)
		if right-left+1 > bestLen {
			bestStart = left
			bestLen = right - left + 1
		}
	}
	return s[bestStart : bestStart+bestLen]
}

// growPalindrome grows outwards from (left, right) while the letters match.
// It returns the first and last index of the palindrome it found.
func growPalindrome(s string, left int, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	// the loop went one step too far on both sides
	return left + 1, right - 1
}
