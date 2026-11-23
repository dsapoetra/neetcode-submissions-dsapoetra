// Idea: same as Longest Palindromic Substring. Try every center and grow
// outwards. Every time both sides match we found one more palindrome.
// Time: O(n^2).
func countSubstrings(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		count += countFromCenter(s, i, i)   // odd length palindromes
		count += countFromCenter(s, i, i+1) // even length palindromes
	}
	return count
}

// countFromCenter counts the palindromes that have this center.
func countFromCenter(s string, left int, right int) int {
	count := 0
	for left >= 0 && right < len(s) && s[left] == s[right] {
		count++
		left--
		right++
	}
	return count
}
