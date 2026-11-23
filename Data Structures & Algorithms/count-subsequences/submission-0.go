// Idea (dynamic programming):
// ways[i][j] = number of ways to build the first j letters of t from the
// first i letters of s.
// Look at letter s[i-1]:
//   - we can always SKIP it:                       ways[i-1][j]
//   - if it equals t[j-1] we can also USE it:      ways[i-1][j-1]
// Building the empty string (j = 0) always has exactly 1 way: take nothing.
// Time: O(len(s) * len(t)).
func numDistinct(s string, t string) int {
	ways := make([][]int, len(s)+1)
	for i := range ways {
		ways[i] = make([]int, len(t)+1)
		ways[i][0] = 1 // the empty t: one way
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			ways[i][j] = ways[i-1][j] // skip s[i-1]
			if s[i-1] == t[j-1] {
				ways[i][j] += ways[i-1][j-1] // use s[i-1] for t[j-1]
			}
		}
	}
	return ways[len(s)][len(t)]
}
