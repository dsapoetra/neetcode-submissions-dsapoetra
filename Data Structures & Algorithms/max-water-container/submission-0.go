// Two pointers: always move the shorter wall inward.
func maxArea(heights []int) int {
	left, right := 0, len(heights)-1
	best := 0
	for left < right {
		h := heights[left]
		if heights[right] < h {
			h = heights[right]
		}
		area := h * (right - left)
		if area > best {
			best = area
		}
		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}
	return best
}
