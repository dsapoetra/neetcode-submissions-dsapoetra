// Idea: the words are sorted in the alien order. Compare each word with
// the next one: the FIRST letter where they differ tells us one rule,
// "letter x comes before letter y".
// The rules form a graph, and a valid alphabet is a topological sort
// of it (same BFS method as Course Schedule). If we cannot place every
// letter, the rules contain a cycle, so there is no valid order.
func foreignDictionary(words []string) string {
	after := map[byte][]byte{} // after[x] = letters that must come after x
	need := map[byte]int{}     // need[y] = how many letters must come before y

	// register every letter that appears
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			need[word[i]] = 0
		}
	}

	// get one rule from each pair of neighboring words
	for i := 0; i < len(words)-1; i++ {
		first := words[i]
		second := words[i+1]

		shorter := len(first)
		if len(second) < shorter {
			shorter = len(second)
		}

		foundDifference := false
		for j := 0; j < shorter; j++ {
			if first[j] != second[j] {
				// first[j] comes before second[j]
				after[first[j]] = append(after[first[j]], second[j])
				need[second[j]]++
				foundDifference = true
				break
			}
		}

		// like "apple" before "app": that is never a valid order
		if !foundDifference && len(first) > len(second) {
			return ""
		}
	}

	// start with the letters that have nothing before them
	queue := []byte{}
	for letter, count := range need {
		if count == 0 {
			queue = append(queue, letter)
		}
	}

	order := []byte{}
	for len(queue) > 0 {
		letter := queue[0]
		queue = queue[1:]
		order = append(order, letter)

		// every letter after this one now waits for one less letter
		for _, next := range after[letter] {
			need[next]--
			if need[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	// some letters were never placed -> there is a cycle
	if len(order) != len(need) {
		return ""
	}
	return string(order)
}
