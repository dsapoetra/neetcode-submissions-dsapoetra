// Idea (2D dynamic programming):
// lcs[i][j] = answer for the first i letters of text1 and the first j
// letters of text2.
//   - if the last letters are the same, use it: lcs[i-1][j-1] + 1
//   - otherwise drop one last letter from either text and take the better
// Time: O(len(text1) * len(text2)).
func longestCommonSubsequence(text1 string, text2 string) int {
	// one extra row and column of zeros for the "empty text" case
	lcs := make([][]int, len(text1)+1)
	for i := 0; i <= len(text1); i++ {
		lcs[i] = make([]int, len(text2)+1)
	}

	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				// same letter: it extends the common subsequence
				lcs[i][j] = lcs[i-1][j-1] + 1
			} else if lcs[i-1][j] > lcs[i][j-1] {
				lcs[i][j] = lcs[i-1][j] // drop last letter of text1
			} else {
				lcs[i][j] = lcs[i][j-1] // drop last letter of text2
			}
		}
	}
	return lcs[len(text1)][len(text2)]
}
