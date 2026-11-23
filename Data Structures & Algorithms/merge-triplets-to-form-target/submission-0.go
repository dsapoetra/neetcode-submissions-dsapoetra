// Idea: merging takes the MAXIMUM of every position, so values can only
// grow. A triplet with any value bigger than the target would break the
// target forever, so we may never use it.
// All the other triplets are safe to merge. So: look only at the safe
// triplets and check that each of the three target values appears in at
// least one of them at the right position.
// Time: O(n).
func mergeTriplets(triplets [][]int, target []int) bool {
	found := [3]bool{} // found[i] = some safe triplet has target[i] at position i

	for _, triplet := range triplets {
		// one value too big: this triplet can never be used
		if triplet[0] > target[0] || triplet[1] > target[1] || triplet[2] > target[2] {
			continue
		}

		for i := 0; i < 3; i++ {
			if triplet[i] == target[i] {
				found[i] = true
			}
		}
	}
	return found[0] && found[1] && found[2]
}
