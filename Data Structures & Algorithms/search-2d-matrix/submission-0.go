// Idea: every row is sorted and each row starts after the previous row
// ends, so the whole matrix is ONE sorted list written row by row.
// Do a normal binary search on positions 0 .. rows*cols-1 and turn a
// position back into (row, column):
//   row = position / cols,  column = position % cols
// Time: O(log(rows * cols)).
func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])

	left := 0
	right := rows*cols - 1

	for left <= right {
		mid := (left + right) / 2
		value := matrix[mid/cols][mid%cols]

		if value == target {
			return true
		}
		if value < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
