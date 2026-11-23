// Idea (sliding window):
//   need[ch] = how many more of ch the window still needs
//              (it becomes negative when the window has extra)
//   missing  = how many letters of t are still missing in total
// Grow the window to the right. When nothing is missing the window is
// valid, so shrink it from the left as long as it stays valid and
// remember the smallest one.
// Time: O(len(s) + len(t)).
func minWindow(s string, t string) string {
	need := map[byte]int{}
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	missing := len(t)

	bestStart := 0
	bestLen := 0 // 0 means we have not found any window yet
	left := 0

	for right := 0; right < len(s); right++ {
		// add s[right] to the window
		if need[s[right]] > 0 {
			missing-- // it was a letter we still needed
		}
		need[s[right]]--

		// the window has everything: try to make it smaller
		for missing == 0 {
			size := right - left + 1
			if bestLen == 0 || size < bestLen {
				bestStart = left
				bestLen = size
			}

			// remove s[left] from the window
			need[s[left]]++
			if need[s[left]] > 0 {
				missing++ // we just lost a letter we need
			}
			left++
		}
	}
	return s[bestStart : bestStart+bestLen]
}
