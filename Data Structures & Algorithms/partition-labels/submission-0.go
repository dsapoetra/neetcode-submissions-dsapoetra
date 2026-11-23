// Idea: a part must contain ALL copies of every letter in it. So first
// remember the last position of each letter.
// Then walk through the string. The current part must reach at least to
// the last position of every letter we have seen in it (end).
// When we arrive at "end", no letter of this part appears later, so we
// can cut here.
// Time: O(n).
func partitionLabels(s string) []int {
	var last [26]int // last[x] = last index of letter x
	for i := 0; i < len(s); i++ {
		last[s[i]-'a'] = i
	}

	res := []int{}
	start := 0 // where the current part begins
	end := 0   // where the current part must at least end

	for i := 0; i < len(s); i++ {
		if last[s[i]-'a'] > end {
			end = last[s[i]-'a']
		}

		if i == end {
			res = append(res, end-start+1) // cut here
			start = i + 1
		}
	}
	return res
}
