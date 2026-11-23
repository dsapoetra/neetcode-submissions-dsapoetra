// Idea (sliding window): a window is OK if we can make all its letters
// the same with at most k changes. The cheapest way is to keep the most
// common letter and change the rest, so:
//   changes needed = window length - count of most common letter
// Grow the window to the right; while it needs more than k changes,
// shrink it from the left.
// Time: O(26 * n).
func characterReplacement(s string, k int) int {
	count := make([]int, 26) // count[i] = how many of letter 'A'+i in the window
	left := 0
	best := 0

	for right := 0; right < len(s); right++ {
		// add the new letter to the window
		count[s[right]-'A']++

		// shrink while the window needs more than k changes
		for (right-left+1)-mostCommon(count) > k {
			count[s[left]-'A']--
			left++
		}

		if right-left+1 > best {
			best = right - left + 1
		}
	}
	return best
}

// mostCommon returns the biggest value in count.
func mostCommon(count []int) int {
	biggest := 0
	for _, c := range count {
		if c > biggest {
			biggest = c
		}
	}
	return biggest
}
