// Sliding window: if a char repeats inside the window, move left past it.
func lengthOfLongestSubstring(s string) int {
	lastSeen := map[byte]int{}
	left, best := 0, 0
	for right := 0; right < len(s); right++ {
		idx, ok := lastSeen[s[right]]
		if ok && idx >= left {
			left = idx + 1
		}
		lastSeen[s[right]] = right
		if right-left+1 > best {
			best = right - left + 1
		}
	}
	return best
}
