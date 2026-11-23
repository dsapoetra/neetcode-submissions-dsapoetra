// Idea: two parts with the same sum means each part is HALF of the total.
// So the question is: can some of the numbers add up to total / 2?
// possible[s] = true if some of the numbers seen so far add up to s.
// For each number, every sum s that was possible makes s + num possible.
// We go through the sums from big to small so that one number is not
// used twice in the same round.
// Time: O(n * total / 2).
func canPartition(nums []int) bool {
	total := 0
	for _, num := range nums {
		total += num
	}
	if total%2 == 1 {
		return false // an odd total cannot be split into two equal parts
	}
	half := total / 2

	possible := make([]bool, half+1)
	possible[0] = true // taking nothing gives sum 0

	for _, num := range nums {
		for sum := half; sum >= num; sum-- {
			if possible[sum-num] {
				possible[sum] = true
			}
		}
	}
	return possible[half]
}
