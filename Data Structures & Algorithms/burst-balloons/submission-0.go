// Idea: thinking about which balloon to burst FIRST is hard, because the
// neighbors change. Think about which balloon is burst LAST instead.
// First put a balloon with value 1 at both ends (they are never burst).
// best[left][right] = the most coins from bursting all balloons strictly
// BETWEEN left and right.
// If balloon k is the last one burst in that range, then at that moment
// its neighbors are exactly left and right, so we get
//   best[left][k] + best[k][right] + value[left]*value[k]*value[right]
// Try every k and keep the best. Short ranges are computed first.
// Time: O(n^3).
func maxCoins(nums []int) int {
	// balloons with a 1 on both ends
	balloons := []int{1}
	balloons = append(balloons, nums...)
	balloons = append(balloons, 1)
	n := len(balloons)

	best := make([][]int, n)
	for i := range best {
		best[i] = make([]int, n)
	}

	// gap = distance between left and right; start with the smallest ranges
	for gap := 2; gap < n; gap++ {
		for left := 0; left+gap < n; left++ {
			right := left + gap

			for k := left + 1; k < right; k++ {
				coins := best[left][k] + best[k][right] + balloons[left]*balloons[k]*balloons[right]
				if coins > best[left][right] {
					best[left][right] = coins
				}
			}
		}
	}
	return best[0][n-1]
}
