// Idea: look at the last bit (n & 1), add it to the count,
// then shift n right by one to look at the next bit. Stop when n is 0.
func hammingWeight(n int) int {
	count := 0
	for n != 0 {
		count += n & 1 // 1 if the last bit is set, otherwise 0
		n = n >> 1     // drop the last bit
	}
	return count
}
