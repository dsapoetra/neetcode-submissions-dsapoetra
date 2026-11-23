// Idea: sort the numbers. The k-th largest is then the k-th number
// counted from the end.
// Time: O(n log n).
func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}
