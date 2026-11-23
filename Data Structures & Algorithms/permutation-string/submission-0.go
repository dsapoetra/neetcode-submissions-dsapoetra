// Idea (sliding window): a permutation of s1 has exactly the same letter
// counts as s1. So slide a window of length len(s1) over s2 and compare
// the letter counts of the window with the letter counts of s1.
// When the window moves one step, one letter enters and one letter leaves.
// Time: O(26 * len(s2)).
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	var need [26]int   // letter counts of s1
	var window [26]int // letter counts of the current window in s2

	// the first window
	for i := 0; i < len(s1); i++ {
		need[s1[i]-'a']++
		window[s2[i]-'a']++
	}
	if need == window { // Go can compare two arrays directly
		return true
	}

	// slide the window: add the letter on the right, remove the one on the left
	for right := len(s1); right < len(s2); right++ {
		window[s2[right]-'a']++
		window[s2[right-len(s1)]-'a']--
		if need == window {
			return true
		}
	}
	return false
}
