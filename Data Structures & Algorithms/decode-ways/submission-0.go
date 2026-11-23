// Idea (dynamic programming): ways[i] = number of ways to decode the
// first i digits. The last letter uses either:
//   - one digit  (must be 1..9)   -> add ways[i-1]
//   - two digits (must be 10..26) -> add ways[i-2]
// Time: O(n).
func numDecodings(s string) int {
	ways := make([]int, len(s)+1)
	ways[0] = 1 // empty string: one way (decode nothing)

	for i := 1; i <= len(s); i++ {
		// last ONE digit as a letter
		oneDigit := int(s[i-1] - '0')
		if oneDigit >= 1 {
			ways[i] += ways[i-1]
		}

		// last TWO digits as a letter
		if i >= 2 {
			twoDigits := int(s[i-2]-'0')*10 + oneDigit
			if twoDigits >= 10 && twoDigits <= 26 {
				ways[i] += ways[i-2]
			}
		}
	}
	return ways[len(s)]
}
