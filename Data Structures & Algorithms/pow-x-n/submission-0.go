// Idea (fast power): x^n = (x^(n/2))^2. So we only need to compute the
// half power once and square it. If n is odd, multiply by one extra x.
// That needs about log2(n) multiplications instead of n.
// A negative n means 1 / x^(-n).
// Time: O(log n).
func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / power(x, -n)
	}
	return power(x, n)
}

// power computes x^n for n >= 0.
func power(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	half := power(x, n/2)
	if n%2 == 0 {
		return half * half
	}
	return half * half * x
}
