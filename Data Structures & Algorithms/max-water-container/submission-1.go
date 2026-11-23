// Idea (two pointers): start with the widest container (first and last
// wall). The water height is limited by the SHORTER wall. Moving the
// taller wall inwards can never help (width shrinks, height cannot grow),
// so always move the shorter wall and hope for a taller one.
// Time: O(n).
func maxArea(heights []int) int {
	left := 0
	right := len(heights) - 1
	best := 0

	for left < right {
		// water height = the shorter of the two walls
		height := heights[left]
		if heights[right] < height {
			height = heights[right]
		}

		area := height * (right - left)
		if area > best {
			best = area
		}

		// move the shorter wall inwards
		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}
	return best
}
