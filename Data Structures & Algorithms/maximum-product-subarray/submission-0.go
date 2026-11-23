// Idea: like Maximum Subarray, but with multiplication a negative number
// can turn the SMALLEST product into the BIGGEST one (negative * negative).
// So for the subarrays ending at the current number we remember both
// the biggest product (high) and the smallest product (low).
// Time: O(n).
func maxProduct(nums []int) int {
	best := nums[0]
	high := 1 // biggest product of a subarray ending at the previous number
	low := 1  // smallest product of a subarray ending at the previous number

	for _, n := range nums {
		// three choices for a subarray ending here:
		// start fresh with n, or extend high, or extend low
		a := high * n
		b := low * n

		high = n
		if a > high {
			high = a
		}
		if b > high {
			high = b
		}

		low = n
		if a < low {
			low = a
		}
		if b < low {
			low = b
		}

		if high > best {
			best = high
		}
	}
	return best
}
