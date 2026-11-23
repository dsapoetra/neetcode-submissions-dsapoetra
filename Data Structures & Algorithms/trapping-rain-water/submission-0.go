// Idea: the water above one bar is limited by the tallest bar on its left
// and the tallest bar on its right. The lower of those two is the water
// level, so:
//   water above bar i = min(tallest left, tallest right) - height[i]
// First compute "tallest bar so far" from the left and from the right,
// then add up the water above every bar.
// Time: O(n), extra space: O(n).
func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	// maxLeft[i] = tallest bar from index 0 to i
	maxLeft := make([]int, n)
	maxLeft[0] = height[0]
	for i := 1; i < n; i++ {
		maxLeft[i] = maxLeft[i-1]
		if height[i] > maxLeft[i] {
			maxLeft[i] = height[i]
		}
	}

	// maxRight[i] = tallest bar from index i to the end
	maxRight := make([]int, n)
	maxRight[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		maxRight[i] = maxRight[i+1]
		if height[i] > maxRight[i] {
			maxRight[i] = height[i]
		}
	}

	water := 0
	for i := 0; i < n; i++ {
		level := maxLeft[i]
		if maxRight[i] < level {
			level = maxRight[i]
		}
		water += level - height[i]
	}
	return water
}
