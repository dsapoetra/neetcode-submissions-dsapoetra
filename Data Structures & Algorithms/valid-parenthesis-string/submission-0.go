// Idea: we do not know what each '*' is, so we track a RANGE for
// "how many open brackets are waiting to be closed":
//   low  = the smallest possible number
//   high = the biggest possible number
//   '('  -> both go up
//   ')'  -> both go down
//   '*'  -> low goes down (it acts as ')'), high goes up (it acts as '(')
// If high drops below 0, there are too many ')' no matter what.
// low can never be below 0 (we just would not use a '*' as ')').
// At the end the string is valid if 0 is possible, that is low == 0.
// Time: O(n).
func checkValidString(s string) bool {
	low := 0
	high := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			low++
			high++
		} else if s[i] == ')' {
			low--
			high--
		} else { // '*'
			low--
			high++
		}

		if high < 0 {
			return false // too many ')' even in the best case
		}
		if low < 0 {
			low = 0
		}
	}
	return low == 0
}
