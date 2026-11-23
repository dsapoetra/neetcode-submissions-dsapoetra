func twoSum(nums []int, target int) []int {
	indexOf := make(map[int]int)

	for i, num := range nums {
		need := target - num

		j, ok := indexOf[need]
		if ok {
			return []int{j, i}
		}

		indexOf[num] = i
	}

	return nil
}