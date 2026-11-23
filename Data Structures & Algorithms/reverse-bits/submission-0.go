// Idea: do this 32 times: take the last bit of n and push it onto the
// end of the result. The first bit we take ends up at the far left,
// so the bits come out reversed.
func reverseBits(n int) int {
	result := 0
	for i := 0; i < 32; i++ {
		lastBit := n & 1
		result = (result << 1) | lastBit // make room, then add the bit
		n = n >> 1                       // drop the bit we just used
	}
	return result
}
