// Idea (two pointers): the list is sorted. Start with the smallest and the
// biggest number.
//   - sum too small -> move the left pointer right (bigger number)
//   - sum too big   -> move the right pointer left (smaller number)
// The answer uses positions that start at 1, not 0.
// Time: O(n), no extra memory.
func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1} // positions start at 1
		}
		if sum < target {
			left++
		} else {
			right--
		}
	}
	return []int{}
}
