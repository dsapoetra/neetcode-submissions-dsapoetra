// Idea (backtracking): build the string one bracket at a time.
//   - we may add "(" while we have used fewer than n of them
//   - we may add ")" only if there is an open bracket waiting for it
//     (fewer ")" than "(" so far)
// Following these two rules, every finished string is valid.
func generateParenthesis(n int) []string {
	res := []string{}

	var build func(current string, open int, closed int)
	build = func(current string, open int, closed int) {
		if len(current) == 2*n {
			res = append(res, current)
			return
		}
		if open < n {
			build(current+"(", open+1, closed)
		}
		if closed < open {
			build(current+")", open, closed+1)
		}
	}

	build("", 0, 0)
	return res
}
