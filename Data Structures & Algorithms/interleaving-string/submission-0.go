// Idea (dynamic programming):
// ok[i][j] = true if the first i letters of s1 and the first j letters of
// s2 can be mixed to build the first i+j letters of s3.
// The last letter of that part of s3 comes either
//   - from s1: then s1[i-1] must equal it and ok[i-1][j] must be true
//   - from s2: then s2[j-1] must equal it and ok[i][j-1] must be true
// Time: O(len(s1) * len(s2)).
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	ok := make([][]bool, len(s1)+1)
	for i := range ok {
		ok[i] = make([]bool, len(s2)+1)
	}
	ok[0][0] = true // two empty strings build the empty string

	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			// take the last letter from s1
			if i > 0 && s1[i-1] == s3[i+j-1] && ok[i-1][j] {
				ok[i][j] = true
			}
			// take the last letter from s2
			if j > 0 && s2[j-1] == s3[i+j-1] && ok[i][j-1] {
				ok[i][j] = true
			}
		}
	}
	return ok[len(s1)][len(s2)]
}
