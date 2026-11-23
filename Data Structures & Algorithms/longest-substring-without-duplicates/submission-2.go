// Idea (sliding window): the window s[left..right] never contains a
// repeated character. Move right one step at a time. If the new
// character is already inside the window, move left to just after
// its previous position, so the window is valid again.
// Time: O(n).
func lengthOfLongestSubstring(s string) int {
	lastSeen := map[byte]int{} // character -> last index where we saw it
	left := 0
	best := 0

	for right := 0; right < len(s); right++ {
		ch := s[right]

		// seen before AND still inside the window -> jump left past it
		index, seen := lastSeen[ch]
		if seen && index >= left {
			left = index + 1
		}

		lastSeen[ch] = right

		size := right - left + 1
		if size > best {
			best = size
		}
	}
	return best
}
