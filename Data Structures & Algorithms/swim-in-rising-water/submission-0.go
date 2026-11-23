// Idea: at time t we can stand on every cell with height <= t.
// If we can reach the end at time t, we can also reach it at any later
// time. So the times look like: no, no, no, YES, YES, YES ...
// and we can binary search the first YES.
// For one time t, a simple DFS tells us if the end is reachable.
// Time: O(n^2 * log(n^2)).
func swimInWater(grid [][]int) int {
	n := len(grid)

	left := 0
	right := n*n - 1 // the heights are 0 .. n*n-1, so this time always works

	for left < right {
		mid := (left + right) / 2

		visited := make([][]bool, n)
		for i := range visited {
			visited[i] = make([]bool, n)
		}

		if canSwim(grid, visited, 0, 0, mid) {
			right = mid // works: maybe an earlier time works too
		} else {
			left = mid + 1 // too early
		}
	}
	return left
}

// canSwim is true if the bottom-right cell can be reached from (r, c)
// using only cells with height <= time.
func canSwim(grid [][]int, visited [][]bool, r int, c int, time int) bool {
	n := len(grid)
	if r < 0 || r >= n || c < 0 || c >= n || visited[r][c] || grid[r][c] > time {
		return false
	}
	if r == n-1 && c == n-1 {
		return true
	}
	visited[r][c] = true

	return canSwim(grid, visited, r+1, c, time) ||
		canSwim(grid, visited, r-1, c, time) ||
		canSwim(grid, visited, r, c+1, time) ||
		canSwim(grid, visited, r, c-1, time)
}
