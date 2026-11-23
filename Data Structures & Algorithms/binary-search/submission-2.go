// Idea: the list is sorted, so look at the middle number.
//   - it is the target      -> found
//   - it is too small       -> the target can only be on the right half
//   - it is too big         -> the target can only be on the left half
// Every step throws away half of the list.
// Time: O(log n).
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1 // search the right half
		} else {
			right = mid - 1 // search the left half
		}
	}
	return -1 // not found
}
