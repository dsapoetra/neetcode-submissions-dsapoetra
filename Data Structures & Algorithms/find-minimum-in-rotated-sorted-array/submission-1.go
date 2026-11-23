// Idea (binary search): compare the middle number with the rightmost one.
//   - middle > right: the "drop" is to the right of middle, so the
//     minimum is in the right half (and it is not middle)
//   - otherwise: the right half is sorted, so the minimum is middle
//     or something to its left
// Time: O(log n).
func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1 // minimum is to the right of mid
		} else {
			right = mid // minimum is mid or to its left
		}
	}

	// left == right: only one candidate remains
	return nums[left]
}
