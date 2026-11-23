// Idea (BFS): BFS finds the shortest way. Start with beginWord. In every
// step, go from each word to all unused words that differ by exactly one
// letter. The first time we reach endWord, the number of steps is the
// shortest possible.
// Time: O(n^2 * word length), fine for the small word lists here.
func ladderLength(beginWord string, endWord string, wordList []string) int {
	used := make([]bool, len(wordList))
	queue := []string{beginWord}
	length := 1 // number of words in the chain so far

	for len(queue) > 0 {
		next := []string{}

		for _, word := range queue {
			if word == endWord {
				return length
			}

			// every unused word one letter away comes in the next round
			for i, candidate := range wordList {
				if !used[i] && differByOne(word, candidate) {
					used[i] = true
					next = append(next, candidate)
				}
			}
		}

		queue = next
		length++
	}
	return 0 // endWord was never reached
}

// differByOne is true if the two words differ in exactly one position.
func differByOne(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}
	different := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			different++
		}
	}
	return different == 1
}
