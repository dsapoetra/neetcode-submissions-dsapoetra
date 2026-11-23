// Kadane: keep a running sum, restart it when it goes negative.
func maxSubArray(nums []int) int {
	best := nums[0]
	cur := 0
	for _, n := range nums {
		if cur < 0 {
			cur = 0
		}
		cur += n
		if cur > best {
			best = cur
		}
	}
	return best
}
