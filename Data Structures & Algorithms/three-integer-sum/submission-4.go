// Idea: sort the numbers. Then fix the first number nums[i] and find
// the other two with two pointers (left and right) that move towards
// each other:
//   - sum too small -> move left forward (bigger number)
//   - sum too big   -> move right back (smaller number)
//   - sum is 0      -> save the triplet
// Skipping equal numbers avoids duplicate triplets.
// Time: O(n^2).
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		// same first number as before -> would give the same triplets
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum < 0 {
				left++ // need a bigger sum
			} else if sum > 0 {
				right-- // need a smaller sum
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				// skip the same second number
				for left < right && nums[left] == nums[left-1] {
					left++
				}
			}
		}
	}
	return res
}
