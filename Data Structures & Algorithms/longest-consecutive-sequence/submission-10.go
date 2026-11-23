func longestConsecutive(nums []int) int {
	seen := make(map[int]bool)
	for _, n := range nums {
		seen[n] = true
	}

	longest := 0
	for n := range seen {
		// Only start counting from the beginning of a run.
		if seen[n-1] {
			continue
		}

		length := 1
		for seen[n+length] {
			length++
		}

		if length > longest {
			longest = length
		}
	}

	return longest
}