// Idea: keep all numbers in a slice that is ALWAYS sorted.
// AddNum puts the new number at the end and moves it left until it is
// in the right place (one step of insertion sort).
// FindMedian just reads the middle of the sorted slice.
//
// AddNum is O(n), FindMedian is O(1). This is the simple version.
// The faster version uses two heaps (O(log n) per AddNum), but it is
// much more code.

type MedianFinder struct {
	nums []int // always sorted
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	this.nums = append(this.nums, num)

	// move the new number left while it is smaller than its left neighbor
	i := len(this.nums) - 1
	for i > 0 && this.nums[i] < this.nums[i-1] {
		this.nums[i], this.nums[i-1] = this.nums[i-1], this.nums[i]
		i--
	}
}

func (this *MedianFinder) FindMedian() float64 {
	n := len(this.nums)
	if n%2 == 1 {
		// odd count: the middle number
		return float64(this.nums[n/2])
	}
	// even count: average of the two middle numbers
	return float64(this.nums[n/2-1]+this.nums[n/2]) / 2.0
}
