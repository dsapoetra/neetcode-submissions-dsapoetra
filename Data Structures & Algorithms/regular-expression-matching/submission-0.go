// Idea (dynamic programming):
// match[i][j] = true if the first i letters of s match the first j
// characters of the pattern p.
// Look at the last pattern character p[j-1]:
//   - a letter or '.':  the last letter of s must fit, and the rest must
//                       match:  match[i-1][j-1]
//   - a '*' (it belongs to the character before it, p[j-2]):
//       use it ZERO times: drop "x*" from the pattern:  match[i][j-2]
//       use it ONE MORE time: the last letter of s must fit p[j-2], and
//       the rest of s must still match the same pattern: match[i-1][j]
// Time: O(len(s) * len(p)).
func isMatch(s string, p string) bool {
	match := make([][]bool, len(s)+1)
	for i := range match {
		match[i] = make([]bool, len(p)+1)
	}
	match[0][0] = true // empty text matches empty pattern

	for i := 0; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if p[j-1] == '*' {
				// zero times
				if match[i][j-2] {
					match[i][j] = true
				}
				// one more time
				if i > 0 && (p[j-2] == '.' || p[j-2] == s[i-1]) && match[i-1][j] {
					match[i][j] = true
				}
			} else {
				// a normal letter or '.'
				if i > 0 && (p[j-1] == '.' || p[j-1] == s[i-1]) && match[i-1][j-1] {
					match[i][j] = true
				}
			}
		}
	}
	return match[len(s)][len(p)]
}
