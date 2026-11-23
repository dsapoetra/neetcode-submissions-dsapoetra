// Idea: keep four walls: top, bottom, left, right. Walk around the
// outside (right, down, left, up), and after finishing each side move
// that wall inwards. Repeat until the walls cross.
// Time: O(rows * cols).
func spiralOrder(matrix [][]int) []int {
	res := []int{}
	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1

	for top <= bottom && left <= right {
		// go right along the top row
		for c := left; c <= right; c++ {
			res = append(res, matrix[top][c])
		}
		top++

		// go down along the right column
		for r := top; r <= bottom; r++ {
			res = append(res, matrix[r][right])
		}
		right--

		// go left along the bottom row (only if there is still a row left)
		if top <= bottom {
			for c := right; c >= left; c-- {
				res = append(res, matrix[bottom][c])
			}
			bottom--
		}

		// go up along the left column (only if there is still a column left)
		if left <= right {
			for r := bottom; r >= top; r-- {
				res = append(res, matrix[r][left])
			}
			left++
		}
	}
	return res
}
