// Idea (dynamic programming): canBreak[i] = true if the first i letters
// of s can be split into dictionary words.
// The first i letters work if some word w ends exactly at i AND the
// letters before that word (the first i-len(w) letters) also work.
// Time: O(len(s) * number of words * word length).
func wordBreak(s string, wordDict []string) bool {
	canBreak := make([]bool, len(s)+1)
	canBreak[0] = true // the empty string needs no words

	for i := 1; i <= len(s); i++ {
		for _, word := range wordDict {
			start := i - len(word) // where this word would begin
			if start >= 0 && canBreak[start] && s[start:i] == word {
				canBreak[i] = true
				break // one way is enough
			}
		}
	}
	return canBreak[len(s)]
}
