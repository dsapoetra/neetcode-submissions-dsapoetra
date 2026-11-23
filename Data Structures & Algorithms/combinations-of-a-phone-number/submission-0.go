// Idea: start with one empty string. For every digit, take all the
// strings we have so far and add each letter of that digit to each of
// them. Example "23": [""] -> ["a","b","c"] -> ["ad","ae","af","bd",...]
// Time: O(4^n), a digit has at most 4 letters.
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	// letters[d] = the letters on phone key d
	letters := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	res := []string{""}
	for i := 0; i < len(digits); i++ {
		options := letters[digits[i]-'0']

		next := []string{}
		for _, prefix := range res {
			for _, letter := range options {
				next = append(next, prefix+string(letter))
			}
		}
		res = next
	}
	return res
}
