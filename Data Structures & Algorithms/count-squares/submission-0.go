// Idea: remember how often each point was added (count), plus a list of
// all added points.
// A square is fixed by two opposite corners. For a query point (x, y),
// try every added point (px, py) as the DIAGONAL corner. It only works if
// the horizontal and vertical distance are the same (and not 0).
// The other two corners are then (x, py) and (px, y). Each of them may
// exist several times, so multiply their counts.
// Going through the list of points (with duplicates) counts a diagonal
// point that was added twice two times, which is what we want.
// Add: O(1). Count: O(number of added points).

type CountSquares struct {
	count  map[[2]int]int // how many times each point was added
	points [][2]int       // every added point, duplicates included
}

func Constructor() CountSquares {
	return CountSquares{count: map[[2]int]int{}}
}

func (this *CountSquares) Add(point []int) {
	p := [2]int{point[0], point[1]}
	this.count[p]++
	this.points = append(this.points, p)
}

func (this *CountSquares) Count(point []int) int {
	x, y := point[0], point[1]
	total := 0

	for _, p := range this.points {
		px, py := p[0], p[1]

		// p must be a diagonal corner: same distance in x and in y, and not 0
		dx := px - x
		dy := py - y
		if dx == 0 || (dx != dy && dx != -dy) {
			continue
		}

		// the two remaining corners
		total += this.count[[2]int{x, py}] * this.count[[2]int{px, y}]
	}
	return total
}
