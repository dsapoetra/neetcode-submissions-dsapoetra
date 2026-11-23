// Idea: add like on paper, but in binary and without using +.
//   a ^ b        adds each bit WITHOUT the carry (1+1 gives 0)
//   (a & b) << 1 is the carry (1+1 carries a 1 to the next position)
// Then we still have to add the carry, so repeat until no carry is left.
func getSum(a int, b int) int {
	for b != 0 {
		carry := (a & b) << 1
		a = a ^ b // sum without carry
		b = carry // the carry still has to be added
	}
	return a
}
