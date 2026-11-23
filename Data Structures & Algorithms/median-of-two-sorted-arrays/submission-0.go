// Idea: both lists are sorted, so we can walk through them together in
// sorted order (like the merge step of merge sort) WITHOUT building the
// merged list. We only need to walk until the middle and remember the
// last two numbers we passed.
//   - odd total:  the median is the middle number
//   - even total: the median is the average of the two middle numbers
//
// This is the simple version: O(m + n).
// The faster version (binary search, O(log(m + n))) is much harder to get right.
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	middle := total / 2

	i, j := 0, 0 // positions in nums1 and nums2
	previous := 0
	current := 0

	// take the next smallest number, middle+1 times
	for step := 0; step <= middle; step++ {
		previous = current

		if j >= len(nums2) || (i < len(nums1) && nums1[i] <= nums2[j]) {
			// nums2 is used up, or the next number of nums1 is smaller
			current = nums1[i]
			i++
		} else {
			current = nums2[j]
			j++
		}
	}

	if total%2 == 1 {
		return float64(current)
	}
	return float64(previous+current) / 2.0
}
