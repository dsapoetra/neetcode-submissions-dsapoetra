// Idea (backtracking): cut the string from left to right. At position
// "start", try every possible first piece s[start..end]. If that piece is
// a palindrome, keep it and cut the rest of the string the same way.
// When we reach the end of the string, the pieces form one answer.
func partition(s string) [][]string {
	res := [][]string{}
	current := []string{}

	var build func(start int)
	build = func(start int) {
		if start == len(s) {
			// the whole string is cut: save a copy of the pieces
			res = append(res, append([]string{}, current...))
			return
		}

		for end := start; end < len(s); end++ {
			if isPalindromePiece(s, start, end) {
				current = append(current, s[start:end+1]) // choose this piece
				build(end + 1)                            // cut the rest
				current = current[:len(current)-1]        // undo the choice
			}
		}
	}

	build(0)
	return res
}

// isPalindromePiece checks if s[left..right] reads the same in both directions.
func isPalindromePiece(s string, left int, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
