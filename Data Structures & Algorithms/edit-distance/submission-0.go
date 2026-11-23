// Idea (dynamic programming):
// dist[i][j] = fewest steps to turn the first i letters of word1 into the
// first j letters of word2.
//   - if the last letters are equal, nothing to do:  dist[i-1][j-1]
//   - otherwise 1 step plus the best of:
//       replace the last letter:  dist[i-1][j-1]
//       delete from word1:        dist[i-1][j]
//       insert into word1:        dist[i][j-1]
// Start values: turning i letters into nothing takes i deletes, turning
// nothing into j letters takes j inserts.
// Time: O(len(word1) * len(word2)).
func minDistance(word1 string, word2 string) int {
	dist := make([][]int, len(word1)+1)
	for i := range dist {
		dist[i] = make([]int, len(word2)+1)
		dist[i][0] = i // delete everything
	}
	for j := 0; j <= len(word2); j++ {
		dist[0][j] = j // insert everything
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dist[i][j] = dist[i-1][j-1] // same letter, no step needed
				continue
			}

			best := dist[i-1][j-1] // replace
			if dist[i-1][j] < best {
				best = dist[i-1][j] // delete
			}
			if dist[i][j-1] < best {
				best = dist[i][j-1] // insert
			}
			dist[i][j] = best + 1
		}
	}
	return dist[len(word1)][len(word2)]
}
