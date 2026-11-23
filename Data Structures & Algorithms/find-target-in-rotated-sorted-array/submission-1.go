// Idea (binary search): when we cut a rotated sorted array in the
// middle, at least ONE half is normally sorted. For the sorted half it
// is easy to check whether the target is inside it (between its ends).
// If yes, search that half; if no, search the other half.
// Time: O(log n).
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[left] <= nums[mid] {
			// the left half is sorted
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1 // target is inside the left half
			} else {
				left = mid + 1
			}
		} else {
			// the right half is sorted
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1 // target is inside the right half
			} else {
				right = mid - 1
			}
		}
	}
	return -1 // not found
}
